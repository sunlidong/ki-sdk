/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package event

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"ki-sdk/test/integration"
	"ki-sdk/test/integration/util/runner"
)

const (
	org1Name = "Org1"
	org1User = "User1"
)

var mainSDK *fabsdk.FabricSDK
var mainTestSetup *integration.BaseSetupImpl
var mainChaincodeID string

func TestMain(m *testing.M) {
	r := runner.NewWithExampleCC()
	r.Initialize()
	mainSDK = r.SDK()
	mainTestSetup = r.TestSetup()
	mainChaincodeID = r.ExampleChaincodeID()

	r.Run(m)
}
