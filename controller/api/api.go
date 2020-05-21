package api

import (
	"ki-sdk/controller/action"
	p "ki-sdk/model/action"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//
func Sendopt(dataName string) interface{} {
	return action.Selectetopt(dataName)
}

// 验证
func Scverify(orgId string, userId string) bool {
	return action.Selectcverify(orgId, userId)
}

// 数据上链
func UpLoad(c *gin.Context) {

	// 数据上链
	list, err := action.CupLoad(c)
	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

/**
@ api  查询区块链首页信息
@ 2019年10月21日16:37:36
@ 描述： 查询各个通道的最新的区块信息来为区块链浏览器进行查询
*/
func QueryBlockInfo(c *gin.Context) {
	// 查询
	result, err := action.CqueryBlockInfo(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

/**
@ api  查询区块链 云图配置
@	时间： 2019年10月21日17:47:24
@ 描述： 查询区块链中的配置 || 链码|| 节点||  交易数量 || 节点||

*/
func QueryConfig(c *gin.Context) {
	// 查询
	result, err := action.CqueryConfig(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

/**
@	函数： 区块链浏览器 - 根据TxID 查询
@	时间： 2019年10月21日18:17:26
@	描述:	根据TXID 查询 数据信息

*/
func QueryBlockByTxId(c *gin.Context) {
	// 查询
	result, err := action.CqueryBlockByTxId(c)

	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

/**
@	函数： 区块链浏览器 - 查询通道动态信息
@	时间： 2019年10月21日19:57:15
@	描述:	通道ID
*/
func QueryBlockDynamic(c *gin.Context) {
	// 查询
	result, err := action.CqueryBlockDynamic(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

/**
@	函数： 区块链浏览器 - 一个区块的交易
@	时间： 2019年10月21日20:26:14
@	描述:	区块高度
*/
func QueryBlockTransaction(c *gin.Context) {
	// 查询
	result, err := action.CqueryBlockTransaction(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

/**
@	函数： 联盟组织	- 添加组织
@	时间： 2019年10月22日10:11:52
@	描述:  对于联盟组织进行注册  ID, USERID  ||| NewAffiliation
*/
func AddOrg(c *gin.Context) {
	// 查询
	result, err := action.CaddOrg(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   result,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

/**
@	函数： 联盟组织	- 注册用户
@	时间： 2019年10月22日10:59:59
@	描述:  对于联盟组织进行用户注册 setUser
*/
func SetUser(c *gin.Context) {
	// 查询
	result, err := action.SetUser(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

// 数据上链
func UpLoadTest(c *gin.Context) {
	list, err := action.CupLoad(c)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

//
func GetById1(c *gin.Context) {
	//
	var args []string

	// 04. 拼接参数
	args = append(args, "getAssetList")                        //函数名称
	args = append(args, "20b362d9-3f1c-4463-b91b-23232qweqwe") //参数 id
	resut, err := p.QueryById("a", "a", "a", "a", args)
	//
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("---", resut)
	//
	datf := []action.Kyc{}
	json.Unmarshal([]byte(resut), &datf)
	//
	for k, _ := range datf {
		log.Println("单个查询", datf[k].KycID)
		log.Println("单个查询", datf[k].KycType)
	}
	//

}

//根据 主键查询元数据
func GetById(c *gin.Context) {
	rep, err := action.QueryMetadataByPrimaryKey(c)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": rep,
			})
		return
	}
}

// 根据主键查询关联数据
func QueryListByID(c *gin.Context) {

	// 数据查询
	list, err := action.QueryListByID(c)
	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

// 根据主键查询标签信息
func GetLabById(c *gin.Context) {
	//
	rep, err := action.QueryLableByPrimaryKey(c)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": rep,
			})
		return
	}

}

// @ 数据上链  应收账款 旧版 ||
func Upload(c *gin.Context) {

	// 数据上链
	list, err := action.UpLoadByAst(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

// V3----------------------------------- 升级

// V3 新增组织
func V3_addOrg(c *gin.Context) {
	// 查询
	result, err := action.V3_addOrg(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

//  @ 旧版@ 应收账款数据上链  2019年11月7日10:43:02
func receUpLoad(c *gin.Context) {

	// 数据上链
	list, err := action.CupLoad(c)
	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

// 微分格数据上链  2019年11月7日18:55:44
func UpLoadByBox(c *gin.Context) {
	//
	// 数据上链
	list, err := action.UpLoadByBox(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

//@微分格@根据 主键查询元数据
func GetByIdByBox(c *gin.Context) {
	rep, err := action.QueryMetadataByPrimaryKeyByBox(c)
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": rep,
			})
		return
	}
}

// SQL -- init
func SqlByInit(c *gin.Context) {
	//
	// 数据上链
	list, err := action.SqlByInit(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

//	大屏@合约处理展示 @ 1
func SqlByConact(c *gin.Context) {
	//
	// 数据上链
	arr, err := action.SqlByConact(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@ 智能合约调用信息表 @ 2
func SqlByLarge(c *gin.Context) {
	// 数据上链
	arr, err := action.SqlByLarge(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@ 共识节点信息表 @ 3
func SqlByServer(c *gin.Context) {
	// 数据上链
	arr, err := action.SqlByServer(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@ 节点信息统计表 @ 4
func SqlBySvg(c *gin.Context) {
	// 数据上链
	arr, err := action.SqlBySvg(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@ 区块信息表 @ 5
func SqlByBlock(c *gin.Context) {
	// 数据上链
	arr, err := action.SqlByBlock(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@ 节点信息交易列表 @ 6
func SqlByDeal(c *gin.Context) {
	// 数据上链
	arr, err := action.SqlByDeal(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@ 资产上链信息表 @ 7
func SqlByAsset(c *gin.Context) {

	arr, err := action.SqlByAsset(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@ 根据区块查询信息 @ 8 BlockData
func SqlQueryBlock(c *gin.Context) {
	// 查询
	result, err := action.SqlQueryBlock(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

//	大屏@ 根据区块查询关联10个区块 @ 9 BlockData
func SqlQueryBlockByHeight(c *gin.Context) {
	// 查询
	result, err := action.SqlQueryBlockByHeight(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

//	大屏@ 节点信息交易列表 @ 6  平均统计
func SqlBySvgForAverage(c *gin.Context) {
	// 数据上链
	arr, err := action.SqlBySvgForAverage(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

// 微分格  -- 数据查询
func QueryLoadByBox(c *gin.Context) {
	//
	// 数据上链
	list, err := action.QueryLoadByBox(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

// 微分格 数据上链  -- 根据主键查询区块信息 2019年12月24日14:10:33
func WQueryData(c *gin.Context) {
	//
	// 数据上链
	list, err := action.WQueryData(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": list,
			})
		return
	}
}

//	大屏@ 地理位置列表 2019年12月26日16:21:33
func SqlBySvgForGeography(c *gin.Context) {
	// 数据上链
	arr, err := action.SqlBySvgForGeography(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

// 	查询通道 区块信息
func QueryBlockDynamicList(c *gin.Context) {
	// 查询
	result, err := action.CqueryBlockDynamicList(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

//  查询后五个
func SqlByGetNextData(c *gin.Context) {
	// 查询
	result, err := action.CqueryNextBlockTransaction(c)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   err,
			})
		return
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": &result,
			})
		return
	}
}

// 查询资产后五条数据
func SqlByAssetNext(c *gin.Context) {

	arr, err := action.SqlByAssetNext(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@  复合查询   2020年1月13日19:30:33
func SqlByCompoundQuery(c *gin.Context) {

	arr, err := action.SqlByCompoundQuery(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func SqlByDealList(c *gin.Context) {

	arr, err := action.SqlByDealList(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//
//	大屏@  抬头 2020年1月17日18:07:41 Rise
func SqlByRise(c *gin.Context) {

	arr, err := action.SqlByRise(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func SqlByDealGeography(c *gin.Context) {

	arr, err := action.SqlByDealGeography(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": action.StatusText(action.StatusFailed),
				"data":   action.StatusText(action.StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": action.StatusText(action.StatusOK),
				"data": arr,
			})
		return
	}
}
