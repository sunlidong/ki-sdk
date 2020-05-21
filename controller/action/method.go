package action

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	model_action "ki-sdk/model/action"
	model_api "ki-sdk/model/api"
	"ki-sdk/util/action"
	"ki-sdk/util/api"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/**
@参数序列化
@时间 ：2019年10月14日15:51:13
@No  011
*/
func opt(dataName string) interface{} {
	//
	var assetStruct interface{}
	switch dataName {

	case LabelRequesRealty:
		assetStruct = &Info{}
		return assetStruct

	case LabelRequestReceivables:
		assetStruct = &Info{}
		return assetStruct
	//	web
	case WebIndex:
		assetStruct = &BlockWordList{}
		return assetStruct

	default:
		log.Println("参数错误")
	}
	return nil
}

// 序列化  数据
func optSerialize(c *gin.Context, data interface{}) (cdata interface{}, code string) {

	err := c.ShouldBindJSON(&data)
	if err != nil {
		//is nil  返回
		//return nil, util(int64(100))
		return nil, ""
	}
	//return data, util.StatusText(int64(0))
	return nil, ""
}

/*
@ 验证请求权限
@ 2019年10月18日10:08:23
@lidong sun
*/
func verify(orgId string, userId string) bool {
	//
	hash := api.GetcurrentHash(action.EncryptText(action.EncryptDsha256), orgId+userId)
	userFile, err := model_action.App.VifUserFull(hash)
	//
	if err != nil {
		log.Println("验证出错")
		return false
	}
	log.Println("验证用户信息=>", userFile)
	if err != nil {
		return false
	} else {
		return true
	}
}

/**
@ 数据上链
@ 2019年10月18日13:50:37
@ lidongsun
@
*/
func upload(data interface{}) {
	// 01. 序列化数据

	//	02.	添加类型
	//	03. 添加标签
	//	04.	数据上链
	//	05.	返回结果

}

/**
@ 序列化数据
@
*/
func mSerialize(c *gin.Context) (data *Info, err error) {
	var cData Info
	// 02. 将参数序列化到结构体 ShouldBind
	if err := c.ShouldBindJSON(&cData); err != nil {
		return &cData, err
	}
	return &cData, nil
}

//
func mSerializeByBox(c *gin.Context) (data *InfoWei, err error) {
	var cData InfoWei
	// 02. 将参数序列化到结构体 ShouldBind
	if err := c.ShouldBindJSON(&cData); err != nil {
		return &cData, err
	}
	return &cData, nil
}

/**
@  分游器
@ 2019年10月18日14:32:06
@ lidongsun
@
*/
func mDispense(funcName string, dataType string, info Info) (list []string, err error) {
	// 请求类型
	switch funcName {
	//新增
	case FuncNew:
		//数据类型
		switch dataType {

		//	房地产数据结构 主表 基础表(一次共同上传)
		case DataDueProjectAll:
			return uploadProjectAll(info)

			//	房地产数据结构表  主表(单次上传)
		case DataDueProject:
			return uploadProject(info)

			//	房地产结构表
		case DataDueProjectFdc:
			return uploadDueProjectFdc(info)

			//	保证人结构表
		case DataDueProjectCemGt:
			return uploadDueProjectCemGt(info)

			//	抵押品集合
		case DataDueProjectCemMg:
			return uploadDueProjectCemMg(info)

			//	质押品集合
		case DataDueProjectCemPg:
			return uploadDueProjectCemPg(info)

			//	 应收账款

			//	钢包

			//	测试
		case Test:
			return uploadDueProjectTest(info)

			//
		case Test12039:
			return uploadTest(info)

			// 找不到方法
		default:
			return nil, err
		}

		//更新
	case FuncUpdate:
		switch dataType {

		//	整体房地产上链 更新
		case DataDueProjectAll:
			return updateProject(info)
		case DataDueProject:
			return uploadProject(info)
		case DataDueProjectFdc:
			return uploadProject(info)
		case DataDueProjectCemGt:
			return uploadProject(info)
		case DataDueProjectCemMg:
			return uploadProject(info)
		case DataDueProjectCemPg:
			return uploadProject(info)
		default:
			log.Println("更新方法没有")
		}
		//查询
	case FuncQuery:
		switch dataType {
		case DataDueProject:
			return uploadProject(info)
		case DataDueProjectFdc:
			return uploadProject(info)
		case DataDueProjectCemGt:
			return uploadProject(info)
		case DataDueProjectCemMg:
			return uploadProject(info)
		case DataDueProjectCemPg:
			return uploadProject(info)
		default:
			log.Println("1")

		}
		//没有查到
	default:
		log.Println("1")
	}
	return nil, err
}

//	V3 上链分发器
func DispenseV3(funcName string, dataType string, info Info) (list []string, err error) {
	// 请求类型
	switch funcName {
	//新增
	case FuncNew:
		//数据类型
		switch dataType {

		//	房地产数据结构 主表 基础表(一次共同上传)
		case Pro_Rea_All:
			return upLoadByAll(info)
			//return upLoadByAllSign(info)

			//	房地产数据结构表  主表(单次上传)
		case Pro_Rea_One:
			return upLoadByProject(info)

			//	房地产结构表
		case Pro_Rea_Fdc:
			return uploadDueProjectFdc(info)

			//	保证人结构表
		case Pro_Rea_Cgt:
			return uploadDueProjectCemGt(info)

			//	抵押品集合
		case Pro_Rea_Cmg:
			return uploadDueProjectCemMg(info)

			//	质押品集合
		case Pro_Rea_Cpg:
			return uploadDueProjectCemPg(info)

			//	 ---------------------------------- 应收账款

			// 应收账款
		case Pro_Rec_Ysk:
			return uploadProjectReceivables(info)

			// 尽调结果
		case Pro_Rec_Yjg:
			return uploadProjectFindings(info)

			// 尽调报告
		case Pro_Rec_Jdb:
			return uploadProjectReport(info)

			//	---------------------------------------钢包 TODO

			//	测试
		case Test:
			return uploadDueProjectTest(info)

		case "100312":
			return uploadTest(info)

			//	 ---------------------------------- 动产质押
		case Pro_Rec_JCZY:
			return uploadProjectByPle(info)
			// 找不到方法
		default:
			return nil, err
		}

		//更新
	case FuncUpdate:
		switch dataType {

		//	整体房地产上链 更新
		case DataDueProjectAll:
			return updateProject(info)
		case DataDueProject:
			return uploadProject(info)
		case DataDueProjectFdc:
			return uploadProject(info)
		case DataDueProjectCemGt:
			return uploadProject(info)
		case DataDueProjectCemMg:
			return uploadProject(info)
		case DataDueProjectCemPg:
			return uploadProject(info)
		default:
			log.Println("1")
		}
		//查询
	case FuncQuery:
		switch dataType {
		case DataDueProject:
			return uploadProject(info)
		case DataDueProjectFdc:
			return uploadProject(info)
		case DataDueProjectCemGt:
			return uploadProject(info)
		case DataDueProjectCemMg:
			return uploadProject(info)
		case DataDueProjectCemPg:
			return uploadProject(info)
		default:
			log.Println("1")

		}
		//没有查到
	default:
		log.Println("1")
	}
	return nil, err
}

/**
@ 主表上链
@ 2019年10月18日15:54:09
@
*/
func uploadProjectAll(Info Info) (list []string, err error) {

	// 01. 获取结构体对象
	data := RealtyInfo{}

	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(&Info.Data)

	// 03. 判断错误
	if err != nil {
		return nil, err
	}

	// 04. 将参数反序列化到结构体
	json.Unmarshal(MarData, &data)
	// 追加  time  id  以及标签
	data.FBase.Ptype = Info.Datatype
	log.Println("data key =>", data.FDueProject.FID)
	// 05. 将参数序列化到结构体
	strcData, err := json.Marshal(&data)
	//	06.	拼接参数
	return api.GetSliceArgs(ChainUploadAsset, Info.Datatype, string(strcData), data.FDueProject.FID)
}

/**
@ 主表上链
@ 2019年10月18日15:54:09
@
*/

// 主表上链
func uploadProject(Info Info) (list []string, err error) {
	// 01. 获取结构体对象
	data := DueProject{}
	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(&Info.Data)
	// 03. 判断错误
	if err != nil {
		return nil, err
	}
	// 04. 将参数反序列化到结构体
	json.Unmarshal(MarData, &data)
	// 追加  time  id  以及标签
	data.FBase.Ptype = Info.Datatype
	// 05. 将参数序列化到结构体
	strcData, err := json.Marshal(data)
	//	06.	拼接参数
	return api.GetSliceArgs(ChainUploadAsset, Info.Datatype, string(strcData), Info.User.Datatime)
}

/**
@ 房地产 数据集上链
@ 2019年10月21日14:33:16
@ lidongsun
*/
func uploadDueProjectFdc(Info Info) (list []string, err error) {
	// 01. 获取结构体对象 || 房地产 数据集合上链
	data := DueProjectFdc{}
	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(&Info.Data)
	// 03. 判断错误
	if err != nil {
		return nil, err
	}
	// 04. 将参数反序列化到结构体
	json.Unmarshal(MarData, &data)
	// 追加  time  id  以及标签  || 房地产标识
	data.FBase.Ptype = Info.Datatype
	// 05. 将参数序列化到结构体
	strcData, err := json.Marshal(data)
	//	06.	拼接参数

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链时间
	*/
	return api.GetSliceArgs(ChainUploadAsset, Info.Datatype, string(strcData), Info.User.Datatime)
}

