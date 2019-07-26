package block

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"testing"
	"time"

	"github.com/ElrondNetwork/elrond-go/core/logger"
	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/data/transaction"
	"github.com/ElrondNetwork/elrond-go/integrationTests"
	"github.com/ElrondNetwork/elrond-go/sharding"
	"github.com/stretchr/testify/assert"
)

var agarioFile = "agarioV2.hex"
var stepDelay = time.Second

func TestShouldProcessBlocksWithScTxsJoinAndRewardTheOwner(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	log := logger.DefaultLogger()
	log.SetLevel(logger.LogDebug)

	scCode, err := ioutil.ReadFile(agarioFile)
	assert.Nil(t, err)

	maxShards := uint32(1)
	numOfNodes := 1
	advertiser := integrationTests.CreateMessengerWithKadDht(context.Background(), "")
	_ = advertiser.Bootstrap()
	advertiserAddr := integrationTests.GetConnectableAddress(advertiser)

	nodes := make([]*integrationTests.TestProcessorNode, numOfNodes)
	for i := 0; i < numOfNodes; i++ {
		nodes[i] = integrationTests.NewTestProcessorNode(maxShards, 0, 0, advertiserAddr)
	}

	idxProposer := 0
	hardCodedSk, _ := hex.DecodeString("5561d28b0d89fa425bbbf9e49a018b5d1e4a462c03d2efce60faf9ddece2af06")
	hardCodedScResultingAddress, _ := hex.DecodeString("000000000000000000005fed9c659422cd8429ce92f8973bba2a9fb51e0eb3a1")
	nodes[idxProposer].LoadTxSignSkBytes(hardCodedSk)

	defer func() {
		_ = advertiser.Close()
		for _, n := range nodes {
			_ = n.Messenger.Close()
		}
	}()

	for _, n := range nodes {
		_ = n.Messenger.Bootstrap()
	}

	fmt.Println("Delaying for nodes p2p bootstrap...")
	time.Sleep(stepDelay)

	round := uint32(0)
	round = incrementAndPrintRound(round)

	initialVal := big.NewInt(10000000)
	topUpValue := big.NewInt(500)
	withdrawValue := big.NewInt(10)
	stepMintAllNodes(nodes, initialVal)

	stepDeployScTx(nodes, idxProposer, string(scCode))
	stepProposeBlock(nodes, idxProposer, round)
	stepSyncBlock(t, nodes, idxProposer, round)
	round = incrementAndPrintRound(round)

	stepNodeDoesJoinGame(nodes, idxProposer, topUpValue, hardCodedScResultingAddress)
	stepProposeBlock(nodes, idxProposer, round)
	stepSyncBlock(t, nodes, idxProposer, round)
	round = incrementAndPrintRound(round)

	stepCheckJoinGameIsDoneCorrectly(
		t,
		nodes,
		idxProposer,
		idxProposer,
		initialVal,
		topUpValue,
		hardCodedScResultingAddress,
	)

	stepNodeCallsRewardAndSend(nodes, idxProposer, idxProposer, withdrawValue, hardCodedScResultingAddress)
	stepProposeBlock(nodes, idxProposer, round)
	stepSyncBlock(t, nodes, idxProposer, round)
	round = incrementAndPrintRound(round)

	stepProposeBlock(nodes, idxProposer, round)
	stepSyncBlock(t, nodes, idxProposer, round)
	round = incrementAndPrintRound(round)

	stepCheckRewardIsDoneCorrectly(
		t,
		nodes,
		idxProposer,
		idxProposer,
		initialVal,
		topUpValue,
		withdrawValue,
		hardCodedScResultingAddress,
	)

	stepCheckRootHashes(t, nodes, []int{0})

	time.Sleep(1 * time.Second)
}

func incrementAndPrintRound(round uint32) uint32 {
	round++
	fmt.Printf("#################################### ROUND %d BEGINS ####################################\n\n", round)

	return round
}

func stepMintAllNodes(nodes []*integrationTests.TestProcessorNode, value *big.Int) {
	for _, n := range nodes {
		if n.ShardCoordinator.SelfId() == sharding.MetachainShardId {
			continue
		}

		for _, n2 := range nodes {
			integrationTests.MintAddress(n.AccntState, n2.PkTxSignBytes, value)
		}
	}
}

