/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package practitioners

import (
	"log"

	"ki-sdk/configless"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	channelID      = "mychannel"
	orgName        = "Org1"
	orgAdmin       = "Admin"
	ordererOrgName = "OrdererOrg"
)

var (
	ccID = "example_cc_e2e"
)

// Run enables testing an end-to-end scenario against the supplied SDK options
func Run(configOpt core.ConfigProvider, sdkOpts ...fabsdk.Option) {
	setupAndRun(true, configOpt, sdkOpts...)
}

// RunWithoutSetup will execute the same way as Run but without creating a new channel and registering a new CC
func RunWithoutSetup(configOpt core.ConfigProvider, sdkOpts ...fabsdk.Option) {
	setupAndRun(false, configOpt, sdkOpts...)
}

// setupAndRun enables testing an end-to-end scenario against the supplied SDK options
// the createChannel flag will be used to either create a channel and the example CC or not(ie run the tests with existing ch and CC)
func setupAndRun(createChannel bool, configOpt core.ConfigProvider, sdkOpts ...fabsdk.Option) {

	if configless.IsLocal() {
		//If it is a local test then add entity mapping to config backend to parse URLs
		configOpt = configless.AddLocalEntityMapping(configOpt)
	}

	sdk, err := fabsdk.New(configOpt, sdkOpts...)
	if err != nil {
		log.Printf("Failed to create new SDK: %s", err)
	}
	// defer sdk.Close()
	log.Println("init------------------------------------")

}

//  peer

func SetupAndRuning(configOpt core.ConfigProvider, sdkOpts ...fabsdk.Option) {

	configOpt = configless.AddLocalEntityMapping(configOpt)

	sdk, err := fabsdk.New(configOpt, sdkOpts...)
	if err != nil {
		log.Printf("----------------------------Failed to create new SDK: %s", err)
	}
	// defer sdk.Close()
	log.Println("init------------------------------------", sdk)

	// if createChannel {
	// 	createChannelAndCC(t, sdk)
	// }
	App = Application{
		SDK: sdk,
	}
	// if createChannel {
	// 	createChannelAndCC(t, sdk)
	// }
	Init_one_sdk()

}
