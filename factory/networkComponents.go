package factory

import (
	"fmt"

	"github.com/ElrondNetwork/elrond-go/config"
	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/p2p/libp2p"
	antifloodFactory "github.com/ElrondNetwork/elrond-go/process/throttle/antiflood/factory"
)

type networkComponentsFactory struct {
	p2pConfig     config.P2PConfig
	mainConfig    config.Config
	statusHandler core.AppStatusHandler
	listenAddress string
}

// NewNetworkComponentsFactory returns a new instance of a network components factory
func NewNetworkComponentsFactory(
	p2pConfig config.P2PConfig,
	mainConfig config.Config,
	statusHandler core.AppStatusHandler,
) (*networkComponentsFactory, error) {
	if check.IfNil(statusHandler) {
		return nil, ErrNilStatusHandler
	}

	return &networkComponentsFactory{
		p2pConfig:     p2pConfig,
		mainConfig:    mainConfig,
		statusHandler: statusHandler,
		listenAddress: libp2p.ListenAddrWithIp4AndTcp,
	}, nil
}

// Create creates and returns the network components
func (ncf *networkComponentsFactory) Create() (*NetworkComponents, error) {
	arg := libp2p.ArgsNetworkMessenger{
		ListenAddress: ncf.listenAddress,
		P2pConfig:     ncf.p2pConfig,
	}

	netMessenger, err := libp2p.NewNetworkMessenger(arg)
	if err != nil {
		return nil, err
	}

	inAntifloodHandler, p2pPeerBlackList, errNewAntiflood := antifloodFactory.NewP2PAntiFloodAndBlackList(
		ncf.mainConfig,
		ncf.statusHandler,
	)
	if errNewAntiflood != nil {
		return nil, errNewAntiflood
	}

	inputAntifloodHandler, ok := inAntifloodHandler.(P2PAntifloodHandler)
	if !ok {
		return nil, fmt.Errorf("%w when casting input antiflood handler to structs/P2PAntifloodHandler", ErrWrongTypeAssertion)
	}

	outAntifloodHandler, errOutputAntiflood := antifloodFactory.NewP2POutputAntiFlood(ncf.mainConfig)
	if errOutputAntiflood != nil {
		return nil, errOutputAntiflood
	}

	outputAntifloodHandler, ok := outAntifloodHandler.(P2PAntifloodHandler)
	if !ok {
		return nil, fmt.Errorf("%w when casting output antiflood handler to structs/P2PAntifloodHandler", ErrWrongTypeAssertion)
	}

	err = netMessenger.SetPeerBlackListHandler(p2pPeerBlackList)
	if err != nil {
		return nil, err
	}

	return &NetworkComponents{
		NetMessenger:           netMessenger,
		InputAntifloodHandler:  inputAntifloodHandler,
		OutputAntifloodHandler: outputAntifloodHandler,
		PeerBlackListHandler:   p2pPeerBlackList,
	}, nil
}