func stepDeployScTx(nodes []*integrationTests.TestProcessorNode, senderIdx int, scCode string) {
	fmt.Println("Deploying SC...")
	txDeploy := createTxDeploy(nodes[senderIdx], scCode)
	nodes[senderIdx].SendTransaction(txDeploy)
	fmt.Println("Delaying for disseminating the deploy tx...")
	time.Sleep(stepDelay)

	fmt.Println(integrationTests.MakeDisplayTable(nodes))
}

func stepProposeBlock(nodes []*integrationTests.TestProcessorNode, idxProposer int, round uint32) {
	fmt.Println("Proposing block...")
	for idx, n := range nodes {
		if idx != idxProposer {
			continue
		}

		body, header := n.ProposeBlockOnlyWithSelf(round)
		n.BroadcastAndCommit(body, header)
	}

	fmt.Println("Delaying for disseminating headers and miniblocks...")
	time.Sleep(stepDelay)
	fmt.Println(integrationTests.MakeDisplayTable(nodes))
}

func stepSyncBlock(t *testing.T, nodes []*integrationTests.TestProcessorNode, idxProposer int, round uint32) {
	fmt.Println("All other shard nodes sync the proposed block...")
	for idx, n := range nodes {
		if idx == idxProposer {
			continue
		}

		err := n.SyncNode(uint64(round))
		if err != nil {
			assert.Fail(t, err.Error())
			return
		}
	}

	time.Sleep(stepDelay)
	fmt.Println(integrationTests.MakeDisplayTable(nodes))
}

func stepNodeDoesJoinGame(
	nodes []*integrationTests.TestProcessorNode,
	idxNode int,
	joinGameVal *big.Int,
	scAddress []byte) {

	fmt.Println("Calling SC.joinGame...")
	txDeploy := createTxJoinGame(nodes[idxNode], joinGameVal, scAddress)
	nodes[idxNode].SendTransaction(txDeploy)
	fmt.Println("Delaying for disseminating SC call tx...")
	time.Sleep(stepDelay)

	fmt.Println(integrationTests.MakeDisplayTable(nodes))
}

func stepNodeCallsRewardAndSend(
	nodes []*integrationTests.TestProcessorNode,
	idxNodeOwner int,
	idxNodeUser int,
	prize *big.Int,
	scAddress []byte) {

	fmt.Println("Calling SC.rewardAndSendToWallet...")
	txDeploy := createTxRewardAndSendToWallet(nodes[idxNodeOwner], nodes[idxNodeUser], prize, scAddress)
	nodes[idxNodeOwner].SendTransaction(txDeploy)
	fmt.Println("Delaying for disseminating SC call tx...")
	time.Sleep(time.Second * 1)

	fmt.Println(integrationTests.MakeDisplayTable(nodes))
}

func createTxDeploy(tn *integrationTests.TestProcessorNode, scCode string) *transaction.Transaction {
	tx := &transaction.Transaction{
		Nonce:    0,
		Value:    big.NewInt(0),
		RcvAddr:  make([]byte, 32),
		SndAddr:  tn.PkTxSignBytes,
		Data:     scCode,
		GasPrice: 0,
		GasLimit: 100000,
	}
	txBuff, _ := integrationTests.TestMarshalizer.Marshal(tx)
	tx.Signature, _ = tn.SingleSigner.Sign(tn.SkTxSign, txBuff)

	return tx
}

func createTxJoinGame(tn *integrationTests.TestProcessorNode, joinGameVal *big.Int, scAddress []byte) *transaction.Transaction {
	tx := &transaction.Transaction{
		Nonce:    0,
		Value:    joinGameVal,
		RcvAddr:  scAddress,
		SndAddr:  tn.PkTxSignBytes,
		Data:     fmt.Sprintf("joinGame@aaaa"),
		GasPrice: 0,
		GasLimit: 100000,
	}
	txBuff, _ := integrationTests.TestMarshalizer.Marshal(tx)
	tx.Signature, _ = tn.SingleSigner.Sign(tn.SkTxSign, txBuff)

	fmt.Printf("Join %s\n", hex.EncodeToString(tn.PkTxSignBytes))

	return tx
}

