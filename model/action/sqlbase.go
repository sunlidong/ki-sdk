package action

import (
	"ki-sdk/util/api"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// SQL --  获取 数据库实例

//
func getSqlDb() (sql *SqlDB) {
	//
	return &SqlDB{
		DbData:             SQL_DB,
		Login:              SQL_Login,
		PassWord:           SQL_PassWord,
		Http:               SQL_Http,
		Ip:                 SQL_Ip,
		DbName:             SQL_DbName,
		Charset:            SQL_Charset,
		ParseTime:          SQL_ParseTime,
		SetMaxIdleConns:    SQL_SetMaxIdleConns,
		SetMaxOpenConns:    SQL_SetMaxOpenConns,
		SetConnMaxLifetime: SQL_DB,
	}
}

// 初始化库表 7 张
func sqlInitByBase() {

	//先删除表
	sqlDel()
	sqlInitByDbContract() //合约处理展示表
	sqlInitByDbLarge()    //智能合约调用信息表
	sqlInitByDbIp()       //共识节点信息表
	sqlInitByDbSvg()      //节点信息表
	sqlInitByDbBlock()    //区块信息表
	sqlInitByDbDeal()     //节点信息交易列表
	sqlInitByDbAsset()    //资产上链信息表
}

// SQL -- 7 张表

// 	SQL -- table  -- 1
func sqlInitByDbContract() {
	if !DB.HasTable(&DbContract{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&DbContract{})
		log.Println("=>DbContract表 创建成功")
	} else {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&DbContract{})
		log.Println("=>DbContract表已经创建")
	}

}

// 	SQL -- table  -- 2
func sqlInitByDbLarge() {
	if !DB.HasTable(&DbLarge{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&DbLarge{})
		log.Println("=>DbLarge 表 创建成功")
	} else {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&DbLarge{})

		log.Println("=> DbLarge 表已经创建")
	}
}

// 	SQL -- table  -- 3
func sqlInitByDbIp() {
	if !DB.HasTable(&DbIp{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&DbIp{})
		log.Println("=>DbIp 表 创建成功")
	} else {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&DbIp{})

		log.Println("=> DbIp 表已经创建")
	}
}

// 	SQL -- table  -- 4
func sqlInitByDbSvg() {
	if !DB.HasTable(&DbSvg{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&DbSvg{})
		log.Println("=>DbSvg 表 创建成功")
	} else {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&DbSvg{})

		log.Println("=> DbSvg 表已经创建")
	}
}

// 	SQL -- table  -- 5
func sqlInitByDbBlock() {
	if !DB.HasTable(&DbBlock{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&DbBlock{})
		log.Println("=>DbBlock 表 创建成功")
	} else {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&DbBlock{})

		log.Println("=> DbBlock 表已经创建")
	}
}

// 	SQL -- table  -- 6
func sqlInitByDbDeal() {
	if !DB.HasTable(&DbDeal{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&DbDeal{})
		log.Println("=>DbDeal 表 创建成功")
	} else {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&DbDeal{})

		log.Println("=> DbDeal 表已经创建")
	}
}

// 	SQL -- table  -- 7
func sqlInitByDbAsset() {
	if !DB.HasTable(&DbAsset{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&DbAsset{})
		log.Println("=>DbAsset 表 创建成功")
	} else {
		DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&DbAsset{})

		log.Println("=> DbAsset 表已经创建")
	}
}

// SQL --  initData DbContract 1
func sqlByDataForDbContract() {
	//
	Pg := []DbContract{
		DbContract{ContractName: "数据导入", ContractType: "数据建模合约"},
		DbContract{ContractName: "分析建模", ContractType: "数据建模合约"},
		DbContract{ContractName: "物联网", ContractType: "数据建模合约"},
		DbContract{ContractName: "实施监控", ContractType: "数据建模合约"},
		DbContract{ContractName: "资产上链合约", ContractType: "数据建模合约"},
		DbContract{ContractName: "应收账款合约", ContractType: "数据建模合约"},
	}

	// 循环插入
	for k, _ := range Pg {
		DB.Create(&Pg[k])
	}
	//
}

