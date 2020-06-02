package practitioners

import (
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/endpoint"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/logging/api"

	"github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
)

////////////////////////////////
var (
	localhostRep = "localhost:"
	dnsMatchRegX = ".*:"
	TLS_PATH     = "/mnt/d/gopath/src/ki-kyc/crypto-config/peerOrganizations/org1.bookstore.com/users/User1@org1.bookstore.com/tls/"
	ORDREER_PAHT = "/mnt/d/gopath/src/ki-kyc/crypto-config/"
	client       = clientConfig{
		Organization:    "org1",
		Logging:         api.LoggingType{Level: "info"},
		CryptoConfig:    msp.CCType{Path: pathvar.Subst("/mnt/d/gopath/src/ki-kyc/crypto-config/")},
		CredentialStore: msp.CredentialStoreType{Path: "/tmp/msp"},
		TLSCerts: endpoint.MutualTLSConfig{Client: endpoint.TLSKeyPair{
			Key:  newTLSConfig(TLS_PATH + "client.key"),
			Cert: newTLSConfig(TLS_PATH + "client.crt")}},
	}

	channelsConfig = map[string]fab.ChannelEndpointConfig{
		"bookchannel": {
			Orderers: []string{
				"orderer1.bookstore.com",
				"orderer2.bookstore.com",
				"orderer3.bookstore.com",
				"orderer4.bookstore.com",
				"orderer5.bookstore.com",
			},
			Peers: map[string]fab.PeerChannelConfig{
				"peer0.org1.bookstore.com": {
					EndorsingPeer:  true,
					ChaincodeQuery: true,
					LedgerQuery:    true,
					EventSource:    true,
				},
				"peer1.org1.bookstore.com": {
					EndorsingPeer:  true,
					ChaincodeQuery: true,
					LedgerQuery:    true,
					EventSource:    true,
				},
				"peer0.org2.bookstore.com": {
					EndorsingPeer:  true,
					ChaincodeQuery: true,
					LedgerQuery:    true,
					EventSource:    true,
				},
				"peer1.org2.bookstore.com": {
					EndorsingPeer:  true,
					ChaincodeQuery: true,
					LedgerQuery:    true,
					EventSource:    true,
				},
			},
			Policies: fab.ChannelPolicies{
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
			},
		},
	}
	orgsConfig = map[string]fab.OrganizationConfig{
		"org1": {
			MSPID:      "Org1MSP",
			CryptoPath: "peerOrganizations/org1.bookstore.com/users/{username}@org1.bookstore.com/msp",
			Peers: []string{
				"peer0.org1.bookstore.com",
				"peer1.org1.bookstore.com",
			},
		},
		"org2": {
			MSPID:      "Org2MSP",
			CryptoPath: "peerOrganizations/org2.bookstore.com/users/{username}@org2.bookstore.com/msp",
			Peers: []string{
				"peer0.org2.bookstore.com",
				"peer1.org2.bookstore.com",
			},
		},
		"ordererorg": {
			MSPID:      "OrdererMSP",
			CryptoPath: "ordererOrganizations/bookstore.com/users/{username}@bookstore.com/msp",
		},
	}

	orderersConfig = map[string]fab.OrdererConfig{
		"orderer1.bookstore.com": {
			URL: "orderer1.bookstore.com:7050",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "orderer1.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "ordererOrganizations/bookstore.com/tlsca/tlsca.bookstore.com-cert.pem"),
		},
		"orderer2.bookstore.com": {
			URL: "orderer2.bookstore.com:8050",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "orderer2.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "ordererOrganizations/bookstore.com/tlsca/tlsca.bookstore.com-cert.pem"),
		},
		"orderer3.bookstore.com": {
			URL: "orderer3.bookstore.com:9050",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "orderer3.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "ordererOrganizations/bookstore.com/tlsca/tlsca.bookstore.com-cert.pem"),
		},
		"orderer4.bookstore.com": {
			URL: "orderer4.bookstore.com:10050",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "orderer4.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "ordererOrganizations/bookstore.com/tlsca/tlsca.bookstore.com-cert.pem"),
		},
		"orderer5.bookstore.com": {
			URL: "orderer5.bookstore.com:11050",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "orderer5.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "ordererOrganizations/bookstore.com/tlsca/tlsca.bookstore.com-cert.pem"),
		},
	}

	peersConfig = map[string]fab.PeerConfig{
		"peer0.org1.bookstore.com": {
			URL: "peer0.org1.bookstore.com:7051",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "peer0.org1.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "peerOrganizations/org1.bookstore.com/tlsca/tlsca.org1.bookstore.com-cert.pem"),
		},
		"peer1.org1.bookstore.com": {
			URL: "peer1.org2.bookstore.com:8051",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "peer1.org2.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "peerOrganizations/org1.bookstore.com/tlsca/tlsca.org1.bookstore.com-cert.pem"),
		},
		"peer0.org2.bookstore.com": {
			URL: "peer0.org2.bookstore.com:9051",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "peer0.org2.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "peerOrganizations/org2.bookstore.com/tlsca/tlsca.org1.bookstore.com-cert.pem"),
		},
		"peer1.org2.bookstore.com": {
			URL: "peer1.org2.bookstore.com:10051",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "peer1.org2.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "peerOrganizations/org2.bookstore.com/tlsca/tlsca.org2.bookstore.com-cert.pem"),
		},
	}

	peersByLocalURL = map[string]fab.PeerConfig{
		"localhost:7051": {
			URL: "localhost:7051",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "peer0.org1.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "peerOrganizations/org1.bookstore.com/tlsca/tlsca.org1.bookstore.com-cert.pem"),
		},
		"localhost:8051": {
			URL: "localhost:8051",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "peer1.org1.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "peerOrganizations/org1.bookstore.com/tlsca/tlsca.org1.bookstore.com-cert.pem"),
		},
		"localhost:9051": {
			URL: "localhost:9051",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "peer0.org2.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "peerOrganizations/org2.bookstore.com/tlsca/tlsca.org2.bookstore.com-cert.pem"),
		},
		"localhost:10051": {
			URL: "localhost:10051",
			GRPCOptions: map[string]interface{}{
				"ssl-target-name-override": "peer1.org2.bookstore.com",
				"keep-alive-time":          0 * time.Second,
				"keep-alive-timeout":       20 * time.Second,
				"keep-alive-permit":        false,
				"fail-fast":                false,
				"allow-insecure":           false,
			},
			TLSCACert: tlsCertByBytes(ORDREER_PAHT + "peerOrganizations/org2.bookstore.com/tlsca/tlsca.org2.bookstore.com-cert.pem"),
		},
	}

	caConfigObj = map[string]caConfig{
		"ca.org1.bookstore.com": {
			ID:  "ca.org1.bookstore.com",
			URL: "https://ca.org1.bookstore.com:7054",
			TLSCACerts: endpoint.MutualTLSConfig{
				Path: pathvar.Subst("${FABRIC_SDK_GO_PROJECT_PATH}/${CRYPTOCONFIG_FIXTURES_PATH}/peerOrganizations/org1.bookstore.com/tlsca/tlsca.org1.bookstore.com-cert.pem"),
				Client: endpoint.TLSKeyPair{
					Key:  newTLSConfig(TLS_PATH + "client.key"),
					Cert: newTLSConfig(TLS_PATH + "client.crt"),
				},
			},
			Registrar: msp.EnrollCredentials{
				EnrollID:     "admin",
				EnrollSecret: "adminpw",
			},
			CAName: "ca.org1.bookstore.com",
		},
		"ca.org2.bookstore.com": {
			ID:  "ca.org2.bookstore.com",
			URL: "https://ca.org2.bookstore.com:8054",
			TLSCACerts: endpoint.MutualTLSConfig{
				Path: pathvar.Subst("${FABRIC_SDK_GO_PROJECT_PATH}/${CRYPTOCONFIG_FIXTURES_PATH}/peerOrganizations/org2.bookstore.com/tlsca/tlsca.org2.bookstore.com-cert.pem"),
				Client: endpoint.TLSKeyPair{
					Key:  newTLSConfig(TLS_PATH + "client.key"),
					Cert: newTLSConfig(TLS_PATH + "client.crt"),
				},
			},
			Registrar: msp.EnrollCredentials{
				EnrollID:     "admin",
				EnrollSecret: "adminpw",
			},
			CAName: "ca.org2.bookstore.com",
		},
	}

	networkConfig = fab.NetworkConfig{
		Channels:      channelsConfig,
		Organizations: orgsConfig,
		Orderers:      newOrderersConfig(),
		Peers:         newPeersConfig(),
		// EntityMatchers are not used in this implementation
		//EntityMatchers: entityMatchers,
	}

	// creating instances of each interface to be referenced in the integration tests:
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
	endpointConfigImpls  = []interface{}{
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
)

type AutoGenerated struct {
	LocalhostRep    string               `json:"localhostRep"`
	DNSMatchRegX    string               `json:"dnsMatchRegX"`
	Client          Client               `json:"client"`
	ChannelsConfig  []ChannelsConfigList `json:"channelsConfigs"`
	OrgsConfig      []OrgsConfig         `json:"orgsConfig"`
	OrderersConfig  []OrderersConfig     `json:"orderersConfig"`
	PeersConfig     []PeersConfig        `json:"peersConfig"`
	PeersByLocalURL []PeersByLocalURL    `json:"peersByLocalURL"`
	CaConfigObj     []CaConfigObj        `json:"caConfigObj"`
}

type Client struct {
	Organization    string   `json:"organization"`
	Logging         string   `json:"logging"`
	CryptoConfig    string   `json:"cryptoConfig"`
	CredentialStore string   `json:"credentialStore"`
	TLSCerts        TLSCerts `json:"tlsCerts"`
}

type TLSCerts struct {
	Key  string `json:"Key"`
	Cert string `json:"Cert"`
}

type ChannelsConfigList struct {
	ChannelName    string         `json:"channelName"`
	ChannelsConfig ChannelsConfig `json:"channelsConfig"`
}

type ChannelsConfig struct {
	Orderers []string `json:"orderers"`
	Peers    []string `json:"peers"`
	Policies Policies `json:"policies"`
}

type Policies struct {
	MinResponses int       `json:"minResponses"`
	MaxTargets   int       `json:"maxTargets"`
	RetryOpts    RetryOpts `json:"retryOpts"`
}

type RetryOpts struct {
	Attempts       int     `json:"attempts"`
	InitialBackoff int     `json:"initialBackoff"`
	MaxBackoff     int     `json:"maxBackoff"`
	BackoffFactor  float64 `json:"backoffFactor"`
}

type OrgsConfig struct {
	Org        string   `json:"org"`
	Type       string   `json:"type"`
	CryptoPath string   `json:"cryptoPath"`
	Peers      []string `json:"peers"`
}

type OrderersConfig struct {
	OrderName        string `json:"orderName"`
	URL              string `json:"url"`
	Ssl              string `json:"ssl"`
	KeepAliveTime    int    `json:"keep-alive-time"`
	KeepAliveTimeout int    `json:"keep-alive-timeout"`
	KeepAlivePermit  bool   `json:"keep-alive-permit"`
	FailFast         bool   `json:"fail-fast"`
	AllowInsecure    bool   `json:"allow-insecure"`
	TLSCaCert        string `json:"tlsCaCert"`
}

type PeersConfig struct {
	PeerURL          string `json:"peerUrl"`
	URL              string `json:"url"`
	Ssl              string `json:"ssl"`
	KeepAliveTime    int    `json:"keep-alive-time"`
	KeepAliveTimeout int    `json:"keep-alive-timeout"`
	KeepAlivePermit  bool   `json:"keep-alive-permit"`
	FailFast         bool   `json:"fail-fast"`
	AllowInsecure    bool   `json:"allow-insecure"`
	TLSCaCert        string `json:"tlsCaCert"`
}

type peersByLocalURL struct {
	PeerURL          string `json:"peerUrl"`
	URL              string `json:"url"`
	Ssl              string `json:"ssl"`
	KeepAliveTime    int    `json:"keep-alive-time"`
	KeepAliveTimeout int    `json:"keep-alive-timeout"`
	KeepAlivePermit  bool   `json:"keep-alive-permit"`
	FailFast         bool   `json:"fail-fast"`
	AllowInsecure    bool   `json:"allow-insecure"`
	TLSCaCert        string `json:"tlsCaCert"`
}

type CaConfigObj struct {
	CaOrgName    string `json:"caOrgName"`
	ID           string `json:"id"`
	URL          string `json:"url"`
	Path         string `json:"path"`
	Key          string `json:"key"`
	Cert         string `json:"cert"`
	EnrollID     string `json:"enrollID"`
	EnrollSecret string `json:"enrollSecret"`
	CaName       string `json:"caName"`
}
