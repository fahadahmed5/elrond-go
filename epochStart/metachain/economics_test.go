package metachain

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/data/block"
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/epochStart"
	"github.com/ElrondNetwork/elrond-go/epochStart/mock"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMockEpochEconomicsArguments() ArgsNewEpochEconomics {
	shardCoordinator := mock.NewMultiShardsCoordinatorMock(1)

	argsNewEpochEconomics := ArgsNewEpochEconomics{
		Hasher:              &mock.HasherMock{},
		Marshalizer:         &mock.MarshalizerMock{},
		Store:               createMetaStore(),
		ShardCoordinator:    shardCoordinator,
		NodesConfigProvider: &mock.NodesCoordinatorStub{},
		RewardsHandler:      &mock.RewardsHandlerStub{},
		RoundTime:           &mock.RoundTimeDurationHandler{},
	}
	return argsNewEpochEconomics
}

func TestEpochEconomics_NewEndOfEpochEconomicsDataCreatorNilMarshalizer(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochEconomicsArguments()
	arguments.Marshalizer = nil

	esd, err := NewEndOfEpochEconomicsDataCreator(arguments)
	require.Nil(t, esd)
	require.Equal(t, epochStart.ErrNilMarshalizer, err)
}

func TestEpochEconomics_NewEndOfEpochEconomicsDataCreatorNilStore(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochEconomicsArguments()
	arguments.Store = nil

	esd, err := NewEndOfEpochEconomicsDataCreator(arguments)
	require.Nil(t, esd)
	require.Equal(t, epochStart.ErrNilStorage, err)
}

func TestEpochEconomics_NewEndOfEpochEconomicsDataCreatorNilShardCoordinator(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochEconomicsArguments()
	arguments.ShardCoordinator = nil

	esd, err := NewEndOfEpochEconomicsDataCreator(arguments)
	require.Nil(t, esd)
	require.Equal(t, epochStart.ErrNilShardCoordinator, err)
}

func TestEpochEconomics_NewEndOfEpochEconomicsDataCreatorNilNodesdCoordinator(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochEconomicsArguments()
	arguments.NodesConfigProvider = nil

	esd, err := NewEndOfEpochEconomicsDataCreator(arguments)
	require.Nil(t, esd)
	require.Equal(t, epochStart.ErrNilNodesConfigProvider, err)
}

func TestEpochEconomics_NewEndOfEpochEconomicsDataCreatorNilRewardsHandler(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochEconomicsArguments()
	arguments.RewardsHandler = nil

	esd, err := NewEndOfEpochEconomicsDataCreator(arguments)
	require.Nil(t, esd)
	require.Equal(t, epochStart.ErrNilRewardsHandler, err)
}

func TestEpochEconomics_NewEndOfEpochEconomicsDataCreatorNilRounder(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochEconomicsArguments()
	arguments.RoundTime = nil

	esd, err := NewEndOfEpochEconomicsDataCreator(arguments)
	require.Nil(t, esd)
	require.Equal(t, epochStart.ErrNilRounder, err)
}

func TestEpochEconomics_NewEndOfEpochEconomicsDataCreatorShouldWork(t *testing.T) {
	t.Parallel()

	arguments := createMockEpochEconomicsArguments()

	esd, err := NewEndOfEpochEconomicsDataCreator(arguments)
	require.NotNil(t, esd)
	require.Nil(t, err)
}

func TestNewEndOfEpochEconomicsDataCreator_NilMarshalizer(t *testing.T) {
	t.Parallel()

	args := getArguments()
	args.Marshalizer = nil
	eoeedc, err := NewEndOfEpochEconomicsDataCreator(args)

	assert.True(t, check.IfNil(eoeedc))
	assert.Equal(t, epochStart.ErrNilMarshalizer, err)
}

func TestNewEndOfEpochEconomicsDataCreator_NilStore(t *testing.T) {
	t.Parallel()

	args := getArguments()
	args.Store = nil
	eoeedc, err := NewEndOfEpochEconomicsDataCreator(args)

	assert.True(t, check.IfNil(eoeedc))
	assert.Equal(t, epochStart.ErrNilStorage, err)
}

func TestNewEndOfEpochEconomicsDataCreator_NilShardCoordinator(t *testing.T) {
	t.Parallel()

	args := getArguments()
	args.ShardCoordinator = nil
	eoeedc, err := NewEndOfEpochEconomicsDataCreator(args)

	assert.True(t, check.IfNil(eoeedc))
	assert.Equal(t, epochStart.ErrNilShardCoordinator, err)
}

