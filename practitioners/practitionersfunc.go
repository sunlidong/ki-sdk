package practitioners

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"ki-sdk/configless"
	"regexp"
	"strings"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/cryptoutil"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/endpoint"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/cryptosuite"
	"github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
	"github.com/pkg/errors"
)

//PeerMSPID  returns the mspID for the given org name in the arg
func PeerMSPID(name string) (string, bool) {
	// Find organisation/msp that peer belongs to
	for _, org := range P.ShiLi.orgsConfig {
		for i := 0; i < len(org.Peers); i++ {
			if strings.EqualFold(org.Peers[i], name) {
				// peer belongs to this org add org msp
				return org.MSPID, true
			}
		}
	}

	return "", false
}

func verifyIsLocalCAsURLs(caConfigs map[string]caConfig) map[string]caConfig {
	re := regexp.MustCompile(P.ShiLi.dnsMatchRegX)
	var newCfg = make(map[string]caConfig)
	// for local integration tests, replace all urls DNS to localhost:
	if configless.IsLocal() {
		for k, caCfg := range caConfigs {
			caCfg.URL = re.ReplaceAllString(caCfg.URL, P.ShiLi.localHostRep)
			newCfg[k] = caCfg
		}
	}
	return newCfg
}

func newCAsConfig() map[string]caConfig {
	c := verifyIsLocalCAsURLs(P.ShiLi.caConfigObj)
	caConfigObj = c
	return c
}

func newPeersConfig() map[string]fab.PeerConfig {
	p := verifyIsLocalPeersURLs(P.ShiLi.peersConfig)
	P.ShiLi.peersConfig = p
	return p
}

func newOrderersConfig() map[string]fab.OrdererConfig {
	o := verifyIsLocalOrderersURLs(P.ShiLi.orderersConfig)
	P.ShiLi.orderersConfig = o
	return o
}

func verifyIsLocalOrderersURLs(oConfig map[string]fab.OrdererConfig) map[string]fab.OrdererConfig {
	re := regexp.MustCompile(P.ShiLi.dnsMatchRegX)
	var newConfig = make(map[string]fab.OrdererConfig)
	// for local integration tests, replace all urls DNS to localhost:
	if configless.IsLocal() {
		for k, orderer := range oConfig {
			orderer.URL = re.ReplaceAllString(orderer.URL, P.ShiLi.localhostRep)
			newConfig[k] = orderer
		}
	}

	if len(newConfig) == 0 {
		return oConfig
	}
	return newConfig
}

//newTLSCACertPool will create a new TLSCACertPool instance with useSystemCertPool bool flag
func newTLSCACertPool(useSystemCertPool bool) *TLSCACertPool {
	m := &TLSCACertPool{}
	var err error
	m.tlsCertPool, err = commtls.NewCertPool(useSystemCertPool)
	if err != nil {
		panic(err)
	}
	return m
}

//newOrderersConfigImpl will create a new OrderersConfig instance with proper ordrerer URLs (local vs normal) tests
// local tests use localhost urls, while the remaining tests use default values as set in orderersConfig var
func newOrderersConfigImpl() *OrderersConfig {
	oConfig := verifyIsLocalOrderersURLs(P.ShiLi.orderersConfig)
	P.ShiLi.orderersConfig = oConfig
	o := &OrderersConfig{}
	return o
}

func verifyIsLocalPeersURLs(pConfig map[string]fab.PeerConfig) map[string]fab.PeerConfig {
	re := regexp.MustCompile(P.ShiLi.dnsMatchRegX)
	var newConfigs = make(map[string]fab.PeerConfig)
	// for local integration tests, replace all urls DNS to localhost:
	if configless.IsLocal() {
		for k, peer := range pConfig {
			peer.URL = re.ReplaceAllString(peer.URL, P.ShiLi.localHostRep)
			newConfigs[k] = peer
		}
	}

	if len(newConfigs) == 0 {
		return pConfig
	}
	return newConfigs
}

