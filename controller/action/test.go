package action

import (
	model_action "ki-sdk/model/action"
	"ki-sdk/model/api"
	model_api "ki-sdk/model/api"
	util_action "ki-sdk/util/action"
	util_api "ki-sdk/util/api"
	"encoding/json"
	"errors"
	"github.com/chainHero/blockdata"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CupLoadTest(c *gin.Context) (list *blockdata.ChainTransaction, err error) {
	return cupLoad(c)
}

//test  1
func CupLoadtest(c *gin.Context) (list *blockdata.ChainTransaction, err error) {
	return cupLoadtest(c)
}

//测试  2
func cupLoadtest(c *gin.Context) (list *blockdata.ChainTransaction, err error) {
	//	接收参数
	data, err := cmSerialize(c)
	if err != nil {
		return nil, err
	}
	//	01.	验证用户权限
	state := cverify(data.User.OrgID, data.User.UserID)
	if !state {
		return nil, err
	}
	//	02.	序列化数据

	//	03.	元数据加密

	//	04.	拼接上链数据结构
	//	05.	拼接标签数据结构
	//	06.	数据上链结构
	//	07.	关联数据
	//	08.	推送数据
	//	09.	查询
	//	10.	返回

	////	02. 验证用户权限
	//state := cverify(data.User.OrgID, data.User.UserID)
	//if !state {
	//	return nil, err
	//}
	//	03. 获取上链参数 args 根据不同参数去请求不同的api
	args, err := mDispense(data.Datafunc, data.Datatype, *data)
	log.Println("args=>", args)
	//
	if err != nil {
		return nil, err
	}

	// 04. 数据上链
	cli := "cli"
	channlName := "AssetToChain"
	peer := "2"
	//
	txid, err := api.UploadAssetTest(cli, channlName, channlName, peer, args)
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

// 元数据加密 Metadata encryption
/**
@	描述：元数据加密
@	时间：2019年10月28日11:21:21
@	备注： 参数： 元数据  私钥
*/
func MetadataEncryption(metadata string, key string) (sign string) {
	//
	decryptCode := AesEncrypt(metadata, key)
	//decryptCode := AesDecrypt(metadata, key)
	return decryptCode
}

//  AesDecrypt  解码
func MetadataAesDecrypt(signtxt, key string) (metadata string) {
	//
	decryptCode := AesDecrypt(signtxt, key)
	return decryptCode
}

/**
@	描述： 获取私钥
@	参数：
*/

func getEncryptionKeytest(orgKey string, userKey string, publicKey string) (SnorgKey string, SuserKey string, SpublicKey string) {
	//
	if orgKey != "" {
		SnorgKey = orgKey[0:16]
	}
	if orgKey != "" {
		SuserKey = userKey[0:16]
	}
	if orgKey != "" {
		SpublicKey = publicKey[0:16]
	}
	// 返回
	return SnorgKey, SuserKey, SpublicKey
}

/**
@	描述：拼接上链数据
@	时间：2019年10月28日14:07:25
@	备注：

// 主键
// 数据类型
// 上链时间
// 内容密文
// 签名key
// 访问权限
//签名用户信息
*/
func concatenateTheUpperLinkDataStructure(pkey string, kycType string, kycTime string, kycString string, signKey string, signPower string, signUser User) (putData string, err error) {
	//
	KycData := Kyc{}
	// 主键

	// TODO  主键重要
	if pkey == "" {
		return "", errors.New("no pkey")
	} else {
		KycData.KycID = pkey
	}
	// 数据类型
	if kycType == "" {
		return "", errors.New("no kycType")
	} else {
		KycData.KycType = kycType
	}
	// 上链时间
	if kycTime == "" {
		return "", errors.New("no kycTime")
	} else {
		KycData.KycTime = kycTime
	}
	// 内容密文
	if kycString == "" {
		return "", errors.New("no kycString")
	} else {
		KycData.KycString = kycString
	}
	// 签名秘钥
	if signKey == "" {
		return "", errors.New("no signKey")
	} else {
		KycData.SignKey = signKey
	}
	// 访问权限
	if signPower == "" {
		return "", errors.New("no signPower")
	} else {
		KycData.SignPower = signPower
	}
	// 用户信息
	if &signUser == nil {
		return "", errors.New("no signUser")
	} else {
		KycData.SignUser.UseName = signUser.UserName
		KycData.SignUser.UseID = signUser.UserID
		KycData.SignUser.UseOrgName = signUser.OrgName
		KycData.SignUser.UseOrgID = signUser.OrgID
		KycData.SignUser.UseCa = signUser.AffiliationId
		KycData.SignUser.UseType = signUser.UserName //TODO  备用字段，用来处理用户类型
	}
	//
	UnmarData, err := json.Marshal(KycData)
	if err != nil {
		return "", errors.New("UnmarData is err")
	}
	return string(UnmarData), nil
}

/**
@	函数： 生成标签结构
@	描述：
@	时间： 2019年10月28日15:19:46
@	备注：
*/
func spliceTagDataStructures(use *Use, mes *[]BycMes, sign *Sign, key string) (bc string, err error) {
	// nil

	if use == nil || mes == nil || sign == nil {
		return "", errors.New("标签数据为空")
	}

	// 拼接数据
	resData := &Byc{
		BycUser: *use,
		BycMes:  *mes,
		BycSign: *sign,
		BycID:   key,
	}

	//
	res, err := json.Marshal(resData)
	if err != nil {
		return "", err
	}
	//
	return string(res), nil
}

//

//BycByKey     string     `json:"bycByKey"`     // 宿主ID
//BycByfMaster string     `json:"bycByfMaster"` // 关联宿主
//BycBySpan    string     `json:"span"`         // 是否跨链码
func testspliceTagDataStructures(use *Use, mes *[]BycMes, sign *Sign, key string, bycByKey string, bycByfMaster string, span string) (bc string, err error) {
	// nil

	if use == nil || mes == nil || sign == nil {
		return "", errors.New("标签数据为空")
	}

	// 拼接数据
	resData := &Byc{
		BycUser:      *use,
		BycMes:       *mes,
		BycSign:      *sign,
		BycID:        key,
		BycByKey:     bycByKey,
		BycByfMaster: bycByfMaster,
		BycBySpan:    span,
	}

	//
	res, err := json.Marshal(resData)
	if err != nil {
		return "", err
	}
	//
	return string(res), nil
}

//
// 微分格数据上链  2019年11月7日18:55:44
func T_UpLoadByBox(c *gin.Context) {
	//
	// 数据上链
	list, err := UpLoadByBox(c)

	//  err
	if err != nil {
		c.JSON(http.StatusOK,
			gin.H{
				"status": StatusText(StatusFailed),
				"data":   StatusText(StatusContinue),
			})
		return
	} else {
		// OK
		c.JSON(
			http.StatusOK,
			gin.H{"status": StatusText(StatusOK),
				"data": list,
			})
		return
	}
}

// wei -- test
func T_caddOrg(c *gin.Context) (rep interface{}, err error) {
	// 次数 和 api  调用次数
	go model_api.SqlByLargeForInserData("cName", "chainName", "funcName", "org", "userName")
	go model_api.SqlBySvgForInserData("txid", "orgName", "userName")
	//	01.	参数序列化 || 数据集合 NewAffiliation
	data, err := unNewAffiliation(c)
	log.Println("01 参数序列化", data)
	if err != nil {
		return "", err
	}

	//TODO
	caName := "ca-org2-msp.dinglian.com"
	caPath := "com.dinglian.org2."
	focle := true
	affilName := data.AffiliationID

	// 添加联盟
	log.Println("添加联盟")
	result, err1 := model_action.App.AddAffiliation(caName, focle, caPath, affilName)
	log.Println(result.CAName)
	if err1 != nil {
		log.Println("注册联盟失败", result.CAName)
		return "注册联盟失败", err1
	}

	log.Println("注册成功")
	//	生成 msp cli

	// TODO
	num := "com.dinglian.org2."
	mspClient, err := model_action.App.CreateNewMspClient(data.AffiliationID, num)
	if err != nil {
		//is nil  返回
		return "注册联盟失败", err
	}
	log.Println("mspClient cli si ok ")

	// 生成   userFFID
	hash := util_api.GetcurrentHash(util_action.EncryptText(util_action.EncryptDsha256), data.OrgID+data.UserId)
	//	添加联盟管理员  || 添加组织
	log.Println("hash=>", hash)
	_, err = model_action.App.MspAffCreateUser(mspClient, data.AffiliationID, hash, data.UserName, data.OrgName, data.OrgID, "com.dinglian.org2.", "123456")
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

func T_setUser(c *gin.Context) (rep string, err error) {

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
	num := "com.dinglian.org2."
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
	mspUser, err := model_action.App.MspAffCreateUser(mspClient, data.RegAffiliationId, hash, data.UserName, data.RegOrgName, data.RegOrgID, "com.dinglian.org2.", "123456")
	// err
	if err != nil {
		log.Println("=>05. 注册用户失败:", err)
		return "", err
	}

	//TODO
	log.Println("=>05. 注册用户")

	//
	return mspUser, nil
}

//	测试- 新链码
func TestUpload(c *gin.Context) (rep interface{}, err error) {
	// 数据上链
	return testUpload(c)
}

//	测试- 新链码
func TestSelect(c *gin.Context) (rep interface{}, err error) {
	// 数据上链
	return testSelect(c)
}

//	测试新链码 --
func testUpload(c *gin.Context) (list *blockdata.ChainTransactionup, err error) {

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

	//
	args, state := testUploadProjectReport(*data)

	if state != nil {
		log.Println("=>数据上链=>03.	获取上链参数失败：", err)
		return nil, err
	}
	log.Println("=>数据上链=>03.	获取上链参数")

	txid, err := model_api.TestAssettest("", "", "", "", args)
	if err != nil {
		log.Println("=>数据上链=>04.	数据上链失败", err)
		return nil, err
	}

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

	return rest[0], err
}

//	测试新链码 --
func testSelect(c *gin.Context) (list *blockdata.ChainTransactionup, err error) {

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

	//
	//args, state := testUploadProjectReport(*data)

	log.Println("=>数据上链=>03.	获取上链参数")
	arr := []string{}
	arr = append(arr, "getHistory")
	arr = append(arr, Mas+data.User.UUID)
	txid, err := model_api.TestSelectChaincode("", "", "", "", arr)
	if err != nil {
		log.Println("=>数据上链=>04.	数据上链失败", err)
		return nil, err
	}

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

	return rest[0], err
}

//
func testUploadProjectReport(Info Info) (list []string, err error) {

	//	01. 数据
	Pg := BaseReport{}

	//	02.	序列化数据
	MarData, err := json.Marshal(&Info.Data)

	log.Println("序列化参数成功")
	// 03. 判断错误
	if err != nil {
		log.Println("序列化参数成功失败")
		return nil, err
	}
	//	04.	元数据操作
	err = json.Unmarshal(MarData, &Pg)
	if err != nil {
		log.Println("尽调报告||元数据操作序列化失败=>", err)
		return nil, err
	}

	strcData, err := json.Marshal(&Pg)
	if err != nil {
		log.Println("尽调报告||06. 数据加密=>", err)
		return nil, err
	}

	//	06.01 获取私钥以及类型
	resKey, useKey, pubKey := getEncryptionKeytest(Info.User.OrgID, Info.User.UserID, Info.User.OrgID+Info.User.UserID)

	//	06.02  元数据加密
	log.Println("尽调报告||获取私钥以及类型 元数据加密=>", resKey, useKey, pubKey)

	//	06.03  元数据加密执行  TODO 默认是用户ID 进行加密
	cipher := MetadataEncryption(string(strcData), useKey)
	log.Println("上链数据数据加密")
	// uuid TODO  uuid 目前是资产包的ID
	uuid := Pg.ReportNo

	//	07.	拼接上链数据结构
	updata, err := concatenateTheUpperLinkDataStructure(uuid, Info.Datatype, Info.User.Datatime, cipher, useKey, signPower04, Info.User)
	if err != nil {
		log.Println("尽调报告||err：07.拼接上链数据结构=>", err)
		return nil, err
	}
	//
	log.Println("尽调报告||07.拼接上链数据结构通过")

	//	08.	拼接标签数据结构

	//拼接 单个数据结合 Byc

	//  用户信息
	sUse := Use{
		Info.User.UserName,
		Info.User.UserID,
		Info.User.OrgName,
		Info.User.OrgID,
		Title_Parent, //TODO  默认是用户类型 ，  管理员 用户 访客  群众
		Info.User.AffiliationId,
	}

	log.Println("尽调报告||用户信息=>", &sUse)

	//  关联数据信息
	sBycMesList := []BycMes{}

	sBycMes := BycMes{
		uuid,
		useKey,
		Info.User.UUID,
		uuid,
		Title_Parent, // TODO  关联关系 type ，数据默认关系是父关系
		State_Yes,
	}
	sBycMesList = append(sBycMesList, sBycMes)

	log.Println("尽调报告||关联数据信息=>", &sUse)

	//	签名信息
	sSign := Sign{
		resKey,
		useKey,
		pubKey,
		useKey, //TODO
	}
	log.Println("尽调报告||签名信息=>", &sSign)

	// 拼接标签信息
	upBiao, err := testspliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid, Info.User.BycByKey, Info.User.BycByfMaster, Info.User.Span)
	if err != nil {
		log.Println("尽调报告||拼接标签信息=>", err)
		return nil, err
	}

	log.Println("尽调报告||拼接上链标签结构通过")

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链标签
	*/
	return util_api.GetSliceArgs(TYUPLOADASSET, KYC, updata, upBiao)
}
