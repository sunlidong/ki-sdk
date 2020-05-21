package action

import "fmt"

/**
contrroller 处理   前期业务
*/
/**
数据上链 函数处理

数据上链函数解析：调用  model层的 函数处理
*/

//	 房地产数据结构定义   房地产标识
/**
@	房地产数据结构定义
@	房地产主表	 DueProject
@	房地产表	 DueProjectFdc
@	保证人	 	 DueProjectCemGt
@	抵押品	 	 DueProjectCemGt
@	质押品	 	 DueProjectCemPg
@
*/

//assetType_DueProject      string = "20" // 房地产主表
//assetType_DueProjectFdc   string = "21" // 房地产信息
//assetType_DueProjectCemGt string = "22" //保证人 结构体
//assetType_DueProjectCemMg string = "23" //抵押品 结构体
//assetType_DueProjectCemPg string = "24" //质押品 结构体

//		公共字段定义    base user  power
/**
@	公共数据结构定义
@	base 标签，标识数据的基础信息     base
@	user 标签，标识数据的用户信息     user
@	power 标签，标识数据的权限信息    power

*/

/*
@	描述： 此函数是 获取 用户的key , getEncryptionKeytest
@	参数：
			orgKey：组织访问Key
			userKey：用户访问Key
			publicKey：公共访问Key
func getEncryptionKeytest(orgKey string,userKey string,publicKey string)(){
	//
}

*/

// 加密解密流程解析：

/**
###  流程解析 ：

概述 ： 数据加解密是为了数据的安全以及隐私，因为无论是通道隔离还是证书校验，只要存在明文数据上链就会
存在 数据泄露,因为fabric 的机制是以通道隔离的，并且不是以单条数据隐私为控制的，他只能做到对数据进行上链保存，不会对数据本身加密，
这样就会存在问题，一些风控系统是需要数据保密的，传统业务系统的中很多企业需要的数据上链，但是数据的保密性也是需要的，权限分级是非常急迫
那么，这种技术就需要自己来处理，后来根据需求定制化需求 ，对上链数据进行上链加密，查询解析，权限标签


基本逻辑 ：

	1. 数据上链

			前段路由请求 >  加密级别（公开，组织，个人，绝密） => 后台根据设置 对数据加密，=> 保存加密key 到 数据 和标签  => 数据上链

	2. 数据查询

			前端路由请求=> 主键查询 = > 先根据主键去查询这个用的 id 和orgid 是否在 该主键的访问组里面，如果在访问，如果没有在|| 是否是数据所有者

	=> 查询数据 返回数据 = > 根据数据进行返回  ||  数据展示, 展示两种情况 || 该数据已经加密 ,您没有访问你权限 || 该数据您有访问权限，直接显示明文数据





核心业务：

			1.  元数据 ： 数据本身要有,一切的操作都是基于原来数据的

			2.	数据 : 数据本身的信息(加密数据，key,类型，用户信息等 )  标签:标识数据的访问策略以及用户等信息() ，数据控制数据，标签控制谁操作数据

			3.	访问组控制策略： 指定谁能来访问数据，由管理员指定

			4.	加密控制策略：加密控制级别

			5.	修改机制策略： 修改添加和删除访问或者数据

			6.	加密key策略：	加密key 的由来， 加密中心  ||  hash || id||

			7.	访问码策略:	 单次访问数据， 访问时间限制，访问次数 等等 细节权限 订制

			8.	容灾备份


*/

// 上链数据标识

/**
1. 应收账款  ||



*/

func test() {
	fmt.Println("test")
}

/*
 type  asset_type


2020年1月17日14:50:45 ：
测试链码名称： chaincodeby01

*/