// SQL --  initData DbContract 2
func sqlByDataForDbLarge() {
	//
	Pg := []DbLarge{
		DbLarge{ChannelName: "channelAssets", ChaincodeName: "chaincodeAssets", ChaincodeFuncName: "updateAsset", Num: int64(1)},
		DbLarge{ChannelName: "channelAssets", ChaincodeName: "chaincodeAssets", ChaincodeFuncName: "updateAsset", Num: int64(1)},
		DbLarge{ChannelName: "channelAssets", ChaincodeName: "chaincodeAssets", ChaincodeFuncName: "updateAsset", Num: int64(1)},
		DbLarge{ChannelName: "channelAssets", ChaincodeName: "chaincodeAssets", ChaincodeFuncName: "updateAsset", Num: int64(1)},
		DbLarge{ChannelName: "channelAssets", ChaincodeName: "chaincodeAssets", ChaincodeFuncName: "updateAsset", Num: int64(1)},
		DbLarge{ChannelName: "channelAssets", ChaincodeName: "chaincodeAssets", ChaincodeFuncName: "updateAsset", Num: int64(1)},
	}

	// 循环插入
	for k, _ := range Pg {
		DB.Create(&Pg[k])
	}
	//
}

// SQL --  initData DbDbIp 3
func sqlByDataForDbIp() {
	//
	Pg := []DbIp{
		DbIp{NodeName: "瑞泰格科技", NodeState: "正常", NodeAddress: "168.128.238.23", NodeCPU: "34%", NodeRAM: "13%"},
		DbIp{NodeName: "扬子保理", NodeState: "正常", NodeAddress: "168.128.238.23", NodeCPU: "64%", NodeRAM: "27%"},
		DbIp{NodeName: "金贸钢宝网", NodeState: "正常", NodeAddress: "168.128.238.22", NodeCPU: "18%", NodeRAM: "48%"},
		DbIp{NodeName: "观微科技", NodeState: "正常", NodeAddress: "168.128.138.24", NodeCPU: "59%", NodeRAM: "48%"},
		DbIp{NodeName: "顶象科技", NodeState: "正常", NodeAddress: "168.128.234.46", NodeCPU: "34%", NodeRAM: "59%"},
		DbIp{NodeName: "百望云科技", NodeState: "正常", NodeAddress: "168.128.583.73", NodeCPU: "56%", NodeRAM: "45%"},
	}

	// 循环插入
	for k, _ := range Pg {
		DB.Create(&Pg[k])
	}
	//
}

// SQL --  initData DbSvg 4
func sqlByDataForDbSvg() {
	//

	time1, err := time.ParseDuration("-1m")
	time2, err := time.ParseDuration("-2m")
	time3, err := time.ParseDuration("-3m")
	time4, err := time.ParseDuration("-4m")
	time5, err := time.ParseDuration("-5m")
	time6, err := time.ParseDuration("-6m")

	//
	log.Println("time=>", err)

	Pg := []DbSvg{
		DbSvg{ConNum: "1443", APINum: "234", TPS: "405", AverageProcessingTime: time.Now().Add(time1).Format("2006-01-02 15:04:05"), FormatTime: time.Now().Add(time1).Format("15:04")},
		DbSvg{ConNum: "2343", APINum: "434", TPS: "467", AverageProcessingTime: time.Now().Add(time2).Format("2006-01-02 15:04:05"), FormatTime: time.Now().Add(time2).Format("15:04")},
		DbSvg{ConNum: "3434", APINum: "458", TPS: "234", AverageProcessingTime: time.Now().Add(time3).Format("2006-01-02 15:04:05"), FormatTime: time.Now().Add(time3).Format("15:04")},
		DbSvg{ConNum: "4434", APINum: "236", TPS: "796", AverageProcessingTime: time.Now().Add(time4).Format("2006-01-02 15:04:05"), FormatTime: time.Now().Add(time4).Format("15:04")},
		DbSvg{ConNum: "4342", APINum: "657", TPS: "243", AverageProcessingTime: time.Now().Add(time5).Format("2006-01-02 15:04:05"), FormatTime: time.Now().Add(time5).Format("15:04")},
		DbSvg{ConNum: "3434", APINum: "343", TPS: "457", AverageProcessingTime: time.Now().Add(time6).Format("2006-01-02 15:04:05"), FormatTime: time.Now().Add(time6).Format("15:04")},
	}

	// 循环插入
	for k, _ := range Pg {
		DB.Create(&Pg[k])
	}
	//
}

