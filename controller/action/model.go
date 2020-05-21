package action

import (
	model_action "ki-sdk/model/action"
	model_api "ki-sdk/model/api"
	util_action "ki-sdk/util/action"
	util_api "ki-sdk/util/api"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/chainHero/blockdata"
	"github.com/gin-gonic/gin"
)

//
func cverify(orgId string, userId string) bool {
	return verify(orgId, userId)
}

//
func cmSerialize(c *gin.Context) (data *Info, err error) {
	return mSerialize(c)
}

/**
@ 数据上链
@ 2019年10月18日14:16:38
@ lidongsun
@
*/
func cupLoad(c *gin.Context) (list *blockdata.ChainTransaction, err error) {

	//	01.	序列化数据
	data, err := cmSerialize(c)
	if err != nil {
		return nil, err
	}

	//	02. 验证用户权限
	//fal:= cverify(data.User.OrgID, data.User.UserID)
	//if !fal {
	//	return nil, err
	//}

	//	03. 获取上链参数 args 根据不同参数去请求不同的api
	args, state := mDispense(data.Datafunc, data.Datatype, *data)
	log.Println("获取上链参数 args=>", args)
	//
	if state != nil {
		return nil, err
	}

	// 04. 数据上链
	cli := "cli"
	channlName := "AssetToChain"
	peer := "2"
	//
	txid, err := model_api.UploadAsset(cli, channlName, channlName, peer, args)
	if err != nil {
		log.Println("数据上链失败")
		return nil, err
	}
	//
	log.Println("上链完成")
	// 根据 txID 查询上链信息
	rest, err := model_action.App.GetnewBlockTxId(txid)

	if err != nil {
		return nil, err
	}

	if len(rest) == 0 {
		return nil, err

	}
	return rest[0], err
}

/**
@  mod 查询区块链首页信息
@	2019年10月21日16:40:25
@	描述： model 二级函数
*/
func cqueryBlockInfo(c *gin.Context) (rep *[]model_action.BlockWordback, err error) {
	//
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//return queryBlockInfo()
	log.Println("=>参数序列化")
	//	01.	参数序列化 || 数据集合 BlockWordlist
	data, err := unBlockWord(c)
	if err != nil {
		return nil, err
	}

	//	02.	权限控制
	//	02. 验证用户权限
	//state := cverify(data.User.OrgID, data.User.UserID)
	//if !state {
	//	return nil, err
	//}

	//	03.	生成对象

	/**
	生成 对应的 APP TODO
	*/
	log.Println("=>查询对象")
	//	04.	查询对象
	channelName := "channelName"
	result, err := model_action.App.GetBlockWord(channelName, data.User.OrgID, data.User.UserID)
	if err != nil {
		return nil, err
	}

	// 	05.	返回
	return result, nil
}

/**
@  mod 查询区块链配置信息
@	时间：2019年10月21日17:52:33
@	描述： model 二级函数
*/
func cqueryConfig(c *gin.Context) (rep *model_action.YunConfig, err error) {
	//
	//return queryBlockInfo()
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.	参数序列化 || 数据集合 BlockWordlist
	data, err := unChannelBlockConfig(c)
	if err != nil {
		return nil, err
	}

	//	02.	权限控制
	//	02. 验证用户权限
	//state := cverify(data.User.OrgID, data.User.UserID)
	//if !state {
	//	return nil, err
	//}

	//	03.	生成对象

	/**
	生成 对应的 APP TODO
	*/

	//	04.	查询对象
	channelName := "channelName"
	//num := "num"
	//channelID := "assetpublish"
	//result, err := models.App.GetBlockMessage(
	//	channelID,
	//	cData.User.AffiliationId,
	//	num,
	//)
	result, err := model_action.App.GetBlockMessage(channelName, data.User.OrgID, data.User.UserID)
	if err != nil {
		return nil, err
	}

	// 	05.	返回
	return result, nil
}

