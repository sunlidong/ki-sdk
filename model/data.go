package action

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/jinzhu/gorm"
)

/**
@ FabricSetup implementation
@ SDK结构体
@ NO 001
*/
type SDK struct {
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

/**
@ app 结构体
@ SDK结构体
@ NO 002
*/
type Application struct {
	gin *gin.Context
	SDK *SDK
}

//channel
const (
	SIGSEGV = 11
	SIGPIPE = 11
	SIGALRM = 12
	SIGTERM = 11
)

// TODO 全局变量
var (
	Current_ChannelName string = "assetpublish" // 当前通道
	Current_OrgName     string = "Org"          // 当前组织
	Current_UserName    string = "user"         // 当前用户
)

//Base Config

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
	TestChain               = "chaincodeby04"         // 测试链码
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

type ChainBlock struct {
	Height       int64 `json:",string"`
	Hash         string
	TimeStamp    string              `json:",omitempty"`
	Transactions []*ChainTransaction `json:"-"`
	TxEvents     []*ChainTxEvents    `json:"-"`
}

type ChainTransaction struct {
	Height                             string `json:",string"`
	Timestamp                          string
	TxID, Chaincode, Method, ChannelId string
	CreatedFlag                        bool
	TxArgs                             [][]byte `json:"-"`
}

type ChainTxEvents struct {
	TxID, Chaincode, Name string
	Status                int
	Payload               []byte `json:"-"`
}
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

const (
	Fabric_ChaincodeNum  = "1"
	Fabric_ChannelList   = "assetpublish,assetpublish"
	Channel_assetpublish = "assetpublish"
)

/**
编码标准列表:
	CT01GA01DT00394853

标准释义：
	A段：CT		：所属哪个平台 2位
	B段：01		：所属哪个公司 2位
	C段：GA		：所属哪个部门 2位
	D段：01		：所属哪种类别 2位
	E段：DT		：所属权限级别 2位
	F段：00		： 切割标识 2位
	G段：394853 ：操作人标识符 6位
*/

//  A 段
const (
	A_ESB_DRC          = "AA" //  A 段：DRC平台
	A_ESB_XIAOMI       = "AA" //  A 段：小米供应链平台
	A_ESB_NANJINGXINQU = "AA" //  A 段：扬子平台
)

//  B 段
const (
	B_CORP_DRC          = "AA" //  B 段：DRC
	B_CORP_XIAOMI       = "AB" //  B 段：小米
	B_CORP_NANJINGXINQU = "AC" //  B 段：扬子科技
)

//  C 段
const (
	C_BRANCH_DRC          = "AA" //  C 段：DRC
	C_BRANCH_XIAOMI       = "AB" //  C 段：小米
	C_BRANCH_NANJINGXINQU = "AC" // C 段：扬子科技
)

//  D 段 :::此为标识sdk 数据类型主键
const (
	D_TYPEC_DRC          = "AA" //  D 段：DRC
	D_TYPEC_XIAOMI       = "AB" //  D 段：小米
	D_TYPEC_NANJINGXINQU = "AC" //  D 段：扬子科技
)

//  E 段
const (
	E_LEVEL_DRC    = "AA" //  E 段：DRC
	E_LEVEL_XIAOMI = "AA" //  E 段：小米金融
	E_LEVEL_NJYZ   = "AA" //  E 段：南京扬子国投
	E_LEVEL_AL     = "AA" //  E 段：test
)

//  F 段
const (
	F_INCISE_DRC    = "00" //  E 段：DRC
	F_INCISE_XIAOMI = "11" //  E 段：小米金融
	F_INCISE_NJYZ   = "22" //  E 段：南京扬子
	F_INCISE_AL     = "33" //  E 段：test
)

//  G 段
const (
	G_USER_DRC    = "002123" //  E 段：DRC
	G_USER_XIAOMI = "002124" //  E 段：小米金融
	G_USER_NJYZ   = "002125" //  E 段：南京扬子
	G_USER_AL     = "002145" //  E 段：test
)

//  子集List 段
const (
	LI_DL_DRC    = "dl"     //  E 段：DRC dl
	lI_DL_XIAOMI = "002124" //  E 段：小米金融
	LI_DL_NJYZ   = "002125" //  E 段：南京扬子
	LI_DL_AL     = "002145" //  E 段：test
)

//controller

const (
	Chaincode_Inc_Name = "IncrementChaincode"
)

//
const (
	Ca_Org1_path string = "com.dinglian.org1."
	//Ca_Org2_path string = "com.dinglian.org2."
	//Ca_Org2_path string = "com.dinglian.org2."
)

//	所有节点
const (
	Org1_peer0 string = "peer0.org1.dinglian.com,"
	Org1_peer1 string = "peer1.org1.dinglian.com,"
	Org2_peer0 string = "peer0.org2.dinglian.com,"
	Org2_peer1 string = "peer1.org2.dinglian.com,"
)

//
const (
	//
	Web_PeerSum = Org1_peer0 + Org1_peer1 + Org2_peer0 + Org2_peer1
)

//
type BlockWordback struct {
	Height         uint64 `json:"height"`
	BlockHash      string `json:"blockHash"`
	PreHash        string `json:"preHash"`
	TransactionNum string `json:"transactionNum"`
	BlockTime      string `json:"blockTime"`
	ChannelID      string `json:"channelId"`
}

// 云图配置结构体
type YunConfig struct {
	ConfigHeight         string `json:"configHeight"`
	ConfigOrgNum         string `json:"configOrgNum"`
	ConfigNodeNum        string `json:"configNodeNum"`
	ConfigTransactionNum string `json:"configTransactionNum"`
	ConfigChannelNum     string `json:"configChannelNum"`
	ConfigChaincodeNum   string `json:"configChaincodeNum"`
	ConfigTime           string `json:"configTime"`
}

type BlockTransactionBack struct {
	Height        int64  `json:"height"`
	BlockHash     string `json:"blockHash"`
	PreHash       string `json:"preHash"`
	TransactionID string `json:"transactionID"`
	BlockTime     string `json:"blockTime"`
	AssetName     string `json:"assetName"`
	AssetUUID     string `json:"assetUuid"`
}

//  CA 数据索引
const (
	CaMsp01 = "ca-org1-msp.dinglian.com"
	CaMsp02 = "ca-org2-msp.dinglian.com"
	CaMsp03 = "ca-org3-msp.dinglian.com"
	CaMsp04 = "ca-org4-msp.dinglian.com"
	CaMsp05 = "ca-org5-msp.dinglian.com"
	CaMsp06 = "ca-org6-msp.dinglian.com"
	CaMsp07 = "ca-org7-msp.dinglian.com"
	CaMsp08 = "ca-org8-msp.dinglian.com"
	CaMsp09 = "ca-org9-msp.dinglian.com"
)

// CA NAME
const (
	Ca01 = "CA01"
	Ca02 = "CA02"
	Ca03 = "CA03"
	Ca04 = "CA04"
	Ca05 = "CA05"
	Ca06 = "CA06"
	Ca07 = "CA07"
	Ca08 = "CA08"
	Ca09 = "CA09"
)

// CA msp
const (
	CaMspPath01 = "com.dinglian.org1."
	CaMspPath02 = "com.dinglian.org1."
	CaMspPath03 = "com.dinglian.org1."
	CaMspPath04 = "com.dinglian.org1."
	CaMspPath05 = "com.dinglian.org1."
	CaMspPath06 = "com.dinglian.org1."
	CaMspPath07 = "com.dinglian.org1."
	CaMspPath08 = "com.dinglian.org1."
	CaMspPath09 = "com.dinglian.org1."
)

//------------------------------------------------------------------------------------- mysql  数据结构展示

//  数据库实例对象
type SqlDB struct {
	DB                 *gorm.DB // DB 实例
	DbData             string   // 数据库名称
	Login              string   // 登录名称
	PassWord           string   // 密码
	Http               string   // http
	Ip                 string   // IP
	DbName             string   // 数据库名称
	Charset            string   // car set
	ParseTime          string   // time
	SetMaxIdleConns    string   //空闲
	SetMaxOpenConns    string   //打开限制
	SetConnMaxLifetime string   // 超时
}

//  -- 数据库基础表
type DbBase struct {
	//
	No             int32  `gorm:"AUTO_INCREMENT" json:"no"` // 自增 // 编号
	CreateOrg      string `json:"-"`                        //创建组织
	CreateUserName string `json:"-"`                        //创建人
	Type           string `json:"-"`                        //类型
	//DBType         string `gorm:"default:dinglian"  json:"-"` // 指定默认值 //库表类型
}

//  SQL -- 合约处理展示表 01
type DbContract struct {
	//
	gorm.Model   `json:"-"` // 主键
	ContractName string     `json:"contName"` // 合约名称
	ContractType string     `json:"contType"` // 合约类型
	//
	No             int32  `gorm:"AUTO_INCREMENT" json:"no"` // 自增 // 编号
	CreateOrg      string `json:"-"`                        //创建组织
	CreateUserName string `json:"-"`                        //创建人
	Type           string `json:"-"`                        //类型
	//DBType         string `gorm:"default:dinglian"  json:"-"` // 指定默认值 //库表类型
}

//  SQL -- 智能合约调用信息表 02
type DbLarge struct {
	//
	gorm.Model        `json:"-"` // 主键
	ChannelName       string     `json:"channelName"`       // 调用通道
	ChaincodeName     string     `json:"chainCodeName"`     // 调用链码
	ChaincodeFuncName string     `json:"chainCodeFuncName"` // 调用函数
	Num               int64      `json:"num"`               // 调用次数

	// 基础字段
	No             int    `gorm:"AUTO_INCREMENT"json:"no"` // 自增 编号
	CreateOrg      string `json:"-"`                       //创建组织
	CreateUserName string `json:"-"`                       //创建人
	Type           string `json:"-"`                       //类型
	//DBType         string `gorm:"default:dinglian"  json:"-"` // 指定默认值 //库表类型
}

//  SQL -- 共识节点信息表 03
type DbIp struct {
	gorm.Model  `json:"-"` // 主键
	NodeName    string     `json:"nodeName"`    // 节点名称
	NodeState   string     `json:"nodeState"`   // 节点状态
	NodeAddress string     `json:"nodeAddress"` // 服务器地址
	NodeCPU     string     `json:"nodeCpu"`     // CPU使用率
	NodeRAM     string     `json:"nodeRam"`     // 内存使用率
	// 基础字段
	No             int    `gorm:"AUTO_INCREMENT"json:"no"` // 自增 编号
	CreateOrg      string `json:"-"`                       //创建组织
	CreateUserName string `json:"-"`                       //创建人
	Type           string `json:"-"`                       //类型
	//DBType         string `gorm:"default:dinglian"  json:"-"` // 指定默认值 //库表类型
}

//  SQL -- 节点信息表 04 计算统计
type DbSvg struct {
	gorm.Model            `json:"-"` // 主键
	ConNum                string     `json:"conNum"` // 合约调用次数
	APINum                string     `json:"apiNum"` // API调用次数
	TPS                   string     `json:"-"`      // 峰值TPS/秒
	AverageProcessingTime string     `json:"-"`      // 平均处理时长/毫秒
	FormatTime            string     `json:"time"`   // 平均处理时长/毫秒
	// 基础字段
	No             int    `gorm:"AUTO_INCREMENT"json:"no"` // 自增 编号
	CreateOrg      string `json:"-"`                       //创建组织
	CreateUserName string `json:"-"`                       //创建人
	Type           string `json:"-"`                       //类型
	//DBType         string `gorm:"default:dinglian"  json:"-"` // 指定默认值 //库表类型
}

//  SQL -- 区块信息表 05
type DbBlock struct {
	gorm.Model       `json:"-"` // 主键
	BlockHeight      string     `json:"height"`    // 区块高度
	BlockHash        string     `json:"hash"`      // 区块hash
	BlockByChannel   string     `json:"channel"`   // 区块所属通道
	BlockByChainCode string     `json:"chainCode"` // 区块所属链码
	// 基础字段
	No             int    `gorm:"AUTO_INCREMENT"json:"no"` // 自增 编号
	CreateOrg      string `json:"-"`                       //创建组织
	CreateUserName string `json:"-"`                       //创建人
	Type           string `json:"-"`                       //类型
	//DBType         string `gorm:"default:dinglian"  json:"-"` // 指定默认值 //库表类型
}

//  SQL -- 节点信息交易列表 06
type DbDeal struct {
	gorm.Model       `json:"-"` // 主键
	BlockHeight      string     `json:"height"`        // 区块高度
	BlockHash        string     `json:"hash"`          // 区块hash
	BlockByChannel   string     `json:"channel"`       // 区块所属通道
	BlockByChainCode string     `json:"chainCode"`     // 区块所属链码
	BlockByTXID      string     `json:"txId"`          // 交易Hash
	UserTxt          string     `json:"userId"`        // 用户信息
	ChainCodeName    string     `json:"chainCodeName"` // 链码名称
	TXIDTime         string     `json:"txIdTime"`      // 交易时间
	TXIDType         string     `json:"txIdType"`      // 交易类型
	BlockByNode      string     `json:"node"`          // 所属节点
	BlockByOrg       string     `json:"org"`           // 所属组织
	// 基础字段
	No             int    `gorm:"AUTO_INCREMENT"json:"no"` // 自增 编号
	CreateOrg      string `json:"-"`                       //创建组织
	CreateUserName string `json:"-"`                       //创建人
	Type           string `json:"-"`                       //类型
	//DBType         string `gorm:"default:dinglian"  json:"-"` // 指定默认值 //库表类型
}

//  SQL -- 资产上链信息表 07
type DbAsset struct {
	gorm.Model      `json:"-"` // 主键
	AssetNo         string     `gorm:"AUTO_INCREMENT" json:"assetNo"` // 资产编号
	ConType         string     `json:"conType"`                       // 合约类型
	DataUpChainCode string     `json:"name"`                          // 数据上链
	UpTime          string     `json:"upTime"`                        // 上链时间
	AssetType       string     `json:"assetType"`                     // 资产类型
	// 基础字段
	No             int    `gorm:"AUTO_INCREMENT"json:"no"` // 自增 编号
	CreateOrg      string `json:"-"`                       //创建组织
	CreateUserName string `json:"-"`                       //创建人
	Type           string `json:"-"`                       //类型
	//DBType         string `gorm:"default:dinglian"  json:"-"` // 指定默认值 //库表类型
}

//  区块链大屏  --- 节点统计返回结果

type SqlJie struct {
	//
	TPS  string      `json:"tps"`
	API  string      `json:"api"`
	Data interface{} `json:"data"`
}

//  区块链大屏  ---  资产信息展示
type SqlAssets struct {
	//
	Sum  string      `json:"sum"`
	Data interface{} `json:"data"`
}

type OrgWeb struct {
	Affiliations string `json:"Affiliations"`
	CAName       string `json:"CAName"`
	Identities   string `json:"Identities"`
	Name         string `json:"Name"`
}

//  地理信息
type WebGeography struct {
	No        string `json:"no"`
	Name      string `json:"name"`
	Longitude string `json:"lon"`
	Latitude  string `json:"lat"`
	Address   string `json:"address"`
	Title     string `json:"title"`
	Text      string `json:"text"`
}

// 	共识节点 标识
const (
	echoByEcologicalNode         string = "eco" // 生态节点
	echoCommonNode               string = "com" // 共识节点
	echoByDistributedStorageNode string = "dis" // 分布式存储节点
)

//	生态节点
type EcologicalNode struct {
	NodeName          string `json:"nodeName"`          // 节点名称
	TradeType         string `json:"tradeType"`         // 行业类型
	Level             string `json:"level"`             // 企业类别
	TradeProfessional string `json:"tradeProfessional"` // 参与业务
	Time              string `json:"time"`              // 加入时间
}

//	共识节点
type CommonNode struct {
	NodeId    string `json:"node"`      // 节点标识
	NodeName  string `json:"nodeName"`  // 节点名称
	NodeType  string `json:"nodeType"`  // 节点类型
	NodeState string `json:"nodeState"` // 节点状态
	Ip        string `json:"ip"`        // ip 地址
	Cpu       string `json:"cpu"`       // CPU使用率
	Mem       string `json:"mem"`       // 内存使用率
	Time      string `json:"time"`      // 时间
}

//	分布式储存节点
type DistributedStorageNode struct {
	NodeName  string `json:"nodeName"`  // 节点名称
	StoreType string `json:"storeType"` // 储存类型
	Address   string `json:"address"`   // 地点
	Percent   string `json:"percent"`   // 容量使用百分比
	Time      string `json:"time"`      // 加入时间
}