func TestNewEndOfEpochEconomicsDataCreator_NilNodesCoordinator(t *testing.T) {
	t.Parallel()

	args := getArguments()
	args.NodesConfigProvider = nil
	eoeedc, err := NewEndOfEpochEconomicsDataCreator(args)

	assert.True(t, check.IfNil(eoeedc))
	assert.Equal(t, epochStart.ErrNilNodesConfigProvider, err)
}

func TestNewEndOfEpochEconomicsDataCreator_NilRewardsHandler(t *testing.T) {
	t.Parallel()

	args := getArguments()
	args.RewardsHandler = nil
	eoeedc, err := NewEndOfEpochEconomicsDataCreator(args)

	assert.True(t, check.IfNil(eoeedc))
	assert.Equal(t, epochStart.ErrNilRewardsHandler, err)
}

func TestNewEndOfEpochEconomicsDataCreator_NilRoundTimeDurationHandler(t *testing.T) {
	t.Parallel()

	args := getArguments()
	args.RoundTime = nil
	eoeedc, err := NewEndOfEpochEconomicsDataCreator(args)

	assert.True(t, check.IfNil(eoeedc))
	assert.Equal(t, epochStart.ErrNilRounder, err)
}

func TestNewEndOfEpochEconomicsDataCreator_ShouldWork(t *testing.T) {
	t.Parallel()

	args := getArguments()
	eoeedc, err := NewEndOfEpochEconomicsDataCreator(args)

	assert.False(t, check.IfNil(eoeedc))
	assert.Nil(t, err)
}

func TestEconomics_ComputeEndOfEpochEconomics_NilMetaBlockShouldErr(t *testing.T) {
	t.Parallel()

	args := getArguments()
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	res, err := ec.ComputeEndOfEpochEconomics(nil)
	assert.Nil(t, res)
	assert.Equal(t, epochStart.ErrNilHeaderHandler, err)
}

func TestEconomics_ComputeEndOfEpochEconomics_NilAccumulatedFeesInEpochShouldErr(t *testing.T) {
	t.Parallel()

	args := getArguments()
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	mb := block.MetaBlock{AccumulatedFeesInEpoch: nil, DevFeesInEpoch: big.NewInt(0)}
	res, err := ec.ComputeEndOfEpochEconomics(&mb)
	assert.Nil(t, res)
	assert.Equal(t, epochStart.ErrNilTotalAccumulatedFeesInEpoch, err)
}

func TestEconomics_ComputeEndOfEpochEconomics_NilDevFeesInEpochShouldErr(t *testing.T) {
	t.Parallel()

	args := getArguments()
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	mb := block.MetaBlock{AccumulatedFeesInEpoch: big.NewInt(0), DevFeesInEpoch: nil}
	res, err := ec.ComputeEndOfEpochEconomics(&mb)
	assert.Nil(t, res)
	assert.Equal(t, epochStart.ErrNilTotalDevFeesInEpoch, err)
}

func TestEconomics_ComputeRewardsForCommunity(t *testing.T) {
	totalRewards := big.NewInt(0).SetUint64(123456)
	args := getArguments()
	args.RewardsHandler = &mock.RewardsHandlerStub{
		CommunityPercentageCalled: func() float64 {
			return 0.1
		},
	}
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	expectedRewards := big.NewInt(12345)
	communityRewards := ec.computeRewardsForCommunity(totalRewards)
	assert.Equal(t, expectedRewards, communityRewards)
}

func TestEconomics_AdjustRewardsPerBlockWithCommunityRewards(t *testing.T) {
	args := getArguments()
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	rwdPerBlock := big.NewInt(0).SetUint64(1000)
	blocksInEpoch := uint64(100)
	communityRewards := big.NewInt(0).SetUint64(10000)

	expectedRewardsCommunityAfterAdjustment := big.NewInt(0).Set(communityRewards)
	expectedRwdPerBlock := big.NewInt(900)

	ec.adjustRewardsPerBlockWithCommunityRewards(rwdPerBlock, communityRewards, blocksInEpoch)

	assert.Equal(t, expectedRewardsCommunityAfterAdjustment, communityRewards)
	assert.Equal(t, expectedRwdPerBlock, rwdPerBlock)
}