/**
@	函数：model   区块链浏览器 - 根据TxID 查询
@	时间：2019年10月21日17:49:20
@	描述： model 二级函数  QueryBlockByTxId
*/
func cqueryBlockByTxId(c *gin.Context) (rep interface{}, err error) {

	// 声明参数
	resData := TxIDrenData{}
	resDataJiu := TxIDrenDataJiu{}
	ChainTrans := ChainTransactionConfig{}
	Mg := Kyc{}
	wByType := WeiByType{}
	var strData string
	var JiuWeifen string
	channelName := "channelName"
	cliName := "cliName"
	peerName := "peerName"
	orgName := "orgName"
	userName := "userName"

	// 插入调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")

	// 序列化数据
	data, err := unBlockQuery(c)

	if err != nil {
		return nil, err
	}

	//	04.	查询对象 TODO

	txid := data.QueTxid

	//	 查询交易人ID
	_, user, _ := model_api.SqlByQueryOntransactionID(txid)

	log.Println("userID", user)
	// 赋值
	ChainTrans.UserId = user
	//	根据TXID查询数据集合
	result, resultByString, err := model_action.App.GetBlockTxId(channelName, cliName, peerName, orgName, userName, txid)

	log.Println("result=>", result)

	if err != nil {
		return nil, err
	}

	log.Println("resultByString (len)=>", len(resultByString))

	if len(resultByString) == 0 {
		return nil, errors.New("resultByString is nil")
	}
	if len(resultByString) < 2 {
		log.Println("微分格类型处理")
		// TODO  微分格的数据  默认第一种
		json.Unmarshal([]byte(result), &ChainTrans)
		resData.ChainTransactionConfig = ChainTrans
		return resData, nil
	}

	//	异常情况处理完毕, 开始进行数据处理

	// TODO  数据解密
	// TODO   kyc 数据在  retstring [1]   retstring [0] 类型      retstring [2] 标签

	for k, v := range resultByString {
		log.Println("====>", k)
		log.Println("==================================>", v)
	}

	// 统一处理  判断是否是历史数据, 加标识判断
	json.Unmarshal([]byte(result), &ChainTrans)
	// 配置信息

	// 如果 查询区块高于 725
	log.Println("统一处理  判断是否是历史数据, 加标识判断", ChainTrans.Height)
	if ChainTrans.Height <= int64(800) {
		JiuWeifen = "v1"
	} else {
		JiuWeifen = "v2"
	}
	ChainTrans.Version = JiuWeifen
	//	数据转码
	err = json.Unmarshal([]byte(resultByString[1]), &Mg)

	if err != nil {
		log.Println("data transcoding failed:", err)
		return "", err
	}
	err = json.Unmarshal([]byte(resultByString[2]), &wByType)
	if err != nil {
		log.Println("data transcoding failed:", err)
		return "", err
	}
	log.Println("wByType:", wByType.CategoryId)
	ChainTrans.PledgeType = wByType.CategoryId
	// TODO  判断是哪一种数据类型

	//	数据判断
	if Mg.KycString != "" || Mg.SignKey != "" {
		// 解密
		strData = MetadataAesDecrypt(Mg.KycString, Mg.SignKey)
		log.Println("strData=>", strData)
	} else {
		// 如果数据为空,那么就确认是微分格数据, 赋值 标签 TODO
		Mg.KycType = Pro_Rec_WI
	}

	// 分发器

	switch Mg.KycType {

	//	房地产
	case Pro_Rea_All:
		label, dataType, data, err := queryDataByFang(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针:", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label
	//	应收账款
	case Pro_Rec_Ysk:
		label, dataType, data, err := queryDataByYing(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resDataJiu.Data = &data
		resDataJiu.DataType = label

		//	旧数据处理 这层必须直接返回,其他 resData 可以后续追加
		json.Unmarshal([]byte(result), &ChainTrans)
		resDataJiu.ChainTransactionConfig = ChainTrans
		return resDataJiu, nil

	//	尽调结果
	case Pro_Rec_Yjg:
		label, dataType, data, err := queryDataByBaseReport(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label

	//	尽调报告
	case Pro_Rec_Jdb:
		label, dataType, data, err := queryDataByBaseSurvey(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label

	//	动产质押
	case Pro_Rec_JCZY:
		label, dataType, data, err := queryDataByBasePle(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
			return "", err
		}

		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label

	//	微分格
	case Pro_Rec_WI:
		log.Println("微分格查询")
		label, dataType, data, err := queryDataByWeiForapp(&Mg, resultByString[1], JiuWeifen)

		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label
	default:
		log.Println("There is no corresponding type")
	}
	resData.ChainTransactionConfig = ChainTrans
	// 返回
	return resData, nil
}

/**
@	函数： 区块链浏览器 - 查询通道动态信息
@	时间： 2019年10月21日19:57:15
@	描述:	通道ID  controller_action.
*/
func cqueryBlockDynamic(c *gin.Context) (rep interface{}, err error) {

	//	01.	参数序列化 || 数据集合 unBlockDynamic
	//data, err := unBlockDynamic(c)
	//if err != nil {
	//	return nil, err
	//}
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	03.	生成对象

	/**
	生成 对应的 APP TODO
	*/

	//	04.	查询对象
	channelName := "assetpublish"

	orgName := "orgName"
	userName := "userName"
	//res := data.ChannelID
	//
	//log.Println("res", res)

	log.Println("channelName=>", channelName)
	//log.Println("channelName=>", data.ChannelID)
	// channelName GetBlockDynamic
	//result, err := model_action.App.GetBlockDynamic(data.ChannelID, orgName, userName)
	//channelName:="assetpublish"
	result, err := model_action.App.GetBlockDynamicByBig(channelName, orgName, userName)
	if err != nil {
		return nil, err
	}
	// 	05.	返回
	return result, nil
}

/**
@	函数： 区块链浏览器 - 一个区块的交易
@	时间： 2019年10月21日20:26:14
@	描述:	区块高度
*/
func cqueryBlockTransaction(c *gin.Context) (rep interface{}, err error) {
	//
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.	参数序列化 || 数据集合 BlockTransaction
	data, err := unBlockTransaction(c)
	if err != nil {
		return nil, err
	}

	/**
	生成 对应的 APP TODO
	*/

	//	04.	查询对象
	channelName := "assetpublish"

	cliName := "cliname"
	orgName := "orgName"
	userName := "userName"
	res := data.BlockNum
	//
	log.Println("res", res)

	// channelName GetBlockDynamic
	result, err := model_action.App.GetBlockTransactionByBig(channelName, cliName, orgName, userName, res)
	if err != nil {
		return nil, err
	}

	// 	05.	返回
	return result, nil
}

/**
@	函数： 联盟组织	- 添加组织
@	时间： 2019年10月22日10:11:52
@	描述:  对于联盟组织进行注册  ID, USERID NewAffiliation
*/
func caddOrg(c *gin.Context) (rep interface{}, err error) {

	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.	参数序列化 || 数据集合 NewAffiliation
	data, err := unNewAffiliation(c)
	log.Println("01 参数序列化", data)
	if err != nil {
		return "", err
	}
	//

	//TODO
	caName := "ca-org1-msp.dinglian.com"
	caPath := "com.dinglian.org1."
	focle := true
	affilName := data.AffiliationID

	result, err1 := model_action.App.AddAffiliation(caName, focle, caPath, affilName)
	log.Println(result.CAName)
	if err1 != nil {
		log.Println("注册联盟失败", result.CAName)
		return "注册联盟失败", err1
	}

	log.Println("注册联盟成功")
	//	生成 msp cli
	// TODO
	num := "com.dinglian.org1."
	mspClient, err := model_action.App.CreateNewMspClient(data.AffiliationID, num)
	if err != nil {
		//is nil  返回
		return "注册联盟失败", nil
	}
	log.Println("mspClient cli si ok ")

	// 生成   userFFID

	hash := util_api.GetcurrentHash(util_action.EncryptText(util_action.EncryptDsha256), data.OrgID+data.UserId)
	//	添加联盟管理员  || 添加组织
	log.Println("hash=>", hash)
	path := "com.dinglian.org1."
	pwd := "123456"
	_, err = model_action.App.MspAffCreateUser(mspClient, data.AffiliationID, hash, data.UserName, data.OrgName, data.OrgID, path, pwd)
	log.Println("注册用户：", err)
	// err
	if err != nil {

		return "", err
	}

	// TODO
	if result.Affiliations == nil && result.Identities == nil {
		return OrgWeb{
			result.Name,
			"",
			"",
			result.CAName,
		}, nil

	} else if result.Identities == nil {
		return OrgWeb{
			result.Name,
			result.Affiliations,
			"",
			result.CAName,
		}, nil

	} else if result.Affiliations == nil {
		return OrgWeb{
			result.Name,
			"",
			result.Identities,
			result.CAName,
		}, nil

	} else {
		return OrgWeb{
			result.Name,
			result.Affiliations,
			result.Identities,
			result.CAName,
		}, nil

	}
}

/**
@	函数： 联盟组织	- 添加用户
@	时间： 2019年10月22日11:01:39
@	描述:  对用户进行注册
*/
func setUser(c *gin.Context) (rep string, err error) {

	//
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.	参数序列化 || 数据集合 Registered
	data, err := unRegistered(c)
	if err != nil {
		log.Println("=>01. 序列化失败", err)
		return "", err
	}
	log.Println("=>01. 序列化")

	//	02. 验证权限 TODO
	state := cverify(data.RegOrgID, data.UserID)
	if !state {
		log.Println("=>01. 验证权限失败")
		return "", err
	}
	log.Println("=>02. 验证权限")

	//	03.	生成 msp 客户端
	num := "com.dinglian.org1."
	mspClient, err := model_action.App.CreateNewMspClient(data.RegAffiliationId, num)
	if err != nil {
		//is nil  返回
		log.Println("=>03. 生成Msp客户端失败:", err)
		return
	}

	log.Println("=>03. 生成Msp客户端")

	//	04.	生成 userFFID
	hash := util_api.GetcurrentHash(util_action.EncryptText(util_action.EncryptDsha256), data.RegOrgID+data.RegUserID)
	log.Println("=>04. 生成用户HASH")

	//	05.	添加联盟管理员  || 添加组织 TODO

	path := "com.dinglian.org2."
	pwd := "123456"
	mspUser, err := model_action.App.MspAffCreateUser(mspClient, data.RegAffiliationId, hash, data.UserName, data.RegOrgName, data.RegOrgID, path, pwd)
	// err
	if err != nil {
		log.Println("=>05. 注册用户失败:", err)
		return "", err
	}

	//TODO
	log.Println("=>05. 注册用户")
	return mspUser, nil
}

//根据主键查询数据关联数据
func queryListByID(c *gin.Context) (rep interface{}, err error) {
	//
	//	01.	序列化数据
	data, err := cmSerialize(c)
	if err != nil {
		return "", err
	}
	log.Println("1")
	//查询 封装
	args, err := qDispense(QueryDataListByID, "", *data)
	if err != nil {
		log.Println("查询 封装=>", err)
		return "", err
	}

	//	查询数据 =>
	//txid, err := api.UploadAsset(cli, channlName, channlName, peer, args)
	//TODO
	cli := "cli"
	channlName := "cli"
	peer := "cli"

	res, err := model_api.QueryAssetTest(cli, channlName, channlName, peer, args)
	if err != nil {
		log.Println("查询数据=>", err)
		return "", nil
	}
	// 数据处理
	log.Println("根据主键查询数据关联数据=>", res)

	Bg := []Kyc{}
	// 反序列化数据

	resData := []LinkedData{}
	//	序列化
	err = json.Unmarshal([]byte(res), &Bg)
	//	解密
	//
	log.Println("arr=>", len(Bg))
	for k, _ := range Bg {
		rData := LinkedData{}
		// 解密
		strData := MetadataAesDecrypt(Bg[k].KycString, Bg[k].SignKey)
		log.Println("解密后的数据是=>", strData)
		// 获取结构体类型
		strPointer, err := PointerToIndex(Bg[k].KycType)
		err = json.Unmarshal([]byte(strData), &strPointer)
		if err != nil {
			log.Println("链码结果反序列化失败:err=>", err)
			return "", err
		}

		// 追加数据集合
		rData.Rel = Title_Sun
		rData.Data = strPointer
		log.Println("strPointer=>", strPointer)
		// 追加
		resData = append(resData, rData)
	}
	return resData, nil
}

//根据主键查询数据关联数据
func queryListByIDbz(c *gin.Context) (rep string, err error) {
	//
	//	01.	序列化数据
	data, err := cmSerialize(c)
	if err != nil {
		return "", err
	}
	log.Println("1")
	//查询 封装
	args, err := qDispense(QueryDataListByID, "", *data)
	if err != nil {
		log.Println("查询 封装=>", err)
		return "", err
	}

	//	查询数据 =>
	//txid, err := api.UploadAsset(cli, channlName, channlName, peer, args)
	//TODO
	cli := "cli"
	channlName := "cli"
	peer := "cli"

	res, err := model_api.QueryAssetTest(cli, channlName, channlName, peer, args)
	if err != nil {
		log.Println("查询数据=>", err)
		return "", nil
	}
	// 数据处理
	log.Println("根据主键查询数据关联数据=>", res)

	Bg := []Kyc{}
	// 反序列化数据
	Infr := RealtyInfo{}
	//	序列化
	err = json.Unmarshal([]byte(res), &Bg)
	//	解密
	//
	log.Println("arr=>", len(Bg))
	for k, _ := range Bg {
		strData := MetadataAesDecrypt(Bg[k].KycString, Bg[k].SignKey)
		log.Println("解密后的数据是=>", strData)
		//
		err = json.Unmarshal([]byte(strData), &Infr)
		if err != nil {
			log.Println("链码结果反序列化失败:err=>", err)
			return "", err
		}
		log.Println("Infr=>", Infr.FDueProject.FID)
		log.Println("Infr data =>", Infr)
	}

	return "", nil
}

// V3 数据上链
func upLoadByAst(c *gin.Context) (list *blockdata.ChainTransactionup, err error) {

	//	01.	序列化数据
	data, err := cmSerialize(c)
	if err != nil {
		log.Println("=>数据上链=>01.	拼接上链标签结构失败", err)
		return nil, err
	}
	log.Println("=>数据上链=>01.	拼接上链标签结构")

	//	02. 验证用户权限
	fal := cverify(data.User.OrgID, data.User.UserID)
	if !fal {
		log.Println("=>数据上链=>02.	验证用户权限失败：", err)
		return nil, err
	}
	log.Println("=>数据上链=>02.	验证用户权限")
	//TODO  SQL 添加调用次数
	go model_api.SqlByLargeForInserData(
		Sql_Channel,
		Sql_ChainCodeName,
		Sql_ChainCodoFunc,
		data.User.OrgName,
		data.User.UserName,
	)

	//	03. 获取上链参数 args 根据不同参数去请求不同的api
	args, state := DispenseV3(data.Datafunc, data.Datatype, *data)
	if state != nil {
		log.Println("=>数据上链=>03.	获取上链参数失败：", err)
		return nil, err
	}
	log.Println("=>数据上链=>03.	获取上链参数")

	//	05 	获取 通道 链码 TODO  通道 和 链码校验
	cli := "cli"
	channlName := ""
	chainCodeName := ""
	peer := "2"

	//查询  合约类型

	//
	txid, err := model_api.UploadAssetTest(cli, channlName, chainCodeName, peer, args)
	if err != nil {
		log.Println("=>数据上链=>04.	数据上链失败", err)
		return nil, err
	}
	//
	log.Println("=>数据上链=>04.	数据上链")

	// 根据 txID 查询上链信息
	rest, err := model_action.App.GetnewBlockTxIdbf(txid)

	if err != nil {
		log.Println("=>数据上链=>05.	根据TXID查询上链信息", err)
		return nil, err
	}

	if len(rest) == 0 {
		log.Println("=>数据上链=>05.	数据上链的数据长度为0")
		return nil, err
	}
	log.Println("=>数据上链=>05. 根据TXID查询上链信息")

	// TODO  合约调用次数插入
	go model_api.SqlBySvgForInserData(rest[0].TxID, data.User.OrgName, data.User.UserName)

	// TODO  区块信息表插入

	go model_api.SqlByBlockForInserData(strconv.FormatInt(rest[0].Height, 10), rest[0].Hash, rest[0].ChannelId, rest[0].Chaincode, data.User.OrgName, data.User.UserName)

	// TODO  插入  节点信息表
	go model_api.SqlByDbDealForInsertData(rest[0].TxID, strconv.FormatInt(rest[0].Height, 10), data.User.UserID, rest[0].Chaincode, data.User.OrgName, data.User.UserName)

	// TODO  插入  资产信息表
	go model_api.SqlByDbAssetForInsertData("001", selectByContract(data.Datatype), selectByAsset(data.Datatype), rest[0].Time, data.User.UUID, data.User.OrgName, data.User.UserID, txid)
	//
	return rest[0], err
}

// V3 根据主键查询元数据
func queryMetadataByPrimaryKey(c *gin.Context) (rep interface{}, err error) {
	//	01.	序列化数据
	data, err := cmSerialize(c)
	if err != nil {
		return "", err
	}
	args, err := MetaLabelByPrimaryKey(*data)

	// 02. 	拼接参数
	if err != nil {
		log.Println("拼接参数=>", err)
		return "", err
	}

	// 03.  链码查询 || getAssetByID
	// TODO  cli  chanel chaincode  peer  args

	cli := "cli"
	chanel := "chanel"
	chaincode := "chaincode"
	peer := "peer"

	//	链码查询
	res, err := model_api.ChainCodeQueryById(cli, chanel, chaincode, peer, args)
	if err != nil {
		log.Println("链码查询 =>", err)
		//
		return "", nil
	}
	//	数据转码
	Mg := Kyc{}

	log.Println("链码查询结果=>", res)

	err = json.Unmarshal([]byte(res), &Mg)
	if err != nil {
		log.Println("数据转码失败=>", err)
		return "", err
	}

	// 解码
	strData := MetadataAesDecrypt(Mg.KycString, Mg.SignKey)
	log.Println("strData=>", strData)

	// 获取对应的对象结构体指针
	strPointer, err := PointerToIndex(Mg.KycType)
	if err != nil {
		log.Println("获取对应的对象结构体指针=>", err)
		return "", err
	}

	err = json.Unmarshal([]byte(strData), strPointer)
	//
	if err != nil {
		log.Println("err=>", err)
		return "", err
	}
	//
	log.Println("RealtyInfo1=>", strPointer)

	//
	return strPointer, nil
}

// V3 根据主键查询元标签
func queryLableByPrimaryKey(c *gin.Context) (rep interface{}, err error) {
	//	01.	序列化数据
	data, err := cmSerialize(c)
	if err != nil {
		return "", err
	}
	args, err := MetaLabelByPrimaryKey(*data)

	// 02. 	拼接参数
	if err != nil {
		log.Println("拼接参数=>", err)
		return "", err
	}

	// 03.  链码查询 || getAssetByID
	// TODO  cli  chanel chaincode  peer  args

	cli := "cli"
	chanel := "chanel"
	chaincode := "chaincode"
	peer := "peer"

	//	链码查询
	res, err := model_api.ChainCodeQueryById(cli, chanel, chaincode, peer, args)
	if err != nil {
		log.Println("链码查询 =>", err)
		//
		return "", nil
	}
	//	数据转码
	Bg := Byc{}

	log.Println("链码查询结果=>", res)

	err = json.Unmarshal([]byte(res), &Bg)
	if err != nil {
		log.Println("数据转码失败=>", err)
		return "", err
	}
	//
	log.Println("Bg=>", Bg.BycSign.Self)
	log.Println("Bg1=>", Bg)

	return Bg, nil
}

// V3 新增组织   默认统一一个CA
func v3_addOrg(c *gin.Context) (rep string, err error) {

	//	01.	参数序列化 || 数据集合 联盟
	data, err := unAffition(c)

	//	02. nil
	if err != nil {
		log.Println("=>1 参数序列化失败:", err)
		return "", err
	}

	log.Println("=>1 参数序列化成功:", data.UserId)

	//	02.	权限控制
	//state := cverify(data.User.OrgID, data.User.UserID)
	//if !state {
	//	return nil, err
	//}

	// 添加联盟

	caName := "caName"
	caPath := "caPath"
	focle := false
	affilName := data.AffiliationID

	// 调用  SDK
	result, err := model_action.App.AddAffiliation(caName, focle, caPath, affilName)

	//TODO  生成 可以 自己的 msp

	if err != nil {
		log.Println("=>2 注册联盟组织失败:", err)

		return "", err
	}
	log.Println("=>0 临时输出:", result.CAName)
	log.Println("=>2 注册联盟组织成功")

	//	生成 msp cli
	mspClient, err := model_action.App.CreateNewMspClient(data.AffiliationID, "1")
	if err != nil {
		//is nil  返回
		log.Println("=>3 生成msp客户端失败:", err)
		return
	}
	log.Println("=>3 生成msp客户端")

	// 生成   user uuid
	hash := util_api.GetcurrentHash(util_action.EncryptText(util_action.EncryptDsha256), data.OrgID+data.UserId)

	log.Println("=>4 生成管理员的hash:", hash)
	//	添加联盟管理员  || 添加组织
	// 第一次注册组织的管理员
	path := "com.dinglian.org2."
	pwd := "123456"
	_, err = model_action.App.MspAffCreateUser(mspClient, data.AffiliationID, hash, data.UserName, data.OrgName, data.OrgID, path, pwd)
	// err
	if err != nil {
		log.Println("=>5 第一次注册组织的管理员失败:", err)
		return "", err
	}
	log.Println("=>5 第一次注册组织的管理员成功:")

	//
	return "ok", err
}

// 微分格 数据上链
func upLoadByBox(c *gin.Context) (list interface{}, err error) {

	
	data, err := mSerializeByBox(c)
	if err != nil {
		log.Println("=>微分格=>01.	拼接上链标签结构失败", err)
		return nil, err
	}
	log.Println("=>微分格=>01.	拼接上链标签结构")

	fmt.Printf("Data:", data.Data)


	//	05 	获取 通道 链码 TODO  通道 和 链码校验
	cli := "cli"
	channlName := data.Chain.ChannelName
	chainCodeName := data.Chain.ChainCodeName
	funcName := data.Chain.FuncName
	args := data.Data
	uuid := data.Chain.UUID
	peer := "peer"

	log.Println("参数 args ：", args)
	log.Println("参数 ChainCodeName：", data.Chain.ChainCodeName)
	log.Println("参数 ChannelName：", data.Chain.ChannelName)
	log.Println("参数 FuncName：", data.Chain.FuncName)

	// nil
	if channlName == "" {
		log.Println("=>微分格=>04. 通道信息名称为空	", err)
		return nil, errors.New("通道信息名称为空")
	}
	//
	if chainCodeName == "" {
		log.Println("=>微分格=>04. 链码信息名称为空	", err)
		return nil, errors.New("链码信息名称为空")
	}

	if funcName == "" {
		log.Println("=>微分格=>04. 函数名称为空	", err)
		return nil, errors.New("函数名称为空")
	}

	//
	result, err := model_api.UploadByBox(cli, channlName, chainCodeName, funcName, peer, args, uuid)
	if err != nil {
		log.Println("=>数据上链=>04.	数据上链失败", err)
		return nil, err
	}

	return result, err
}

// 微分格 数据上链  -- 查询
func queryLoadByBox(c *gin.Context) (list interface{}, err error) {

	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")

	//	01.	序列化数据
	data, err := mSerializeByBox(c)
	if err != nil {
		log.Println("=>微分格=>01.	拼接上链标签结构失败", err)
		return nil, err
	}
	log.Println("微分格=>01.拼接上链标签结构")

	fmt.Printf("Data:", data.Data)
	//	02. 验证用户权限
	fal := cverify(data.User.OrgID, data.User.UserID)
	if !fal {
		log.Println("=>微分格=>02.	验证用户权限失败：", err)
		return nil, err
	}
	log.Println("微分格=>02.验证用户权限")

	//	05 	获取 通道 链码 TODO  通道 和 链码校验
	cli := "cli"
	channlName := data.Chain.ChannelName
	chainCodeName := data.Chain.ChainCodeName
	funcName := data.Chain.FuncName
	args := data.Data
	uuid := data.Chain.UUID
	peer := "peer"

	log.Println("转换后的参数是：", args)

	// nil
	if channlName == "" {
		log.Println("=>微分格=>04. 通道信息名称为空	", err)
		return nil, errors.New("通道信息名称为空")
	}
	//
	if chainCodeName == "" {
		log.Println("=>微分格=>04. 链码信息名称为空	", err)
		return nil, errors.New("链码信息名称为空")
	}

	if funcName == "" {
		log.Println("=>微分格=>04. 函数名称为空	", err)
		return nil, errors.New("函数名称为空")
	}

	//
	result, err := model_api.QueryloadByBox(cli, channlName, chainCodeName, funcName, peer, args, uuid)
	if err != nil {
		log.Println("=>数据上链=>04.	数据上链失败", err)
		return nil, err
	}

	// TODO  插入  资产信息表

	return result, err
}

// @微分格@ 根据主键查询元数据
func queryMetadataByPrimaryKeyByBox(c *gin.Context) (rep interface{}, err error) {
	//	01.	序列化数据
	data, err := cmSerializeByBox(c)
	if err != nil {
		log.Println("微分格@查询数据@序列化数据失败", err)
		return "", err
	}
	log.Println("1@微分格@查询数据@序列化数据成功")
	args, err := MetadataByPrimaryKeyByBox(*data)

	// 02. 	拼接参数
	if err != nil {
		log.Println("2@微分格@查询数据@拼接参数失败", err)
		return "", err
	}
	log.Println("2@微分格@查询数据@拼接参数成功", args)
	// 03.  链码查询 || getAssetByID
	// TODO  cli  chanel chaincode  peer  args

	cli := "cli"
	chanel := "chanel"
	chaincode := data.Chain.ChainCodeName
	peer := "peer"

	//	链码查询
	res, err := model_api.ChainCodeQueryByIdByBox(cli, chanel, chaincode, peer, args)
	if err != nil {
		log.Println("链码查询 =>", err)
		//
		return "", nil
	}

	//
	return res, nil
}

// 微分格数据 请求参数序列化
func cmSerializeByBox(c *gin.Context) (data *InfoWei, err error) {
	return mSerializeByBox(c)
}

// SQL  --init
func sqlByInit(c *gin.Context) (list interface{}, err error) {
	//	01.	序列化数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	log.Println("=>1. 数据库表初始化")

	//	 02. 连接数据库
	model_api.SqlByInit()

	//返回
	result := "ok"
	return result, err
}

//	大屏@合约处理展示 @ 1
func sqlByConact(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlByConact()
}

//	大屏@ 智能合约调用信息表 @ 2
func sqlByLarge(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlByLarge()
}

//	大屏@ 共识节点信息表 @ 3
func sqlByServer(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlByServer()
}

//	大屏@ 节点信息统计表 @ 4
func sqlBySvg(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlBySvg()
}

//	大屏@ 区块信息表 @ 5
func sqlByBlock(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlByBlock()
}

//	大屏@ 节点信息交易列表 @ 6 DbDeal
func sqlByDeal(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlByDeal()
}

//	大屏@ 资产上链信息表 @ 7
func sqlByAsset(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlByAsset()
}

//	大屏@ 根据区块查询信息 @ 8 BlockData
func sqlQueryBlock(c *gin.Context) (rep interface{}, err error) {
	//
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.	参数序列化 || 数据集合 BlockTransaction
	data, err := unBlockTransaction(c)
	if err != nil {
		return nil, err
	}
	// TODO
	log.Println("查询到的区块高度是：", data)
	log.Println("查询到的区块高度是：", data.BlockNum)
	return model_action.SqlQueryByBlock(data.BlockNum)
}

//	大屏@ 根据区块查询关联10个区块 @ 9 BlockData
func sqlQueryBlockByHeight(c *gin.Context) (rep interface{}, err error) {

	//
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.	参数序列化 || 数据集合 BlockTransaction
	data, err := unBlockTransaction(c)
	if err != nil {
		return nil, err
	}
	return model_action.SqlQueryBlockByHeight(data.BlockNum)
}

// -- 根据交易ID 查询 数据类型
func QueryDataByBase(x string, Mg *Kyc, strData string) (label string, dataByType string, rep interface{}, err error) {
	return queryDataByBase(x, Mg, strData)
}

//	大屏@ 节点信息交易列表 @ 6 DbDeal
func sqlBySvgForAverage(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// TODO
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlBySvgForBig()
}

// 微分格 数据上链  -- 根据主键查询区块信息 2019年12月24日14:10:33
func queryWQueryData(c *gin.Context) (list interface{}, err error) {
	ChainTrans := ChainTransactionConfig{}
	resData := TxIDrenData{}
	ResInfo := ResInfo{}
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")

	//	01.	序列化数据
	data, err := mSerializeByBox(c)
	if err != nil {
		log.Println("=>微分格=>01.	拼接上链标签结构失败", err)
		return nil, err
	}
	log.Println("微分格=>01.拼接上链标签结构")

	fmt.Printf("Data:", data.Data)
	//	02. 验证用户权限
	fal := cverify(data.User.OrgID, data.User.UserID)
	if !fal {
		log.Println("=>微分格=>02.	验证用户权限失败：", err)
		return nil, err
	}
	log.Println("微分格=>02.验证用户权限")

	//	05 	获取 通道 链码 TODO  通道 和 链码校验
	cli := "cli"
	channlName := data.Chain.ChannelName
	chainCodeName := data.Chain.ChainCodeName
	funcName := data.Chain.FuncName
	// 参数拼接
	sool := "txid-"
	if len(data.Data) <= 0 {
		return nil, errors.New("data 数据为空")
	}
	data.Data[0] = sool + data.Data[0]
	args := data.Data
	uuid := data.Chain.UUID
	peer := "peer"

	log.Println("转换后的参数是：", args)

	// nil
	if channlName == "" {
		log.Println("=>微分格=>04. 通道信息名称为空	", err)
		return nil, errors.New("通道信息名称为空")
	}
	//
	if chainCodeName == "" {
		log.Println("=>微分格=>04. 链码信息名称为空	", err)
		return nil, errors.New("链码信息名称为空")
	}

	if funcName == "" {
		log.Println("=>微分格=>04. 函数名称为空	", err)
		return nil, errors.New("函数名称为空")
	}

	//
	resultID, err := model_api.QueryloadByBox(cli, channlName, chainCodeName, funcName, peer, args, uuid)
	if err != nil {
		log.Println("=>数据上链=>04.	数据上链失败", err)
		return nil, err
	}

	log.Println("result:", resultID)

	err = json.Unmarshal([]byte(resultID), &ResInfo)
	if ResInfo.Msg == "" {
		log.Println("查询的数据的区块信息为空")
		return resData, nil
	}
	content := ResInfo.Msg[1 : len(ResInfo.Msg)-1]

	result, _, err1 := model_action.App.GetBlockTxId("", "", "", "", "", content)

	//log.Println("config:", result, resultByString)
	if err1 != nil {
		log.Println("查询交易信息:", err1)
		return nil, err1
	}

	//	序列化数据
	err = json.Unmarshal([]byte(result), &ChainTrans)
	if err != nil {
		log.Println("序列化数据:", err)
		return nil, err
	}

	TXid := content
	ChainTrans.DataTxid = TXid

	log.Println("ChainTrans.TxID:", ChainTrans.DataTxid)
	resData.ChainTransactionConfig = ChainTrans

	return resData, nil
}

func CqueryBlockByTxIdTest(txid string) (rep interface{}, err error) {

	// 声明参数
	resData := TxIDrenData{}
	resDataJiu := TxIDrenDataJiu{}
	ChainTrans := ChainTransactionConfig{}
	Mg := Kyc{}
	wByType := WeiByType{}
	var strData string
	var JiuWeifen string
	channelName := "channelName"
	cliName := "cliName"
	peerName := "peerName"
	orgName := "orgName"
	userName := "userName"

	// 插入调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")

	//	04.	查询对象 TODO

	//	根据TXID查询数据集合
	result, resultByString, err := model_action.App.GetBlockTxId(channelName, cliName, peerName, orgName, userName, txid)

	log.Println("result=>", result)

	if err != nil {
		return nil, err
	}

	log.Println("resultByString (len)=>", len(resultByString))

	if len(resultByString) == 0 {
		return nil, errors.New("resultByString is nil")
	}
	if len(resultByString) < 2 {
		log.Println("微分格类型处理")
		// TODO  微分格的数据  默认第一种
		json.Unmarshal([]byte(result), &ChainTrans)
		resData.ChainTransactionConfig = ChainTrans
		return resData, nil
	}

	//	异常情况处理完毕, 开始进行数据处理
	// 统一处理  判断是否是历史数据, 加标识判断
	json.Unmarshal([]byte(result), &ChainTrans)
	// 配置信息
	resData.ChainTransactionConfig = ChainTrans

	// 如果 查询区块高于 725
	if ChainTrans.Height <= int64(800) {
		//
		JiuWeifen = "v1"
	} else {
		JiuWeifen = "v2"
	}
	ChainTrans.Version = JiuWeifen
	// TODO  数据解密
	// TODO   kyc 数据在  retstring [1]   retstring [0] 类型      retstring [2] 标签

	for k, _ := range resultByString {
		log.Println("====>", k)
	}

	//	数据转码
	err = json.Unmarshal([]byte(resultByString[1]), &Mg)

	if err != nil {
		log.Println("data transcoding failed:", err)
		return "", err
	}
	err = json.Unmarshal([]byte(resultByString[2]), &wByType)
	if err != nil {
		log.Println("data transcoding failed:", err)
		return "", err
	}
	log.Println("wByType:", wByType.CategoryId)
	ChainTrans.PledgeType = wByType.CategoryId
	// TODO  判断是哪一种数据类型

	//	数据判断
	if Mg.KycString != "" || Mg.SignKey != "" {
		// 解密
		strData = MetadataAesDecrypt(Mg.KycString, Mg.SignKey)
		log.Println("strData=>", strData)
	} else {
		// 如果数据为空,那么就确认是微分格数据, 赋值 标签 TODO
		Mg.KycType = Pro_Rec_WI
	}

	// 分发器

	switch Mg.KycType {

	//	房地产
	case Pro_Rea_All:
		label, dataType, data, err := queryDataByFang(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针:", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label
	//	应收账款
	case Pro_Rec_Ysk:
		label, dataType, data, err := queryDataByYing(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resDataJiu.Data = &data
		resDataJiu.DataType = label

		//	旧数据处理 这层必须直接返回,其他 resData 可以后续追加
		json.Unmarshal([]byte(result), &ChainTrans)
		resDataJiu.ChainTransactionConfig = ChainTrans
		return resDataJiu, nil

	//	尽调结果
	case Pro_Rec_Yjg:
		label, dataType, data, err := queryDataByBaseReport(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label

	//	尽调报告
	case Pro_Rec_Jdb:
		label, dataType, data, err := queryDataByBaseSurvey(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label

	//	动产质押
	case Pro_Rec_JCZY:
		label, dataType, data, err := queryDataByBasePle(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
			return "", err
		}

		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label

	//	微分格
	case Pro_Rec_WI:
		log.Println("微分格查询")
		label, dataType, data, err := queryDataByWeiForapp(&Mg, resultByString[1], JiuWeifen)

		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label
	default:
		log.Println("There is no corresponding type")
	}

	resData.ChainTransactionConfig = ChainTrans
	// 返回
	return resData, nil
}

//	大屏@ 地理位置列表 2019年12月26日16:21:33
func sqlBySvgForGeography(c *gin.Context) (rep interface{}, err error) {

	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.  查询数据返回
	return model_api.SqlBySvgForGeography()
}

//
func cqueryBlockDynamicList(c *gin.Context) (rep interface{}, err error) {
	channelName := "assetpublish"
	orgName := "orgName"
	userName := "userName"
	ResultChannelData := []ChannelData{}

	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")

	log.Println("channelName=>", channelName)

	channelList := getChannel()
	if len(channelList) > 0 {
		for k, _ := range channelList {
			result, err := model_action.App.GetBlockDynamicByBig(channelList[k].Channel, orgName, userName)

			if err != nil {
				log.Println("查询通道出错：", err)
				channelList[k].Data = ""
				ResultChannelData = append(ResultChannelData, channelList[k])
				continue
			}
			channelList[k].Data = result
			ResultChannelData = append(ResultChannelData, channelList[k])
		}

	}

	// 	05.	返回
	return ResultChannelData, nil
}

//	查询后五个
func cqueryNextBlockTransaction(c *gin.Context) (rep interface{}, err error) {
	//
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.	参数序列化 || 数据集合 BlockTransaction
	data, err := unBlockTransaction(c)
	if err != nil {
		return nil, err
	}

	/**
	生成 对应的 APP TODO
	*/

	//	04.	查询对象
	channelName := "assetpublish"

	cliName := "cliname"
	orgName := "orgName"
	res := data.BlockNum
	//
	log.Println("res", res)

	// channelName GetBlockDynamic  // 查询后五个
	result, err := model_action.App.GetNetBlockDynamicByBig(channelName, cliName, orgName, res)
	if err != nil {
		return nil, err
	}

	// 	05.	返回
	return result, nil
}

//	大屏@ 资产上链信息表 @ 7 后五条
func sqlByAssetNext(c *gin.Context) (rep interface{}, err error) {

	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData(

		"cName",
		"chainName",
		"funcName",
		"org",
		"userName",
	)
	go model_api.SqlBySvgForInserData(
		"txid",
		"orgName",
		"userName",
	)

	//获取参数
	//	01.  查询数据返回
	data, err := unBlockTransactionNext(c)
	if err != nil {
		log.Println("=>获取参数 拼接上链标签结构失败", err)
		return nil, err
	}

	return model_api.SqlByAssetNext(data.No, data.Type)
}

//	大屏@  复合查询   2020年1月13日19:30:33
func sqlByCompoundQuery(c *gin.Context) (rep interface{}, err error) {

	var result interface{}
	var selectType bool
	var Pg Compound

	//	01.  查询数据返回
	data, err := unBlockCompound(c)
	if err != nil {
		log.Println("=>获取参数 拼接上链标签结构失败", err)
		return nil, err
	}

	if data.Id == "" {
		log.Println("主键为空")
		return nil, errors.New("主键为空")
	}

	// 主键||交易ID|| 区块高度 || 区块hash ||

	// 交易ID
	if !selectType {
		result, err = compoundQueryByTxid(data.Id)
		if err != nil {
			log.Println("复合查询-交易ID查询失败：", err)
		} else {
			log.Println("查询- 交易ID 通过")
			selectType = true
			Pg.Compound = FhByBlockTxid
			Pg.CompoundData = result
		}
	}

	if !selectType {
		// 	- 根据区块HASH 查询数据
		result, err = compoundQueryByBlockHash(data.Id)
		if err != nil {
			log.Println("复合查询-根据区块HASH 失败：", err)
		} else {
			log.Println("查询- 根据区块HASH 通过")
			selectType = true
			Pg.Compound = FhByBlockHash
			Pg.CompoundData = result
		}
	}

	if !selectType {
		// 	- 根据区块高度查询
		result, err = compoundQueryByBlockHeight(data.Id)
		if err != nil {
			log.Println("复合查询-根据区块高度查询 失败：", err)
		} else {
			log.Println("查询- 根据区块高度查询 通过")
			log.Println("查询- 根据区块高度查询 通过", result)
			selectType = true
			Pg.Compound = FhByBlockHeight
			Pg.CompoundData = result
		}
	}

	if !selectType {
		// 	- 根据主键查询数据
		result, err = compoundQueryByID(data.Id)
		if err != nil {
			log.Println("复合查询-根据主键查询数据 失败：", err)
		} else {
			log.Println("查询- 根据主键查询数据 通过")
			selectType = true
			Pg.Compound = FhByBlockID
			Pg.CompoundData = result
		}
	}

	if !selectType {
		Pg.Compound = FhByBlockNil
		Pg.CompoundData = "nil"

	}
	return Pg, nil
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func sqlByDeaLlList(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")

	// 序列化数据
	data, err := unEcho(c)
	if err != nil || data.Type == "" {
		log.Println("序列化失败：", err)
		return nil, err
	}

	//	01.  查询数据返回
	return model_api.SqlByDealList(data.Type)
}

//	大屏@  抬头 2020年1月17日18:07:41 Ri
func sqlByRise(c *gin.Context) (rep interface{}, err error) {

	result := RiseInfo{}
	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")

	Pg, err := model_action.App.GetBlockMessage("", "", "")

	if err != nil {
		log.Println("err:", err)
	}
	// 赋值 区块高度
	result.RisBlock = Pg.ConfigHeight

	//赋值 节点数量
	result.RisNode = Pg.ConfigNodeNum

	//赋值 链码数量
	result.RisChaincode = "2"

	transactionNum, err := model_api.SqlByLarge()
	//
	if err != nil {
		log.Println("赋值交易数量:", err)
	}

	//赋值交易数量
	result.RisTransaction = transactionNum

	assetSum, err := model_api.SqlByAssetSum()
	if err != nil {
		log.Println("资产列表:", err)
	}
	//赋值 资产列表
	result.RisAsset = assetSum

	return result, nil
}

//
//	大屏@  地理位置  三个Tab  2020年1月16日09:45:26
func sqlByDealGeography(c *gin.Context) (rep interface{}, err error) {

	// 处理数据
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")

	// 序列化数据
	data, err := unEcho(c)
	if err != nil || data.Type == "" {
		log.Println("序列化失败：", err)
		return nil, err
	}

	//	01.  查询数据返回
	return model_api.SqlByDealGeography(data.Type)
}
