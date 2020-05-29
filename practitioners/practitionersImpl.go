package practitioners

import (
	"sync"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/endpoint"
	logApi "github.com/hyperledger/fabric-sdk-go/pkg/core/logging/api"
)

var defaultTypes = map[fab.TimeoutType]time.Duration{
	fab.PeerConnection:           time.Second * 10,
	fab.PeerResponse:             time.Minute * 3,
	fab.DiscoveryGreylistExpiry:  time.Second * 10,
	fab.EventReg:                 time.Second * 15,
	fab.OrdererConnection:        time.Second * 15,
	fab.OrdererResponse:          time.Minute * 2,
	fab.DiscoveryConnection:      time.Second * 15,
	fab.DiscoveryResponse:        time.Second * 15,
	fab.Query:                    time.Minute * 3,
	fab.Execute:                  time.Minute * 3,
	fab.ResMgmt:                  time.Minute * 3,
	fab.ConnectionIdle:           time.Second * 30,
	fab.EventServiceIdle:         time.Minute * 2,
	fab.ChannelConfigRefresh:     time.Minute * 90,
	fab.ChannelMembershipRefresh: time.Second * 60,
	fab.DiscoveryServiceRefresh:  time.Second * 10,
	fab.SelectionServiceRefresh:  time.Minute * 15,
	// EXPERIMENTAL - do we need this to be configurable?
	fab.CacheSweepInterval: time.Second * 15,
}

var P Practitioners

// init

func NewPractitioners(Name sring, Version int64, Open bool) {

	return &Practitioners{
		Name:    Name,
		Version: Version,
		Open:    Open,
		ShiLi: NewEndpointConfig(),
	}
}

func NewEndpointConfig(){
	return &EndpointConfig{
		localHostRep:"",
		dnsMatchRegX:"", 
		tlsPath:"",
		ordererPath:"", 
		client: func()clientConfig{clientBy1:=clientConfig{};return clientBy1}(), 
		channelsConfig:func()map[string]fab.ChannelEndpointConfig{channelsConfigBy1:=map[string]fab.ChannelEndpointConfig{};return channelsConfigBy1}(),
		orgsConfig:func()map[string]fab.OrganizationConfig{orgsConfigBy1:=map[string]fab.OrganizationConfig{};return orgsConfigBy1}(),
		orderersConfig:func()map[string]fab.OrdererConfig{
			orderersConfigBy1:=map[string]fab.OrdererConfig{"1":fab.OrdererConfig{
				
			}}
			return  orderersConfigBy1
		},
		peersConfig:func()map[string]fab.PeerConfig{
			peersConfigBy1:=map[string]fab.PeerConfig{}
			return peersConfigBy1
		},
		peersByLocalURL:func()map[string]fab.PeerConfig{peersByLocalURLBy1:=map[string]fab.PeerConfig{};return peersByLocalURLBy1},
		caConfigObj:func()map[string]caConfig{caConfigObjBy1:=map[string]caConfig{}; return caConfigObjBy1},
		networkConfig:func()fab.NetworkConfig{ return &fab.NetworkConfig{}},
		endpointConfigImpls:func()[]interface{}{endpointConfigImplsBy1:=[]interface{};return endpointConfigImplsBy1}(),
	}

}
type Practitioners struct {
	Name   string         `json:"name"`
	Verson int64          `json:"version"`
	Open   bool           `json:"open"`
	ShiLi  EndpointConfig `json:"ShiLi"`
}

type EndpointConfig struct {
	localHostRep        string       `json:"localhost"`
	dnsMatchRegX        string       `json:"dnsMatchRegX"`
	tlsPath             string       `json:"tlsPath"`
	ordererPath         string       `json:"ordererPath"`
	client              clientConfig `json:"client"`
	channelsConfig      map[string]fab.ChannelEndpointConfig
	orgsConfig          map[string]fab.OrganizationConfig
	orderersConfig      map[string]fab.OrdererConfig
	peersConfig         map[string]fab.PeerConfig
	peersByLocalURL     map[string]fab.PeerConfig
	caConfigObj         map[string]caConfig
	networkConfig       fab.NetworkConfig
	endpointConfigImpls []interface{}
}

type clientConfig struct {
	Organization    string
	Logging         logApi.LoggingType
	CryptoConfig    msp.CCType
	TLSCerts        endpoint.MutualTLSConfig
	TLSKey          []byte
	TLSCert         []byte
	CredentialStore msp.CredentialStoreType
}
type caConfig struct {
	ID         string
	URL        string
	TLSCACerts endpoint.MutualTLSConfig
	Registrar  msp.EnrollCredentials
	CAName     string
}

type Timeout struct{}

type OrderersConfig struct {
	isSystemCertPool bool
}
type OrdererConfig struct{}

type PeersConfig struct {
	isSystemCertPool bool
}

type PeerConfig struct{}

type NetworkConfig struct{}

type NetworkPeers struct {
	isSystemCertPool bool
}

type ChannelConfig struct{}

type ChannelPeers struct {
	isSystemCertPool bool
}
type ChannelOrderers struct{}

type TLSCACertPool struct {
	tlsCertPool commtls.CertPool
}

type TLSClientCerts struct {
	RWLock sync.RWMutex
}

type CryptoConfigPath struct{}

//  new  endpointConfigImpls

func (p *EndpointConfig) setEndpointConfigImpls() {

	//  new
	timeoutImpl := &Timeout{}
	orderersConfigImpl := newOrderersConfigImpl()
	ordererConfigImpl := &OrdererConfig{}
	peersConfigImpl := newPeersConfigImpl()
	peerConfigImpl := &PeerConfig{}
	networkConfigImpl := &NetworkConfig{}
	networkPeersImpl := &NetworkPeers{}
	channelConfigImpl := &ChannelConfig{}
	channelPeersImpl := &ChannelPeers{}
	channelOrderersImpl := &ChannelOrderers{}
	tlsCACertPoolImpl := newTLSCACertPool(false)
	tlsClientCertsImpl := &TLSClientCerts{}
	cryptoConfigPathImpl := &CryptoConfigPath{}

	endpointConfigImpls := []interface{}{
		timeoutImpl,
		orderersConfigImpl,
		ordererConfigImpl,
		peersConfigImpl,
		peerConfigImpl,
		networkConfigImpl,
		networkPeersImpl,
		channelConfigImpl,
		channelPeersImpl,
		channelOrderersImpl,
		tlsCACertPoolImpl,
		tlsClientCertsImpl,
		cryptoConfigPathImpl,
	}

	//
	p.endpointConfigImpls = endpointConfigImpls

}