//newPeersConfigImpl will create a new PeersConfig instance with proper peers URLs (local vs normal) tests
// local tests use localhost urls, while the remaining tests use default values as set in peersConfig var
func newPeersConfigImpl() *PeersConfig {
	pConfig := verifyIsLocalPeersURLs(P.ShiLi.peersConfig)
	P.ShiLi.peersConfig = pConfig
	p := &PeersConfig{}
	return p
}

func newTLSConfig(path string) endpoint.TLSConfig {
	config := endpoint.TLSConfig{Path: pathvar.Subst(path)}
	if err := config.LoadBytes(); err != nil {
		panic(fmt.Sprintf("error loading bytes: %s", err))
	}
	return config
}

func tlsCertByBytes(path string) *x509.Certificate {

	bytes, err := ioutil.ReadFile(pathvar.Subst(path))
	if err != nil {
		return nil
	}

	block, _ := pem.Decode(bytes)

	if block != nil {
		pub, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			panic(err)
		}

		return pub
	}

	//no cert found and there is no error
	return nil
}

//newTLSCACertPool will create a new TLSCACertPool instance with useSystemCertPool bool flag
func newTLSCACertPool(useSystemCertPool bool) *TLSCACertPool {
	m := &TLSCACertPool{}
	var err error
	m.tlsCertPool, err = commtls.NewCertPool(useSystemCertPool)
	if err != nil {
		panic(err)
	}
	return m
}

//Timeout overrides EndpointConfig's Timeout function which returns the timeout for the given timeoutType in the arg
func (m *Timeout) Timeout(tType fab.TimeoutType) time.Duration {
	t, ok := defaultTypes[tType]
	if !ok {
		return time.Second * 30 // general default if type is not found
	}
	return t
}

//OrderersConfig overrides EndpointConfig's OrderersConfig function which returns the ordererConfigs list
func (m *OrderersConfig) OrderersConfig() []fab.OrdererConfig {
	orderers := []fab.OrdererConfig{}

	for _, orderer := range P.ShiLi.orderersConfig {

		if orderer.TLSCACert == nil && !m.isSystemCertPool {
			return nil
		}
		orderers = append(orderers, orderer)
	}

	return orderers
}

//OrdererConfig overrides EndpointConfig's OrdererConfig function which returns the ordererConfig instance for the name/URL arg
func (m *OrdererConfig) OrdererConfig(ordererNameOrURL string) (*fab.OrdererConfig, bool, bool) {
	orderer, ok := P.ShiLi.networkConfig.Orderers[strings.ToLower(ordererNameOrURL)]
	if !ok {
		// EntityMatchers are not used in this implementation, below is an  of how to use them if needed, see default implementation for live
		//matchingOrdererConfig := m.tryMatchingOrdererConfig(networkConfig, strings.ToLower(ordererNameOrURL))
		//if matchingOrdererConfig == nil {
		//	return nil, errors.WithStack(status.New(status.ClientStatus, status.NoMatchingOrdererEntity.ToInt32(), "no matching orderer config found", nil))
		//}
		//orderer = *matchingOrdererConfig
		return nil, false, false
	}

	return &orderer, true, false
}

//PeersConfig overrides EndpointConfig's PeersConfig function which returns the peersConfig list
func (m *PeersConfig) PeersConfig(org string) ([]fab.PeerConfig, bool) {
	orgPeers := P.ShiLi.orgsConfig[strings.ToLower(org)].Peers
	peers := []fab.PeerConfig{}

	for _, peerName := range orgPeers {
		p := P.ShiLi.networkConfig.Peers[strings.ToLower(peerName)]
		if err := m.verifyPeerConfig(p, peerName, endpoint.IsTLSEnabled(p.URL)); err != nil {
			// EntityMatchers are not used in this implementation, below is an  of how to use them if needed
			//matchingPeerConfig := m.tryMatchingPeerConfig(networkConfig, peerName)
			//if matchingPeerConfig == nil {
			//	continue
			//}
			//
			//p = *matchingPeerConfig
			return nil, false
		}
		peers = append(peers, p)
	}
	return peers, true
}