/**
@ 保证人 数据集上链
@ 2019年10月21日14:38:34
@ lidongsun
*/
type Hero struct {
	Name     string `json:"hero_name"` //起别名为：hero_name
	Age      int    `json:"hero_age"`
	Birthday string
	Sal      float64
	Skill    string
}

func uploadDueProjectCemGt(Info Info) (list []string, err error) {
	// 01. 获取结构体对象 || 保证人 数据集合上链
	data := DueProjectCemGt{}
	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(Info.Data)
	// 03. 判断错误
	if err != nil {
		return nil, err
	}
	log.Println("1序列化后数据是：", string(MarData))

	// 04. 将参数反序列化到结构体
	json.Unmarshal(MarData, &data)
	// 追加  time  id  以及标签  || 保证人 标识
	data.FBase.Ptype = Info.Datatype

	//将monster序列化
	mraData, err := json.Marshal(&data)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("序列化后=%v\n", string(mraData))
	//	06.	拼接参数

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链时间
	*/
	log.Println("pkey", data.FID)
	return api.GetSliceArgs(ChainUploadAsset, Info.Datatype, string(mraData), Info.User.Datatime)
}

/**
@ 抵押品 数据集上链
@ 2019年10月21日14:38:34
@ lidongsun
*/
func uploadDueProjectCemMg(Info Info) (list []string, err error) {
	// 01. 获取结构体对象 || 抵押品 数据集合上链
	data := DueProjectCemMg{}
	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(&Info.Data)
	// 03. 判断错误
	if err != nil {
		return nil, err
	}
	// 04. 将参数反序列化到结构体
	json.Unmarshal(MarData, &data)
	// 追加  time  id  以及标签  || 抵押品 标识
	data.FBase.Ptype = Info.Datatype
	// 05. 将参数序列化到结构体
	strcData, err := json.Marshal(data)
	//	06.	拼接参数

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链时间
	*/
	return api.GetSliceArgs(ChainUploadAsset, Info.Datatype, string(strcData), Info.User.Datatime)
}

/**
@ 质押品数据集上链
@ 2019年10月21日14:38:34
@ lidongsun
*/
func uploadDueProjectCemPg(Info Info) (list []string, err error) {
	// 01. 获取结构体对象 || 质押品 数据集合上链
	data := DueProjectCemPg{}
	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(&Info.Data)
	// 03. 判断错误
	if err != nil {
		return nil, err
	}
	// 04. 将参数反序列化到结构体
	json.Unmarshal(MarData, &data)
	// 追加  time  id  以及标签  || 质押品 标识
	data.FBase.Ptype = Info.Datatype
	// 05. 将参数序列化到结构体
	strcData, err := json.Marshal(data)
	//	06.	拼接参数

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链时间
	*/
	return api.GetSliceArgs(ChainUploadAsset, Info.Datatype, string(strcData), Info.User.Datatime)
}

func uploadDueProjectTest1(Info Info) (list []string, err error) {
	// 01. 获取结构体对象 || 质押品 数据集合上链
	data := DueProjectCemPg{}
	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(&Info.Data)
	// 03. 判断错误
	if err != nil {
		return nil, err
	}
	// 04. 将参数反序列化到结构体
	json.Unmarshal(MarData, &data)
	// 追加  time  id  以及标签  || 质押品 标识
	data.FBase.Ptype = Info.Datatype
	//
	// 05. 将参数序列化到结构体
	strcData, err := json.Marshal(data)
	//	06.	拼接参数

	// 数据序列化

	// 声明上链结构体
	Kyc := Kyc{}

	Kyc.KycID = data.FID
	Kyc.SignUser.UseID = "张三"
	// key
	Singkey := Info.User.UserID
	sskey := Singkey[0:16]
	//
	log.Println("ssk", sskey)
	//key := "dj9d9d9d9d9d9d9d"

	//log.Println("原文：", orig)
	encryptCode := AesEncrypt(string(strcData), sskey)
	log.Println("密文：", encryptCode)
	decryptCode := AesDecrypt(encryptCode, sskey)
	log.Println("解密结果：", decryptCode)

	//sdsd := "dj9d9d9d9d9d9d9d"
	//key 数据加密
	//skystring := kycapi.Kyc(kycaction.KYC_DES_E, sdsd, string(strcData))
	// 序列化   上链
	Kyc.KycString = encryptCode
	Kyc.KycType = Info.User.Anchor
	//赋值 父子关系
	if Info.User.UUID != "" {
		//Kyc..Sun = append(Kyc.KycMes.Sun, Sun{SID: Info.User.UUID, Stype: "3", Skey: "sss"})
	}
	//
	fdata, err := json.Marshal(Kyc)

	// 拼接结构体

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链时间
	*/
	dataType111 := "1"
	return api.GetSliceArgs(ChainUploadAsset, dataType111, string(fdata), Info.User.Datatime)
}

// test
func uploadDueProjectTest(Info Info) (list []string, err error) {

	//
	Pg := DueProjectCemPg{}
	//	02.	序列化数据
	//
	log.Println("Info=Data>", Info.Data)
	log.Println("Info=Datatype>", Info.Datatype)
	log.Println("Info=Datafunc>", Info.Datafunc)
	// 将参数序列化到结构体
	MarData, err := json.Marshal(&Info.Data)

	log.Println("序列化参数成功")
	// 03. 判断错误
	if err != nil {
		log.Println("序列化参数成功失败")
		return nil, err
	}
	//*****************************************************元数据操作
	// 04. 将参数反序列化到结构体
	err = json.Unmarshal(MarData, &Pg)
	if err != nil {
		log.Println("err 04：>", err)
		return nil, err
	}
	// 追加  time  id  以及标签
	Pg.FBase.Ptype = Info.Datatype

	//
	log.Println("data=>", Pg)
	log.Println("data=>", Pg.FID)
	//***************************************************** 数据加密
	// 05. 将参数序列化到结构体
	strcData, err := json.Marshal(&Pg)
	if err != nil {
		return nil, err
	}
	//获取私钥以及类型
	uuid := Pg.FID
	log.Println("将参数序列化到结构体=>", Pg.FID)
	resKey, useKey, pubKey := getEncryptionKeytest(Info.User.OrgID, Info.User.UserID, Info.User.OrgID+Info.User.UserID)
	//	03.	元数据加密
	log.Println("元数据加密=>", resKey, pubKey)

	//	数据加密
	cipher := MetadataEncryption(string(strcData), useKey)
	log.Println("数据加密", cipher)

	//	04.	拼接上链数据结构
	updata, err := concatenateTheUpperLinkDataStructure(Pg.FID, Info.Datatype, Info.User.Datatime, cipher, useKey, signPower04, Info.User)
	if err != nil {
		log.Println("err=>", err)
		return nil, err
	}

	log.Println("拼接上链数据结构")
	//
	log.Println("上链数据是=>", updata)
	//	05.	拼接标签数据结构

	//拼接 单个数据结合
	//Byc{}

	sUse := Use{
		Info.User.UserName,
		Info.User.UserID,
		Info.User.OrgName,
		Info.User.OrgID,
		Info.User.Anchor, //TODO  type
		Info.User.AffiliationId,
	}
	log.Println("", &sUse)
	//
	sBycMesList := []BycMes{}

	sBycMes := BycMes{
		Pg.FID,
		useKey,
		Info.User.UUID,
		Pg.FID,
		Info.User.Anchor,
		State_Yes,
	}
	sBycMesList = append(sBycMesList, sBycMes)
	//
	sSign := Sign{
		resKey,
		useKey,
		pubKey,
		useKey, //TODO
	}
	log.Println("", &sSign)
	//

	//BycData:=Byc{}
	upBiao, err := spliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid)
	if err != nil {
		return nil, err
	}
	//
	log.Println("拼接上链标签结构")

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链时间
	*/
	return api.GetSliceArgs(TYUPUPDATEASSET, "KYC", updata, upBiao)
}

// test
func uploadTest(Info Info) (list []string, err error) {
	DataKyc := Kyc{}
	//
	DataKyc.KycID = Info.User.UUID
	//

	ByteMar, err := json.Marshal(&DataKyc)

	//
	if err != nil {
		return nil, err
	}
	log.Println("=>", string(ByteMar))
	//
	return api.GetSliceArgs(TYUPUPDATEASSET, "KYC", string(ByteMar), string(ByteMar))
}

//
func queryBlockInfo() (text string) {
	//
	return "1"
}

