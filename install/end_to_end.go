/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package install

import (
	"ki-sdk/e2e"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// this test mimics the original e2e test with the difference of injecting interface functions implementations
// to programmatically supply configs instead of using a yaml file. With this change, application developers can fetch
// configs from any source as long as they provide their own implementations.
func InitApi() {
	e2e.SetupAndRuning(false, nil,
		fabsdk.WithEndpointConfig(endpointConfigImpls...),
		fabsdk.WithCryptoSuiteConfig(cryptoConfigImpls...),
		fabsdk.WithIdentityConfig(identityConfigImpls...),
		fabsdk.WithMetricsConfig(operationsConfigImpls...),
	)

}
