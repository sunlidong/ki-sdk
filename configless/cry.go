/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package configless

import (
	"io/ioutil"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
	"github.com/pkg/errors"
)

// identityconfig_override_test.go is an example of programmatically configuring the client by injecting instances that implement IdentityConfig's functions (representing the client's msp configs)
// for the sake of overriding IdentityConfig integration tests, the structure variables below are similar to what is found in /test/fixtures/config/config_e2e.yaml
// application developers can fully override these functions to load configs in any way that suit their application need

var (

	// creating instances of each interface to be referenced in the integration tests:
	ClientImpl              = &ExampleClient{}
	CaConfigImpl            = &ExampleCaConfig{}
	CaServerCertsImpl       = &ExampleCaServerCerts{}
	CaClientKeyImpl         = &ExampleCaClientKey{}
	CaClientCertImpl        = &ExampleCaClientCert{}
	CaKeyStorePathImpl      = &ExampleCaKeyStorePath{}
	CredentialStorePathImpl = &ExampleCredentialStorePath{}

	IdentityConfigImpls = []interface{}{
		ClientImpl,
		CaConfigImpl,
		CaServerCertsImpl,
		CaClientKeyImpl,
		CaClientCertImpl,
		CaKeyStorePathImpl,
		CredentialStorePathImpl,
	}
)

type ExampleClient struct {
}

func (m *ExampleClient) Client() *msp.ClientConfig {

	return &msp.ClientConfig{
		Organization:    strings.ToLower(client.Organization),
		Logging:         client.Logging,
		CryptoConfig:    client.CryptoConfig,
		CredentialStore: client.CredentialStore,
		TLSKey:          client.TLSCerts.Client.Key.Bytes(),
		TLSCert:         client.TLSCerts.Client.Cert.Bytes(),
	}
}

type ExampleCaConfig struct{}

func (m *ExampleCaConfig) CAConfig(org string) (*msp.CAConfig, bool) {
	return getCAConfig(&networkConfig, org)
}

func getMSPCAConfig(caConfig *caConfig) (*msp.CAConfig, error) {

	serverCerts, err := getServerCerts(caConfig)
	if err != nil {
		return nil, err
	}

	return &msp.CAConfig{
		ID:               caConfig.ID,
		URL:              caConfig.URL,
		Registrar:        caConfig.Registrar,
		CAName:           caConfig.CAName,
		TLSCAClientCert:  caConfig.TLSCACerts.Client.Cert.Bytes(),
		TLSCAClientKey:   caConfig.TLSCACerts.Client.Key.Bytes(),
		TLSCAServerCerts: serverCerts,
	}, nil

}

func getServerCerts(caConfig *caConfig) ([][]byte, error) {

	var serverCerts [][]byte

	//check for pems first
	pems := caConfig.TLSCACerts.Pem
	if len(pems) > 0 {
		serverCerts = make([][]byte, len(pems))
		for i, pem := range pems {
			serverCerts[i] = []byte(pem)
		}
		return serverCerts, nil
	}

	//check for files if pems not found
	certFiles := strings.Split(caConfig.TLSCACerts.Path, ",")
	serverCerts = make([][]byte, len(certFiles))
	for i, certPath := range certFiles {
		bytes, err := ioutil.ReadFile(pathvar.Subst(certPath))
		if err != nil {
			return nil, errors.WithMessage(err, "failed to load server certs")
		}
		serverCerts[i] = bytes
	}

	return serverCerts, nil
}

// the below function is used in multiple implementations, this is fine because networkConfig is the same for all of them
func getCAConfig(networkConfig *fab.NetworkConfig, org string) (*msp.CAConfig, bool) {
	if len(networkConfig.Organizations[strings.ToLower(org)].CertificateAuthorities) == 0 {
		return nil, false
	}
	//for now, we're only loading the first Cert Authority by default. TODO add logic to support passing the Cert Authority ID needed by the client.
	caID := networkConfig.Organizations[strings.ToLower(org)].CertificateAuthorities[0]

	if caID == "" {
		return nil, false
	}

	caConfigs := newCAsConfig()
	caConfig, ok := caConfigs[strings.ToLower(caID)]
	if !ok {
		// EntityMatchers are not supported in this implementation. If needed, uncomment the below lines
		//caConfig, mappedHost := m.tryMatchingCAConfig(networkConfig, strings.ToLower(certAuthorityName))
		//if mappedHost == "" {
		return nil, false
		//}
		//return caConfig, nil
	}

	mspCAConfig, err := getMSPCAConfig(&caConfig)
	if err != nil {
		return nil, false
	}
	return mspCAConfig, true
}

type ExampleCaServerCerts struct{}

func (m *ExampleCaServerCerts) CAServerCerts(org string) ([][]byte, bool) {
	caConfig, ok := getCAConfig(&networkConfig, org)
	if !ok {
		return nil, false
	}

	return caConfig.TLSCAServerCerts, true
}

type ExampleCaClientKey struct{}

func (m *ExampleCaClientKey) CAClientKey(org string) ([]byte, bool) {
	caConfig, ok := getCAConfig(&networkConfig, org)
	if !ok {
		return nil, false
	}

	return caConfig.TLSCAClientKey, true
}

type ExampleCaClientCert struct{}

func (m *ExampleCaClientCert) CAClientCert(org string) ([]byte, bool) {
	caConfig, ok := getCAConfig(&networkConfig, org)
	if !ok {
		return nil, false
	}

	return caConfig.TLSCAClientCert, true
}

type ExampleCaKeyStorePath struct{}

func (m *ExampleCaKeyStorePath) CAKeyStorePath() string {
	return "/tmp/msp"
}

type ExampleCredentialStorePath struct{}

func (m *ExampleCredentialStorePath) CredentialStorePath() string {
	return "/tmp/state-store"
}