/**
@	函数：	浏览器首页展示
@	时间：	2019年10月21日16:57:31
@	描述：	返回 首页结构体数据
*/
func unBlockWord(c *gin.Context) (cdata *BlockWordList, err error) {

	//	01. 浏览器展示结构体
	data := BlockWordList{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

/**
@	函数：	浏览器区块配置结构体序列化
@	时间：	时间
@	描述：	返回 首页结构体数据
*/
func unChannelBlockConfig(c *gin.Context) (cdata *ChannelBlockConfig, err error) {

	//	01. 浏览器展示结构体
	data := ChannelBlockConfig{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

/**
@	函数：model   区块链浏览器 - 根据TxID 查询
@	时间：2019年10月21日17:49:20
@	描述： model 二级函数  QueryBlockByTxId
*/
func unBlockQuery(c *gin.Context) (cdata *BlockQuery, err error) {

	//	01. 浏览器展示结构体
	data := BlockQuery{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

/**
@	函数： 区块链浏览器 - 查询通道动态信息
@	时间： 2019年10月21日19:57:15
@	描述:	通道ID
*/
func unBlockDynamic(c *gin.Context) (cdata *ChannelConfigRe, err error) {

	//	01. 浏览器展示结构体
	data := ChannelConfigRe{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

/**
@	函数： 区块链浏览器 - 一个区块的交易
@	时间： 2019年10月21日20:26:14
@	描述:	区块高度
*/
func unBlockTransaction(c *gin.Context) (cdata *BlockTransaction, err error) {

	//	01. 浏览器展示结构体 ||一个区块的交易
	data := BlockTransaction{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

/**
@	函数： 联盟组织	- 添加组织
@	时间： 2019年10月22日10:11:52
@	描述:  对于联盟组织进行注册  ID, USERID NewAffiliation
*/
func unNewAffiliation(c *gin.Context) (cdata *NewAffiliation, err error) {

	//	01. 浏览器展示结构体 ||添加组织
	data := NewAffiliation{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

/**
@	函数： 联盟组织	- 添加用户
@	时间： 2019年10月22日11:01:39
@	描述:  对用户进行注册
*/
func unRegistered(c *gin.Context) (cdata *Registered, err error) {

	//	01. 展示结构体 ||添加用户
	data := Registered{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

/**
@ 主表上链
@ 2019年10月18日15:54:09
@
*/
func updateProject(Info Info) (list []string, err error) {
	// 01. 获取结构体对象
	data := RealtyInfo{}
	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(&Info.Data)
	// 03. 判断错误
	if err != nil {
		return nil, err
	}
	// 04. 将参数反序列化到结构体
	json.Unmarshal(MarData, &data)
	// 追加  time  id  以及标签
	data.FBase.Ptype = Info.Datatype
	// 05. 将参数序列化到结构体
	strcData, err := json.Marshal(data)
	//	06.	拼接参数
	return api.GetSliceArgs(ChainUpDataAsset, Info.Datatype, string(strcData), data.FDueProject.FID)
}

//-----------------------------------//
func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	//
	log.Println("数据加密成功")
	return base64.StdEncoding.EncodeToString(cryted)
}

func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	//
	log.Println("length=>", length)
	//if length-1>0 {
	//	unpadding := int(origData[length-1])
	//}
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 查询分发器
func qDispense(funcName string, dataType string, info Info) (list []string, err error) {

	switch funcName {
	//	查询 list
	case QueryDataListByID:
		return queryDataListById(info)
	default:
		return nil, nil
	}
}

//	查询 list
func queryDataListById(info Info) (arr []string, err error) {
	//	01.	查询序列结构
	data := UseDataById{}
	// 02. 将参数序列化到结构体
	MarData, err := json.Marshal(&info.Data)
	// 03. 判断错误
	if err != nil {
		return nil, err
	}
	// 04. 将参数反序列化到结构体
	err = json.Unmarshal(MarData, &data)

	//
	if err != nil {
		log.Println("err=>", err)
		return nil, nil
	}
	log.Println("=>", data.ById)
	log.Println("=>", data.ByTime)
	//	05.	PG
	return api.GetSliceArgsById(GetAssetList, data.ById)
}

//	房地产数据整体上链 V3
func upLoadByAll(Info Info) (list []string, err error) {

	//	01. 房地产数据结构体
	Pg := RealtyInfo{}

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
		log.Println("元数据操作序列化失败=>", err)
		return nil, err
	}

	//	05. 追加  time  id 以及标签
	Pg.FBase.Ptype = Info.Datatype

	// 06. 数据加密
	strcData, err := json.Marshal(&Pg)
	if err != nil {
		log.Println("err:06. 数据加密=>", err)
		return nil, err
	}

	//	06.01 获取私钥以及类型
	resKey, useKey, pubKey := getEncryptionKeytest(Info.User.OrgID, Info.User.UserID, Info.User.OrgID+Info.User.UserID)

	//	06.02  元数据加密
	log.Println("获取私钥以及类型 元数据加密=>", resKey, useKey, pubKey)

	//	06.03  元数据加密执行  TODO 默认是用户ID 进行加密
	cipher := MetadataEncryption(string(strcData), useKey)
	log.Println("上链数据数据加密")

	// uuid TODO  uuid 目前是资产包的ID
	uuid := Pg.FDueProject.FID

	//TODO SQL  添加  资产合约

	//	07.	拼接上链数据结构
	updata, err := concatenateTheUpperLinkDataStructure(uuid, Info.Datatype, Info.User.Datatime, cipher, useKey, signPower04, Info.User)
	if err != nil {
		log.Println("err：07.拼接上链数据结构=>", err)
		return nil, err
	}
	//
	log.Println("07.拼接上链数据结构通过")

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

	log.Println("用户信息=>", &sUse)

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

	log.Println("关联数据信息=>", &sUse)

	//	签名信息
	sSign := Sign{
		resKey,
		useKey,
		pubKey,
		useKey, //TODO
	}
	log.Println("签名信息=>", &sSign)

	// 拼接标签信息
	upBiao, err := spliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid)
	if err != nil {
		log.Println("拼接标签信息=>", err)
		return nil, err
	}

	log.Println("拼接上链标签结构通过")

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链标签
	*/

	//	调用链码
	return api.GetSliceArgs(TYUPLOADASSET, KYC, updata, upBiao)
}

//
//	房地产数据上链 V3
func upLoadByProject(Info Info) (list []string, err error) {
	//	01. 房地产数据结构体
	Pg := DueProject{}

	//	02.	序列化数据
	MarData, err := json.Marshal(&Info.Data)
	// 03. 判断错误
	if err != nil {
		log.Println("=>01.	序列化数据失败：", err)
		return nil, err
	}

	log.Println("=>01.	序列化数据")

	//	04.	元数据操作
	err = json.Unmarshal(MarData, &Pg)
	if err != nil {
		log.Println("=>02.	元数据操作序列化失败", err)
		return nil, err
	}
	log.Println("=>02.	元数据操作序列化")

	//	05. 追加  time  id 以及标签
	Pg.FBase.Ptype = Info.Datatype
	log.Println("=>03.	追加数据标签")
	// 06. 数据加密
	strcData, err := json.Marshal(&Pg)
	if err != nil {
		log.Println("=>04.	数据加密序列化失败：", err)
		return nil, err
	}
	log.Println("=>04.	数据加密序列化")
	//	06.01 获取私钥以及类型
	resKey, useKey, pubKey := getEncryptionKeytest(Info.User.OrgID, Info.User.UserID, Info.User.OrgID+Info.User.UserID)

	//	06.02  元数据加密
	fmt.Printf("=>05. 获取私钥以及类型：组织key:%s,用户key:%s,公共key:%s", resKey, useKey, pubKey)

	//	06.03  元数据加密执行  TODO 默认是用户ID 进行加密

	// 判断 是用什么加密数据 TODO

	cipher := MetadataEncryption(string(strcData), useKey)

	log.Println("上链数据数据加密")

	// uuid TODO  uuid 目前是资产包的ID
	uuid := Pg.FID
	if uuid == "" {
		log.Println("=>07.	获取UUID是空")
		return nil, errors.New("主键为空")
	}
	log.Println("=>07.	获取UUID是:", uuid)
	//	07.	拼接上链数据结构
	updata, err := concatenateTheUpperLinkDataStructure(uuid, Info.Datatype, Info.User.Datatime, cipher, useKey, signPower04, Info.User)
	if err != nil {
		log.Println("=>08.	拼接上链数据结构:", err)
		return nil, err
	}
	//
	log.Println("=>08.	拼接上链数据结构")
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

	log.Println("用户信息=>", &sUse)

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

	log.Println("关联数据信息=>", &sUse)

	//	签名信息
	sSign := Sign{
		resKey,
		useKey,
		pubKey,
		useKey, //TODO
	}
	log.Println("签名信息=>", &sSign)

	// 拼接标签信息
	upBiao, err := spliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid)
	if err != nil {
		log.Println("=>09.	拼接上链标签结构失败：", err)
		return nil, err
	}

	log.Println("=>09.	拼接上链标签结构")

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链标签
	*/
	log.Println("=>10.	返回")
	return api.GetSliceArgs(TYUPLOADASSET, KYC, updata, upBiao)
}

//	根据主键查询元数据
func MetadataByPrimaryKey(info Info) (arr []string, err error) {
	//
	Pg := UseDataById{}
	//	01.	序列化数据
	Byteshal, err := json.Marshal(info.Data)
	if err != nil {
		log.Println("2:反序列化数据：=>", string(Byteshal))
		return nil, err
	}
	log.Println("获取数据集=>", string(Byteshal))
	//
	err = json.Unmarshal(Byteshal, &Pg)
	if err != nil {
		log.Println("3. Unmarshal err:", err)
		return nil, err
	}

	// 获取uuid 即主键  || 验证空
	log.Println("uuid=>", Pg.ById)
	uuid := Pg.ById
	if uuid == "" {
		return nil, errors.New("uuid 为空")

	}
	//03.  拼接上链结构体 查询 TODO
	return api.GetSliceArgsById(GETASSEETBYID, uuid)
}

//	根据主键查询标签
func MetaLabelByPrimaryKey(info Info) (arr []string, err error) {
	//
	Pg := UseDataById{}
	//	01.	序列化数据
	Byteshal, err := json.Marshal(info.Data)
	if err != nil {
		log.Println("2:反序列化数据：=>", string(Byteshal))
		return nil, err
	}
	log.Println("获取数据集=>", string(Byteshal))
	//
	err = json.Unmarshal(Byteshal, &Pg)
	if err != nil {
		log.Println("3. Unmarshal err:", err)
		return nil, err
	}

	// 获取uuid 即主键  || 验证空
	log.Println("uuid=>", Pg.ById)
	uuid := Pg.ById
	if uuid == "" {
		return nil, errors.New("uuid 为空")

	}
	//03.  拼接上链结构体 查询 TODO
	return api.GetSliceArgsById(GETASSEETBYID, BYC+uuid)
}

//  根据索引返回结构体指针  ||
func pointerToIndex(pointer string) (inter interface{}, err error) {

	switch pointer {

	//-- 	房地产 ---
	case Pro_Rea_All:
		inter = &RealtyInfoWeb{} //	房地产主表(整体)
	case Pro_Rea_One:
		inter = &RealtyInfoWeb{} //	主表标识
	case Pro_Rea_Fdc:
		inter = &RealtyInfoWeb{} //	房地产标识
	case Pro_Rea_Cgt:
		inter = &DueProjectCemGt{} //	保证人标识
	case Pro_Rea_Cmg:
		inter = &DueProjectCemMg{} //	抵押品标识
	case Pro_Rea_Cpg:
		inter = &DueProjectCemPg{} //	质押品标识
	case Pro_Rec_Yjg:
		inter = &BaseSurvey{} //	尽调结果
	case Pro_Rec_Jdb:
		inter = &BaseReport{} //	尽调报告
	//-- 	应收账款 ---
	case Pro_Rec_Ysk:
		inter = &AstInfoWeb{} //	质押品标识
	case Pro_Rec_JCZY:
		inter = &PleInfoWeb{} //	动产质押标识
	//-- 	微分格  ---  创建 质押 解压
	case Pro_Rec_CR:
		inter = &WeiByCr{} //	动产质押标识
	case Pro_Rec_ZH:
		inter = &WeiByCr{} //	动产质押标识
	case Pro_Rec_JI:
		inter = &WeiByCr{} //	动产质押标识
	case Pro_Rec_Jiu:
		inter = &WeiByCrbf{} //	 微分格旧数据
	// default
	default:
		return nil, errors.New("no struct")
	}
	return inter, nil
}

//  根据索引返回结构体指针  ||
func pointerToType(pointer string) (inter string, dataByType string, err error) {
	//PROJECT:资产内容上链;REPORT:尽调报告上链；DILIGENCE_RESULT:尽调结果;PERFECT:补充资产信息;UPDATE:更新资产信息
	switch pointer {

	//-- 	房地产 ---
	case Pro_Rea_All:
		inter = "PROJECT" //	房地产主表(整体)
		dataByType = "fdc"
	case Pro_Rea_One:
		inter = "PROJECT" //	主表标识
	case Pro_Rea_Fdc:
		inter = "PROJECT" //	房地产标识

	case Pro_Rea_Cgt:
		inter = "bzr" //	保证人标识

	case Pro_Rea_Cmg:
		inter = "dyp" //	抵押品标识

	case Pro_Rea_Cpg:
		inter = "zyp" //	质押品标识

	case Pro_Rec_Yjg:
		inter = "DILIGENCE_RESULT" //	尽调结果
		dataByType = "jdjg"

	case Pro_Rec_Jdb:
		inter = "REPORT" //	尽调报告
		dataByType = "jdbg"
	//-- 	应收账款 ---

	case Pro_Rec_Ysk:
		inter = "PROJECT" //	应收账款
		dataByType = "ar"
	// default

	case Pro_Rec_JCZY:
		inter = "PROJECT" //	动产质押
		dataByType = "dc"

		//-- 	微分格  ---  创建 质押 解压
	case Pro_Rec_CR:
		inter = "PROJECT" //	动产质押
		dataByType = "wcr"
	case Pro_Rec_ZH:
		inter = "PROJECT" //	动产质押
		dataByType = "wzh"
	case Pro_Rec_JI:
		inter = "PROJECT" //	动产解压
		dataByType = "whr"
	default:
		return "", "", errors.New("no struct")
	}
	return inter, dataByType, nil
}

// V3 ----------------- 升级

// 添加组织序列化
func unAffition(c *gin.Context) (cdata *UnionWeb, err error) {

	//	01. 浏览器展示结构体 ||添加组织
	data := UnionWeb{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

//@ 应收账款数据上链
func uploadProjectReceivables(Info Info) (list []string, err error) {

	//	01. 数据
	Pg := AstInfo{}

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
		log.Println("元数据操作序列化失败=>", err)
		return nil, err
	}

	//	05. 追加  time  id 以及标签

	// 06. 数据加密
	strcData, err := json.Marshal(&Pg)
	if err != nil {
		log.Println("err:06. 数据加密=>", err)
		return nil, err
	}

	//	06.01 获取私钥以及类型
	resKey, useKey, pubKey := getEncryptionKeytest(Info.User.OrgID, Info.User.UserID, Info.User.OrgID+Info.User.UserID)

	//	06.02  元数据加密
	log.Println("获取私钥以及类型 元数据加密=>", resKey, useKey, pubKey)

	//	06.03  元数据加密执行  TODO 默认是用户ID 进行加密
	cipher := MetadataEncryption(string(strcData), useKey)
	log.Println("上链数据数据加密")

	// uuid TODO  uuid 目前是资产包的ID
	var uuid string
	if len(Pg.AstAssetsInfo.AstAssetsList) < 0 {
		return nil, errors.New("数据主键为空")
	}
	uuid = Pg.AstAssetsInfo.AstAssetsList[0].AstAssetsuuid

	//	07.	拼接上链数据结构
	updata, err := concatenateTheUpperLinkDataStructure(uuid, Info.Datatype, Info.User.Datatime, cipher, useKey, signPower04, Info.User)
	if err != nil {
		log.Println("err：07.拼接上链数据结构=>", err)
		return nil, err
	}
	//
	log.Println("07.拼接上链数据结构通过")

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

	log.Println("用户信息=>", &sUse)

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

	log.Println("关联数据信息=>", &sUse)

	//	签名信息
	sSign := Sign{
		resKey,
		useKey,
		pubKey,
		useKey, //TODO
	}
	log.Println("签名信息=>", &sSign)

	// 拼接标签信息
	upBiao, err := spliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid)
	if err != nil {
		log.Println("拼接标签信息=>", err)
		return nil, err
	}

	log.Println("拼接上链标签结构通过")

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链标签
	*/
	return api.GetSliceArgs(TYUPLOADASSET, KYC, updata, upBiao)
}

//@ 动产融资数据上链
func uploadProjectByPle(Info Info) (list []string, err error) {

	//	01. 数据
	Pg := PleInfo{}

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
		log.Println("元数据操作序列化失败=>", err)
		return nil, err
	}

	//	05. 追加  time  id 以及标签

	// 06. 数据加密
	strcData, err := json.Marshal(&Pg)
	if err != nil {
		log.Println("err:06. 数据加密=>", err)
		return nil, err
	}

	//	06.01 获取私钥以及类型
	resKey, useKey, pubKey := getEncryptionKeytest(Info.User.OrgID, Info.User.UserID, Info.User.OrgID+Info.User.UserID)

	//	06.02  元数据加密
	log.Println("获取私钥以及类型 元数据加密=>", resKey, useKey, pubKey)

	//	06.03  元数据加密执行  TODO 默认是用户ID 进行加密
	cipher := MetadataEncryption(string(strcData), useKey)
	log.Println("上链数据数据加密")

	// uuid TODO  uuid 目前是资产包的ID
	var uuid string
	if len(Pg.AstAssetsInfo.AstAssetsListByPle) < 0 {
		return nil, errors.New("数据主键为空")
	}
	uuid = Pg.AstAssetsInfo.AstAssetsListByPle[0].Id

	//	07.	拼接上链数据结构
	updata, err := concatenateTheUpperLinkDataStructure(uuid, Info.Datatype, Info.User.Datatime, cipher, useKey, signPower04, Info.User)
	if err != nil {
		log.Println("err：07.拼接上链数据结构=>", err)
		return nil, err
	}
	//
	log.Println("07.拼接上链数据结构通过")

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

	log.Println("用户信息=>", &sUse)

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

	log.Println("关联数据信息=>", &sUse)

	//	签名信息
	sSign := Sign{
		resKey,
		useKey,
		pubKey,
		useKey, //TODO
	}
	log.Println("签名信息=>", &sSign)

	// 拼接标签信息
	upBiao, err := spliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid)
	if err != nil {
		log.Println("拼接标签信息=>", err)
		return nil, err
	}

	log.Println("拼接上链标签结构通过")

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链标签
	*/
	return api.GetSliceArgs(TYUPLOADASSET, KYC, updata, upBiao)
}

//@ 尽调结果上链数据上链
func uploadProjectFindings(Info Info) (list []string, err error) {

	//	01. 数据
	Pg := BaseSurvey{}

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
		log.Println("元数据操作序列化失败=>", err)
		return nil, err
	}

	//	05. 追加  time  id 以及标签 TODO 目前标签追加是在 数据上

	// 06. 数据加密
	strcData, err := json.Marshal(&Pg)
	if err != nil {
		log.Println("err:06. 数据加密=>", err)
		return nil, err
	}

	//	06.01 获取私钥以及类型
	resKey, useKey, pubKey := getEncryptionKeytest(Info.User.OrgID, Info.User.UserID, Info.User.OrgID+Info.User.UserID)

	//	06.02  元数据加密
	log.Println("获取私钥以及类型 元数据加密=>", resKey, useKey, pubKey)

	//	06.03  元数据加密执行  TODO 默认是用户ID 进行加密
	cipher := MetadataEncryption(string(strcData), useKey)
	log.Println("上链数据数据加密")

	// uuid TODO  uuid 目前是资产包的ID
	uuid := Pg.IpfsHash

	//	07.	拼接上链数据结构
	updata, err := concatenateTheUpperLinkDataStructure(uuid, Info.Datatype, Info.User.Datatime, cipher, useKey, signPower04, Info.User)
	if err != nil {
		log.Println("err：07.拼接上链数据结构=>", err)
		return nil, err
	}
	//
	log.Println("07.拼接上链数据结构通过")

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

	log.Println("用户信息=>", &sUse)

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

	log.Println("关联数据信息=>", &sUse)

	//	签名信息
	sSign := Sign{
		resKey,
		useKey,
		pubKey,
		useKey, //TODO
	}
	log.Println("签名信息=>", &sSign)

	// 拼接标签信息
	upBiao, err := spliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid)
	if err != nil {
		log.Println("拼接标签信息=>", err)
		return nil, err
	}

	log.Println("拼接上链标签结构通过")

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链标签
	*/
	return api.GetSliceArgs(TYUPLOADASSET, KYC, updata, upBiao)
}

//@ 尽调报告 上链
func uploadProjectReport(Info Info) (list []string, err error) {

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

	//	05. 追加  time  id 以及标签 TODO 目前标签追加是在 数据上

	// 06. 数据加密
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
	upBiao, err := spliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid)
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
	return api.GetSliceArgs(TYUPLOADASSET, KYC, updata, upBiao)
}

// 微分格 V4
func DispenseV4(dataType string, info Info) (list []string, err error) {
	// 请求类型
	switch dataType {
	case Wei_Rea_01:
		return upLoadByBoxInBox(info)
	case Wei_Rea_02:
		return upLoadByBoxInBox(info)
	case Wei_Rea_03:
		return upLoadByBoxInBox(info)
	case Wei_Rea_04:
		return upLoadByBoxInBox(info)
	default:
		log.Println("no func")
	}
	return nil, errors.New("no func")
}

//  1类数据上链 || 微分格
func upLoadByBoxInBox(Info Info) (list []string, err error) {

	Pg := Box{}

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
		log.Println("元数据操作序列化失败=>", err)
		return nil, err
	}

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链标签
	*/
	return api.GetSliceArgs(TYUPLOADASSET, KYC, "1", "23")
}

//	微分格 根据主键查询元数据
func MetadataByPrimaryKeyByBox(info InfoWei) (arr []string, err error) {

	//03.  拼接上链结构体 查询 TODO
	return api.GetSliceArgsById(info.Chain.FuncName, info.Data[0])
}

//
//	房地产数据整体上链 V3 加密Sign
func upLoadByAllSign(Info Info) (list []string, err error) {

	//	01. 房地产数据结构体
	Pg := RealtyInfo{}

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
		log.Println("元数据操作序列化失败=>", err)
		return nil, err
	}

	//	05. 追加  time  id 以及标签
	Pg.FBase.Ptype = Info.Datatype

	// 06. 数据加密
	strcData, err := json.Marshal(&Pg)
	if err != nil {
		log.Println("err:06. 数据加密=>", err)
		return nil, err
	}

	//	06.01 获取私钥以及类型
	resKey, useKey, pubKey := getEncryptionKeytest(Info.User.OrgID, Info.User.UserID, Info.User.OrgID+Info.User.UserID)

	//	06.02  元数据加密
	log.Println("获取私钥以及类型 元数据加密=>", resKey, useKey, pubKey)

	//	06.03  元数据加密执行  TODO 默认是用户ID 进行加密
	cipher := MetadataEncryption(string(strcData), useKey)
	log.Println("上链数据数据加密")

	// uuid TODO  uuid 目前是资产包的ID
	uuid := Pg.FDueProject.FID

	//	07.	拼接上链数据结构
	updata, err := concatenateTheUpperLinkDataStructure(uuid, Info.Datatype, Info.User.Datatime, cipher, useKey, signPower04, Info.User)
	if err != nil {
		log.Println("err：07.拼接上链数据结构=>", err)
		return nil, err
	}
	//
	log.Println("07.拼接上链数据结构通过")

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

	log.Println("用户信息=>", &sUse)

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

	log.Println("关联数据信息=>", &sUse)

	//	签名信息
	sSign := Sign{
		resKey,
		useKey,
		pubKey,
		useKey, //TODO
	}
	log.Println("签名信息=>", &sSign)

	// 拼接标签信息
	upBiao, err := spliceTagDataStructures(&sUse, &sBycMesList, &sSign, uuid)
	if err != nil {
		log.Println("拼接标签信息=>", err)
		return nil, err
	}

	log.Println("拼接上链标签结构通过")

	/**
	@	参数解析：
		1. 链码函数名称
		2. 数据类型标识
		3. 上链数据集
		4. 上链标签
	*/
	return api.GetSliceArgs(TYUPLOADASSET, KYC, updata, upBiao)
}

//  查询合约  和 查询 上链 || 合约转换
func selectByContract(inx string) (rep string) {

	//
	switch inx {
	//
	case Pro_Rea_All:
		return "房地产合约"
	case Pro_Rec_Ysk:
		return "应收账款合约"
	case Pro_Rec_JCZY:
		return "动产质押合约"
	default:
		return "第三方资产合约"
	}

}

//  查询合约  和 查询 上链  资产转换
func selectByAsset(inx string) (rep string) {

	//
	switch inx {
	//
	case Pro_Rea_All:
		return "房地产资产"
	case Pro_Rec_Ysk:
		return "应收账资产"
	case Pro_Rec_JCZY:
		return "动产质押资产"
	default:
		return "第三方资产"
	}

}

// -- 根据交易ID 查询  房地产信息
func queryDataByFang(Mg *Kyc, strData string) (label string, dataByType string, rep interface{}, err error) {
	// 获取对应的对象结构体指针
	strPointer, err := PointerToIndex(Mg.KycType)
	if err != nil {
		log.Println("获取对应的对象结构体指针=>", err)
		return "", "", &Mg, err
	}
	strType, dataByType, err := PointerToType(Mg.KycType)
	err = json.Unmarshal([]byte(strData), strPointer)
	//
	if err != nil {
		log.Println("err=>", err)
		return strType, dataByType, &Mg, err
	}
	//
	log.Println("strPointer=>", strPointer)
	return strType, dataByType, strPointer, nil
}

// -- 根据交易ID 查询  应收账款信息
func queryDataByYing(Mg *Kyc, strData string) (label string, dataByType string, rep interface{}, err error) {
	// 获取对应的对象结构体指针
	strPointer, err := PointerToIndex(Mg.KycType)
	if err != nil {
		log.Println("获取对应的对象结构体指针=>", err)
		return "ar", "", &Mg, err
	}
	strType, dataByType, err := PointerToType(Mg.KycType)
	err = json.Unmarshal([]byte(strData), strPointer)

	// TODO  旧数据处理

	data := AstInfo{}
	err = json.Unmarshal([]byte(strData), &data)
	//

	if err != nil {
		//
		log.Println("err=>", err)
		return "", "", nil, nil
	}

	res := ParameterTransformation(data)

	//

	if err != nil {
		log.Println("err=>", err)
		return strType, dataByType, &Mg, err
	}
	//
	log.Println("strPointer=>", strPointer)
	return strType, dataByType, res, nil
}

// -- 根据交易ID 查询  尽调结果信息
func queryDataByBasePle(Mg *Kyc, strData string) (label string, dataByType string, rep interface{}, err error) {
	// 获取对应的对象结构体指针
	strPointer, err := PointerToIndex(Mg.KycType)
	if err != nil {
		log.Println("获取对应的对象结构体指针=>", err)
		return "", "", &Mg, err
	}
	strType, dataByType, err := PointerToType(Mg.KycType)
	err = json.Unmarshal([]byte(strData), strPointer)
	//
	if err != nil {
		log.Println("err=>", err)
		return strType, dataByType, &Mg, err
	}
	//TODO 旧数据处理
	//data := BaseSurvey{}
	//ParameterTransformationByBaseSurvey(data)

	log.Println("strPointer=>", strPointer)
	return strType, dataByType, strPointer, nil
}

// -- 根据交易ID 查询  尽调报告信息
func queryDataByBaseReport(Mg *Kyc, strData string) (label string, dataByType string, rep interface{}, err error) {
	// 获取对应的对象结构体指针
	strPointer, err := PointerToIndex(Mg.KycType)
	if err != nil {
		log.Println("获取对应的对象结构体指针=>", err)
		return "", "", &Mg, err
	}
	strType, dataByType, err := PointerToType(Mg.KycType)
	err = json.Unmarshal([]byte(strData), strPointer)
	//
	if err != nil {
		log.Println("err=>", err)
		return strType, dataByType, &Mg, err
	}
	//

	//TODO 旧数据处理
	//data := BaseReport{}
	//ParameterTransformationByBaseReport(data)

	log.Println("strPointer=>", strPointer)
	return strType, dataByType, strPointer, nil
}

// -- 根据交易ID 查询  应收账款信息
func queryDataByAdjoin(Mg *Kyc, strData string) (label string, dataType string, rep interface{}, err error) {
	// 获取对应的对象结构体指针
	strPointer, err := PointerToIndex(Mg.KycType)
	if err != nil {
		log.Println("获取对应的对象结构体指针=>", err)
		return "ys", "", &Mg, err
	}
	strType, dataType, err := PointerToType(Mg.KycType)
	err = json.Unmarshal([]byte(strData), strPointer)
	//
	if err != nil {
		log.Println("err=>", err)
		return "ys", "", &Mg, err
	}
	//
	log.Println("strPointer=>", strPointer)

	return strType, dataType, strPointer, nil
}

// -- 根据交易ID 查询  动产质押信息
func queryDataByBaseSurvey(Mg *Kyc, strData string) (label string, dataByType string, rep interface{}, err error) {
	// 获取对应的对象结构体指针
	strPointer, err := PointerToIndex(Mg.KycType)
	if err != nil {
		log.Println("获取对应的对象结构体指针=>", err)
		return "", "", &Mg, err
	}
	strType, dataByType, err := PointerToType(Mg.KycType)
	err = json.Unmarshal([]byte(strData), strPointer)
	//
	if err != nil {
		log.Println("err=>", err)
		return strType, dataByType, &Mg, err
	}
	//TODO 旧数据处理
	log.Println("strPointer=>", strPointer)
	return strType, dataByType, strPointer, nil
}

// -- 根据交易ID 查询  - 微分格
func queryDataByWeiForapp(Mg *Kyc, strData string, sType string) (strType string, dataType string, rep interface{}, err error) {
	// 获取对应的对象结构体指针
	var strPointer interface{}
	//
	log.Println("数据类型：", sType)
	if sType == "v1" {
		strPointer, err = PointerToIndex(Pro_Rec_Jiu)
	} else if sType == "v2" {
		strPointer, err = PointerToIndex(Pro_Rec_CR)
	}

	if err != nil {
		log.Println("获取对应的对象结构体指针=>", err)
		return "ys", "", &Mg, err
	}
	strType, dataType, err = PointerToType(Pro_Rec_CR)
	if err != nil {
		log.Println("序列化类型=>", err)
		return "ys", "", &Mg, err
	}

	err = json.Unmarshal([]byte(strData), strPointer)
	//
	if err != nil {
		log.Println("数据序列化失败=>", err)
		return "ys", "", &Mg, err
	}

	return strType, dataType, strPointer, nil
}

//  参数转换啊

func parameterTransformation(res AstInfo) (repdata *AstInfoWebJiu) {
	data := AstInfoWebJiu{}
	data.AstAssetsInfo.AstPackageuuid = res.AstAssetsInfo.AstPackageuuid
	data.AstAssetsInfo.AstPackageuuid = res.AstAssetsInfo.AstPackageuuid
	data.AstAssetsInfo.AstPackageNo = res.AstAssetsInfo.AstPackageNo
	data.AstAssetsInfo.AstPackagePlatForm = res.AstAssetsInfo.AstPackagePlatForm
	data.AstAssetsInfo.AstPackagePlatFormID = res.AstAssetsInfo.AstPackagePlatFormID
	data.AstAssetsInfo.AstPackageName = res.AstAssetsInfo.AstPackageName
	data.AstAssetsInfo.AstPackageNumber = res.AstAssetsInfo.AstPackageNumber
	data.AstAssetsInfo.AstPackageOwnerName = res.AstAssetsInfo.AstPackageOwnerName
	data.AstAssetsInfo.AstPackageOwnerId = res.AstAssetsInfo.AstPackageOwnerId
	data.AstAssetsInfo.AstPackageSplit = res.AstAssetsInfo.AstPackageSplit
	data.AstAssetsInfo.AstPackageEvaluation = res.AstAssetsInfo.AstPackageEvaluation

	//融资信息
	data.AstAssetsInfo.AstFinancingInfo.AstFinuuid = res.AstAssetsInfo.AstFinancingInfo.AstFinuuid
	data.AstAssetsInfo.AstFinancingInfo.AstFinfMasterID = res.AstAssetsInfo.AstFinancingInfo.AstFinfMasterID
	data.AstAssetsInfo.AstFinancingInfo.AstFinPrice = res.AstAssetsInfo.AstFinancingInfo.AstFinPrice
	data.AstAssetsInfo.AstFinancingInfo.AstCreGuaName = res.AstAssetsInfo.AstFinancingInfo.AstCreGuaName
	data.AstAssetsInfo.AstFinancingInfo.AstCreGuaManner = res.AstAssetsInfo.AstFinancingInfo.AstCreGuaManner
	data.AstAssetsInfo.AstFinancingInfo.AstCreGuaEnsureId = res.AstAssetsInfo.AstFinancingInfo.AstCreGuaEnsureId

	//增信措施信息
	if len(res.AstAssetsInfo.AstCreditInfo.AstCreEnsureList) > 0 {
		AstCreEnsureListJiu := AstCreEnsureListJiu{}
		for k, _ := range res.AstAssetsInfo.AstCreditInfo.AstCreEnsureList {
			AstCreEnsureListJiu.AstCreEnsfMasterID = res.AstAssetsInfo.AstCreditInfo.AstCreEnsureList[k].AstCreEnsfMasterID
			AstCreEnsureListJiu.AstCreEnsName = res.AstAssetsInfo.AstCreditInfo.AstCreEnsureList[k].AstCreEnsName
			AstCreEnsureListJiu.AstCreEnsOwner = res.AstAssetsInfo.AstCreditInfo.AstCreEnsureList[k].AstCreEnsOwner
			AstCreEnsureListJiu.AstCreEnsType = res.AstAssetsInfo.AstCreditInfo.AstCreEnsureList[k].AstCreEnsType
			AstCreEnsureListJiu.AstCreEnsuuid = res.AstAssetsInfo.AstCreditInfo.AstCreEnsureList[k].AstCreEnsuuid

			data.AstAssetsInfo.AstCreditInfo.AstCreEnsureList = append(data.AstAssetsInfo.AstCreditInfo.AstCreEnsureList, AstCreEnsureListJiu)

		}
	}

	//增信措施信息

	// 增信措施类型

	if len(res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList) > 0 {
		AstCreGuarantyListJiu := AstCreGuarantyListJiu{}
		for k, _ := range res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList {
			AstCreGuarantyListJiu.AstCreGuaEnsureId = res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList[k].AstCreGuaEnsureId
			AstCreGuarantyListJiu.AstCreGuaEnsureName = res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList[k].AstCreGuaEnsureName
			AstCreGuarantyListJiu.AstCreGuafMasterID = res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList[k].AstCreGuafMasterID
			AstCreGuarantyListJiu.AstCreGuaManner = res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList[k].AstCreGuaManner
			AstCreGuarantyListJiu.AstCreGuaName = res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList[k].AstCreGuaName
			AstCreGuarantyListJiu.AstCreGuauuid = res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList[k].AstCreGuauuid
			AstCreGuarantyListJiu.AstCreGuaType = res.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList[k].AstCreGuaType

			data.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList = append(data.AstAssetsInfo.AstCreditInfo.AstCreGuarantyList, AstCreGuarantyListJiu)
		}

	}

	// 增信措施类型
	if len(res.AstAssetsInfo.AstCreditInfo.AstCrePledgeList) > 0 {
		AstCrePledgeListJiu := AstCrePledgeListJiu{}
		for k, _ := range res.AstAssetsInfo.AstCreditInfo.AstCrePledgeList {
			AstCrePledgeListJiu.AstCrePleuuid = res.AstAssetsInfo.AstCreditInfo.AstCrePledgeList[k].AstCrePleuuid
			AstCrePledgeListJiu.AstCrePlefMasterID = res.AstAssetsInfo.AstCreditInfo.AstCrePledgeList[k].AstCrePlefMasterID
			AstCrePledgeListJiu.AstCrePleType = res.AstAssetsInfo.AstCreditInfo.AstCrePledgeList[k].AstCrePleType
			AstCrePledgeListJiu.AstCrePleName = res.AstAssetsInfo.AstCreditInfo.AstCrePledgeList[k].AstCrePleName
			AstCrePledgeListJiu.AstCrePleOwner = res.AstAssetsInfo.AstCreditInfo.AstCrePledgeList[k].AstCrePleOwner

			//
			data.AstAssetsInfo.AstCreditInfo.AstCrePledgeList = append(data.AstAssetsInfo.AstCreditInfo.AstCrePledgeList, AstCrePledgeListJiu)
		}
	}

	//
	data.AstAssetsInfo.AstSendInfo.AstSenduuid = res.AstAssetsInfo.AstSendInfo.AstSenduuid
	data.AstAssetsInfo.AstSendInfo.AstSendfMasterID = res.AstAssetsInfo.AstSendInfo.AstSendfMasterID
	data.AstAssetsInfo.AstSendInfo.AstSendName = res.AstAssetsInfo.AstSendInfo.AstSendName
	data.AstAssetsInfo.AstSendInfo.AstSendid = res.AstAssetsInfo.AstSendInfo.AstSendid
	data.AstAssetsInfo.AstSendInfo.AstSendContact = res.AstAssetsInfo.AstSendInfo.AstSendContact
	data.AstAssetsInfo.AstSendInfo.AstSendTime = res.AstAssetsInfo.AstSendInfo.AstSendTime

	if len(res.AstAssetsInfo.AstAssetsList) > 0 {

		// TODO
		AstAssetsListJiu := AstAssetsListJiu{}
		for k, _ := range res.AstAssetsInfo.AstAssetsList {
			AstAssetsListJiu.AstAssetsfMasterID = res.AstAssetsInfo.AstAssetsList[k].AstAssetsfMasterID
			AstAssetsListJiu.AstAssetsType = res.AstAssetsInfo.AstAssetsList[k].AstAssetsType
			AstAssetsListJiu.AstAssetsIntroduce = res.AstAssetsInfo.AstAssetsList[k].AstAssetsIntroduce
			AstAssetsListJiu.AstAssetsCreditorName = res.AstAssetsInfo.AstAssetsList[k].AstAssetsCreditorName
			AstAssetsListJiu.AstAssetsCreditorTaxNum = res.AstAssetsInfo.AstAssetsList[k].AstAssetsCreditorTaxNum
			AstAssetsListJiu.AstAssetsDebtorName = res.AstAssetsInfo.AstAssetsList[k].AstAssetsDebtorName
			AstAssetsListJiu.AstAssetsDebtorTaxNum = res.AstAssetsInfo.AstAssetsList[k].AstAssetsDebtorTaxNum
			AstAssetsListJiu.AstAssetsValuation = res.AstAssetsInfo.AstAssetsList[k].AstAssetsValuation
			AstAssetsListJiu.AstAssetsHonour = res.AstAssetsInfo.AstAssetsList[k].AstAssetsHonour
			AstAssetsListJiu.AstAssetsState = res.AstAssetsInfo.AstAssetsList[k].AstAssetsState
			AstAssetsListJiu.AstAssetsPrimeval = res.AstAssetsInfo.AstAssetsList[k].AstAssetsPrimeval

			//
			data.AstAssetsInfo.AstAssetsList = append(data.AstAssetsInfo.AstAssetsList, AstAssetsListJiu)

			if len(res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList) > 0 {
				AstConJiu := AstConJiu{}
				for j, _ := range res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList {
					AstConJiu.AstConuuid = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConuuid
					AstConJiu.AstConfMasterID = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConfMasterID
					AstConJiu.AstConNo = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConNo
					AstConJiu.AstConName = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConName
					AstConJiu.AstConPrice = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConPrice
					AstConJiu.AstConPriceUnit = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConPriceUnit
					AstConJiu.AstConType = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConType
					AstConJiu.AstConPayerTaxNum = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConPayerTaxNum
					AstConJiu.AstConPayerName = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConPayerName
					AstConJiu.AstConTime = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConTime
					AstConJiu.AstConDays = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConDays
					AstConJiu.AstConCount = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConCount
					// TODO

					data.AstAssetsInfo.AstAssetsList[k].AstContractInfoList = append(data.AstAssetsInfo.AstAssetsList[k].AstContractInfoList, AstConJiu)

					AttachmentListJiu := AttachmentListJiu{}

					if len(res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList) > 0 {
						for w, _ := range res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList {
							AttachmentListJiu.IpfsHash = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].IpfsHash
							AttachmentListJiu.AstAttachAddr = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachAddr
							AttachmentListJiu.AstAttachfMasterID = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachfMasterID
							AttachmentListJiu.AstAttachMD = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachMD
							AttachmentListJiu.AstAttachName = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachName
							AttachmentListJiu.AstAttachNo = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachNo
							AttachmentListJiu.AstAttachType = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachType

							//TODO
							data.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList = append(data.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList, AttachmentListJiu)
							//
						}
					}
				}
			}

			//
			if len(res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList) > 0 {
				AstInvJiu := AstInvJiu{}
				for j, _ := range res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList {
					AstInvJiu.AstInvuuid = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvuuid
					AstInvJiu.AstInvfMasterID = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvfMasterID
					AstInvJiu.AstInvType = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvType
					AstInvJiu.AstInvNum = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvNum
					AstInvJiu.AstInvCode = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvCode
					AstInvJiu.AstInvChecksum = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvChecksum
					AstInvJiu.AstInvPrice = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvPrice
					AstInvJiu.AstInvUnit = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvUnit
					AstInvJiu.AstInvTime = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvTime
					AstInvJiu.AstInvBuyerTaxNum = res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvBuyerTaxNum

					//
					data.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList = append(data.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList, AstInvJiu)

					if len(res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvAttachmentList) > 0 {

						AttachmentListJiua := AttachmentList{}
						for w, _ := range res.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvAttachmentList {
							AttachmentListJiua.IpfsHash = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].IpfsHash
							AttachmentListJiua.AstAttachAddr = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachAddr
							AttachmentListJiua.AstAttachfMasterID = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachfMasterID
							AttachmentListJiua.AstAttachMD = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachMD
							AttachmentListJiua.AstAttachName = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachName
							AttachmentListJiua.AstAttachNo = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachNo
							AttachmentListJiua.AstAttachType = res.AstAssetsInfo.AstAssetsList[k].AstContractInfoList[j].AstConAttachmentList[w].AstAttachType
							// TODO
							data.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvAttachmentList = append(data.AstAssetsInfo.AstAssetsList[k].AstInvoiceInfoList[j].AstInvAttachmentList, AttachmentListJiua)
						}
					}
				}
			}
		}
	}

	return &data
}

// -- 旧数据转换
func parameterTransformationByBaseSurvey(res BaseSurvey) (repdata *BaseSurveyByJiu) {

	data := BaseSurveyByJiu{}
	data.IpfsHash = res.IpfsHash
	return &data
}

// -- 旧数据转换
func parameterTransformationByBaseReport(res BaseReport) (repdata *BaseReportByJiu) {

	data := BaseReportByJiu{}
	data.ReportName = res.ReportName
	data.ReportType = res.ReportType
	data.ReportNo = res.ReportNo
	data.ReportTx = res.ReportTx
	data.ReportMD = res.ReportMD
	data.ReportAddr = res.ReportAddr
	data.FMasterID = res.FMasterID
	return &data
}

// -- queryDataByBase

//转换数据 --
func queryDataByBase(x string, Mg *Kyc, strData string) (label string, dataByType string, rep interface{}, err error) {

	//
	switch x {

	// 房地产
	case Pro_Rea_All:
		label, dataByType, rep, err = queryDataByFang(Mg, strData)

		// 应收账款
	case Pro_Rec_Ysk:
		label, dataByType, rep, err = queryDataByYing(Mg, strData)

		// 尽调结果
	case Pro_Rec_Yjg:
		label, dataByType, rep, err = queryDataByBaseReport(Mg, strData)

		// 尽调报告
	case Pro_Rec_Jdb:
		label, dataByType, rep, err = queryDataByBaseSurvey(Mg, strData)

		// -- 测试
	case Pro_Rea_Cpg:
		label, dataByType, rep, err = queryDataByAdjoin(Mg, strData)

		// 房地产
	default:
		label, dataByType, rep, err = queryDataByAdjoin(Mg, strData)
		log.Println("queryDataByBase no func")
	}

	return label, dataByType, rep, err
}

//  通道名称
func getChannel() []ChannelData {
	return []ChannelData{
		ChannelData{UUID: "001", Channel: Chan_assetpublish},
	}
}

//	微分格 通知系统
func NotificationSystem(assetState string, assetCode string) (err error) {
	// 拼接数据
	str, err := notificationSystemforDataCon(assetState)

	log.Println("str:", str)

	if err != nil {
		log.Println("err:", err)
		return err
	}

	// 发送请求
	return callTheNotification(str, assetCode)
}

//  微分格 通知系统 数据转换
func notificationSystemforDataCon(assetState string) (str string, err error) {
	log.Println("数据转换")
	if assetState != "" {
		switch assetState {
		case "invoke":
			str = "new"
		case "update":
			str = "update"
		case "query":
			str = "query"
			err = errors.New("is query")
		case "delete":
			str = "delete"
			err = errors.New("is delete")
		default:
			str = ""
			err = errors.New("not func")
		}
	}
	return str, err
}

// 资产后五个查询
func unBlockTransactionNext(c *gin.Context) (cdata *BlockAssetQuery, err error) {
	//	01. 浏览器展示结构体 ||一个区块的交易
	data := BlockAssetQuery{}
	// 02.	序列化
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 资产后五个查询
func unBlockCompound(c *gin.Context) (cdata *QueryBlockCompound, err error) {

	data := QueryBlockCompound{}

	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func unEcho(c *gin.Context) (cdata *Echo, err error) {
	data := Echo{}
	if err := c.ShouldBindJSON(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

//	复合查询  -根据txid  查询
func compoundQueryByTxid(uuid string) (rep interface{}, err error) {

	//	声明参数
	Mg := Kyc{}
	resData := TxIDrenData{}
	resDataJiu := TxIDrenDataJiu{}
	ChainTrans := ChainTransactionConfig{}
	wByType := WeiByType{}
	var strData string
	var JiuWeifen string

	// 根据txid查询数据
	result, resultByString, err := model_action.App.GetBlockTxId("", "", "", "", "", uuid)

	if err != nil {
		return nil, err
	}
	// 反序列化
	json.Unmarshal([]byte(result), &ChainTrans)

	//	拼接返回体
	ChainTrans.DataTxid = uuid
	resData.ChainTransactionConfig = ChainTrans
	log.Println("根据txid查询数据:", result)

	//	判断
	if len(resultByString) == 0 {
		return nil, errors.New("resultByString is nil")
	}
	if len(resultByString) < 2 {
		log.Println("根据txid查询到的数据是微分格数据类型", 2)

		return resData, nil
	}

	//	异常情况处理完毕, 开始进行数据处理  kyc 数据在  retstring [1]   retstring [0] 类型      retstring [2] 标签

	// 数据解密

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
	//	数据转码

	//	数据本体
	err = json.Unmarshal([]byte(resultByString[1]), &Mg)
	if err != nil {
		log.Println("data transcoding failed:", err)
		return resData, nil
	}

	//	数据配置
	err = json.Unmarshal([]byte(resultByString[2]), &wByType)
	if err != nil {
		log.Println("data transcoding failed:", err)
		return resData, nil
	}

	log.Println("数据配置:", wByType.CategoryId)

	// TODO  判断是哪一种数据类型

	//	数据判断

	//	判断是否是加密类型数据上链类型
	if Mg.KycString != "" || Mg.SignKey != "" {
		//	如果是 加解密类型,开始解密
		strData = MetadataAesDecrypt(Mg.KycString, Mg.SignKey)
		log.Println("解密成功")
	} else {
		//	如果解密后的类型没有包含 kyc signkey,那就是微分格数据，赋值类型  微分格 数据类型 49
		Mg.KycType = Pro_Rec_WI
	}

	// 分发器 根据 KycType 解析对应的数据结构
	switch Mg.KycType {

	//	房地产
	case Pro_Rea_All:
		log.Println("房地产 解析")
		label, dataType, data, err := queryDataByFang(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针:", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label
	//	应收账款 旧数据处理
	case Pro_Rec_Ysk:
		log.Println("应收账款 解析")
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
		log.Println("尽调结果 解析")
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
		log.Println("尽调报告 解析")
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
		log.Println("动产质押 解析")
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
		log.Println("微分格数据解析")
		label, dataType, data, err := queryDataByWeiForapp(&Mg, resultByString[1], JiuWeifen)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label
	default:
		log.Println("There is no corresponding type")
		ChainTrans.Type = "default"
	}
	resData.ChainTransactionConfig = ChainTrans

	// 返回
	return resData, nil
}

//	复合查询  -根据txid  查询
func compoundQueryByTxidForUser(uuid string, th string, userId string) (rep interface{}, err error) {

	//	声明参数
	Mg := Kyc{}
	resData := TxIDrenData{}
	resDataJiu := TxIDrenDataJiu{}
	ChainTrans := ChainTransactionConfig{}
	wByType := WeiByType{}
	var strData string
	var JiuWeifen string

	// 根据txid查询数据
	result, resultByString, err := model_action.App.GetBlockTxId("", "", "", "", "", uuid)

	if err != nil {
		return nil, err
	}
	// 反序列化
	json.Unmarshal([]byte(result), &ChainTrans)

	//	拼接返回体
	ChainTrans.DataTxid = uuid
	ChainTrans.UserId = userId
	ChainTrans.Incident = th

	resData.ChainTransactionConfig = ChainTrans
	log.Println("根据txid查询数据:", result)

	//	判断
	if len(resultByString) == 0 {
		return nil, errors.New("resultByString is nil")
	}
	if len(resultByString) < 2 {
		log.Println("根据txid查询到的数据是微分格数据类型", 2)

		return resData, nil
	}

	//	异常情况处理完毕, 开始进行数据处理  kyc 数据在  retstring [1]   retstring [0] 类型      retstring [2] 标签

	// 数据解密

	//// 统一处理  判断是否是历史数据, 加标识判断
	//json.Unmarshal([]byte(result), &ChainTrans)
	//// 配置信息
	//resData.ChainTransactionConfig = ChainTrans

	// 如果 查询区块高于 725
	if ChainTrans.Height <= int64(800) {
		//
		JiuWeifen = "v1"
	} else {
		JiuWeifen = "v2"
	}
	//	数据转码

	//	数据本体
	err = json.Unmarshal([]byte(resultByString[1]), &Mg)
	if err != nil {
		log.Println("data transcoding failed:", err)
		return resData, nil
	}

	//	数据配置
	err = json.Unmarshal([]byte(resultByString[2]), &wByType)
	if err != nil {
		log.Println("data transcoding failed:", err)
		return resData, nil
	}

	log.Println("数据配置:", wByType.CategoryId)

	// TODO  判断是哪一种数据类型

	//	数据判断

	//	判断是否是加密类型数据上链类型
	if Mg.KycString != "" || Mg.SignKey != "" {
		//	如果是 加解密类型,开始解密
		strData = MetadataAesDecrypt(Mg.KycString, Mg.SignKey)
		log.Println("解密成功")
	} else {
		//	如果解密后的类型没有包含 kyc signkey,那就是微分格数据，赋值类型  微分格 数据类型 49
		Mg.KycType = Pro_Rec_WI
	}

	// 分发器 根据 KycType 解析对应的数据结构
	switch Mg.KycType {

	//	房地产
	case Pro_Rea_All:
		log.Println("房地产 解析")
		label, dataType, data, err := queryDataByFang(&Mg, strData)
		if err != nil {
			log.Println("获取对应的对象结构体指针:", err)
			return "", err
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label
	//	应收账款 旧数据处理
	case Pro_Rec_Ysk:
		log.Println("应收账款 解析")
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
		log.Println("尽调结果 解析")
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
		log.Println("尽调报告 解析")
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
		log.Println("动产质押 解析")
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
		log.Println("微分格数据解析")
		label, dataType, data, err := queryDataByWeiForapp(&Mg, resultByString[1], JiuWeifen)
		if err != nil {
			log.Println("获取对应的对象结构体指针=>", err)
		}
		ChainTrans.Type = dataType
		resData.Data = &data
		resData.DataType = label
	default:
		log.Println("There is no corresponding type")
		ChainTrans.Type = "default"
	}
	resData.ChainTransactionConfig = ChainTrans

	// 返回
	return resData, nil
}

//	复合查询  -根据区块hash查询  查询
func compoundQueryByBlockHash(uuid string) (rep interface{}, err error) {

	//  查询区块Hash
	return model_action.App.GetBlockHash(Chan_assetpublish, uuid)
	//  判断数据 GetBlockHash

	//  返回数据
}

//	复合查询  -根据区块高度查询  查询
func compoundQueryByBlockHeight(uuid string) (rep interface{}, err error) {

	// 查询区块高度数据
	return model_action.App.GetBlockHeight(Chan_assetpublish, "", "", uuid)
	//  判断数据

	// 返回数据
}

//	复合查询  -根据主键查询  查询
func compoundQueryByID(uuid string) (rep interface{}, err error) {

	// 根据主键查询
	log.Println("根据主键查询开始")
	ResInfo := ResInfo{}
	var row []FinanceList
	Resl := ResInfoDataByID{}
	var err3 error
	//  根据主键查询 微分格

	resultby1Up, err1 := model_action.QueryloadByBox("", Chan_assetpublish, Chan_weiFinanceTest2Up, Chan_weiFuncName, "", []string{uuid}, "")
	if err1 != nil {
		log.Println("根据主键查询 微分格 1 err", err)
	}
	log.Println("resultby1Up", resultby1Up)
	resultby2Up, err2 := model_action.QueryloadByBox("", Chan_assetpublish, Chan_weiFinanceTest1Up, Chan_weiFuncName, "", []string{uuid}, "")
	if err2 != nil {
		log.Println("根据主键查询 微分格 2 err", err)
	}
	log.Println("resultby2Up", resultby2Up)
	if err1 != nil && err2 != nil {
		// 如果两个都失败，说明没有这个key关联的value
		log.Println("两个都失败")
		return nil, err1
	}
	// 序列化数据
	if len(resultby1Up) > 0 {
		err3 = json.Unmarshal([]byte(resultby1Up), &ResInfo)
	} else {
		err3 = json.Unmarshal([]byte(resultby2Up), &ResInfo)
	}

	if err3 != nil {
		log.Println("两个都失败")
		return nil, err3
	}
	log.Println("ResInfo", ResInfo.Msg)
	log.Println("ResInfo", ResInfo.Status)
	//
	err5 := json.Unmarshal([]byte(ResInfo.Msg), &row)
	if err5 != nil {
		log.Println("二次序列化失败", err5)
		return nil, err5
	}
	// 根据区块信息获取交易数据
	for k, _ := range row {
		//	用户信息
		th, userId, err := model_api.SqlByQueryOntransactionID(row[k].TxId)
		if err != nil {
			log.Println("根据交易ID查询对应的用户信息失败", err)
		}
		log.Println("th========>", th)
		log.Println("userId========>", userId)
		//  根据txid 查询交易区块配置信息

		result, err := compoundQueryByTxidForUser(row[k].TxId, th, userId)
		log.Println("k=============>", result)
		if err != nil {
			log.Println("根据区块信息获取交易数据", err)
			return nil, err
		}

		Resl.Config = append(Resl.Config, result)
	}
	return Resl.Config, nil
}
