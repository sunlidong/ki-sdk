package action

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//	数据集合

//	全局SDK
var App Application

// SDK 结构体
type Setup struct {
	OrdererID       string
	ChaincodeGoPath string
	GoPath          string
	Version         string
	ConfigFile      string //sdk配置文件所在路径
	ChannelID       string //应用通道名称
	ChannelConfig   string //应用通道交易配置文件所在路径
	OrgAdmin        string // 组织管理员名称
	OrgName         string //组织名称
	Initialized     bool   //是否初始化
	Orgmsp          []string
	Args            string
	UserName        string
	ChainCodeID     string
	MspClient       *msp.Client
	SDK             *fabsdk.FabricSDK //SDK实例
	client          *channel.Client   //cli
	resmgmt         *resmgmt.Client   //admian
}

//	APP
type Application struct {
	gin   *gin.Context
	Setup *Setup
}

//
const (
	Conf = "/root/go/src/dlchain/sdk/conf/"
	//Conf                = "/root/dlchain/sdk/conf/"
	ConfChannelConfig   = "/root/ABS_chaincode/fabric-asset/network/config/"
	ConfChaincodeGoPath = "/root/ABS_chaincode/goChainCode/"
	//
	ConfFile      = "config.yaml"
	ConfChannelID = "assetpublish"
	//ConfChannelID           = "threeorgschannel"
	ConfChannelfile = "channel.tx"
	//ConfChainCodeID         = "AssetToChain_realty14"//
	ConfChainCodeID         = "AssetToChain_realty16" //房地产
	ConfChainCodeIDTest     = "AssetToChain_realty34" //统一链码
	ConfChaincodeGoPathName = "underlying"
	ConfOrgAdmin            = "Admin"
	ConfOrgName             = "org1"
	ConfUserName            = "Admin"
	ConfVersion             = "1.0"
	ConfOrdererID           = "orderer0.cmbfae.com"
	Peerfactoring           = "peer0.org1.dinglian.com"
	Peerproduct             = "peer0.org2.dinglian.com"
	Peerorg3                = "peer0.org3.dinglian.com"
	CAfactoring             = "ca-org1-msp.dinglian.com"
	CAfactoring_ca2         = "ca-org2-msp.dinglian.com"
	CAfactoring_ca3         = "ca-org3-msp.dinglian.com"
	ConfArgs                = "init"
	Port                    = ":10081"
)

//Init
const (
	Init_ConfigFile      = Conf + ConfFile
	Init_ChannelID       = ConfChannelID
	Init_ChannelConfig   = ConfChannelConfig + ConfChannelfile
	Init_ChainCodeID     = ConfChainCodeID
	Init_ChaincodeGoPath = ConfChaincodeGoPath + ConfChaincodeGoPathName
	Init_OrgAdmin        = ConfOrgAdmin
	Init_OrgName         = ConfOrgName
	Init_UserName        = ConfUserName
	Init_Version         = ConfVersion
	Init_GoPath          = ConfChaincodeGoPath + ConfChaincodeGoPathName
	Init_OrdererID       = ConfOrdererID
	Init_Args            = ConfArgs
)

//  SQL  -- mysql
const (
	SQL_DB                 = "SQL_DB"
	SQL_DbData             = "mysql"
	SQL_Login              = "root"
	SQL_PassWord           = "123456"
	SQL_Http               = "tcp"
	SQL_Ip                 = "(127.0.0.1:3306)"
	SQL_DbName             = "drc"
	SQL_Charset            = "utf8mb4"
	SQL_ParseTime          = "true"
	SQL_SetMaxIdleConns    = "300"
	SQL_SetMaxOpenConns    = "500"
	SQL_SetConnMaxLifetime = "30"
)
