package model

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

/**
@
@ SDK结构体
@ NO 001
*/
type SDK struct {
	OrdererID        string
	ChaincodeGoPath  string
	GoPath           string
	ChainCodeVersion string
	ConfigFile       string //sdk配置文件所在路径
	ChannelID        string //应用通道名称
	ChannelConfig    string //应用通道交易配置文件所在路径
	OrgAdmin         string // 组织管理员名称
	OrgName          string //组织名称
	Initialized      bool   //是否初始化
	Orgmsp           []string
	Args             string
	UserName         string
	ChainCodeID      string
	MspClient        *msp.Client
	SDK              *fabsdk.FabricSDK //SDK实例
	Client           *channel.Client   //cli
	Resmgmt          *resmgmt.Client   //admian
}

/**
@ app 结构体
@ SDK结构体
@ NO 002
*/
type Application struct {
	gin *gin.Context
	SDK *SDK
}

//Init
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
