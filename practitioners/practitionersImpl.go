package practitioners

import (
	"sync"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

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

func NewPractitioners(Name string, Version int64, Open bool) *Practitioners {

	return &Practitioners{
		Name:    Name,
		Version: Version,
		Open:    Open,
		ShiLi:   NewEndpointConfig(),
	}
}

func NewEndpointConfig() *EndpointConfig {

	return &EndpointConfig{
		localHostRep: func() string {
			return "localHostRep"
		}(),
		dnsMatchRegX: func() string {
			return "dnsMatchRegX"
		}(),
		tlsPath: func() string {
			return "tlsPath"
		}(),
		ordererPath: func() string {
			return "ordererPath"
		}(),
		client: func() clientConfig {
			clientBy1 := clientConfig{}
			return clientBy1
		}(),
		channelsConfig: func() map[string]fab.ChannelEndpointConfig {
			channelsConfigBy1 := map[string]fab.ChannelEndpointConfig{}
			return channelsConfigBy1
		}(),
		orgsConfig: func() map[string]fab.OrganizationConfig {
			orgsConfigBy1 := map[string]fab.OrganizationConfig{}
			return orgsConfigBy1
		}(),
		orderersConfig: func() map[string]fab.OrdererConfig {
			orderersConfigBy1 := map[string]fab.OrdererConfig{}
			return orderersConfigBy1
		}(),
		peersConfig: func() map[string]fab.PeerConfig {
			peersConfigBy1 := map[string]fab.PeerConfig{}
			return peersConfigBy1
		}(),
		peersByLocalURL: func() map[string]fab.PeerConfig {
			peersByLocalURLBy1 := map[string]fab.PeerConfig{}
			return peersByLocalURLBy1
		}(),
		caConfigObj: func() map[string]caConfig {
			caConfigObjBy1 := map[string]caConfig{}
			return caConfigObjBy1
		}(),
		networkConfig: func() fab.NetworkConfig {
			NetworkConfigBy1 := fab.NetworkConfig{}
			return NetworkConfigBy1
		}(),
		endpointConfigImpls: func() []interface{} {
			return []interface{}{}
		}(),
	}
}

 
// DB 
type Practitioners struct {
	Name     	string          	`json:"name"`
	Version  	int64           	`json:"version"`
	Open     	bool            	`json:"open"`
	ShiLi    	*EndpointConfig 	`json:"ShiLi"`
	CretPath  	string  		 	`json:"CretPath"`
	ChannelName []string     		`json:"channelName"`
	Peer 		[]Peer       		`json:"peer"`
	Orderer 	[]Orderer     		`json:"orderer"`
	Ca   		[]CA    			`json:"ca"`
}


// peer info
type  Peer struct {
	Com   		string 			 `json:"com"`
	Name   		string           `json:"name"`
	Domain  	string           `json:"domain"`
	Template  	int64            `json:"template"`
	Users  		string           `json:"users"`
	MspID  		string           `json:"mspid"`
	PeerName 	string			`json:"peerName"`
}

// Orderer info
type  Orderer struct {
	Name   		string           `json:"name"`
	Domain  	string           `json:"domain"`
	Com   		string 			 `json:"com"`
	Template  	[]string         `json:"template"`
	Users  		string           `json:"users"`
	MspID  		string           `json:"mspid"`
}


// CA 
type  CA struct {
	Name   			string           `json:"name"`
	Domain  		string           `json:"domain"`
	EnrollID 		string           `json:"enrollid"`
	EnrollSecret 	string           `json:"enrollSecret"`
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

func (p *Practitioners) setEndpointConfigImpls() {

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

func (p *Practitioners) InitApi() {
	p.initApi()

}
func (p *Practitioners) initApi() {
	SetupAndRuning(false, nil,
		fabsdk.WithEndpointConfig(endpointConfigImpls...),
		fabsdk.WithCryptoSuiteConfig(cryptoConfigImpls...),
		fabsdk.WithIdentityConfig(identityConfigImpls...),
		fabsdk.WithMetricsConfig(operationsConfigImpls...),
	)

}

//  new
func (p *Practitioners) setPeerName(peerlist []string) {

	// 校验参数
func (p *Practitioners) setPeerName(peerlist []string) err error{
	if len(peerlist)<=0{
		return  errors.New("err: peerlist len  is 0 ")
	}

	for k,v:=range peerlist{
		if k!=""{
			p.PeerName=append(p.PeerName,v)
		}else{
			log.Printf("peerlist %s is nil",k)
		}
	}
}

func (p *Practitioners) setOrdererName(ordererlist []string) err error{
	if len(ordererlist)<=0{
		return  errors.New("err: ordererlist len  is 0 ")
	}

	for k,v:=range ordererlist{
		if k!=""{
			p.=append(p.OrdererName,v)
		}else{
			log.Printf("OrdererName %s is nil",k)
		}
	}
}


//  获取 通道 排序节点
func (p *Practitioners) getOrdererNameBychannel() []string{

		if  len(p.Orderer)>0{
		return func()(res []string){
				if len(p.Orderer.Template)>0{
					for k,v:=range  p.Orderer.Template{
						res= append(res,func()(string){
							return p.Orderer.Domain+"."+v+"."+p.Orderer.Domain+"."+p.Orderer.Com
						}())
					}
				}
				return res
			}()
		}
		return nil
}

//  channel peer   
func (p *Practitioners) getPeerNameBychannel() (map[string]fab.PeerChannelConfig){

	if len(p.Peer) >0 {

		return func()(res map[string]fab.PeerChannelConfig){

				for k,v:=range p.Peer{

					for i:=0;i<p.Peer[k].Template{
						
						res[p.Peer[k].PeerName+i+"."+p.Peer[k].Name+"."+p.Peer[k].Domain+"."+p.Peer[k].Com]=fab.PeerChannelConfig{
							EndorsingPeer:  true,
							ChaincodeQuery: true,
							LedgerQuery:    true,
							EventSource:    true,
						}	
					}
				}
				return  res 
		}()
	}
	log.Pritln("getPeerNameBychannel len  is nil ")
	return nil 

 
}

//  new
func (p *Practitioners) getPoliciesBychannel() (fab.ChannelPolicies){
	return fab.ChannelPolicies{
		QueryChannelConfig: fab.QueryChannelConfigPolicy{
			MinResponses: 1,
			MaxTargets:   1,
			RetryOpts: retry.Opts{
				Attempts:       5,
				InitialBackoff: 500 * time.Millisecond,
				MaxBackoff:     5 * time.Second,
				BackoffFactor:  2.0,
			},
		},
		EventService: fab.EventServicePolicy{
			ResolverStrategy:                 fab.MinBlockHeightStrategy,
			MinBlockHeightResolverMode:       fab.ResolveByThreshold,
			BlockHeightLagThreshold:          5,
			ReconnectBlockHeightLagThreshold: 10,
			PeerMonitorPeriod:                5 * time.Second,
		},
	}
 

}


//  set channel name 
func (p *Practitioners) setChannelName(channellist []string) err error{
	if len(channellist)<=0{
		return  errors.New("err: channellist len  is 0 ")
	}

	for k,v:=range channellist{
		if k!=""{
			p.ChannelName=append(p.ChannelName,v)
		}else{
			log.Printf("channellist %s is nil",k)
		}
	}
}







//-----------------------------------------------------------------------------



//  new 
func (p *Practitioners) setChannelsConfig() {

	// 校验参数
	p.Shili.channelsConfig = func()(map[string]fab.ChannelEndpointConfig){
		
		return &map[string]fab.ChannelEndpointConfig{
			p.Channel[0]: {
				Orderers: p.getOrdererNameBychannel(),
				Peers: p.getPeerNameBychannel(),
				Policies: p.getPoliciesBychannel(),
			},
		}
	}()
}


//  set setorgsConfig
func (p *Practitioners) setOrgsConfig() {

	if  len(p.Peer)>0{
		func()(res map[string]fab.OrganizationConfig){
			
		// 循环 
// peer info
// type  Peer struct {
// 	Com   		string 			 `json:"com"`
// 	Name   		string           `json:"name"`
// 	Domain  	string           `json:"domain"`
// 	Template  	int64            `json:"template"`
// 	Users  		string           `json:"users"`
// 	MspID  		string           `json:"mspid"`
// 	PeerName 	string			`json:"peerName"`
// }

			for k,v :=range  p.Peer {
				
			res[p.Peer[k].Name]=fab.OrganizationConfig{
				MSPID:      p.Peer[k].MspID,
				CryptoPath: "peerOrganizations/"+p.Peer[k].Name+"."+p.Peer[k].Domain+"."+p.Peer[k].Com+"/"+"users/{username}"+"@"+p.Peer[k].Name+"."+p.Peer[k].Domain+"."+p.Peer[k].Com+"/msp",
				Peers: []string{
					"peer0.org1.bookstore.com",
					"peer1.org1.bookstore.com",

			}


			}
		}()


	}


}

	 

//  set setOrderersConfig
func (p *Practitioners) setOrderersConfig() {

	p.Shili.orderersConfig = func() res map[string]fab.OrdererConfig{
		if len(p.OrdererName) >0{
			for k,v:=range  p.OrdererName {
				res[v] =fab.OrdererConfig{
					URL: "orderer.example.com:7050",
					GRPCOptions: map[string]interface{}{
						"ssl-target-name-override": "orderer.example.com",
						"keep-alive-time":          0 * time.Second,
						"keep-alive-timeout":       20 * time.Second,
						"keep-alive-permit":        false,
						"fail-fast":                false,
						"allow-insecure":           false,
					},
					TLSCACert: tlsCertByBytes("${FABRIC_SDK_GO_PROJECT_PATH}/${CRYPTOCONFIG_FIXTURES_PATH}/ordererOrganizations/example.com/tlsca/tlsca.example.com-cert.pem"),
				}
			}
		}
		return 
	}()
}



//  peer  config  set  
func (p *Practitioners)setPeersConfig(){

	p.Shili.peersConfig = func() res map[string]fab.PeerConfig{
		if len(p.PeerName)>0{
			for k,v:=range p.PeerName{
				res[v]= fab.PeerConfig{
					URL: "peer0.org1.example.com:7051",
					GRPCOptions: map[string]interface{}{
						"ssl-target-name-override": "peer0.org1.example.com",
						"keep-alive-time":          0 * time.Second,
						"keep-alive-timeout":       20 * time.Second,
						"keep-alive-permit":        false,
						"fail-fast":                false,
						"allow-insecure":           false,
					}
				}
			}
		}
		return 
	}()

}
 


// set setPeersByLocalURL 
func (p *Practitioners)setPeersByLocalURL(){

	p.Shili.peersByLocalURL = func() res map[string]fab.PeerConfig{
		if len(p.PeerName)>0{
			for k,v:=range p.PeerName{
				res[v]= fab.PeerConfig{
					URL: "peer0.org1.example.com:7051",
					GRPCOptions: map[string]interface{}{
						"ssl-target-name-override": "peer0.org1.example.com",
						"keep-alive-time":          0 * time.Second,
						"keep-alive-timeout":       20 * time.Second,
						"keep-alive-permit":        false,
						"fail-fast":                false,
						"allow-insecure":           false,
					}
				}
			}
		}
		return 
	}()
}



func(p *Practitioners)setCaConfigObj(){

	p.Shili.caConfigObj =  func() res map[string]caConfig {

		if  len(p.CA) >0{
			for k,v:=range p.CA {
				res[v] = caConfig{
					ID:  "ca.org1.example.com",
					URL: "https://ca.org1.example.com:7054",
					TLSCACerts: endpoint.MutualTLSConfig{
						Path: pathvar.Subst("${FABRIC_SDK_GO_PROJECT_PATH}/${CRYPTOCONFIG_FIXTURES_PATH}/peerOrganizations/org1.example.com/tlsca/tlsca.org1.example.com-cert.pem"),
						Client: endpoint.TLSKeyPair{
							Key:  newTLSConfig("${FABRIC_SDK_GO_PROJECT_PATH}/${CRYPTOCONFIG_FIXTURES_PATH}/peerOrganizations/tls.example.com/users/User1@tls.example.com/tls/client.key"),
							Cert: newTLSConfig("${FABRIC_SDK_GO_PROJECT_PATH}/${CRYPTOCONFIG_FIXTURES_PATH}/peerOrganizations/tls.example.com/users/User1@tls.example.com/tls/client.crt"),
						},
					},
					Registrar: msp.EnrollCredentials{
						EnrollID:     "admin",
						EnrollSecret: "adminpw",
					},
					CAName: "ca.org1.example.com",	

				}
			}
		}
		return  
	}()
	// caConfigObj         map[string]caConfig
}

 
//  set  setNetworkConfig
func(p *Practitioners)setNetworkConfig(){

	p.Shili.networkConfig = func()(res fab.NetworkConfig){
		return fab.NetworkConfig{
			Channels:      channelsConfig,
			Organizations: orgsConfig,
			Orderers:      newOrderersConfig(),
			Peers:         newPeersConfig(),
		}
	}
}

func(p *Practitioners)setEndpointConfigImpls(){

	timeoutImpl          = &exampleTimeout{}
	orderersConfigImpl   = newOrderersConfigImpl()
	ordererConfigImpl    = &exampleOrdererConfig{}
	peersConfigImpl      = newPeersConfigImpl()
	peerConfigImpl       = &examplePeerConfig{}
	networkConfigImpl    = &exampleNetworkConfig{}
	networkPeersImpl     = &exampleNetworkPeers{}
	channelConfigImpl    = &exampleChannelConfig{}
	channelPeersImpl     = &exampleChannelPeers{}
	channelOrderersImpl  = &exampleChannelOrderers{}
	tlsCACertPoolImpl    = newTLSCACertPool(false)
	tlsClientCertsImpl   = &exampleTLSClientCerts{}
	cryptoConfigPathImpl = &exampleCryptoConfigPath{}


	p.Shili.endpointConfigImpls =  []interface{}{
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
}