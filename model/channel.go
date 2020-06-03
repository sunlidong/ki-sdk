package model

import (
	"fmt"
	"log"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	sdkConfig "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/pkg/errors"
)

// Model

func CreateChannel(path string, channelName string, org string, user string, orderers []string) error {
	log.Println("model =>开始创建通道 ")
	// 创建SDK
	sdk, err := fabsdk.New(sdkConfig.FromFile(path))
	if err != nil {
		fmt.Println("创建FabricSDK失败，错误为", err)
		return errors.Errorf("创建FabricSDK失败，错误为：%s", err)
	}

	//创建MSP客户端
	mspClient, err := mspclient.New(sdk.Context(), mspclient.WithOrg(org))

	if err != nil {
		return errors.Errorf("创建MSP客户端失败，错误为：%s", err)
	}

	resourceManagerClientContext := sdk.Context(
		fabsdk.WithUser(user),
		fabsdk.WithOrg(org),
	)
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.Errorf("创建组织：%s 的资源管理客户端失败，错误为：%s")
	}
	adminIdentity, err := mspClient.GetSigningIdentity(user)
	if err != nil {
		return errors.Errorf("获取签名身份失败，错误为：%s", err)
	}

	// channel tx path
	channelTxFile := fmt.Sprintf("./config/%s.tx", channelName)

	req := resmgmt.SaveChannelRequest{
		ChannelID:         channelName,
		ChannelConfigPath: channelTxFile,
		SigningIdentities: []msp.SigningIdentity{adminIdentity},
	}

	orderersOpt := []resmgmt.RequestOption{}

	for _, v := range orderers {
		orderersOpt = append(orderersOpt, resmgmt.WithOrdererEndpoint(v))
	}

	txID, err := resMgmtClient.SaveChannel(req, orderersOpt...)
	if err != nil || txID.TransactionID == "" {
		return errors.Errorf("创建channel失败，错误为：%s", err)
	}
	fmt.Printf("创建名为%s的channel成功\n", channelName)
	return nil
}

func JoinChannel(path string, channelName string, org string, user string, orderers []string) error {

	orderersOpt := []resmgmt.RequestOption{}
	// 创建SDK
	sdk, err := fabsdk.New(sdkConfig.FromFile(path))
	if err != nil {
		fmt.Println("创建FabricSDK失败，错误为", err)
		return errors.Errorf("创建FabricSDK失败，错误为：%s", err)
	}
	resourceManagerClientContext := sdk.Context(
		fabsdk.WithUser(user),
		fabsdk.WithOrg(org),
	)
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.Errorf("创建组织：%s 的资源管理客户端失败，错误为：%s")
	}

	for _, v := range orderers {
		orderersOpt = append(orderersOpt, resmgmt.WithOrdererEndpoint(v))
	}

	err = resMgmtClient.JoinChannel(channelName, orderersOpt...)

	if err != nil {
		return errors.Errorf("创建组织：%s 的资源管理客户端失败，错误为：%s")
	}

	fmt.Printf("加入mychannel成功\n")
	return nil
}