func (m *PeersConfig) verifyPeerConfig(p fab.PeerConfig, peerName string, tlsEnabled bool) error {
	if p.URL == "" {
		return errors.Errorf("URL does not exist or empty for peer %s", peerName)
	}
	if tlsEnabled && p.TLSCACert == nil && !m.isSystemCertPool {
		return errors.Errorf("tls.certificate does not exist or empty for peer %s", peerName)
	}
	return nil
}

// PeerConfig overrides EndpointConfig's PeerConfig function which returns the peerConfig instance for the name/URL arg
func (m *PeerConfig) PeerConfig(nameOrURL string) (*fab.PeerConfig, bool) {
	pcfg, ok := P.ShiLi.peersConfig[nameOrURL]
	if ok {
		return &pcfg, true
	}

	if configless.IsLocal() {
		pcfg, ok := P.ShiLi.peersByLocalURL[nameOrURL]
		if ok {
			return &pcfg, true
		}
	}

	i := strings.Index(nameOrURL, ":")
	if i > 0 {
		return m.PeerConfig(nameOrURL[0:i])
	}

	return nil, false
}

// NetworkConfig overrides EndpointConfig's NetworkConfig function which returns the full network Config instance
func (m *NetworkConfig) NetworkConfig() *fab.NetworkConfig {
	return &P.ShiLi.networkConfig
}

//NetworkPeers overrides EndpointConfig's NetworkPeers function which returns the networkPeers list
func (m *NetworkPeers) NetworkPeers() []fab.NetworkPeer {
	netPeers := []fab.NetworkPeer{}
	// referencing another interface to call PeerMSPID to match config yaml content

	for name, p := range P.ShiLi.networkConfig.Peers {

		if err := m.verifyPeerConfig(p, name, endpoint.IsTLSEnabled(p.URL)); err != nil {
			return nil
		}

		mspID, ok := PeerMSPID(name)
		if !ok {
			return nil
		}

		netPeer := fab.NetworkPeer{PeerConfig: p, MSPID: mspID}
		netPeers = append(netPeers, netPeer)
	}

	return netPeers
}

func (m *NetworkPeers) verifyPeerConfig(p fab.PeerConfig, peerName string, tlsEnabled bool) error {
	if p.URL == "" {
		return errors.Errorf("URL does not exist or empty for peer %s", peerName)
	}
	if tlsEnabled && p.TLSCACert == nil && !m.isSystemCertPool {
		return errors.Errorf("tls.certificate does not exist or empty for peer %s", peerName)
	}
	return nil
}

// ChannelConfig overrides EndpointConfig's ChannelConfig function which returns the channelConfig instance for the channel name arg
func (m *ChannelConfig) ChannelConfig(channelName string) *fab.ChannelEndpointConfig {
	ch, ok := P.ShiLi.channelsConfig[strings.ToLower(channelName)]
	if !ok {
		// EntityMatchers are not used in this implementation, below is an  of how to use them if needed
		//matchingChannel, _, matchErr := m.tryMatchingChannelConfig(channelName)
		//if matchErr != nil {
		//	return nil, errors.WithMessage(matchErr, "channel config not found")
		//}
		//return matchingChannel, nil
		return &fab.ChannelEndpointConfig{}
	}

	return &ch
}

// ChannelPeers overrides EndpointConfig's ChannelPeers function which returns the list of peers for the channel name arg
func (m *ChannelPeers) ChannelPeers(channelName string) []fab.ChannelPeer {
	peers := []fab.ChannelPeer{}

	chConfig, ok := P.ShiLi.channelsConfig[strings.ToLower(channelName)]
	if !ok {
		// EntityMatchers are not used in this implementation, below is an  of how to use them if needed
		//matchingChannel, _, matchErr := m.tryMatchingChannelConfig(channelName)
		//if matchErr != nil {
		//	return peers, nil
		//}
		//
		//// reset 'name' with the mappedChannel as it's referenced further below
		//chConfig = *matchingChannel
		return nil
	}

	for peerName, chPeerConfig := range chConfig.Peers {

		// Get generic peer configuration
		p, ok := P.ShiLi.peersConfig[strings.ToLower(peerName)]
		if !ok {
			// EntityMatchers are not used in this implementation, below is an  of how to use them if needed
			//matchingPeerConfig := m.tryMatchingPeerConfig(networkConfig, strings.ToLower(peerName))
			//if matchingPeerConfig == nil {
			//	continue
			//}
			//p = *matchingPeerConfig
			return nil
		}

		if err := m.verifyPeerConfig(p, peerName, endpoint.IsTLSEnabled(p.URL)); err != nil {
			return nil
		}

		mspID, ok := PeerMSPID(peerName)
		if !ok {
			return nil
		}

		networkPeer := fab.NetworkPeer{PeerConfig: p, MSPID: mspID}

		peer := fab.ChannelPeer{PeerChannelConfig: chPeerConfig, NetworkPeer: networkPeer}

		peers = append(peers, peer)
	}

	return peers

}

