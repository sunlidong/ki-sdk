package model

import (
	"fmt"
	"ki-sdk/vendor/github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"ki-sdk/vendor/github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"ki-sdk/vendor/github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"ki-sdk/vendor/github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
)

// Action
func (setup *SDK) Initialize() error {
	// 01. 读取配置文件config.yaml
	configProvider := config.FromFile(
		setup.ConfigFile,
	)
	// 02. 创建sdk对象
	sdk, err := fabsdk.New(configProvider)
	// 03. 验错
	if err != nil {
		return fmt.Errorf("SDK实例化失败:%v", err)
	}
	// 04. 参数转换至结构体对象 setup
	setup.SDK = sdk
	// 05. 返回
	return nil
}

/**
@ 初始化 msg
@
*/
func (setup *SDK) CreateresMgmtClient() error {
	// 01. 创建资源管理客户端上下文
	resourceManagerClientContext :=
		setup.SDK.Context(fabsdk.WithUser(setup.OrgAdmin),
			fabsdk.WithOrg(setup.OrgName))
	// 02. 创建资源管理客户端
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	// 03. 验错
	if err != nil {
		return fmt.Errorf("创建资源管理客户端失败:%v", err)
	}
	// 04. 参数转换至结构体对象 setup
	setup.resmgmt = resMgmtClient
	// 05. 返回
	return nil
}

/**
@ 初始化 channel cli
@
*/
func (setup *SDK) CreateChannelCli() error {

	// 01. 封装数据Channle cli
	clientContext := setup.SDK.ChannelContext(
		setup.ChannelID,
		fabsdk.WithUser(setup.UserName))

	// 02. 创建Channle cli
	channelCli, err := channel.New(clientContext)

	// 03. 验错
	if err != nil {
		return fmt.Errorf("创建通道管理客户端失败:%v", err)
	}

	// 04. 参数转换至结构体对象 setup
	setup.client = channelCli
	// 05. 返回
	return nil
}

/**
@ 初始化 msp cli
@
*/
func (setup *SDK) CreateMspClient() error {

	// 01. 创建资源管理客户端上下文
	clientCTX := setup.SDK.Context(
		fabsdk.WithUser(setup.OrgAdmin),
		fabsdk.WithOrg(setup.OrgName),
	)

	// 02. 创建资源实例
	c, err := msp.New(clientCTX)

	if err != nil {
		return fmt.Errorf("创建msp管理客户端失败:%v", err)
	}

	if c == nil {
		return fmt.Errorf("创建msp管理客户端为空:%v", c)
	}

	setup.MspClient = c

	return nil
}
