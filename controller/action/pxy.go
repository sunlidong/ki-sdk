package action

import (
	modelapp "ki-sdk/model/action"
	serverapi "ki-sdk/server/api"

	"github.com/chainHero/blockdata"
	"github.com/gin-gonic/gin"
)

/**
@ 返回结构体
@ 2019年10月14日17:55:05
*/
func Selectetopt(dataName string) interface{} {
	return opt(dataName)
}

/*
@ 调用序列化数据
@ 2019年10月18日10:06:30
@ lidong sun
*/
func SelectoptSerialize(c *gin.Context, data interface{}) (cdata interface{}, code string) {
	return optSerialize(c, data)
}

/*
@ 验证用户信息
@ 2019年10月18日10:06:30
@ lidong sun
*/
func Selectcverify(orgId string, userId string) bool {
	return cverify(orgId, userId)
}

/**
@ 数据上链  ||房地产数据上链
@ 2019年10月21日12:01:38
@
*/
func CupLoad(c *gin.Context) (list *blockdata.ChainTransaction, err error) {
	return cupLoad(c)
}

// V3 数据上链
func UpLoadByAst(c *gin.Context) (list *blockdata.ChainTransactionup, err error) {
	return upLoadByAst(c)
}

/**
@	函数：pxy 查询区块链首页信息
@	时间：2019年10月21日16:44:05
@	描述： model 一级函数
*/
func CqueryBlockInfo(c *gin.Context) (rep *[]modelapp.BlockWordback, err error) {
	return cqueryBlockInfo(c)
}

/**
@	函数：pxy 查询区块链配置信息
@	时间：2019年10月21日17:49:20
@	描述： pxy 一级函数
*/
func CqueryConfig(c *gin.Context) (rep *modelapp.YunConfig, err error) {
	return cqueryConfig(c)
}

/**
@	函数：pxy   区块链浏览器 - 根据TxID 查询
@	时间：2019年10月21日17:49:20
@	描述： pxy 一级函数  QueryBlockByTxId
*/
func CqueryBlockByTxId(c *gin.Context) (list interface{}, err error) {
	return cqueryBlockByTxId(c)
}

/**
@	函数： 区块链浏览器 - 查询通道动态信息
@	时间： 2019年10月21日19:57:15
@	描述:	通道ID
*/
func CqueryBlockDynamic(c *gin.Context) (rep interface{}, err error) {
	return cqueryBlockDynamic(c)
}

/**
@	函数： 区块链浏览器 - 一个区块的交易
@	时间： 2019年10月21日20:26:14
@	描述:	区块高度
*/
func CqueryBlockTransaction(c *gin.Context) (rep interface{}, err error) {
	return cqueryBlockTransaction(c)
}

/**
@	函数： 联盟组织	- 添加组织
@	时间： 2019年10月22日10:11:52
@	描述:  对于联盟组织进行注册  ID, USERID
*/
func CaddOrg(c *gin.Context) (rep interface{}, err error) {
	return caddOrg(c)
}

/**
@	函数： 联盟组织	- 添加用户
@	时间： 2019年10月22日11:01:39
@	描述:  对用户进行注册
*/
func SetUser(c *gin.Context) (rep string, err error) {
	return setUser(c)
}

// 根据主键查询数据关联数据
func QueryListByID(c *gin.Context) (rep interface{}, err error) {
	return queryListByID(c)
}

// 根据主键查询元数据
func QueryMetadataByPrimaryKey(c *gin.Context) (rep interface{}, err error) {
	return queryMetadataByPrimaryKey(c)
}

// 根据结构体标识获取结果体字段
func PointerToIndex(pointer string) (inter interface{}, err error) {
	return pointerToIndex(pointer)
}

// 根据结构体标识获取结果体类型
func PointerToType(pointer string) (inter string, dataBype string, err error) {
	return pointerToType(pointer)
}

// 根据主键查询元标签
func QueryLableByPrimaryKey(c *gin.Context) (rep interface{}, err error) {
	return queryLableByPrimaryKey(c)
}

// V3 新增组织
func V3_addOrg(c *gin.Context) (rep string, err error) {
	return v3_addOrg(c)
}

// 微分格 数据上链
func UpLoadByBox(c *gin.Context) (list interface{}, err error) {
	return upLoadByBox(c)
}

// @微分格@根据主键查询元数据
func QueryMetadataByPrimaryKeyByBox(c *gin.Context) (rep interface{}, err error) {
	return queryMetadataByPrimaryKeyByBox(c)
}

