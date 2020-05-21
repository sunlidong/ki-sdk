package action

import (
	"errors"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// sql  处理函数合约
type UserTable1 struct {
	gorm.Model // 主键
	Name       string
	Age        int64
	Birthday   string
	Text       string
	Num        int `gorm:"AUTO_INCREMENT"` // 自增
}

//  `gorm:"type:varchar(100);unique"`

type UserInfo231 struct {
	gorm.Model // 主键
	Name       string
	Age        int64
	Birthday   string
	Text       string
	Num        int `gorm:"AUTO_INCREMENT"` // 自增
}

// 初始化数据库 test
func sqlByInit() {
	//	1. 连接数据库
	err := sqlByInitByDB()
	if err != nil {
		log.Println("@初始化数据库=>创建数据库全局对象失败")
	}
	log.Println("@初始化数据库=>创建数据库全局对象成功")

	// 	2.  初始化库表
	sqlInitByBase()
	log.Println("@初始化库表=>初始化库表成功")

	// 	3.  初始化Demo 数据
	sqlInitByData()
	log.Println("@初始化Demo=>初始化Demo成功")
}

func checkErr(err error) {
	if err != nil {
		log.Println("err=>", err)
		err = nil
	}

}

//
func valfal(state bool) {
	if state == false {
		log.Println("state=>", state)
	} else {
		log.Println("state=>", state)
	}

}

//  SQL -- 初始化 Mysql 全局对象  DB
func sqlByInitByDB() (err error) {
	//
	//Pig := getSqlDbByone()
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/drc?charset=utf8mb4&parseTime=true")
	//
	if err != nil {
		return err
	}
	db.LogMode(true)
	//设置连接池
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
	//空闲
	db.DB().SetMaxIdleConns(300)
	//打开
	db.DB().SetMaxOpenConns(500)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	log.Println("@=> 初始化 Mysql 全局对象  DB")
	//
	return nil
}

// SQL -- 初始化数据
func sqlInitByData() {
	//
	sqlByDataForDbContract()
	sqlByDataForDbLarge()
	sqlByDataForDbIp()
	sqlByDataForDbSvg()
	sqlByDataForDbBlock()
	sqlByDataForDbDeal()
	sqlByDataForDbAsset()
}

//	大屏@合约处理展示 @ 1
func sqlByConact() (rep interface{}, err error) {

	//	01. DB
	db_contracts := []DbContract{}

	//	02. 查询数据
	DB.Order(("id desc")).Limit(5).Find(&db_contracts)
	//
	log.Println("@大屏@合约处理展示查询")

	//返回
	return db_contracts, nil
}

//	大屏@ 智能合约调用信息表 @ 2
func sqlByLarge() (rep interface{}, err error) {

	//	01. DB
	var count string
	DB.Table("db_larges").Count(&count)

	//	02. 查询数据

	//
	log.Println("2 @大屏@智能合约调用信息表")

	//返回
	return count, nil
}

//	大屏@ 共识节点信息表 @ 3
func sqlByServer() (rep interface{}, err error) {

	//	01. DB
	db_ips := []DbIp{}

	//	02. 查询数据
	DB.Order(("id desc")).Limit(5).Find(&db_ips)
	//
	log.Println("3 @大屏@ 共识节点信息表")
	//返回
	return db_ips, nil
}

//	大屏@ 节点信息统计表 @ 4
func sqlBySvg() (rep interface{}, err error) {
	//TODO
	db_svgsList := []DbSvg{}

	log.Println("start svg sql")
	index := 1
	for i := 0; i < 5; i++ {
		db_svgs := []DbSvg{}
		log.Println("i=>", i)
		//
		timeSp, err := time.ParseDuration("-" + strconv.Itoa(i) + "m")
		timeIn, err := time.ParseDuration("-" + strconv.Itoa(i+index) + "m")
		//

		// TODO 时间
		log.Println("小于=>", time.Now().Add(timeSp).Format("2006-01-02 15:04:05"))
		log.Println("大于=>", time.Now().Add(timeIn).Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Println("err=>", err)
		}
		DB.Find(&db_svgs, " average_processing_time < ?  AND average_processing_time > ?", time.Now().Add(timeSp).Format("2006-01-02 15:04:05"), time.Now().Add(timeIn).Format("2006-01-02 15:04:05"))
		db_svgsList = append(db_svgsList, DbSvg{
			FormatTime: time.Now().Add(timeSp).Format("15:04"),
			ConNum:     strconv.Itoa(len(db_svgs)),
			APINum:     strconv.Itoa(len(db_svgs)),
		})
	}

	log.Println("end  svg sql")

	//
	a := "68"
	t := "2000"
	log.Println("4 @大屏@ 节点信息统计表")
	//返回
	return SqlJie{
		API:  a,
		TPS:  t,
		Data: db_svgsList,
	}, nil
}

//	大屏@ 区块信息表 @ 5
func sqlByBlock() (rep interface{}, err error) {

	//	01. DB
	db_blocks := []DbBlock{}

	//	02. 查询数据
	DB.Order(("id desc")).Limit(500).Find(&db_blocks)
	//
	log.Println("5 @大屏@ 区块信息表")
	//返回
	return db_blocks, nil
}

//	大屏@ 节点信息交易列表 @ 6
func sqlByDeal() (rep interface{}, err error) {

	//	01. DB
	db_deals := []DbDeal{}

	//	02. 查询数据
	DB.Order(("id desc")).Limit(5).Find(&db_deals)
	//
	log.Println("6 @大屏@ 节点信息交易列表")
	//返回
	return db_deals, nil
}

//	大屏@ 资产上链信息表 @ 7
func sqlByAsset() (rep interface{}, err error) {

	//	01. DB
	db_assets := []DbAsset{}

	Pg := SqlAssets{}
	var count string

	//	02. 查询数据

	// 统计
	DB.Table("db_assets").Count(&count)

	//	5条记录
	DB.Order(("id desc")).Limit(5).Find(&db_assets)

	//
	log.Println("7 @大屏@ 资产上链信息表")
	//返回count
	Pg.Sum = count
	Pg.Data = db_assets

	return Pg, nil
}

// SQL --- 智能合约 调用次数
func sqlByLargeForInserData(cName string, chainName string, funcName string, org string, userName string) {
	//
	Pg := DbLarge{
		ChannelName:       cName,
		ChaincodeName:     chainName,
		ChaincodeFuncName: funcName,
		Num:               int64(1),
		CreateOrg:         org,
		CreateUserName:    userName,
	}

	DB.Create(&Pg)

	//
	log.Println("@sql## 调用次数 调用成功")
}

// SQL --- 区块信息表
func sqlByBlockForInserData(height string, hash string, channel string, chaincode string, orgName string, userName string) {
	//
	Pg := DbBlock{
		BlockHeight:      height,
		BlockHash:        hash,
		BlockByChannel:   channel,
		BlockByChainCode: chaincode,
		CreateOrg:        orgName,
		CreateUserName:   userName,
	}

	DB.Create(&Pg)

	//
	log.Println("@sql## 区块信息表 插入成功")
}

// SQL --- 合约调用次数  插入
func sqlBySvgForInsertData(txId string, orgName string, userName string) {
	//
	Pg := DbSvg{
		ConNum:                "1",
		APINum:                "1",
		CreateOrg:             orgName,
		CreateUserName:        userName,
		Type:                  txId,
		AverageProcessingTime: time.Now().Format("2006-01-02 15:04:05"), // TODO  时间
	}

	DB.Create(&Pg)

	//
	log.Println("@sql## 区块信息表 插入成功")
}

// SQL --- 节点交易信息表
func sqlByDbDealForInsertData(txid string, height string, userID string, chainCode string, orgName string, userName string) {
	//
	Pg := DbDeal{
		BlockByTXID:      txid,
		BlockHeight:      height,
		UserTxt:          userID,
		BlockByChainCode: chainCode,
		CreateOrg:        orgName,
		CreateUserName:   userName,
		TXIDTime:         time.Now().Format("2006-01-02 15:04:05"),
	}
	DB.Create(&Pg)
	//
	log.Println("@sql## 节点交易信息表 插入成功")
}

// SQL --- 资产信息表 插入
func sqlByDbAssetForInsertData(assetNo string, conType string, chainCode string, upTime string, assetType string, orgName string, userName string, txid string) {
	//

	rep, err := getDbAssetCount()
	if err != nil {
		log.Println("SQL --- 资产信息表 插入=>", rep)
		return
	}
	Pg := DbAsset{
		AssetNo:         rep, // TODO  资产数量统计
		ConType:         conType,
		DataUpChainCode: chainCode,
		UpTime:          time.Now().Format("2006-01-02 15:04:05"),
		AssetType:       txid,
		CreateOrg:       orgName,
		CreateUserName:  userName,
	}
	DB.Create(&Pg)
	//
	log.Println("@sql## 资产信息表 插入成功")
}

//	大屏@ 节点信息交易列表 @ 6 平局值 2019年12月17日13:48:16
func sqlBySvgForBig() (rep interface{}, err error) {

	//TODO
	db_svgsList := []DbSvg{}
	index := 1
	//
	channel := "assetpublish"
	org := "d"
	user := "d"
	currentTime, err := App.GetsTheCreationBlockTime(channel, org, user)
	// 初始块时间
	if err != nil {
		return nil, err
	}
	log.Println("初始块的时间=>", currentTime)
	timeCurrentTime, err := time.Parse("2006-01-02 15:04:05", currentTime)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 5; i++ {
		db_svgs := []DbSvg{}
		log.Println("i=>", i)
		//
		timeSp, err := time.ParseDuration("-" + strconv.Itoa(i) + "m")
		timeIn, err := time.ParseDuration("-" + strconv.Itoa(i+index) + "m")

		// TODO 时间

		floatsum := time.Now().Add(timeSp).Sub(timeCurrentTime)

		//
		log.Println("减去的时间结果：=>", floatsum)

		log.Println("小于=>", time.Now().Add(timeSp).Format("2006-01-02 15:04:05"))
		log.Println("大于=>", time.Now().Add(timeIn).Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Println("err=>", err)
		}
		//DB.Find(&db_svgs, " average_processing_time < ?  AND average_processing_time > ?", time.Now().Add(timeSp).Format("2006-01-02 15:04:05"), time.Now().Add(timeIn).Format("2006-01-02 15:04:05"))
		DB.Find(&db_svgs, " average_processing_time < ? ", time.Now().Add(timeSp).Format("2006-01-02 15:04:05"))

		//TODO
		log.Println("单项查询的数据 =>", len(db_svgs))
		db_svgsList = append(db_svgsList, DbSvg{
			FormatTime: time.Now().Add(timeSp).Format("15:04"),
			ConNum:     strconv.Itoa(len(db_svgs)),
			APINum:     strconv.Itoa(len(db_svgs)),
		})
		//
		floatsvg, _ := strconv.ParseFloat(strconv.Itoa(len(db_svgs)), 64)
		//
		log.Println("转换的条数=>", floatsvg)
		result := floatsvg / floatsum.Minutes()

		log.Println("result=>", result)
	}

	a := "68"
	t := "2000"
	log.Println("4 @大屏@ 节点信息统计表")
	//返回
	return SqlJie{
		API:  a,
		TPS:  t,
		Data: db_svgsList,
	}, nil
}

//	大屏@ 地理位置列表 2019年12月26日16:21:33
func sqlBySvgForGeography() (rep interface{}, err error) {
	return getCoordinate(), nil
}

// 查询资产后五个
func sqlByAssetNext(no string, t string) (rep interface{}, err error) {

	log.Println("查询编号：", no)
	log.Println("查询方向：", t)
	db_assets := []DbAsset{}
	var rp int

	var rps int

	Pg := SqlAssets{}
	var count string
	var stateLabel string

	//
	switch t {
	case "pre":
		stateLabel = "p" // 前
	case "next":
		stateLabel = "n" //后
	default:
		stateLabel = "f" // 第一次查询
		log.Println("第一次查询")
	}

	DB.Table("db_assets").Count(&count)
	// 首先判断是否是最新的五个
	rpcount, err := strconv.Atoi(count)
	if err != nil {
		log.Println("rp err:", err)
		return nil, nil
	}
	if no != "" {
		rp, err = strconv.Atoi(no)
		if err != nil {
			log.Println("rp err:", err)
			return nil, nil
		}
	}

	log.Println("rpcount:", rpcount)
	log.Println("rp:", rp)
	log.Println(" (rpcount-rp) <= 5", rpcount-rp)

	// 对比 如果对比结果小于5，返回最新的五个数据
	if stateLabel == "f" {
		log.Println("初始查询")
		DB.Order(("id desc")).Limit(5).Find(&db_assets)
	} else if rpcount-rp <= int(5) && stateLabel == "p" {
		log.Println("最新的数据小于5，返回最新的5条数据，上一个判断")
		DB.Order(("id desc")).Limit(5).Find(&db_assets)
	} else if rpcount-rp >= rpcount-5 && stateLabel == "n" {
		log.Println("查询创世块前5个，只要小于5，就直接返回5个")
		DB.Order(("id asc")).Limit(5).Find(&db_assets)
	} else if rpcount-rp >= 5 || (rpcount-rp) <= rpcount-5 {
		log.Println("获取的是中间的数据，正常查询五条数据")
		for i := 1; i <= 5; i++ {
			db_assets_row := DbAsset{}
			switch stateLabel {
			case "p":
				//  6不是5
				rps = rp + 6 - i
			case "n":
				rps = rp - i
			default:
				rps = rp
			}
			log.Println("查询的资产编号：", strconv.Itoa(rps))
			DB.Where("asset_no = ?", strconv.Itoa(rps)).Find(&db_assets_row)
			db_assets = append(db_assets, db_assets_row)
			//
			log.Println("row:", db_assets_row.AssetNo)
		}
	}

	log.Println("7 @大屏@ 资产上链信息表")
	//返回count
	Pg.Sum = count
	Pg.Data = db_assets

	return Pg, nil
}

//	大屏@  共识节点  三个Tab  2020年1月16日09:45:26
func sqlByDealList(id string) (rep interface{}, err error) {

	Pg := SqlAssets{}
	var count string

	//	获取总条数
	DB.Table("db_assets").Count(&count)
	Pg.Sum = count

	// 获取服务器参数
	cpu, mem := UpdateByIpBack()

	//	搅拌机
	switch id {
	//	生态节点
	case echoByEcologicalNode:
		//	共识节点
		Pg.Data = getByEcologicalNode()
	case echoCommonNode:
		//	分布式存储节点
		Pg.Data = getByCommonNode(cpu, mem)
	case echoByDistributedStorageNode:
		Pg.Data = getByDistributedStorageNode()
	default:
		log.Println("There is no alternative")
		return nil, errors.New("There is no alternative")
	}

	return Pg, nil
}

//	 根据交易ID  查询 用户信息以及  事件
func sqlByQueryOntransactionID(txid string) (things string, userId string, err error) {

	row := DbAsset{}
	DB.Where("asset_type = ?", txid).Find(&row)
	if row.AssetType == "" {
		return "", "", errors.New("AssetType is nil")
	}

	return row.DataUpChainCode, row.CreateUserName, nil
}

// 查询上链资产数量
func sqlByAssetSum() (rep interface{}, err error) {
	var count string
	DB.Table("db_assets").Count(&count)

	return count, nil
}

//	大屏@  地理位置  三个Tab  2020年1月16日09:45:26
func sqlByDealGeography(id string) (rep interface{}, err error) {

	Pg := SqlAssets{}
	var count string

	//	获取总条数
	DB.Table("db_assets").Count(&count)
	Pg.Sum = count

	//	搅拌机
	switch id {
	//	生态节点
	case echoByEcologicalNode:
		//	共识节点 getCoordinateByeco
		Pg.Data = getCoordinateByeco()
	case echoCommonNode:
		//	分布式存储节点
		Pg.Data = getCoordinateBycom()
	case echoByDistributedStorageNode:

		Pg.Data = getCoordinateBydis()
	default:
		log.Println("There is no alternative")
		return nil, errors.New("There is no alternative")
	}

	return Pg, nil
}