func TestEconomics_AdjustRewardsPerBlockWithLeaderPercentage(t *testing.T) {
	args := getArguments()
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	rwdPerBlock := big.NewInt(0).SetUint64(1000)
	blocksInEpoch := uint64(100)
	accumulatedFeesInEpoch := big.NewInt(0).SetUint64(10000)

	expectedRwdPerBlock := big.NewInt(900)
	ec.adjustRewardsPerBlockWithLeaderPercentage(rwdPerBlock, accumulatedFeesInEpoch, blocksInEpoch)
	assert.Equal(t, expectedRwdPerBlock, rwdPerBlock)
}

func TestEconomics_AdjustRewardsPerBlockWithDeveloperFees(t *testing.T) {
	args := getArguments()
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	rwdPerBlock := big.NewInt(0).SetUint64(1000)
	blocksInEpoch := uint64(100)
	developerFees := big.NewInt(0).SetUint64(500)

	expectedRwdPerBlock := big.NewInt(995)
	ec.adjustRewardsPerBlockWithDeveloperFees(rwdPerBlock, developerFees, blocksInEpoch)
	assert.Equal(t, expectedRwdPerBlock, rwdPerBlock)
}

func TestEconomics_ComputeEndOfEpochEconomics_NotEpochStartShouldErr(t *testing.T) {
	t.Parallel()

	args := getArguments()
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	mb1 := block.MetaBlock{
		Epoch:                  0,
		AccumulatedFeesInEpoch: big.NewInt(10),
		DevFeesInEpoch:         big.NewInt(0),
	}
	res, err := ec.ComputeEndOfEpochEconomics(&mb1)
	assert.Nil(t, res)
	assert.Equal(t, epochStart.ErrNotEpochStartBlock, err)
}

