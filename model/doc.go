package action

import "fmt"

/**
model  链码相关的业务

*/

/**
@ 数据上链  || 客户端||通道 || 链码 || 节点 || 参数 ||
@ 2019年10月18日16:47:56
@ No:6002
@ 数据上链操作需要参数：   客户端 ,通道 || 链码  || 节点 || 参数

@ 首先客户端用sdk ，然后通道是通过参数指定的，但是目前来看通道仅仅只有几个，所以也写成几个参数写固定参数
链码名称 ，然后请求发送到哪个节点 ，以及最后参数


*/

/**
@	函数： 查看云图配置
@
*/

/**
### 数据库表  列表

### 数据库实例对象  SqlDB
### 数据库基础表  DbBase
### 1. 合约处理展示表  DbContract
### 2. 智能合约调用信息表  DbLarge
### 3. 共识节点信息表  DbLarge
### 4. 节点信息表  DbSvg
### 5. 区块信息表  DbBlock
### 6. 节点信息交易列表  DbDeal
### 7. 资产上链信息表  DbAsset

/**
| db_assets     |
| db_blocks     |
| db_contracts  |
| db_deals      |
| db_ips        |
| db_larges     |
| db_svgs

| db_assets     |
| db_blocks     |
| db_contracts  |
| db_deals      |
| db_ips        |
| db_larges     |
| db_svgs
*/
func test() {
	fmt.Println("test")
}

/*
@ 轮询器展示


1. 由于有很多数据去展示, 那么查询轮询去展示
2. 2019年12月17日16:32:30


###
生态节点：
	节点名称
	行业类型
	企业类别
	参与业务
	加入时间



*/
