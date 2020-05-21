package api

import (
	"ki-sdk/model/action"
)

/**
api ||  SDK 初始化
*/
func GetSDK() error {
	return action.GetInitSdk()
}

// 数据上链
func UploadAsset(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	return action.UploadAsset(clientName, channelName, chaincodeName, peer, args)
}

// 数据上链
func UploadAssetTest(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	return action.UploadAssetTest(clientName, channelName, chaincodeName, peer, args)
}

// 数据查询 	||
func QueryAssetTest(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	return action.QueryAssetTest(clientName, channelName, chaincodeName, peer, args)
}

// 数据查询||  根据主键查询元数据 ||getAssetByID
func ChainCodeQueryById(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	return action.ChainCodeQueryById(clientName, channelName, chaincodeName, peer, args)
}

//  微分格 数据上链
func UploadByBox(clientName string, channelName string, chaincodeName string, funcName string, peer string, args []string, uuid string) (result string, err error) {
	return action.UploadByBox(clientName, channelName, chaincodeName, funcName, peer, args, uuid)
}

//  微分格 数据上链 --查询
func QueryloadByBox(clientName string, channelName string, chaincodeName string, funcName string, peer string, args []string, uuid string) (result string, err error) {
	return action.QueryloadByBox(clientName, channelName, chaincodeName, funcName, peer, args, uuid)
}

// 微分格 数据查询||  根据主键查询元数据 ||getAssetByID
func ChainCodeQueryByIdByBox(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	return action.ChainCodeQueryByIdByBox(clientName, channelName, chaincodeName, peer, args)
}

// SQL  --init
func SqlByInit() {
	action.SqlByInit()
}

//	大屏@合约处理展示 @ 1
func SqlByConact() (rep interface{}, err error) {
	return action.SqlByConact()
}

//	大屏@ 智能合约调用信息表 @ 2
func SqlByLarge() (rep interface{}, err error) {
	return action.SqlByLarge()
}

//	大屏@ 共识节点信息表 @ 3
func SqlByServer() (rep interface{}, err error) {
	return action.SqlByServer()
}

//	大屏@ 节点信息统计表 @ 4
func SqlBySvg() (rep interface{}, err error) {
	return action.SqlBySvg()
}

//	大屏@ 区块信息表 @ 5
func SqlByBlock() (rep interface{}, err error) {
	return action.SqlByBlock()
}

//	大屏@ 节点信息交易列表 @ 6
func SqlByDeal() (rep interface{}, err error) {
	return action.SqlByDeal()
}

//	大屏@ 资产上链信息表 @ 7
func SqlByAsset() (rep interface{}, err error) {
	return action.SqlByAsset()
}

// SQL -- DB 全局 对象
func SqlByInitByDB() {
	//
	action.SqlByInitByDB()
}

// SQL -- 查询 --
func SqlQueryByBlock(blockHeight string) (rep interface{}, err error) {
	return action.SqlQueryByBlock(blockHeight)
}

// SQL -- 查询 --调用次数
func SqlByLargeForInserData(cName string, chainName string, funcName string, org string, userName string) {
	action.SqlByLargeForInserData(cName, chainName, funcName, org, userName)
}

// SQL -- 查询 --调用次数
func SqlByBlockForInserData(height string, hash string, channel string, chaincode string, orgName string, userName string) {
	action.SqlByBlockForInserData(height, hash, channel, chaincode, orgName, userName)
}

// SQL -- 插入  -- 合约调用个次数
func SqlBySvgForInserData(txId string, orgName string, userName string) {
	action.SqlBySvgForInsertData(txId, orgName, userName)
}

// SQL --- 插入 节点交易信息表
func SqlByDbDealForInsertData(txid string, height string, userID string, chainCode string, orgName string, userName string) {
	action.SqlByDbDealForInsertData(txid, height, userID, chainCode, orgName, userName)
}

// SQL --- 资产信息表 插入// SQL --- 插入 节点交易信息表
func SqlByDbAssetForInsertData(assetNo string, conType string, chainCode string, upTime string, assetType string, orgName string, userName string, txid string) {
	action.SqlByDbAssetForInsertData(assetNo, conType, chainCode, upTime, assetType, orgName, userName, txid)
}

// SQL -- 更新 服务器配置信息
func UpdateByIp() {
	action.UpdateByIp()
}

// SQL -- 节点信息列表统计

//	大屏@ 节点信息交易列表 @ 6
func SqlBySvgForBig() (rep interface{}, err error) {
	return action.SqlBySvgForBig()
}

// Poller
func Poller() {
	action.Poller()
}

// 微分格 数据上链  -- 根据主键查询区块信息 2019年12月24日14:10:33
func GetBlockTxIdForBox(channelName string, cliName string, peerName string, orgName string, userName string, txid string) (config string, data []string, err error) {
	return action.GetBlockTxIdForBox(channelName, cliName, peerName, orgName, userName, txid)
}

//	大屏@ 地理位置列表 2019年12月26日16:21:33
func SqlBySvgForGeography() (rep interface{}, err error) {
	return action.SqlBySvgForGeography()
}

//	大屏@ 资产上链信息表 @ 7
func SqlByAssetNext(no string, t string) (rep interface{}, err error) {
	return action.SqlByAssetNext(no, t)
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func SqlByDealList(id string) (rep interface{}, err error) {
	return action.SqlByDealList(id)
}

func SqlByQueryOntransactionID(txid string) (things string, userId string, err error) {
	return action.SqlByQueryOntransactionID(txid)
}

//
func TestAssettest(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {

	return action.TestAssettest(clientName, channelName, chaincodeName, peer, args)
}

func TestSelectChaincode(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {

	return action.TestSelectChaincode(clientName, channelName, chaincodeName, peer, args)
}

//

func SqlByAssetSum() (rep interface{}, err error) {
	return action.SqlByAssetSum()
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func SqlByDealGeography(id string) (rep interface{}, err error) {
	return action.SqlByDealGeography(id)
}
