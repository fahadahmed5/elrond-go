package topicResolverSender_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/dataRetriever/mock"
	"github.com/ElrondNetwork/elrond-go/dataRetriever/resolvers/topicResolverSender"
	"github.com/ElrondNetwork/elrond-go/p2p"
	"github.com/stretchr/testify/assert"
)

var defaultHashes = [][]byte{[]byte("hash")}

func createMockArgTopicResolverSender() topicResolverSender.ArgTopicResolverSender {
	return topicResolverSender.ArgTopicResolverSender{
		Messenger:          &mock.MessageHandlerStub{},
		TopicName:          "topic",
		PeerListCreator:    &mock.PeerListCreatorStub{},
		Marshalizer:        &mock.MarshalizerMock{},
		Randomizer:         &mock.IntRandomizerStub{},
		TargetShardId:      0,
		OutputAntiflooder:  &mock.P2PAntifloodHandlerStub{},
		NumIntraShardPeers: 2,
		NumCrossShardPeers: 2,
	}
}

//------- NewTopicResolverSender

func TestNewTopicResolverSender_NilMessengerShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.Messenger = nil
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.True(t, check.IfNil(trs))
	assert.Equal(t, dataRetriever.ErrNilMessenger, err)
}

func TestNewTopicResolverSender_NilPeersListCreatorShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.PeerListCreator = nil
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.True(t, check.IfNil(trs))
	assert.Equal(t, dataRetriever.ErrNilPeerListCreator, err)
}

func TestNewTopicResolverSender_NilMarshalizerShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.Marshalizer = nil
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.True(t, check.IfNil(trs))
	assert.Equal(t, dataRetriever.ErrNilMarshalizer, err)
}

func TestNewTopicResolverSender_NilRandomizerShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.Randomizer = nil
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.True(t, check.IfNil(trs))
	assert.Equal(t, dataRetriever.ErrNilRandomizer, err)
}

func TestNewTopicResolverSender_NilOutputAntiflooderShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.OutputAntiflooder = nil
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.True(t, check.IfNil(trs))
	assert.Equal(t, dataRetriever.ErrNilAntifloodHandler, err)
}

func TestNewTopicResolverSender_InvalidNumIntraShardPeersShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.NumIntraShardPeers = -1
	arg.NumCrossShardPeers = 100
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.True(t, check.IfNil(trs))
	assert.True(t, errors.Is(err, dataRetriever.ErrInvalidValue))
}

func TestNewTopicResolverSender_InvalidNumCrossShardPeersShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.NumCrossShardPeers = -1
	arg.NumIntraShardPeers = 100
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.True(t, check.IfNil(trs))
	assert.True(t, errors.Is(err, dataRetriever.ErrInvalidValue))
}

func TestNewTopicResolverSender_InvalidNumberOfPeersShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.NumIntraShardPeers = 1
	arg.NumCrossShardPeers = 0
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.True(t, check.IfNil(trs))
	assert.True(t, errors.Is(err, dataRetriever.ErrInvalidValue))
}

func TestNewTopicResolverSender_OkValsShouldWork(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.False(t, check.IfNil(trs))
	assert.Nil(t, err)
	assert.Equal(t, uint32(0), trs.TargetShardID())
}

func TestNewTopicResolverSender_OkValsWithNumZeroShouldWork(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.NumIntraShardPeers = 0
	trs, err := topicResolverSender.NewTopicResolverSender(arg)

	assert.False(t, check.IfNil(trs))
	assert.Nil(t, err)
	assert.Equal(t, uint32(0), trs.TargetShardID())
}

//------- SendOnRequestTopic

