/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package configless

// cryptoconfig_override_test.go is an example of programmatically configuring the client by injecting instances that implement CryptoSuiteConfig's functions (representing the client's crypto configs).
// For the sake of overriding CryptoSuiteConfig in the integration tests, the implementations below return similar values to what is found in /test/fixtures/config/config_e2e.yaml
// application developers can fully override these functions to load configs in any way that suit their application need

var (
	// creating instances of each interface to be referenced in the integration tests:
	IsSecurityEnabledImpl       = &ExampleIsSecurityEnabled{}
	SecurityAlgorithmImpl       = &ExampleSecurityAlgorithm{}
	SecurityLevelImpl           = &ExampleSecurityLevel{}
	SecurityProviderImpl        = &ExampleSecurityProvider{}
	SoftVerifyImpl              = &ExampleSoftVerify{}
	SecurityProviderLibPathImpl = &ExampleSecurityProviderLibPath{}
	SecurityProviderPinImpl     = &ExampleSecurityProviderPin{}
	SecurityProviderLabelImpl   = &ExampleSecurityProviderLabel{}
	ExampleKeyStorePathImpl     = &ExampleKeyStorePath{}
	CryptoConfigImpls           = []interface{}{
		IsSecurityEnabledImpl,
		SecurityAlgorithmImpl,
		SecurityLevelImpl,
		SecurityProviderImpl,
		SoftVerifyImpl,
		SecurityProviderLibPathImpl,
		SecurityProviderPinImpl,
		SecurityProviderLabelImpl,
		ExampleKeyStorePathImpl,
	}
)

type ExampleIsSecurityEnabled struct{}

func (m *ExampleIsSecurityEnabled) IsSecurityEnabled() bool {
	return true
}

type ExampleSecurityAlgorithm struct{}

func (m *ExampleSecurityAlgorithm) SecurityAlgorithm() string {
	return "SHA2"
}

type ExampleSecurityLevel struct{}

func (m *ExampleSecurityLevel) SecurityLevel() int {
	return 256
}

type ExampleSecurityProvider struct{}

func (m *ExampleSecurityProvider) SecurityProvider() string {
	return "sw"
}

type ExampleSoftVerify struct{}

func (m *ExampleSoftVerify) SoftVerify() bool {
	return true
}

type ExampleSecurityProviderLibPath struct{}

func (m *ExampleSecurityProviderLibPath) SecurityProviderLibPath() string {
	return ""
	// below is an example implementation with real libraries path pulled from /test/fixtures/config/config_e2e_pkcs11.yaml
	// It is not used in this e2e configless test since it uses config_e2e.yaml
	/*configuredLibs := "/usr/lib/x86_64-linux-gnu/softhsm/libsofthsm2.so, /usr/lib/softhsm/libsofthsm2.so ,/usr/lib/s390x-linux-gnu/softhsm/libsofthsm2.so, /usr/lib/powerpc64le-linux-gnu/softhsm/libsofthsm2.so, /usr/local/Cellar/softhsm/2.1.0/lib/softhsm/libsofthsm2.so"
	libPaths := strings.Split(configuredLibs, ",")
	var lib string
	for _, path := range libPaths {
		if _, err := os.Stat(strings.TrimSpace(path)); !os.IsNotExist(err) {
			lib = strings.TrimSpace(path)
			break
		}
	}

	return lib*/
}

type ExampleSecurityProviderPin struct{}

func (m *ExampleSecurityProviderPin) SecurityProviderPin() string {
	return ""
}

type ExampleSecurityProviderLabel struct{}

func (m *ExampleSecurityProviderLabel) SecurityProviderLabel() string {
	return ""
}

type ExampleKeyStorePath struct{}

func (m *ExampleKeyStorePath) KeyStorePath() string {
	return "/tmp/msp/keystore"
}
