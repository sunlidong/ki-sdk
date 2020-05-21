package action

import (
	"log"
)

/**
@ 对外 函数 || SDK 单个 实例化
*/
func initSDK() error {
	var err error

	//	01.	sdk
	SDK := initsdk()

	//	02.	init
	err = SDK.Initialize()
	if err != nil {
		return err
	}
	log.Println("01 || 客戶端初始化成功")
	//	03.	msg
	err = SDK.CreateresMgmtClient()
	if err != nil {
		return err
	}
	log.Println("02 || 资源客戶端初始化成功")

	//	04.	cli
	err = SDK.CreateChannelCli()
	if err != nil {
		return err
	}
	log.Println("03 || 通道客戶端初始化成功")

	//	05.	msp
	err = SDK.CreateMspClient()
	if err != nil {
		return err
	}
	log.Println("03 || 证书客戶端初始化成功")

	return nil
}

// get  sdk
func GetInitSdk() error {
	return initSDK()
}

func Cpeer() []string {
	return cpeer()
}

/**
@ 数据上链
@ 2019年10月21日09:55:34
@
*/
func UploadAsset(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	return uploadAsset(clientName, channelName, chaincodeName, peer, args)
}

// SQL  --init
func SqlByInit() {
	sqlByInit()
}

//
func SqlByConact() (rep interface{}, err error) {
	return sqlByConact()
}

//	大屏@ 智能合约调用信息表 @ 2
func SqlByLarge() (rep interface{}, err error) {
	return sqlByLarge()
}

//	大屏@ 共识节点信息表 @ 3
func SqlByServer() (rep interface{}, err error) {

	// TODO  更新 服务器  性能信息
	UpdateByIp()

	//

	return sqlByServer()
}

//	大屏@ 节点信息统计表 @ 4
func SqlBySvg() (rep interface{}, err error) {
	return sqlBySvg()
}

//	大屏@ 区块信息表 @ 5
func SqlByBlock() (rep interface{}, err error) {
	return sqlByBlock()
}

//	大屏@ 节点信息交易列表 @ 6
func SqlByDeal() (rep interface{}, err error) {
	return sqlByDeal()
}

//	大屏@ 资产上链信息表 @ 7
func SqlByAsset() (rep interface{}, err error) {
	return sqlByAsset()
}

// SQL  初始化  SQL 全局 DB

func SqlByInitByDB() {
	sqlByInitByDB()
}

// SQL --查询 block
func SqlQueryByBlock(blockHeight string) (rep interface{}, err error) {
	return sqlQueryByBlock(blockHeight)
}

// SQL  -- 智能合约 插入 调用次数
func SqlByLargeForInserData(cName string, chainName string, funcName string, org string, userName string) {
	sqlByLargeForInserData(cName, chainName, funcName, org, userName)
}

// SQL --- 区块信息表  插入
func SqlByBlockForInserData(height string, hash string, channel string, chaincode string, orgName string, userName string) {
	sqlByBlockForInserData(height, hash, channel, chaincode, orgName, userName)
}

// SQL --- 合约调用次数  插入
func SqlBySvgForInsertData(txId string, orgName string, userName string) {
	sqlBySvgForInsertData(txId, orgName, userName)
}

// SQL --- 节点交易信息表
func SqlByDbDealForInsertData(txid string, height string, userID string, chainCode string, orgName string, userName string) {
	sqlByDbDealForInsertData(txid, height, userID, chainCode, orgName, userName)
}

// SQL --- 资产信息表 插入
func SqlByDbAssetForInsertData(assetNo string, conType string, chainCode string, upTime string, assetType string, orgName string, userName string, txid string) {
	sqlByDbAssetForInsertData(assetNo, conType, chainCode, upTime, assetType, orgName, userName, txid)
}

// SQL -- 更新 服务器配置信息
func UpdateByIp() {
	updateByIp()
}
func UpdateByIpBack() (cpu string, mem string) {
	return updateByIpBack()
}

//	大屏@ 根据区块查询关联10个区块 @ 9 BlockData
func SqlQueryBlockByHeight(blockHeight string) (rep interface{}, err error) {
	return sqlQueryBlockByHeight(blockHeight)
}

// 	获取 Peer 背书节点

func Peer() (arr []string) {
	return peer()
}

//	大屏@ 节点信息交易列表 @ 6 2019年12月17日13:47:31 统计
func SqlBySvgForBig() (rep interface{}, err error) {
	return sqlBySvgForBig()
}

//	轮询器
func Poller() {
	poller()
}

// SQL --- 资产信息表 插入
func GetBlockTxIdForBox(channelName string, cliName string, peerName string, orgName string, userName string, txid string) (config string, data []string, err error) {
	return getBlockTxIdForBox(channelName, cliName, peerName, orgName, userName, txid)
}

//	大屏@ 节点信息交易列表 @ 6 2019年12月17日13:47:31 统计
func SqlBySvgForGeography() (rep interface{}, err error) {
	return sqlBySvgForGeography()
}

//
//	大屏@ 资产上链信息表 @ 7
func SqlByAssetNext(no string, t string) (rep interface{}, err error) {
	return sqlByAssetNext(no, t)
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func SqlByDealList(id string) (rep interface{}, err error) {
	return sqlByDealList(id)
}

//	 根据交易ID  查询 用户信息以及  事件
func SqlByQueryOntransactionID(txid string) (things string, userId string, err error) {
	return sqlByQueryOntransactionID(txid)
}

// sqlByAssetSum
func SqlByAssetSum() (rep interface{}, err error) {
	return sqlByAssetSum()
}

//
//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func SqlByDealGeography(id string) (rep interface{}, err error) {
	return sqlByDealGeography(id)
}