func TestEconomics_ComputeEndOfEpochEconomics(t *testing.T) {
	t.Parallel()

	args := getArguments()
	args.Store = &mock.ChainStorerStub{
		GetStorerCalled: func(unitType dataRetriever.UnitType) storage.Storer {
			return &mock.StorerStub{GetCalled: func(key []byte) ([]byte, error) {
				hdr := block.MetaBlock{
					Round: 10,
					Nonce: 5,
					EpochStart: block.EpochStart{
						Economics: block.Economics{
							TotalSupply:            big.NewInt(100000),
							TotalToDistribute:      big.NewInt(10),
							TotalNewlyMinted:       big.NewInt(109),
							RewardsPerBlockPerNode: big.NewInt(10),
							NodePrice:              big.NewInt(10),
						},
					},
				}
				hdrBytes, _ := json.Marshal(hdr)
				return hdrBytes, nil
			}}
		},
	}
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	mb := block.MetaBlock{
		Round: 15000,
		EpochStart: block.EpochStart{
			LastFinalizedHeaders: []block.EpochStartShardData{
				{ShardID: 0, Round: 2, Nonce: 3},
				{ShardID: 1, Round: 2, Nonce: 3},
			},
			Economics: block.Economics{},
		},
		Epoch:                  2,
		AccumulatedFeesInEpoch: big.NewInt(10000),
		DevFeesInEpoch:         big.NewInt(0),
	}
	res, err := ec.ComputeEndOfEpochEconomics(&mb)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestEconomics_VerifyRewardsPerBlock_DifferentHitRates(t *testing.T) {
	t.Parallel()

	commAddress := "communityAddress"
	totalSupply := big.NewInt(20000000000) // 20B
	accFeesInEpoch := big.NewInt(0)
	devFeesInEpoch := big.NewInt(0)
	roundDur := 4
	args := getArguments()
	args.RewardsHandler = &mock.RewardsHandlerStub{
		MaxInflationRateCalled: func() float64 {
			return 0.1
		},
		CommunityAddressCalled: func() string {
			return commAddress
		},
		CommunityPercentageCalled: func() float64 {
			return 0.1
		},
	}
	args.RoundTime = &mock.RoundTimeDurationHandler{
		TimeDurationCalled: func() time.Duration {
			return time.Duration(roundDur) * time.Second
		},
	}
	hdrPrevEpochStart := block.MetaBlock{
		Round: 0,
		Nonce: 0,
		Epoch: 0,
		EpochStart: block.EpochStart{
			Economics: block.Economics{
				TotalSupply:            totalSupply,
				TotalToDistribute:      big.NewInt(10),
				TotalNewlyMinted:       big.NewInt(10),
				RewardsPerBlockPerNode: big.NewInt(10),
				NodePrice:              big.NewInt(10),
				RewardsForCommunity:    big.NewInt(10),
			},
			LastFinalizedHeaders: []block.EpochStartShardData{
				{ShardID: 0, Nonce: 0},
				{ShardID: 1, Nonce: 0},
			},
		},
	}
	args.Store = &mock.ChainStorerStub{
		GetStorerCalled: func(unitType dataRetriever.UnitType) storage.Storer {
			// this will be the previous epoch meta block. It has initial 0 values so it can be considered at genesis
			return &mock.StorerStub{GetCalled: func(key []byte) ([]byte, error) {
				hdrBytes, _ := json.Marshal(hdrPrevEpochStart)
				return hdrBytes, nil
			}}
		},
	}
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	expRwdPerBlock := 84 // based on 0.1 inflation

	numBlocksInEpochSlice := []int{
		numberOfSecondsInDay / roundDur,       // 100 % hit rate
		(numberOfSecondsInDay / roundDur) / 2, // 50 % hit rate
		(numberOfSecondsInDay / roundDur) / 4, // 25 % hit rate
		(numberOfSecondsInDay / roundDur) / 8, // 12.5 % hit rate
		1,                                     // only the metablock was committed in that epoch
		37,                                    // random
		63,                                    // random
	}

	hdrPrevEpochStartHash, _ := core.CalculateHash(&mock.MarshalizerMock{}, &mock.HasherMock{}, hdrPrevEpochStart)
	for _, numBlocksInEpoch := range numBlocksInEpochSlice {
		expectedTotalToDistribute := big.NewInt(int64(expRwdPerBlock * numBlocksInEpoch * 3)) // 2 shards + meta
		expectedTotalNewlyMinted := big.NewInt(0).Sub(expectedTotalToDistribute, accFeesInEpoch)
		expectedTotalSupply := big.NewInt(0).Add(totalSupply, expectedTotalNewlyMinted)
		expectedCommunityRewards := big.NewInt(0).Div(expectedTotalToDistribute, big.NewInt(10))
		commRewardPerBlock := big.NewInt(0).Div(expectedCommunityRewards, big.NewInt(int64(numBlocksInEpoch*3)))
		adjustedRwdPerBlock := big.NewInt(0).Sub(big.NewInt(int64(expRwdPerBlock)), commRewardPerBlock)

		mb := block.MetaBlock{
			Round: uint64(numBlocksInEpoch),
			Nonce: uint64(numBlocksInEpoch),
			EpochStart: block.EpochStart{
				LastFinalizedHeaders: []block.EpochStartShardData{
					{ShardID: 0, Round: uint64(numBlocksInEpoch), Nonce: uint64(numBlocksInEpoch)},
					{ShardID: 1, Round: uint64(numBlocksInEpoch), Nonce: uint64(numBlocksInEpoch)},
				},
				Economics: block.Economics{
					TotalSupply:            expectedTotalSupply,
					TotalToDistribute:      expectedTotalToDistribute,
					TotalNewlyMinted:       expectedTotalNewlyMinted,
					RewardsPerBlockPerNode: adjustedRwdPerBlock,
					NodePrice:              big.NewInt(10),
					PrevEpochStartHash:     hdrPrevEpochStartHash,
					RewardsForCommunity:    expectedCommunityRewards,
				},
			},
			Epoch:                  1,
			AccumulatedFeesInEpoch: accFeesInEpoch,
			DevFeesInEpoch:         devFeesInEpoch,
		}

		err := ec.VerifyRewardsPerBlock(&mb)
		assert.Nil(t, err)
	}
}

type testInput struct {
	blockPerEpochOneShard  uint64
	accumulatedFeesInEpoch *big.Int
	devFeesInEpoch         *big.Int
}

func TestComputeEndOfEpochEconomics(t *testing.T) {
	t.Parallel()

	totalSupply, _ := big.NewInt(0).SetString("20000000000000000000000000000", 10) // 20 Billions ERD
	nodePrice, _ := big.NewInt(0).SetString("1000000000000000000000", 10)          // 1000 ERD
	roundDuration := 4

	args := createArgsForComputeEndOfEpochEconomics(roundDuration, totalSupply, nodePrice)
	ec, _ := NewEndOfEpochEconomicsDataCreator(args)

	epochDuration := numberOfSecondsInDay
	roundsPerEpoch := uint64(epochDuration / roundDuration)

	testInputs := []testInput{
		{blockPerEpochOneShard: roundsPerEpoch, accumulatedFeesInEpoch: intToErd(1000), devFeesInEpoch: intToErd(1000)},
		{blockPerEpochOneShard: roundsPerEpoch / 2, accumulatedFeesInEpoch: intToErd(1000), devFeesInEpoch: intToErd(1000)},
		{blockPerEpochOneShard: roundsPerEpoch / 4, accumulatedFeesInEpoch: intToErd(1000), devFeesInEpoch: intToErd(1000)},
		{blockPerEpochOneShard: roundsPerEpoch / 8, accumulatedFeesInEpoch: intToErd(1000), devFeesInEpoch: intToErd(1000)},
		{blockPerEpochOneShard: roundsPerEpoch / 16, accumulatedFeesInEpoch: intToErd(1000), devFeesInEpoch: intToErd(1000)},
		{blockPerEpochOneShard: roundsPerEpoch / 32, accumulatedFeesInEpoch: intToErd(1000), devFeesInEpoch: intToErd(1000)},
		{blockPerEpochOneShard: roundsPerEpoch / 64, accumulatedFeesInEpoch: intToErd(100000000000000), devFeesInEpoch: intToErd(10000000)},
		{blockPerEpochOneShard: roundsPerEpoch, accumulatedFeesInEpoch: intToErd(100000000000000), devFeesInEpoch: intToErd(30000000000000)},
	}

	rewardsPerBlock, _ := big.NewInt(0).SetString("84559445290038908043", 10) // *based on 0.1 inflation
	for _, input := range testInputs {
		meta := &block.MetaBlock{
			AccumulatedFeesInEpoch: input.accumulatedFeesInEpoch,
			DevFeesInEpoch:         input.devFeesInEpoch,
			Epoch:                  1,
			Round:                  roundsPerEpoch,
			Nonce:                  input.blockPerEpochOneShard,
			EpochStart: block.EpochStart{
				LastFinalizedHeaders: []block.EpochStartShardData{
					{ShardID: 0, Round: roundsPerEpoch, Nonce: input.blockPerEpochOneShard},
					{ShardID: 1, Round: roundsPerEpoch, Nonce: input.blockPerEpochOneShard},
				},
			},
		}

		economicsBlock, err := ec.ComputeEndOfEpochEconomics(meta)
		assert.Nil(t, err)

		verifyEconomicsBlock(t, economicsBlock, input, rewardsPerBlock, nodePrice, totalSupply, roundsPerEpoch, args.RewardsHandler)
	}

}

func createArgsForComputeEndOfEpochEconomics(
	roundDuration int,
	totalSupply *big.Int,
	nodePrice *big.Int,
) ArgsNewEpochEconomics {
	commAddress := "communityAddress"

	args := getArguments()
	args.RewardsHandler = &mock.RewardsHandlerStub{
		MaxInflationRateCalled: func() float64 {
			return 0.1
		},
		CommunityAddressCalled: func() string {
			return commAddress
		},
		CommunityPercentageCalled: func() float64 {
			return 0.1
		},
		LeaderPercentageCalled: func() float64 {
			return 0.1
		},
	}
	args.RoundTime = &mock.RoundTimeDurationHandler{
		TimeDurationCalled: func() time.Duration {
			return time.Duration(roundDuration) * time.Second
		},
	}
	hdrPrevEpochStart := block.MetaBlock{
		Round: 0,
		Nonce: 0,
		Epoch: 0,
		EpochStart: block.EpochStart{
			Economics: block.Economics{
				TotalSupply:            totalSupply,
				TotalToDistribute:      big.NewInt(10),
				TotalNewlyMinted:       big.NewInt(10),
				RewardsPerBlockPerNode: big.NewInt(10),
				NodePrice:              nodePrice,
				RewardsForCommunity:    big.NewInt(10),
			},
			LastFinalizedHeaders: []block.EpochStartShardData{
				{ShardID: 0, Nonce: 0},
				{ShardID: 1, Nonce: 0},
			},
		},
	}
	args.Store = &mock.ChainStorerStub{
		GetStorerCalled: func(unitType dataRetriever.UnitType) storage.Storer {
			// this will be the previous epoch meta block. It has initial 0 values so it can be considered at genesis
			return &mock.StorerStub{GetCalled: func(key []byte) ([]byte, error) {
				hdrBytes, _ := json.Marshal(hdrPrevEpochStart)
				return hdrBytes, nil
			}}
		},
	}

	return args
}

func verifyEconomicsBlock(
	t *testing.T,
	economicsBlock *block.Economics,
	input testInput,
	rewardsPerBlock *big.Int,
	nodePrice *big.Int,
	totalSupply *big.Int,
	roundsPerEpoch uint64,
	rewardsHandler process.RewardsHandler,
) {
	totalBlocksPerEpoch := int64(input.blockPerEpochOneShard * 3)
	hitRate := float64(input.blockPerEpochOneShard) / float64(roundsPerEpoch) * 100
	printEconomicsData(economicsBlock, hitRate, totalBlocksPerEpoch)

	expectedTotalRewardsToBeDistributed := big.NewInt(0).Mul(rewardsPerBlock, big.NewInt(totalBlocksPerEpoch))
	expectedNewTokens := big.NewInt(0).Sub(expectedTotalRewardsToBeDistributed, input.accumulatedFeesInEpoch)
	if expectedNewTokens.Cmp(big.NewInt(0)) < 0 {
		expectedNewTokens = big.NewInt(0)
		expectedTotalRewardsToBeDistributed = input.accumulatedFeesInEpoch
	}

	adjustedRewardsPerBlock := big.NewInt(0).Div(expectedTotalRewardsToBeDistributed, big.NewInt(totalBlocksPerEpoch))

	// subtract developer rewards per block
	developerFeesPerBlock := big.NewInt(0).Div(input.devFeesInEpoch, big.NewInt(totalBlocksPerEpoch))
	adjustedRewardsPerBlock.Sub(adjustedRewardsPerBlock, developerFeesPerBlock)
	// subtract leader percentage per block
	rewardsForLeader := core.GetPercentageOfValue(input.accumulatedFeesInEpoch, rewardsHandler.LeaderPercentage())
	rewardsForLeaderPerBlock := big.NewInt(0).Div(rewardsForLeader, big.NewInt(totalBlocksPerEpoch))
	adjustedRewardsPerBlock.Sub(adjustedRewardsPerBlock, rewardsForLeaderPerBlock)
	// communityPercentage
	expectedCommunityRewards := core.GetPercentageOfValue(expectedTotalRewardsToBeDistributed, rewardsHandler.CommunityPercentage())
	// subtract community percentage per block
	communityRewardsPerBlock := big.NewInt(0).Div(expectedCommunityRewards, big.NewInt(totalBlocksPerEpoch))
	adjustedRewardsPerBlock.Sub(adjustedRewardsPerBlock, communityRewardsPerBlock)

	assert.Equal(t, expectedNewTokens, economicsBlock.TotalNewlyMinted)
	assert.Equal(t, big.NewInt(0).Add(totalSupply, expectedNewTokens), economicsBlock.TotalSupply)
	assert.Equal(t, expectedTotalRewardsToBeDistributed, economicsBlock.TotalToDistribute)
	assert.Equal(t, expectedCommunityRewards, economicsBlock.RewardsForCommunity)
	assert.Equal(t, nodePrice, economicsBlock.NodePrice)
	assert.Equal(t, adjustedRewardsPerBlock, economicsBlock.RewardsPerBlockPerNode)
}

func printEconomicsData(eb *block.Economics, hitRate float64, numBlocksTotal int64) {
	fmt.Printf("Hit rate per shard %.4f%%, Total block produced: %d \n", hitRate, numBlocksTotal)
	fmt.Printf("Total supply: %vERD, TotalToDistribute %vERD, "+
		"TotalNewlyMinted %vERD, RewardsPerBlockPerNode %vERD, RewardsForCommunity %vERD, NodePrice: %vERD",
		denomination(eb.TotalSupply), denomination(eb.TotalToDistribute), denomination(eb.TotalNewlyMinted),
		denomination(eb.RewardsPerBlockPerNode), denomination(eb.RewardsForCommunity), denomination(eb.NodePrice))
	fmt.Println()
}

func intToErd(value int) *big.Int {
	denomination, _ := big.NewInt(0).SetString("1000000000000000000", 10)

	return big.NewInt(0).Mul(denomination, big.NewInt(int64(value)))
}

func denomination(value *big.Int) string {
	denomination, _ := big.NewInt(0).SetString("1000000000000000000", 10)
	cpValue := big.NewInt(0).Set(value)

	return cpValue.Div(cpValue, denomination).String()
}

func getArguments() ArgsNewEpochEconomics {
	return ArgsNewEpochEconomics{
		Marshalizer:         &mock.MarshalizerMock{},
		Hasher:              mock.HasherMock{},
		Store:               &mock.ChainStorerStub{},
		ShardCoordinator:    mock.NewMultipleShardsCoordinatorMock(),
		NodesConfigProvider: &mock.NodesCoordinatorStub{},
		RewardsHandler:      &mock.RewardsHandlerStub{},
		RoundTime:           &mock.RoundTimeDurationHandler{},
	}
}