// SQL --  initData DbSvg 5
func sqlByDataForDbBlock() {
	//
	Pg := []DbBlock{
		DbBlock{BlockHeight: "8499", BlockHash: "0x69229c4c5433bb1f19426e76f18120c1256b12e5cfd6d0441622a85b8dc41a0e", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
		DbBlock{BlockHeight: "8498", BlockHash: "0xbfa05f9c6609513c94fb032bf85266e2d3ad85e4e2cad1ebddb2930a94723185", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
		DbBlock{BlockHeight: "8497", BlockHash: "0x69229c4c5433bb1f19426e76f18120c1256b12e5cfd6d0441622a85b8dc41a0e", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
		DbBlock{BlockHeight: "8496", BlockHash: "0xb29ffc2b58a8cd48fb70405a61edff0427eaa1b3ff099d21ca6dec5052afe618", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
		DbBlock{BlockHeight: "8495", BlockHash: "0x3247cb0221face77e7541ecd9bf2dbb3ccf4f1fb19c28b80ca5bb3c9d72d9911", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
		DbBlock{BlockHeight: "8494", BlockHash: "0x1c5dcaf6cea0b91369d9012177f40949a8cf888060d529ab2f34572144f049ac", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
		DbBlock{BlockHeight: "8493", BlockHash: "0x1c5dcaf6cea0b91369d9012177f40949a8cf888060d529ab2f34572144f049ac", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
		DbBlock{BlockHeight: "8492", BlockHash: "0x1c5dcaf6cea0b91369d9012177f40949a8cf888060d529ab2f34572144f049ac", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
		DbBlock{BlockHeight: "8491", BlockHash: "0x1c5dcaf6cea0b91369d9012177f40949a8cf888060d529ab2f34572144f049ac", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests"},
	}

	// 循环插入
	for k, _ := range Pg {
		DB.Create(&Pg[k])
	}
	//
}

// SQL --  initData DbDeal 6
func sqlByDataForDbDeal() {
	//
	Pg := []DbDeal{
		DbDeal{BlockHeight: "8499", BlockHash: "0x69229c4c5433bb1f19426e76f18120c1256b12e5cfd6d0441622a85b8dc41a0e", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests", BlockByTXID: "channelAssets", UserTxt: "ac78ab9d-7577-416d-9557-9d410490b06b", ChainCodeName: "channelAssets", TXIDTime: time.Now().Format("2006-01-02 15:04:05"), TXIDType: "房地产", BlockByNode: "Peer1", BlockByOrg: "瑞泰格组织"},
		DbDeal{BlockHeight: "8498", BlockHash: "0xbfa05f9c6609513c94fb032bf85266e2d3ad85e4e2cad1ebddb2930a94723185", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests", BlockByTXID: "channelAssets", UserTxt: "d5d3bb8e-7363-4577-aec4-40ccc9b95ec7", ChainCodeName: "channelAssets", TXIDTime: time.Now().Format("2006-01-02 15:04:05"), TXIDType: "房地产", BlockByNode: "Peer1", BlockByOrg: "瑞泰格组织"},
		DbDeal{BlockHeight: "8497", BlockHash: "0x69229c4c5433bb1f19426e76f18120c1256b12e5cfd6d0441622a85b8dc41a0e", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests", BlockByTXID: "channelAssets", UserTxt: "127d429a-6c35-4da8-8570-a562b76491b1", ChainCodeName: "channelAssets", TXIDTime: time.Now().Format("2006-01-02 15:04:05"), TXIDType: "房地产", BlockByNode: "Peer1", BlockByOrg: "瑞泰格组织"},
		DbDeal{BlockHeight: "8496", BlockHash: "0xb29ffc2b58a8cd48fb70405a61edff0427eaa1b3ff099d21ca6dec5052afe618", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests", BlockByTXID: "channelAssets", UserTxt: "fe53c03b-75ef-4e33-ab06-f149f22774a5", ChainCodeName: "channelAssets", TXIDTime: time.Now().Format("2006-01-02 15:04:05"), TXIDType: "账款", BlockByNode: "Peer1", BlockByOrg: "瑞泰格组织"},
		DbDeal{BlockHeight: "8495", BlockHash: "0x3247cb0221face77e7541ecd9bf2dbb3ccf4f1fb19c28b80ca5bb3c9d72d9911", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests", BlockByTXID: "channelAssets", UserTxt: "fa4908f8-0e18-46a4-a343-fd5f5aec7f96", ChainCodeName: "channelAssets", TXIDTime: time.Now().Format("2006-01-02 15:04:05"), TXIDType: "账款", BlockByNode: "Peer1", BlockByOrg: "瑞泰格组织"},
		DbDeal{BlockHeight: "8494", BlockHash: "0x1c5dcaf6cea0b91369d9012177f40949a8cf888060d529ab2f34572144f049ac", BlockByChannel: "channelAssets", BlockByChainCode: "chainByAeests", BlockByTXID: "channelAssets", UserTxt: "57d7b43f-2bf7-4de8-b132-41ae1bd4f9c5", ChainCodeName: "channelAssets", TXIDTime: time.Now().Format("2006-01-02 15:04:05"), TXIDType: "账款", BlockByNode: "Peer1", BlockByOrg: "瑞泰格组织"},
	}

	// 循环插入
	for k, _ := range Pg {
		DB.Create(&Pg[k])
	}
}

// SQL --  initData DbAsset 7
func sqlByDataForDbAsset() {
	//
	Pg := []DbAsset{
		DbAsset{AssetNo: "3848", ConType: "房地产合约", DataUpChainCode: "上海房地产", UpTime: time.Now().Format("2006-01-02 15:04:05"), AssetType: "房产"},
		DbAsset{AssetNo: "3847", ConType: "应收账款合约", DataUpChainCode: "保理应收账款", UpTime: time.Now().Format("2006-01-02 15:04:05"), AssetType: "应收账款"},
		DbAsset{AssetNo: "3846", ConType: "动产融资合约", DataUpChainCode: "钢宝动产质押", UpTime: time.Now().Format("2006-01-02 15:04:05"), AssetType: "动产融资"},
		DbAsset{AssetNo: "3845", ConType: "动产融资合约", DataUpChainCode: "保理应收账款", UpTime: time.Now().Format("2006-01-02 15:04:05"), AssetType: "动产融资"},
		DbAsset{AssetNo: "3844", ConType: "动产融资合约", DataUpChainCode: "钢宝动产质押", UpTime: time.Now().Format("2006-01-02 15:04:05"), AssetType: "动产融资"},
		DbAsset{AssetNo: "3843", ConType: "房地产合约", DataUpChainCode: "上海房地产", UpTime: time.Now().Format("2006-01-02 15:04:05"), AssetType: "房地产"},
	}

	// 循环插入
	for k, _ := range Pg {
		DB.Create(&Pg[k])
	}
}

// SQL  -- DEL
func sqlDel() {
	DB.DropTable(&DbContract{})
	DB.DropTable(&DbLarge{})
	DB.DropTable(&DbIp{})
	DB.DropTable(&DbSvg{})
	DB.DropTable(&DbBlock{})
	DB.DropTable(&DbDeal{})
	DB.DropTable(&DbAsset{})
	//	### 1. 合约处理展示表  DbContract
	//	### 2. 智能合约调用信息表  DbLarge
	//	### 3. 共识节点信息表  DbIp
	//	### 4. 节点信息表  DbSvg
	//	### 5. 区块信息表  DbBlock
	//	### 6. 节点信息交易列表  DbDeal
	//	### 7. 资产上链信息表  DbAsset
}

// SQL  -- query block
func sqlQueryByBlock(blockHeight string) (rep interface{}, err error) {
	//
	sr := "ds"
	//log.Println("查询的区块高度是=>", blockHeight)
	//db_deals := []DbDeal{}
	//DB.Where(&DbDeal{BlockHeight: blockHeight}).Find(&db_deals)
	////TODO  判断是否查询到数据如果没有 去链上查询

	return sr, nil
}

// GetCpuSize 更新 服务器性能配置
func updateByIp() {
	// cpu
	cpu := api.GetCpuSize()

	// mem
	mem := api.GetMemSize()

	//
	log.Println("cpu=>", cpu)
	log.Println("mem=>", mem)

	// 更新表
	DB.Model(&DbIp{}).Updates(map[string]interface{}{"node_cpu": cpu + "%", "node_ram": mem + "%"})
}

func updateByIpBack() (cpu string, mem string) {
	// cpu
	cpu = api.GetCpuSize()

	// mem
	mem = api.GetMemSize()
	return cpu, mem
}

//	大屏@ 根据区块查询关联10个区块 @ 9 BlockData
func sqlQueryBlockByHeight(blockHeight string) (rep interface{}, err error) {
	//

	// TODO
	log.Println("查询的区块高度是=>", blockHeight)
	db_deals := []DbDeal{}
	DB.Where(&DbDeal{BlockHeight: blockHeight}).Find(&db_deals)
	//
	return db_deals, nil
}

//  SQL 查询 资产列表数据总量
func getDbAssetCount() (rep string, err error) {
	// TODO
	var count string
	DB.Table("db_assets").Count(&count)
	IntByCount, err := strconv.Atoi(count)
	if err != nil {
		log.Println("SQL 查询 资产列表数据总量=>", err)
		return "", err
	} else {
		rep = strconv.Itoa(IntByCount + int(1))
	}

	//
	return rep, nil
}

//  轮询器
func poller() {

	for {
		//
		randNum := rand.Intn(9)
		if randNum > 0 {
			for i := 0; i < randNum; i++ {
				SqlByLargeForInserData("poller", "chainName", "funcName", "org", "userName")
				SqlBySvgForInsertData("poller", "orgName", "userName")
			}

		}
		// 次数 和 api  调用次数
		time.Sleep(60 * time.Second)
	}

}

//	获取地理位置
func getCoordinate() interface{} {
	return []WebGeography{
		//WebGeography{No: "001", Name: "北京", Longitude: "116.46", Latitude: "39.92", Address: "北京", Title: "观微科技,百望云科技", Text: "暂无"},
		WebGeography{No: "002", Name: "南京", Longitude: "118.78", Latitude: "32.04", Address: "南京", Title: "金贸钢宝网,扬子保理", Text: "暂无"},
		//WebGeography{No: "003", Name: "杭州", Longitude: "120.19", Latitude: "30.26", Address: "杭州", Title: "顶象科技", Text: "暂无"},
	}
}

// 获取 生态节点列表
func getByEcologicalNode() (rep interface{}) {

	return []EcologicalNode{
		EcologicalNode{
			NodeName:          "扬子保理",
			TradeType:         "金融",
			Level:             "国企",
			TradeProfessional: "资金方",
			Time:              "2020年1月16日10:48:32",
		},
		EcologicalNode{
			NodeName:          "钢宝股份",
			TradeType:         "信息服务",
			Level:             "股份有限公司",
			TradeProfessional: "资金方",
			Time:              "2019年12月16日12:00:00",
		},
		EcologicalNode{
			NodeName:          "瑞泰格",
			TradeType:         "科技推广和应用",
			Level:             "有限责任公司",
			TradeProfessional: "监管方",
			Time:              "2019年12月16日12:00:00",
		},
		EcologicalNode{
			NodeName:          "百望云科技",
			TradeType:         "科技推广和应用",
			Level:             "有限责任公司",
			TradeProfessional: "参与方",
			Time:              "2019年12月16日12:00:00",
		},
		EcologicalNode{
			NodeName:          "顶象科技",
			TradeType:         "科技推广和应用",
			Level:             "有限责任公司",
			TradeProfessional: "参与方",
			Time:              "2019年12月16日12:00:00",
		},
		EcologicalNode{
			NodeName:          "观微科技",
			TradeType:         "科技推广和应用",
			Level:             "有限责任公司",
			TradeProfessional: "参与方",
			Time:              "2019年12月16日12:00:00",
		},
		EcologicalNode{
			NodeName:          "微分格",
			TradeType:         "科技推广和应用",
			Level:             "有限责任公司",
			TradeProfessional: "参与方",
			Time:              "2019年12月16日12:00:00",
		},
	}
}

// 获取 共识节点列表
func getByCommonNode(cpu string, mem string) (rep interface{}) {
	cpub := cpu + "%"
	memb := mem + "%"
	return []CommonNode{
		CommonNode{NodeId: "orderer0.dinglian.com", NodeName: "orderer", NodeType: "排序节点", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"},
		CommonNode{NodeId: "orderer1.dinglian.com", NodeName: "orderer", NodeType: "排序节点", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"},
		CommonNode{NodeId: "orderer2.dinglian.com", NodeName: "orderer", NodeType: "排序节点", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"},
		CommonNode{NodeId: "orderer3.dinglian.com", NodeName: "orderer", NodeType: "排序节点", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"},
		CommonNode{NodeId: "orderer4.dinglian.com", NodeName: "orderer", NodeType: "排序节点", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"},
		CommonNode{NodeId: "peer0.org1.dinglian.com", NodeName: "org1", NodeType: "背书", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"},
		CommonNode{NodeId: "peer1.org1.dinglian.com", NodeName: "org1", NodeType: "背书", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"},
		CommonNode{NodeId: "peer0.org2.dinglian.com", NodeName: "org2", NodeType: "背书", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"},
		CommonNode{NodeId: "peer1.org2.dinglian.com", NodeName: "org2", NodeType: "背书", NodeState: "在线", Ip: "161.117.0.57", Cpu: cpub, Mem: memb, Time: "2020年1月16日11:03:42"}}

}

// 获取 分布式存储节点列表
func getByDistributedStorageNode() (rep interface{}) {
	//	分布式储存节点
	return []DistributedStorageNode{
		DistributedStorageNode{NodeName: "金山云", StoreType: "分布式对象存储", Address: "161.117.0.57", Percent: "24%", Time: "2020年1月16日11:03:42"},
		DistributedStorageNode{NodeName: "金山云", StoreType: "分布式对象存储", Address: "161.117.0.57", Percent: "39%", Time: "2020年1月16日11:03:42"},
		DistributedStorageNode{NodeName: "金山云", StoreType: "IPFS", Address: "161.117.0.57", Percent: "65%", Time: "2020年1月16日11:03:42"},
		DistributedStorageNode{NodeName: "金山云", StoreType: "分布式对象存储", Address: "161.117.0.57", Percent: "45%", Time: "2020年1月16日11:03:42"},
		DistributedStorageNode{NodeName: "金山云", StoreType: "IPFS", Address: "161.117.0.57", Percent: "73%", Time: "2020年1月16日11:03:42"},
		DistributedStorageNode{NodeName: "金山云", StoreType: "IPFS", Address: "161.117.0.57", Percent: "27%", Time: "2020年1月16日11:03:42"},
		DistributedStorageNode{NodeName: "金山云", StoreType: "分布式对象存储", Address: "161.117.0.57", Percent: "16%", Time: "2020年1月16日11:03:42"},
		DistributedStorageNode{NodeName: "金山云", StoreType: "IPFS", Address: "161.117.0.57", Percent: "29%", Time: "2020年1月16日11:03:42"},
		DistributedStorageNode{NodeName: "金山云", StoreType: "分布式对象存储", Address: "161.117.0.57", Percent: "38%", Time: "2020年1月16日11:03:42"},
	}

}

//
func getCoordinateByeco() interface{} {
	return []WebGeography{
		//WebGeography{No: "001", Name: "北京", Longitude: "116.46", Latitude: "39.92", Address: "北京", Title: "观微科技,百望云科技", Text: "暂无"},
		WebGeography{No: "002", Name: "南京", Longitude: "118.78", Latitude: "32.04", Address: "南京", Title: "金贸钢宝网,扬子保理", Text: "暂无"},
		//WebGeography{No: "003", Name: "杭州", Longitude: "120.19", Latitude: "30.26", Address: "杭州", Title: "顶象科技", Text: "暂无"},
	}
}

func getCoordinateBycom() interface{} {
	return []WebGeography{
		//WebGeography{No: "001", Name: "北京", Longitude: "116.46", Latitude: "39.92", Address: "北京", Title: "观微科技,百望云科技", Text: "暂无"},
		WebGeography{No: "002", Name: "南京", Longitude: "118.78", Latitude: "32.04", Address: "南京", Title: "金贸钢宝网,扬子保理", Text: "暂无"},
		//WebGeography{No: "003", Name: "杭州", Longitude: "120.19", Latitude: "30.26", Address: "杭州", Title: "顶象科技", Text: "暂无"},
	}
}

func getCoordinateBydis() interface{} {
	return []WebGeography{
		//WebGeography{No: "001", Name: "北京", Longitude: "116.46", Latitude: "39.92", Address: "北京", Title: "观微科技,百望云科技", Text: "暂无"},
		WebGeography{No: "002", Name: "南京", Longitude: "118.78", Latitude: "32.04", Address: "南京", Title: "金贸钢宝网,扬子保理", Text: "暂无"},
		//WebGeography{No: "003", Name: "杭州", Longitude: "120.19", Latitude: "30.26", Address: "杭州", Title: "顶象科技", Text: "暂无"},
	}
}
