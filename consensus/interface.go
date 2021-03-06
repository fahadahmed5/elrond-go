package consensus

import (
	"time"

	"github.com/ElrondNetwork/elrond-go/data"
	"github.com/ElrondNetwork/elrond-go/p2p"
)

// BlsConsensusType specifies the signature scheme used in the consensus
const BlsConsensusType = "bls"

// Rounder defines the actions which should be handled by a round implementation
type Rounder interface {
	Index() int64
	// UpdateRound updates the index and the time stamp of the round depending of the genesis time and the current time given
	UpdateRound(time.Time, time.Time)
	TimeStamp() time.Time
	TimeDuration() time.Duration
	RemainingTime(startTime time.Time, maxTime time.Duration) time.Duration
	IsInterfaceNil() bool
}

// SubroundHandler defines the actions which should be handled by a subround implementation
type SubroundHandler interface {
	// DoWork implements of the subround's job
	DoWork(rounder Rounder) bool
	// Previous returns the ID of the previous subround
	Previous() int
	// Next returns the ID of the next subround
	Next() int
	// Current returns the ID of the current subround
	Current() int
	// StartTime returns the start time, in the rounder time, of the current subround
	StartTime() int64
	// EndTime returns the top limit time, in the rounder time, of the current subround
	EndTime() int64
	// Name returns the name of the current rounder
	Name() string
	// ConsensusChannel returns the consensus channel
	ConsensusChannel() chan bool
	// IsInterfaceNil returns true if there is no value under the interface
	IsInterfaceNil() bool
}

// ChronologyHandler defines the actions which should be handled by a chronology implementation
type ChronologyHandler interface {
	Close() error
	AddSubround(SubroundHandler)
	RemoveAllSubrounds()
	// StartRounds starts rounds in a sequential manner, one after the other
	StartRounds()
	IsInterfaceNil() bool
}

// BroadcastMessenger defines the behaviour of the broadcast messages by the consensus group
type BroadcastMessenger interface {
	BroadcastBlock(data.BodyHandler, data.HeaderHandler) error
	BroadcastHeader(data.HeaderHandler) error
	BroadcastMiniBlocks(map[uint32][]byte) error
	BroadcastTransactions(map[string][][]byte) error
	BroadcastConsensusMessage(*Message) error
	SetDataForDelayBroadcast(headerHash []byte, miniBlocks map[uint32][]byte, transactions map[string][][]byte) error
	IsInterfaceNil() bool
}

// P2PMessenger defines a subset of the p2p.Messenger interface
type P2PMessenger interface {
	Broadcast(topic string, buff []byte)
	IsInterfaceNil() bool
}

// NetworkShardingCollector defines the updating methods used by the network sharding component
// The interface assures that the collected data will be used by the p2p network sharding components
type NetworkShardingCollector interface {
	UpdatePeerIdPublicKey(pid p2p.PeerID, pk []byte)
	UpdatePublicKeyShardId(pk []byte, shardId uint32)
	UpdatePeerIdShardId(pid p2p.PeerID, shardId uint32)
	IsInterfaceNil() bool
}

// P2PAntifloodHandler defines the behavior of a component able to signal that the system is too busy (or flooded) processing
// p2p messages
type P2PAntifloodHandler interface {
	CanProcessMessage(message p2p.MessageP2P, fromConnectedPeer p2p.PeerID) error
	CanProcessMessagesOnTopic(peer p2p.PeerID, topic string, numMessages uint32) error
	ResetForTopic(topic string)
	SetMaxMessagesForTopic(topic string, maxNum uint32)
	IsInterfaceNil() bool
}

// HeadersPoolSubscriber can subscribe for notifications when a new block header is added to the headers pool
type HeadersPoolSubscriber interface {
	RegisterHandler(handler func(headerHandler data.HeaderHandler, headerHash []byte))
	IsInterfaceNil() bool
}