func createTxRewardAndSendToWallet(tnOwner *integrationTests.TestProcessorNode, tnUser *integrationTests.TestProcessorNode, prizeVal *big.Int, scAddress []byte) *transaction.Transaction {
	tx := &transaction.Transaction{
		Nonce:    0,
		Value:    big.NewInt(0),
		RcvAddr:  scAddress,
		SndAddr:  tnOwner.PkTxSignBytes,
		Data:     fmt.Sprintf("rewardAndSendToWallet@aaaa@%s@%X", hex.EncodeToString(tnUser.PkTxSignBytes), prizeVal),
		GasPrice: 0,
		GasLimit: 100000,
	}
	txBuff, _ := integrationTests.TestMarshalizer.Marshal(tx)
	tx.Signature, _ = tnOwner.SingleSigner.Sign(tnOwner.SkTxSign, txBuff)

	fmt.Printf("Reward %s\n", hex.EncodeToString(tnUser.PkTxSignBytes))

	return tx
}

func stepCheckJoinGameIsDoneCorrectly(
	t *testing.T,
	nodes []*integrationTests.TestProcessorNode,
	idxNodeScExists int,
	idxNodeCallerExists int,
	initialVal *big.Int,
	topUpVal *big.Int,
	scAddressBytes []byte,
) {

	nodeWithSc := nodes[idxNodeScExists]
	nodeWithCaller := nodes[idxNodeCallerExists]

	fmt.Println("Checking SC account received topUp val...")
	accnt, _ := nodeWithSc.AccntState.GetExistingAccount(integrationTests.CreateAddresFromAddrBytes(scAddressBytes))
	assert.NotNil(t, accnt)
	assert.Equal(t, topUpVal, accnt.(*state.Account).Balance)

	fmt.Println("Checking sender has initial-topUp val...")
	expectedVal := big.NewInt(0).Set(initialVal)
	expectedVal.Sub(expectedVal, topUpVal)
	fmt.Printf("Checking %s\n", hex.EncodeToString(nodeWithCaller.PkTxSignBytes))
	accnt, _ = nodeWithCaller.AccntState.GetExistingAccount(integrationTests.CreateAddresFromAddrBytes(nodeWithCaller.PkTxSignBytes))
	assert.NotNil(t, accnt)
	assert.Equal(t, expectedVal, accnt.(*state.Account).Balance)
}

func stepCheckRewardIsDoneCorrectly(
	t *testing.T,
	nodes []*integrationTests.TestProcessorNode,
	idxNodeScExists int,
	idxNodeCallerExists int,
	initialVal *big.Int,
	topUpVal *big.Int,
	withdraw *big.Int,
	scAddressBytes []byte,
) {

	nodeWithSc := nodes[idxNodeScExists]
	nodeWithCaller := nodes[idxNodeCallerExists]

	fmt.Println("Checking SC account has topUp-withdraw val...")
	accnt, _ := nodeWithSc.AccntState.GetExistingAccount(integrationTests.CreateAddresFromAddrBytes(scAddressBytes))
	assert.NotNil(t, accnt)
	expectedSC := big.NewInt(0).Set(topUpVal)
	expectedSC.Sub(expectedSC, withdraw)
	assert.Equal(t, expectedSC, accnt.(*state.Account).Balance)

	fmt.Println("Checking sender has initial-topUp+withdraw val...")
	expectedSender := big.NewInt(0).Set(initialVal)
	expectedSender.Sub(expectedSender, topUpVal)
	expectedSender.Add(expectedSender, withdraw)
	fmt.Printf("Checking %s\n", hex.EncodeToString(nodeWithCaller.PkTxSignBytes))
	accnt, _ = nodeWithCaller.AccntState.GetExistingAccount(integrationTests.CreateAddresFromAddrBytes(nodeWithCaller.PkTxSignBytes))
	assert.NotNil(t, accnt)
	assert.Equal(t, expectedSender, accnt.(*state.Account).Balance)
}

func stepCheckRootHashes(t *testing.T, nodes []*integrationTests.TestProcessorNode, idxProposers []int) {
	for _, idx := range idxProposers {
		checkRootHashInShard(t, nodes, idx)
	}
}

func checkRootHashInShard(t *testing.T, nodes []*integrationTests.TestProcessorNode, idxProposer int) {
	proposerNode := nodes[idxProposer]
	proposerRootHash, _ := proposerNode.AccntState.RootHash()

	for i := 0; i < len(nodes); i++ {
		node := nodes[i]

		if node.ShardCoordinator.SelfId() != proposerNode.ShardCoordinator.SelfId() {
			continue
		}

		fmt.Printf("Testing roothash for node index %d, shard ID %d...\n", i, node.ShardCoordinator.SelfId())
		nodeRootHash, _ := node.AccntState.RootHash()
		assert.Equal(t, proposerRootHash, nodeRootHash)
	}
}