func TestTopicResolverSender_SendOnRequestTopicMarshalizerFailsShouldErr(t *testing.T) {
	t.Parallel()

	errExpected := errors.New("expected error")

	arg := createMockArgTopicResolverSender()
	arg.Marshalizer = &mock.MarshalizerStub{
		MarshalCalled: func(obj interface{}) (bytes []byte, e error) {
			return nil, errExpected
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.SendOnRequestTopic(&dataRetriever.RequestData{}, defaultHashes)

	assert.Equal(t, errExpected, err)
}

func TestTopicResolverSender_SendOnRequestTopicNoOneToSendShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	arg.PeerListCreator = &mock.PeerListCreatorStub{
		PeerListCalled: func() []p2p.PeerID {
			return make([]p2p.PeerID, 0)
		},
		IntraShardPeerListCalled: func() []p2p.PeerID {
			return make([]p2p.PeerID, 0)
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.SendOnRequestTopic(&dataRetriever.RequestData{}, defaultHashes)

	assert.True(t, errors.Is(err, dataRetriever.ErrSendRequest))
}

func TestTopicResolverSender_SendOnRequestTopicShouldWork(t *testing.T) {
	t.Parallel()

	pID1 := p2p.PeerID("peer1")
	pID2 := p2p.PeerID("peer2")
	sentToPid1 := false
	sentToPid2 := false

	arg := createMockArgTopicResolverSender()
	arg.Messenger = &mock.MessageHandlerStub{
		SendToConnectedPeerCalled: func(topic string, buff []byte, peerID p2p.PeerID) error {
			if bytes.Equal(peerID.Bytes(), pID1.Bytes()) {
				sentToPid1 = true
			}
			if bytes.Equal(peerID.Bytes(), pID2.Bytes()) {
				sentToPid2 = true
			}

			return nil
		},
	}
	arg.PeerListCreator = &mock.PeerListCreatorStub{
		PeerListCalled: func() []p2p.PeerID {
			return []p2p.PeerID{pID1}
		},
		IntraShardPeerListCalled: func() []p2p.PeerID {
			return []p2p.PeerID{pID2}
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.SendOnRequestTopic(&dataRetriever.RequestData{}, defaultHashes)

	assert.Nil(t, err)
	assert.True(t, sentToPid1)
	assert.True(t, sentToPid2)
}

func TestTopicResolverSender_SendOnRequestShouldStopAfterSendingToRequiredNum(t *testing.T) {
	t.Parallel()

	pIDs := []p2p.PeerID{"pid1", "pid2", "pid3", "pid4", "pid5"}

	numSent := 0
	arg := createMockArgTopicResolverSender()
	arg.Messenger = &mock.MessageHandlerStub{
		SendToConnectedPeerCalled: func(topic string, buff []byte, peerID p2p.PeerID) error {
			numSent++

			return nil
		},
	}
	arg.PeerListCreator = &mock.PeerListCreatorStub{
		PeerListCalled: func() []p2p.PeerID {
			return pIDs
		},
		IntraShardPeerListCalled: func() []p2p.PeerID {
			return pIDs
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.SendOnRequestTopic(&dataRetriever.RequestData{}, defaultHashes)

	assert.Nil(t, err)
	assert.Equal(t, arg.NumCrossShardPeers+arg.NumIntraShardPeers, numSent)
}

func TestTopicResolverSender_SendOnRequestNoIntraShardShouldNotCallIntraShard(t *testing.T) {
	t.Parallel()

	pIDs := []p2p.PeerID{"pid1", "pid2", "pid3", "pid4", "pid5"}
	pIDNotCalled := p2p.PeerID("pid not called")

	numSent := 0
	arg := createMockArgTopicResolverSender()
	arg.Messenger = &mock.MessageHandlerStub{
		SendToConnectedPeerCalled: func(topic string, buff []byte, peerID p2p.PeerID) error {
			if peerID == pIDNotCalled {
				assert.Fail(t, fmt.Sprintf("should not have called pid %s", peerID))
			}
			numSent++

			return nil
		},
	}
	arg.NumIntraShardPeers = 0
	arg.PeerListCreator = &mock.PeerListCreatorStub{
		PeerListCalled: func() []p2p.PeerID {
			return pIDs
		},
		IntraShardPeerListCalled: func() []p2p.PeerID {
			return []p2p.PeerID{pIDNotCalled}
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.SendOnRequestTopic(&dataRetriever.RequestData{}, defaultHashes)

	assert.Nil(t, err)
	assert.Equal(t, arg.NumCrossShardPeers, numSent)
}

func TestTopicResolverSender_SendOnRequestNoCrossShardShouldNotCallCrossShard(t *testing.T) {
	t.Parallel()

	pIDs := []p2p.PeerID{"pid1", "pid2", "pid3", "pid4", "pid5"}
	pIDNotCalled := p2p.PeerID("pid not called")

	numSent := 0
	arg := createMockArgTopicResolverSender()
	arg.Messenger = &mock.MessageHandlerStub{
		SendToConnectedPeerCalled: func(topic string, buff []byte, peerID p2p.PeerID) error {
			if peerID == pIDNotCalled {
				assert.Fail(t, fmt.Sprintf("should not have called pid %s", peerID))
			}
			numSent++

			return nil
		},
	}
	arg.NumCrossShardPeers = 0
	arg.PeerListCreator = &mock.PeerListCreatorStub{
		PeerListCalled: func() []p2p.PeerID {
			return []p2p.PeerID{pIDNotCalled}
		},
		IntraShardPeerListCalled: func() []p2p.PeerID {
			return pIDs
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.SendOnRequestTopic(&dataRetriever.RequestData{}, defaultHashes)

	assert.Nil(t, err)
	assert.Equal(t, arg.NumIntraShardPeers, numSent)
}

func TestTopicResolverSender_SendOnRequestTopicErrorsShouldReturnError(t *testing.T) {
	t.Parallel()

	pID1 := p2p.PeerID("peer1")
	sentToPid1 := false

	expectedErr := errors.New("expected error")
	arg := createMockArgTopicResolverSender()
	arg.Messenger = &mock.MessageHandlerStub{
		SendToConnectedPeerCalled: func(topic string, buff []byte, peerID p2p.PeerID) error {
			if bytes.Equal(peerID.Bytes(), pID1.Bytes()) {
				sentToPid1 = true
			}

			return expectedErr
		},
	}
	arg.PeerListCreator = &mock.PeerListCreatorStub{
		PeerListCalled: func() []p2p.PeerID {
			return []p2p.PeerID{pID1}
		},
		IntraShardPeerListCalled: func() []p2p.PeerID {
			return make([]p2p.PeerID, 0)
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.SendOnRequestTopic(&dataRetriever.RequestData{}, defaultHashes)

	assert.True(t, errors.Is(err, dataRetriever.ErrSendRequest))
	assert.True(t, sentToPid1)
}

//------- Send

func TestTopicResolverSender_SendOutputAntiflooderErrorsShouldNotSendButError(t *testing.T) {
	t.Parallel()

	pID1 := p2p.PeerID("peer1")
	buffToSend := []byte("buff")

	expectedErr := errors.New("can not send to peer")
	arg := createMockArgTopicResolverSender()
	arg.Messenger = &mock.MessageHandlerStub{
		SendToConnectedPeerCalled: func(topic string, buff []byte, peerID p2p.PeerID) error {
			assert.Fail(t, "send shouldn't have been called")

			return nil
		},
	}
	arg.OutputAntiflooder = &mock.P2PAntifloodHandlerStub{
		CanProcessMessageCalled: func(message p2p.MessageP2P, fromConnectedPeer p2p.PeerID) error {
			if fromConnectedPeer == pID1 {
				return expectedErr
			}

			assert.Fail(t, "wrong peer provided, should have been called with the destination peer")
			return nil
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.Send(buffToSend, pID1)

	assert.True(t, errors.Is(err, expectedErr))
}

func TestTopicResolverSender_SendShouldWork(t *testing.T) {
	t.Parallel()

	pID1 := p2p.PeerID("peer1")
	sentToPid1 := false
	buffToSend := []byte("buff")

	arg := createMockArgTopicResolverSender()
	arg.Messenger = &mock.MessageHandlerStub{
		SendToConnectedPeerCalled: func(topic string, buff []byte, peerID p2p.PeerID) error {
			if bytes.Equal(peerID.Bytes(), pID1.Bytes()) &&
				bytes.Equal(buff, buffToSend) {
				sentToPid1 = true
			}

			return nil
		},
	}
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.Send(buffToSend, pID1)

	assert.Nil(t, err)
	assert.True(t, sentToPid1)
}

func TestTopicResolverSender_Topic(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	assert.Equal(t, arg.TopicName+topicResolverSender.TopicRequestSuffix, trs.RequestTopic())
}

func TestTopicResolverSender_ResolverDebugHandler(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	handler := &mock.ResolverDebugHandler{}

	err := trs.SetResolverDebugHandler(handler)
	assert.Nil(t, err)

	assert.True(t, handler == trs.ResolverDebugHandler()) //pointer testing
}

func TestTopicResolverSender_SetResolverDebugHandlerNilShouldErr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	err := trs.SetResolverDebugHandler(nil)
	assert.Equal(t, dataRetriever.ErrNilResolverDebugHandler, err)
}

func TestTopicResolverSender_NumPeersToQueryr(t *testing.T) {
	t.Parallel()

	arg := createMockArgTopicResolverSender()
	trs, _ := topicResolverSender.NewTopicResolverSender(arg)

	intra := 1123
	cross := 2143

	trs.SetNumPeersToQuery(intra, cross)
	recoveredIntra, recoveredCross := trs.NumPeersToQuery()

	assert.Equal(t, intra, recoveredIntra)
	assert.Equal(t, cross, recoveredCross)
}

// ------- FisherYatesShuffle

func TestFisherYatesShuffle_EmptyShouldReturnEmpty(t *testing.T) {
	indexes := make([]int, 0)
	randomizer := &mock.IntRandomizerStub{}

	resultIndexes := topicResolverSender.FisherYatesShuffle(indexes, randomizer)

	assert.Empty(t, resultIndexes)
}

func TestFisherYatesShuffle_OneElementShouldReturnTheSame(t *testing.T) {
	indexes := []int{1}
	randomizer := &mock.IntRandomizerStub{
		IntnCalled: func(n int) int {
			return n - 1
		},
	}

	resultIndexes := topicResolverSender.FisherYatesShuffle(indexes, randomizer)

	assert.Equal(t, indexes, resultIndexes)
}

func TestFisherYatesShuffle_ShouldWork(t *testing.T) {
	indexes := []int{1, 2, 3, 4, 5}
	randomizer := &mock.IntRandomizerStub{
		IntnCalled: func(n int) int {
			return 0
		},
	}

	//this will cause a rotation of the first element:
	//i = 4: 5, 2, 3, 4, 1 (swap 1 <-> 5)
	//i = 3: 4, 2, 3, 5, 1 (swap 5 <-> 4)
	//i = 2: 3, 2, 4, 5, 1 (swap 3 <-> 4)
	//i = 1: 2, 3, 4, 5, 1 (swap 3 <-> 2)

	resultIndexes := topicResolverSender.FisherYatesShuffle(indexes, randomizer)
	expectedResult := []int{2, 3, 4, 5, 1}

	assert.Equal(t, expectedResult, resultIndexes)
}
