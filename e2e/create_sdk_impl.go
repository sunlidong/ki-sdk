package e2e

import (
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	ConfigFile  = "./config/org1_peer0_admin.yaml"
	ChannelID   = "bookchannel"
	OrgName     = "Org1"
	ChainCodeID = "bookstorechain"
	OrgAdmin    = "Admin"
	UserName    = "Admin"
	Version     = "1.0"
	OrdererID   = "orderer1.bookstore.com:7050"
)

var APP *fabsdk.FabricSDK

func (swp APP) CreateresMgmtClient() error {
	// 01. 创建资源管理客户端上下文
	resourceManagerClientContext :=
		swp.Context(fabsdk.WithUser(OrgAdmin),
			fabsdk.WithOrg(OrgName))
	// 02. 创建资源管理客户端
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	// 03. 验错
	if err != nil {
		return fmt.Errorf("创建资源管理客户端失败:%v", err)
	}

	// 05. 返回
	return nil
}

/**
@ 初始化 channel cli
@
*/
func (swp APP) CreateChannelCli() error {

	// 01. 封装数据Channle cli
	clientContext := swp.SDK.ChannelContext(
		ChannelID,
		fabsdk.WithUser(UserName))

	// 02. 创建Channle cli
	channelCli, err := channel.New(clientContext)

	// 03. 验错
	if err != nil {
		return fmt.Errorf("创建通道管理客户端失败:%v", err)
	}

	// // 04. 参数转换至结构体对象 swp
	// swp.Client = channelCli
	// // 05. 返回
	return nil
}

/**
@ 初始化 msp cli
@
*/
func (swp APP) CreateMspClient() error {

	// 01. 创建资源管理客户端上下文
	clientCTX := swp.Context(
		fabsdk.WithUser(OrgAdmin),
		fabsdk.WithOrg(OrgName),
	)

	// 02. 创建资源实例
	c, err := msp.New(clientCTX)

	if err != nil {
		return fmt.Errorf("创建msp管理客户端失败:%v", err)
	}

	if c == nil {
		return fmt.Errorf("创建msp管理客户端为空:%v", c)
	}

	swp.MspClient = c

	return nil
}

func Init_one_sdk() {
	//	02.	init
	err = APP.CreateresMgmtClient()
	if err != nil {
		log.Println("2err:", err)
		return err
	}
	log.Println("02 || 资源客戶端初始化成功")

	//	04.	cli
	err = APP.CreateChannelCli()
	if err != nil {
		log.Println("3err:", err)
		return err
	}
	log.Println("03 || 通道客戶端初始化成功")

	//	05.	msp
	err = APP.CreateMspClient()
	if err != nil {
		log.Println("4err:", err)
		return err
	}
	log.Println("04 || 证书客戶端初始化成功")

	return nil

}
