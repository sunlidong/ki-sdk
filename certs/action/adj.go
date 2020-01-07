package action

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/pkg/errors"
	"log"
)

// 方法

// 初始化 SDK
func (setup *Setup) Initialize() error {
	// 01. 读取配置文件config.yaml
	configProvider := config.FromFile(setup.ConfigFile)
	// 02. 创建sdk对象
	sdkstring, err := fabsdk.New(configProvider)
	// 03. 验错
	if err != nil {
		return fmt.Errorf("SDK实例化失败:%v", err)
	}
	// 04. 参数转换至结构体对象 setup
	setup.SDK = sdkstring

	// 05. 返回
	return nil
}

//	创建 Msg
func (setup *Setup) CreateMsgClient() error {
	// 01. 创建资源管理客户端上下文
	resourceManagerClientContext := setup.SDK.Context(fabsdk.WithUser(setup.OrgAdmin), fabsdk.WithOrg(setup.OrgName))

	// 02. 创建资源管理客户端
	MsgClient, err := resmgmt.New(resourceManagerClientContext)
	// 03. 验错
	if err != nil {
		return errors.WithMessage(err, "创建资源管理客户端失败!")
	}

	// 04. 参数转换至结构体对象 setup
	setup.resmgmt = MsgClient

	// 05. 返回
	return nil
}

//	创建 Channel cli
func (setup *Setup) CreateChannelCli() error {
	// 01. 封装数据Channle cli
	clientContext := setup.SDK.ChannelContext(
		setup.ChannelID,
		fabsdk.WithUser(setup.UserName))

	// 02. 创建Channle cli
	channelCli, err := channel.New(clientContext)
	// 03. 验错
	if err != nil {
		return errors.WithMessage(err, "Channel Cli 创建失败！")
	}
	// 04. 参数转换至结构体对象 setup
	setup.client = channelCli
	log.Println("Channel Cli 创建成功")
	// 05. 返回
	return nil
}

//	创建 msp
func (setup *Setup) CreateMspClient() error {
	// 01. 创建资源管理客户端上下文

	clientCTX := setup.SDK.Context(
		fabsdk.WithUser(setup.OrgAdmin),
		fabsdk.WithOrg(setup.OrgName),
	)

	c, err := msp.New(clientCTX)

	if err != nil {
		return errors.WithMessage(err, "msp.client is error")
	}
	if c != nil {
		fmt.Println("msp client created is ok")
	}
	setup.MspClient = c
	return nil
}