// SQL -- Init
func SqlByInit(c *gin.Context) (list interface{}, err error) {
	return sqlByInit(c)
}

//	大屏@合约处理展示 @ 1
func SqlByConact(c *gin.Context) (rep interface{}, err error) {
	return sqlByConact(c)
}

//	大屏@ 智能合约调用信息表 @ 2
func SqlByLarge(c *gin.Context) (rep interface{}, err error) {
	return sqlByLarge(c)
}

//	大屏@ 共识节点信息表 @ 3
func SqlByServer(c *gin.Context) (rep interface{}, err error) {
	return sqlByServer(c)
}

//	大屏@ 节点信息统计表 @ 4
func SqlBySvg(c *gin.Context) (rep interface{}, err error) {
	return sqlBySvg(c)
}

//	大屏@ 区块信息表 @ 5
func SqlByBlock(c *gin.Context) (rep interface{}, err error) {
	return sqlByBlock(c)
}

//	大屏@ 节点信息交易列表 @ 6 DbDeal
func SqlByDeal(c *gin.Context) (rep interface{}, err error) {
	return sqlByDeal(c)
}

//	大屏@ 资产上链信息表 @ 7
func SqlByAsset(c *gin.Context) (rep interface{}, err error) {
	return sqlByAsset(c)
}

//	大屏@ 根据区块查询信息 @ 8 BlockData
func SqlQueryBlock(c *gin.Context) (rep interface{}, err error) {
	return sqlQueryBlock(c)
}

//	大屏@ 根据区块查询关联10个区块 @ 9 BlockData
func SqlQueryBlockByHeight(c *gin.Context) (rep interface{}, err error) {
	return sqlQueryBlockByHeight(c)
}

//
func ParameterTransformation(res AstInfo) (rep *AstInfoWebJiu) {
	return parameterTransformation(res)
}

// -- 尽调结果 处理
func ParameterTransformationByBaseSurvey(res BaseSurvey) (rep *BaseSurveyByJiu) {
	return parameterTransformationByBaseSurvey(res)
}

// -- 尽调报告 处理
func ParameterTransformationByBaseReport(res BaseReport) (rep *BaseReportByJiu) {
	return parameterTransformationByBaseReport(res)
}

//	大屏@ 节点信息交易列表 @ 6 DbDeal
func SqlBySvgForAverage(c *gin.Context) (rep interface{}, err error) {
	return sqlBySvgForAverage(c)
}

// 微分格 数据上链  -- 查询 2019年12月18日10:16:21
func QueryLoadByBox(c *gin.Context) (list interface{}, err error) {
	return queryLoadByBox(c)
}

// 微分格 数据上链  -- 根据主键查询区块信息 2019年12月24日14:10:33
func WQueryData(c *gin.Context) (list interface{}, err error) {
	return queryWQueryData(c)
}

//	大屏@ 地理位置列表 2019年12月26日16:21:33
func SqlBySvgForGeography(c *gin.Context) (rep interface{}, err error) {
	return sqlBySvgForGeography(c)
}

func CqueryBlockDynamicList(c *gin.Context) (rep interface{}, err error) {
	return cqueryBlockDynamicList(c)
}

//	调用通知  参数： assetState	，assetCode
func callTheNotification(assetState string, assetCode string) (err error) {
	return serverapi.NoticePost(assetState, assetCode)
}

//	查询后五个
func CqueryNextBlockTransaction(c *gin.Context) (rep interface{}, err error) {
	return cqueryNextBlockTransaction(c)
}

//	大屏@ 资产上链信息表 @ 7 后五条
func SqlByAssetNext(c *gin.Context) (rep interface{}, err error) {
	return sqlByAssetNext(c)
}

//	大屏@  复合查询   2020年1月13日19:30:33
func SqlByCompoundQuery(c *gin.Context) (rep interface{}, err error) {
	return sqlByCompoundQuery(c)
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func SqlByDealList(c *gin.Context) (rep interface{}, err error) {
	return sqlByDeaLlList(c)
}

//	大屏@  抬头 2020年1月17日18:07:41 Rise
func SqlByRise(c *gin.Context) (rep interface{}, err error) {
	return sqlByRise(c)
}

//	大屏@  地理位置  三个Tab  2020年1月16日09:45:26
func SqlByDealGeography(c *gin.Context) (rep interface{}, err error) {
	return sqlByDealGeography(c)
}