func (m *ChannelPeers) verifyPeerConfig(p fab.PeerConfig, peerName string, tlsEnabled bool) error {
	if p.URL == "" {
		return errors.Errorf("URL does not exist or empty for peer %s", peerName)
	}
	if tlsEnabled && p.TLSCACert == nil && !m.isSystemCertPool {
		return errors.Errorf("tls.certificate does not exist or empty for peer %s", peerName)
	}
	return nil
}

// ChannelOrderers overrides EndpointConfig's ChannelOrderers function which returns the list of orderers for the channel name arg
func (m *ChannelOrderers) ChannelOrderers(channelName string) []fab.OrdererConfig {
	// referencing other interfaces to call ChannelConfig and OrdererConfig to match config yaml content
	chCfg := &ChannelConfig{}
	oCfg := &OrdererConfig{}

	orderers := []fab.OrdererConfig{}
	channel := chCfg.ChannelConfig(channelName)

	for _, chOrderer := range channel.Orderers {
		orderer, ok, _ := oCfg.OrdererConfig(chOrderer)
		if !ok || orderer == nil {
			return nil
		}
		orderers = append(orderers, *orderer)
	}

	return orderers
}

// TLSCACertPool overrides EndpointConfig's TLSCACertPool function which will add the list of cert args to the cert pool and return it
func (m *TLSCACertPool) TLSCACertPool() commtls.CertPool {
	return m.tlsCertPool
}

// TLSClientCerts overrides EndpointConfig's TLSClientCerts function which will return the list of configured client certs
func (m *TLSClientCerts) TLSClientCerts() []tls.Certificate {
	var clientCerts tls.Certificate
	cb := P.ShiLi.client.TLSCerts.Client.Cert.Bytes()

	if len(cb) == 0 {
		// if no cert found in the config, return empty cert chain
		return []tls.Certificate{clientCerts}
	}

	// Load private key from cert using default crypto suite
	cs := cryptosuite.GetDefault()
	pk, err := cryptoutil.GetPrivateKeyFromCert(cb, cs)

	// If CryptoSuite fails to load private key from cert then load private key from config
	if err != nil || pk == nil {
		m.RWLock.Lock()
		defer m.RWLock.Unlock()
		ccs, err := m.loadPrivateKeyFromConfig(&P.ShiLi.client, clientCerts, cb)
		if err != nil {
			return nil
		}
		return ccs
	}

	// private key was retrieved from cert
	clientCerts, err = cryptoutil.X509KeyPair(cb, pk, cs)
	if err != nil {
		return nil
	}

	return []tls.Certificate{clientCerts}
}
func (m *TLSClientCerts) loadPrivateKeyFromConfig(clientConfig *clientConfig, clientCerts tls.Certificate, cb []byte) ([]tls.Certificate, error) {

	kb := clientConfig.TLSCerts.Client.Key.Bytes()

	// load the key/cert pair from []byte
	clientCerts, err := tls.X509KeyPair(cb, kb)
	if err != nil {
		return nil, errors.Errorf("Error loading cert/key pair as TLS client credentials: %s", err)
	}

	return []tls.Certificate{clientCerts}, nil
}

func (m *CryptoConfigPath) CryptoConfigPath() string {
	return P.ShiLi.client.CryptoConfig.Path
}
