package node

import (
	"io"

	"github.com/ElrondNetwork/elrond-go/p2p"
)

// P2PMessenger defines a subset of the p2p.Messenger interface
type P2PMessenger interface {
	io.Closer
	Bootstrap() error
	Broadcast(topic string, buff []byte)
	BroadcastOnChannel(channel string, topic string, buff []byte)
	BroadcastOnChannelBlocking(channel string, topic string, buff []byte) error
	CreateTopic(name string, createChannelForTopic bool) error
	HasTopic(name string) bool
	HasTopicValidator(name string) bool
	RegisterMessageProcessor(topic string, handler p2p.MessageProcessor) error
	PeerAddress(pid p2p.PeerID) string
	IsConnectedToTheNetwork() bool
	ID() p2p.PeerID
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
	ApplyConsensusSize(size int)
	IsInterfaceNil() bool
}

// Accumulator defines the interface able to accumulate data and periodically evict them
type Accumulator interface {
	AddData(data interface{})
	OutputChannel() <-chan []interface{}
	IsInterfaceNil() bool
}

// HardforkTrigger defines the behavior of a hardfork trigger
type HardforkTrigger interface {
	TriggerReceived(payload []byte, data []byte, pkBytes []byte) (bool, error)
	RecordedTriggerMessage() ([]byte, bool)
	Trigger(epoch uint32) error
	CreateData() []byte
	IsSelfTrigger() bool
	IsInterfaceNil() bool
}

// Throttler can monitor the number of the currently running go routines
type Throttler interface {
	CanProcess() bool
	StartProcessing()
	EndProcessing()
	IsInterfaceNil() bool
}
