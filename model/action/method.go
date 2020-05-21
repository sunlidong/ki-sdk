package action

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/chainHero/blockdata"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric/protos/common"
	futils "github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
	"time"
)

/**
  结构分为两部分
		||  方法实现
		||
*/

/**
@ 初始化 SDK
@
*/
func (setup *SDK) Initialize() error {
	// 01. 读取配置文件config.yaml
	configProvider := config.FromFile(
		setup.ConfigFile,
	)
	// 02. 创建sdk对象
	sdk, err := fabsdk.New(configProvider)
	// 03. 验错
	if err != nil {
		return fmt.Errorf("SDK实例化失败:%v", err)
	}
	// 04. 参数转换至结构体对象 setup
	setup.SDK = sdk
	// 05. 返回
	return nil
}

/**
@ 初始化 msg
@
*/
func (setup *SDK) CreateresMgmtClient() error {
	// 01. 创建资源管理客户端上下文
	resourceManagerClientContext :=
		setup.SDK.Context(fabsdk.WithUser(setup.OrgAdmin),
			fabsdk.WithOrg(setup.OrgName))
	// 02. 创建资源管理客户端
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	// 03. 验错
	if err != nil {
		return fmt.Errorf("创建资源管理客户端失败:%v", err)
	}
	// 04. 参数转换至结构体对象 setup
	setup.resmgmt = resMgmtClient
	// 05. 返回
	return nil
}

/**
@ 初始化 channel cli
@
*/
func (setup *SDK) CreateChannelCli() error {

	// 01. 封装数据Channle cli
	clientContext := setup.SDK.ChannelContext(
		setup.ChannelID,
		fabsdk.WithUser(setup.UserName))

	// 02. 创建Channle cli
	channelCli, err := channel.New(clientContext)

	// 03. 验错
	if err != nil {
		return fmt.Errorf("创建通道管理客户端失败:%v", err)
	}

	// 04. 参数转换至结构体对象 setup
	setup.client = channelCli
	// 05. 返回
	return nil
}

/**
@ 初始化 msp cli
@
*/
func (setup *SDK) CreateMspClient() error {

	// 01. 创建资源管理客户端上下文
	clientCTX := setup.SDK.Context(
		fabsdk.WithUser(setup.OrgAdmin),
		fabsdk.WithOrg(setup.OrgName),
	)

	// 02. 创建资源实例
	c, err := msp.New(clientCTX)

	if err != nil {
		return fmt.Errorf("创建msp管理客户端失败:%v", err)
	}

	if c == nil {
		return fmt.Errorf("创建msp管理客户端为空:%v", c)
	}

	setup.MspClient = c

	return nil
}

/**
@ 数据上链  || 客户端||通道 || 链码 || 节点 || 参数 ||
@ 2019年10月18日16:47:56
@ No:6002
*/
func uploadAsset(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	//
	//	01. client TODO
	//	02. channel TODO
	//	03. chaincode TODO
	//	04. peer

	// peer
	//peerlist := Cpeer(1)
	var peer111 []string
	// 02. 设置背书节点
	peer111 = append(peer111, Peerproduct)
	peer111 = append(peer111, Peerfactoring)
	//fcn, arglist := action.Cargsplict(args)
	//log.Println("fcn=>", fcn)
	//log.Println("arg1:=>", string(arglist[0]))

	//	 args split
	// 03. 拼接请求对象 request
	//request := channel.Request{
	//	ChaincodeID: ,
	//	Fcn:         fcn,
	//	Args:        arglist,
	//}
	request := channel.Request{
		ChaincodeID: ConfChainCodeID,
		Fcn:         args[0],
		Args: [][]byte{
			[]byte(args[1]),
			[]byte(args[2]),
			[]byte(args[3]),
		}}
	//  client =>App.SDK.client

	response, err := App.SDK.client.Execute(
		request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peer111...),
	)
	// 05. 判断错误
	if err != nil {
		//
		log.Println("err:", err)

		return "", err
	}
	// 06. 返回调用结果
	return string(response.Payload), nil

}

func uploadAssettest(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	//
	//	01. client TODO
	//	02. channel TODO
	//	03. chaincode TODO
	//	04. peer

	// peer
	//peerlist := Cpeer(1)
	var peer111 []string
	// 02. 设置背书节点
	peer111 = append(peer111, Peerproduct)
	peer111 = append(peer111, Peerfactoring)
	//fcn, arglist := action.Cargsplict(args)
	//log.Println("fcn=>", fcn)
	//log.Println("arg1:=>", string(arglist[0]))

	//	 args split
	// 03. 拼接请求对象 request
	//request := channel.Request{
	//	ChaincodeID: ,
	//	Fcn:         fcn,
	//	Args:        arglist,
	//}
	request := channel.Request{
		ChaincodeID: ConfChainCodeID,
		Fcn:         args[0],
		Args: [][]byte{
			[]byte(args[1]),
			[]byte(args[2]),
			[]byte(args[3]),
		}}
	//  client =>App.SDK.client

	response, err := App.SDK.client.Execute(
		request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peer111...),
	)
	// 05. 判断错误
	if err != nil {
		//
		log.Println("err:", err)

		return "", err
	}
	// 06. 返回调用结果
	return string(response.Payload), nil

}

/**
@ 数据查询  || 客户端||通道 || 链码 || 节点 || 参数 ||
@ 2019年10月18日16:47:56
@ No:6002
*/
func QueryById(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	//
	//	01. client TODO
	//	02. channel TODO
	//	03. chaincode TODO
	//	04. peer

	// peer
	//peerlist := Cpeer(1)
	var peer111 []string
	// 02. 设置背书节点
	peer111 = append(peer111, Peerproduct)
	peer111 = append(peer111, Peerfactoring)
	//fcn, arglist := action.Cargsplict(args)
	//log.Println("fcn=>", fcn)
	//log.Println("arg1:=>", string(arglist[0]))

	//	 args split
	// 03. 拼接请求对象 request
	//request := channel.Request{
	//	ChaincodeID: ,
	//	Fcn:         fcn,
	//	Args:        arglist,
	//}
	request := channel.Request{
		ChaincodeID: ConfChainCodeID,
		Fcn:         args[0],
		Args: [][]byte{
			[]byte(args[1]),
		}}
	//  client =>App.SDK.client

	response, err := App.SDK.client.Execute(
		request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peer111...),
	)
	// 05. 判断错误
	if err != nil {
		//
		log.Println("err:", err)

		return "", err
	}
	// 06. 返回调用结果
	return string(response.Payload), nil

}

/**
@	函数： 根据TXID获取 区块链属于交易信息
@	时间：	2019年10月21日17:58:36
@	描述：	参数：  txid
*/
func (setup *Application) GetnewBlockTxId(txid string) ([]*blockdata.ChainTransaction, error) {

	// 声明  peer 对等节点
	var peer []string
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)

	// 交易
	var transaction = make([]*blockdata.ChainTransaction, 0)

	//  sdk 本身  调用 sdk   通道 || 组织名称 || user
	ctx := setup.SDK.SDK.ChannelContext(setup.SDK.ChannelID, fabsdk.WithOrg(setup.SDK.OrgName), fabsdk.WithUser(setup.SDK.UserName))

	// cli 客户端
	cli, err := ledger.New(ctx)
	if err != nil {
		return nil, err
	}

	// 查询 txid 所在区块
	block, err := cli.QueryBlockByTxID(fab.TransactionID(txid), ledger.WithTargetEndpoints(peer...))
	if err != nil {
		//
		log.Println("查询交易ID出错", err)
		return nil, err
	}

	// 查询当前最新区块信息
	lastBlock, err1 := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))

	if err1 != nil {
		log.Println("查询最新区块出错", err)
		return nil, err
	}

	//这里要做下处理 。就是根据交易ID查询的时候，要判断是否是最新的区块，如果是最新的区块，区块Hash就是要变化
	//此处应该遍历block.Data.Data
	indexOne, _ := strconv.ParseUint(strconv.Itoa(1), 10, 64)
	for _, data := range block.Data.Data {
		env, err := futils.GetEnvelopeFromBlock(data)
		row, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		times := time.Unix(row.Timestamp, 0).Format("2006-01-02 15:04:05")
		//时间
		row.Time = times
		row.TxID = txid
		log.Println("txid=============================>", row.TxID)
		// 区块高度
		row.Height = int64(block.Header.Number)
		//上一个区块hash
		row.PreHash = hex.EncodeToString(block.Header.PreviousHash)

		//区块hash
		if lastBlock.BCI.Height == block.Header.Number+indexOne {
			// 如果相等，就说明查询的是最新区块，那么本区块hash
			//区块hash
			log.Println("最新区块hash赋值", row.Hash)
			row.Hash = hex.EncodeToString(lastBlock.BCI.CurrentBlockHash)
		} else {
			// 查询本区块hash
			log.Println("区块不是最新")
			nextBlock, err := cli.QueryBlock(block.Header.Number+indexOne, ledger.WithTargetEndpoints(peer...))
			if err != nil {
				log.Println("查询出错")
			}
			row.Hash = hex.EncodeToString(nextBlock.Header.PreviousHash)
			row.TxID = txid
			log.Println("区块高度安装完毕")
		}
		if err != nil {
			continue
		}
		transaction = append(transaction, row)
	}

	// 返回
	return transaction, nil
}

func (setup *Application) GetnewBlockTxIdbf(txid string) ([]*blockdata.ChainTransactionup, error) {

	// 声明  peer 对等节点
	var peer []string
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)

	// 交易
	var transaction = make([]*blockdata.ChainTransactionup, 0)

	//  sdk 本身  调用 sdk   通道 || 组织名称 || user
	ctx := setup.SDK.SDK.ChannelContext(setup.SDK.ChannelID, fabsdk.WithOrg(setup.SDK.OrgName), fabsdk.WithUser(setup.SDK.UserName))

	// cli 客户端
	cli, err := ledger.New(ctx)
	if err != nil {
		return nil, err
	}

	// 查询 txid 所在区块
	block, err := cli.QueryBlockByTxID(fab.TransactionID(txid), ledger.WithTargetEndpoints(peer...))
	if err != nil {
		//
		log.Println("查询交易ID出错", err)
		return nil, err
	}

	// 查询当前最新区块信息
	lastBlock, err1 := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))

	if err1 != nil {
		log.Println("查询最新区块出错", err)
		return nil, err
	}

	//这里要做下处理 。就是根据交易ID查询的时候，要判断是否是最新的区块，如果是最新的区块，区块Hash就是要变化
	//此处应该遍历block.Data.Data
	indexOne, _ := strconv.ParseUint(strconv.Itoa(1), 10, 64)
	for _, data := range block.Data.Data {
		env, err := futils.GetEnvelopeFromBlock(data)
		row, err := blockdata.EnvelopeToTrasactionup((*common.Envelope)(env))
		times := time.Unix(row.Timestamp, 0).Format("2006-01-02 15:04:05")
		//时间
		row.Time = times
		// 区块高度
		row.Height = int64(block.Header.Number)
		//上一个区块hash
		row.PreHash = hex.EncodeToString(block.Header.PreviousHash)

		//区块hash
		if lastBlock.BCI.Height == block.Header.Number+indexOne {
			// 如果相等，就说明查询的是最新区块，那么本区块hash
			//区块hash
			log.Println("最新区块hash赋值", row.Hash)
			row.Hash = hex.EncodeToString(lastBlock.BCI.CurrentBlockHash)
		} else {
			// 查询本区块hash
			log.Println("区块不是最新")
			nextBlock, err := cli.QueryBlock(block.Header.Number+indexOne, ledger.WithTargetEndpoints(peer...))
			if err != nil {
				log.Println("查询出错")
			}
			row.Hash = hex.EncodeToString(nextBlock.Header.PreviousHash)
			log.Println("区块高度安装完毕")
		}
		if err != nil {
			continue
		}
		transaction = append(transaction, row)
	}

	// 返回
	return transaction, nil
}

/**
@  mod 查询区块链首页信息
@	2019年10月21日16:40:25
@	描述： setup
*/
func (setup *Application) GetBlockWord(channelName string, orgName string, userName string) (*[]BlockWordback, error) {
	//	01.
	log.Println("开始查询默认通道最新区块信息")
	//
	list := []BlockWordback{}
	var peer []string
	// 02. 设置背书节点
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)

	// 生成上下文实例   参数 ： 通道名称   组织  用户
	ctx := setup.SDK.SDK.ChannelContext(Channel_assetpublish, fabsdk.WithOrg(setup.SDK.OrgName), fabsdk.WithUser(setup.SDK.UserName))

	// 生成 客户端实例
	cli, err := ledger.New(ctx)
	if err != nil {
		return nil, err
	}

	//查询此通道上各种有用的区块链信息
	resp, err := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))
	if err != nil {
		return nil, err
		log.Println("err:", err)
	}
	log.Println("最新的区块高度是=>", resp.BCI.Height)

	blockNumun := resp.BCI.Height

	// 2. 对应区块高度信息
	for i := 1; i < 2; i++ {
		row := BlockWordback{}
		num, err := strconv.ParseUint(strconv.Itoa(i), 10, 64)
		if err != nil {
			return nil, err
		}
		// 查询区块信息
		block, err := cli.QueryBlock(blockNumun-num, ledger.WithTargetEndpoints(peer...))
		log.Println("要查询的区块编号是：", blockNumun-num)
		//
		if err != nil {
			return nil, err
		}

		if (resp.BCI.Height - num) > 0 {
			log.Println("resp.BCI.Height - num")
		} else {
			log.Println("continue")
			continue
		}
		//	循环查询
		for _, data := range block.Data.Data {
			env, err := futils.GetEnvelopeFromBlock(data)
			chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
			//区块高度
			row.Height = uint64(block.Header.Number)

			//区块时间
			times := time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")
			row.BlockTime = times
			//区块hash
			row.BlockHash = hex.EncodeToString(block.Header.DataHash)
			//上一个区块hash
			row.PreHash = hex.EncodeToString(block.Header.PreviousHash)
			//区块交易数据量
			row.TransactionNum = strconv.Itoa(len(block.Data.Data))
			//通道ID
			row.ChannelID = Channel_assetpublish
			//
			if err != nil {
				continue
			}
			list = append(list, row)
		}
	}
	return &list, nil
}

/**
@   mod 查询区块链配置信息
@	2019年10月21日17:59:47
@	描述： setup
*/
func (setup *Application) GetBlockMessage(channelName string, orgName string, userName string) (*YunConfig, error) {
	//
	var peer []string
	config := YunConfig{}
	// 02. 设置背书节点
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)
	//
	ctx := setup.SDK.SDK.ChannelContext(
		setup.SDK.ChannelID,
		fabsdk.WithOrg(setup.SDK.OrgName),
		fabsdk.WithUser(setup.SDK.UserName),
	)

	cli, err := ledger.New(ctx)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	resp, err := cli.QueryInfo(
		ledger.WithTargetEndpoints(peer...))
	if err != nil {
		return nil, err
		log.Println("err:", err)
	}
	// 1. gaodu
	s_height := resp.BCI.Height
	y_height := strconv.FormatUint(s_height-1, 10)
	//2. zuzhi
	s_orglist, err := App.GetAllAffiliations("default")
	orgNum := 0
	//
	if err == nil {
		if len(s_orglist.AffiliationInfo.Affiliations) > 0 {
			//
			log.Println("一级组织数量：", s_orglist.AffiliationInfo.Name)
			//
			for k, _ := range s_orglist.AffiliationInfo.Affiliations {
				//	如果下级有，查询
				if len(s_orglist.AffiliationInfo.Affiliations[k].Affiliations) >= 0 {
					log.Println("二级组织数量：", s_orglist.AffiliationInfo.Affiliations[k].Name)
					for w, _ := range s_orglist.AffiliationInfo.Affiliations[k].Affiliations {
						//
						if len(s_orglist.AffiliationInfo.Affiliations[k].Affiliations[w].Affiliations) >= 0 {
							log.Println("三级组织数量：", s_orglist.AffiliationInfo.Affiliations[k].Affiliations[w].Name)
							orgNum += 1
						}
						for h, _ := range s_orglist.AffiliationInfo.Affiliations[k].Affiliations[w].Affiliations {
							//
							if len(s_orglist.AffiliationInfo.Affiliations[k].Affiliations[w].Affiliations[h].Affiliations) >= 0 {

								log.Println("四级组织数量：", s_orglist.AffiliationInfo.Affiliations[k].Affiliations[w].Affiliations[h].Name)
								orgNum += 1
							}
						}
					}
				}
			}
		}
	}

	//查询 所有peer 节点
	PeerNum := getConfigureNodesTheNetwork()

	// 查询 当前节点加入通道
	channelslist, err := App.SDK.resmgmt.QueryChannels(
		resmgmt.WithTargetEndpoints(Peerfactoring),
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
	)
	if err != nil {
		log.Println("查询当前节点加入通道:", err)
	}
	//
	log.Println("通道中的所有节点:", channelslist)
	//var Num int64
	x_height, err := strconv.Atoi(y_height)
	//
	if err != nil {
		return nil, err

	}
	var rowNum int
	for i := 1; i < x_height; i++ {
		num, err := strconv.ParseUint(strconv.Itoa(i), 10, 64)
		if err != nil {
			log.Println("err:", err)
			return nil, nil
		}
		block, err := cli.QueryBlock(s_height-num, ledger.WithTargetEndpoints(peer...))
		if err != nil {
			log.Println("err:", err)
			return nil, nil
		}
		rowNum += len(block.Data.Data)
	}
	block, err := cli.QueryBlock(s_height-1, ledger.WithTargetEndpoints(peer...))
	if err != nil {
		return nil, err
	}
	for _, data := range block.Data.Data {
		env, err := futils.GetEnvelopeFromBlock(data)
		chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		chainTransaction.Height = int64(block.Header.Number) //区块高度
		times := time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")
		config.ConfigTime = times
		if err != nil {
			continue
		}
	}

	// 区块高度
	config.ConfigHeight = y_height
	//组织数量
	OrgNum1 := getConfigureOrgTheNetwork()
	config.ConfigOrgNum = strconv.Itoa(OrgNum1)
	//配置节点
	config.ConfigNodeNum = strconv.Itoa(PeerNum)
	//事务数量
	config.ConfigTransactionNum = strconv.Itoa(rowNum)
	//通道数量

	//链码数量
	config.ConfigChaincodeNum = Fabric_ChaincodeNum
	if channelslist != nil {
		config.ConfigChannelNum = strconv.Itoa(len(channelslist.Channels))
	} else {
		log.Println("channelslist is null")
	}

	return &config, nil

}

/**
@   mod 根据TxID 查询  上链数据信息
@	2019年10月21日17:59:47
@	描述： setup , ||  通道名称  客户端实例  , peer 节点 ， 组织名称 用户名称 || 数据类型
*/
func (setup *Application) GetBlockTxId(channelName string, cliName string, peerName string, orgName string, userName string, txid string) (config string, data []string, err error) {

	var resultRow []string
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	//var sr interface{}
	var peer []string
	// 02. 设置背书节点
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)

	// ctx
	ctx := setup.SDK.SDK.ChannelContext(
		setup.SDK.ChannelID,
		fabsdk.WithOrg(setup.SDK.OrgName),
		fabsdk.WithUser(setup.SDK.UserName),
	)

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		return "", nil, err
	}
	block, err := cli.QueryBlockByTxID(fab.TransactionID(txid), ledger.WithTargetEndpoints(peer...))

	if err != nil {
		return "", nil, err
	}

	var transaction = make([]*blockdata.ChainTransactionQuery, 0)
	//此处应该遍历block.Data.Data？
	for _, data := range block.Data.Data {

		env, err := futils.GetEnvelopeFromBlock(data)
		row, _ := blockdata.EnvelopeToTrasactionup((*common.Envelope)(env))
		chainTransaction, err := blockdata.EnvelopeToTrasactionQuery((*common.Envelope)(env))
		chainTransaction.ChainTransactionConfig.Height = int64(block.Header.Number)              //区块高度
		chainTransaction.ChainTransactionConfig.Hash = hex.EncodeToString(block.Header.DataHash) //hash
		times := time.Unix(chainTransaction.ChainTransactionConfig.Timestamp, 0).Format("2006-01-02 15:04:05")

		loc, _ := time.LoadLocation("Local") //获取时区
		tmp, _ := time.ParseInLocation(times, timeLayout, loc)
		sd := tmp.UnixNano() / 1e6
		stringTime := strconv.FormatInt(sd, 10)
		log.Println("", stringTime)
		chainTransaction.ChainTransactionConfig.Time = times
		chainTransaction.ChainTransactionConfig.Timestamp = chainTransaction.ChainTransactionConfig.Timestamp * 1000
		chainTransaction.ChainTransactionConfig.Timestampes = chainTransaction.ChainTransactionConfig.Timestamp * 1000
		chainTransaction.ChainTransactionConfig.Chaincode = row.Chaincode
		chainTransaction.ChainTransactionConfig.ChannelId = row.ChannelId
		chainTransaction.ChainTransactionConfig.Method = row.Method

		//
		log.Println("row================================>", row.Method, row.ChannelId)
		chainTransaction.ChainTransactionConfig.PreHash = hex.EncodeToString(block.Header.PreviousHash)
		//
		for _, args := range chainTransaction.ChainTransactionConfig.TxArgs {
			chainTransaction.Data = append(chainTransaction.Data, string(args))
			// TODO 根据反射去获取字段类型
			resultRow = append(resultRow, string(args))
		}
		chainTransaction.ChainTransactionConfig.TxID = txid

		transaction = append(transaction, chainTransaction)
		if err != nil {
			continue
		}
	}

	if err != nil {
	}
	ByteRes, err := json.Marshal(transaction[0].ChainTransactionConfig)

	return string(ByteRes), resultRow, nil

}

func (setup *Application) GetBlockTxIdbf(channelName string, cliName string, peerName string, orgName string, userName string, txid string) (config string, data []string, err error) {

	var resultRow []string

	//var sr interface{}
	var peer []string
	// 02. 设置背书节点
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)

	// ctx
	ctx := setup.SDK.SDK.ChannelContext(
		setup.SDK.ChannelID,
		fabsdk.WithOrg(setup.SDK.OrgName),
		fabsdk.WithUser(setup.SDK.UserName),
	)

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		log.Println("614")
		return "", nil, err
	}
	block, err := cli.QueryBlockByTxID(fab.TransactionID(txid), ledger.WithTargetEndpoints(peer...))

	if err != nil {
		log.Println("620")
		return "", nil, err
	}

	var transaction = make([]*blockdata.ChainTransactionQuery, 0)
	//此处应该遍历block.Data.Data？
	for _, data := range block.Data.Data {

		env, err := futils.GetEnvelopeFromBlock(data)
		chainTransaction, err := blockdata.EnvelopeToTrasactionQuery((*common.Envelope)(env))
		chainTransaction.ChainTransactionConfig.Height = int64(block.Header.Number)              //区块高度
		chainTransaction.ChainTransactionConfig.Hash = hex.EncodeToString(block.Header.DataHash) //hash
		times := time.Unix(chainTransaction.ChainTransactionConfig.Timestamp, 0).Format("2006-01-02 15:04:05")
		chainTransaction.ChainTransactionConfig.Time = times
		chainTransaction.ChainTransactionConfig.PreHash = hex.EncodeToString(block.Header.PreviousHash)
		//
		for _, args := range chainTransaction.ChainTransactionConfig.TxArgs {
			chainTransaction.Data = append(chainTransaction.Data, string(args))
			// TODO 根据反射去获取字段类型
			resultRow = append(resultRow, string(args))
		}
		chainTransaction.ChainTransactionConfig.TxID = txid

		transaction = append(transaction, chainTransaction)
		if err != nil {
			continue
		}
	}

	if err != nil {
	}
	ByteRes, err := json.Marshal(transaction[0].ChainTransactionConfig)

	return string(ByteRes), resultRow, nil

}

//	根据区块高度查询数据
func (setup *Application) GetBlockHeight(ChannleID string, OrgName string, UserID string, num string) (*blockdata.Block, error) {
	//

	blo := &blockdata.Block{}
	peer := Peer()
	ctx := setup.SDK.SDK.ChannelContext(ChannleID, fabsdk.WithOrg(setup.SDK.OrgName), fabsdk.WithUser(setup.SDK.UserName))

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		return nil, err
	}

	//	查询指定区块高度
	u, err := strconv.ParseUint(num, 10, 64)

	//	默认+1 区块
	Index, _ := strconv.ParseUint("1", 10, 64)

	//	根据指定区块高度去查询
	block, err := cli.QueryBlock(u, ledger.WithTargetEndpoints(peer...))
	if err != nil {
		log.Println("根据指定区块高度去查询出错:", err)
		return &blockdata.Block{}, errors.New("根据指定区块高度去查询出错")
	}
	NextBlock, err := cli.QueryBlock(u+Index, ledger.WithTargetEndpoints(peer...))
	if err != nil {
		log.Println("根据指定区块高度+1去查询出错:", err)
		return &blockdata.Block{}, errors.New("根据指定区块高度+1去查询出错")
	}
	//	当前区块Hash 要去另一个区块查询
	blo.BlockHash = hex.EncodeToString(NextBlock.Header.PreviousHash)
	blo.PreHash = hex.EncodeToString(block.Header.PreviousHash)
	//	当前区块交易条数
	blo.TransactionNumber = len(block.Data.Data)
	//此处应该遍历block.Data.Data
	var transaction = make([]*blockdata.ChainTransaction, 0)

	//	 交易详情
	for _, data := range block.Data.Data {
		env, err := futils.GetEnvelopeFromBlock(data)
		chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		chainTransactionup, _ := blockdata.EnvelopeToTrasactionup((*common.Envelope)(env))
		chainTransaction.Height = int64(block.Header.Number) //区块高度
		blo.Timestamp = chainTransaction.Timestamp
		times := time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")
		blo.Height = chainTransaction.Height
		blo.Time = times
		chainTransaction.TxID = chainTransactionup.TxID
		chainTransaction.Time = times
		chainTransaction.Hash = hex.EncodeToString(NextBlock.Header.PreviousHash)
		chainTransaction.PreHash = hex.EncodeToString(block.Header.PreviousHash)

		// 临时展示 txid
		log.Println("临时展示", chainTransactionup.TxID)

		if chainTransaction.TxID != "" {
			transaction = append(transaction, chainTransaction)
		}
		if err != nil {
			continue
		}
	}

	blo.Transaction = transaction

	return blo, nil
}

//	根据区块hash查询数据
func (setup *Application) GetBlockHash(channelId string, hash string) (*blockdata.Block, error) {
	//
	peer := Peer()
	Pg := &blockdata.Block{}

	// ctx
	ctx := setup.SDK.SDK.ChannelContext(channelId, fabsdk.WithOrg(setup.SDK.OrgName), fabsdk.WithUser(setup.SDK.UserName))

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		return nil, err
	}

	//	查询
	query, _ := hex.DecodeString(hash)
	block, err := cli.QueryBlockByHash(query, ledger.WithTargetEndpoints(peer...))
	if err != nil {
		return nil, err
	}
	//	当前区块Hash 要去另一个区块查询
	//	当前区块交易条数
	Index, _ := strconv.ParseUint("1", 10, 64)

	//	查询区块hash,下一个区块Hash 查询
	NextBlock, err := cli.QueryBlock(block.Header.Number+Index, ledger.WithTargetEndpoints(peer...))

	if err != nil {
		return nil, err
	}

	//	区块Hash
	Pg.BlockHash = hex.EncodeToString(NextBlock.Header.PreviousHash)
	Pg.PreHash = hex.EncodeToString(block.Header.PreviousHash)

	//	交易条数 len
	Pg.TransactionNumber = len(block.Data.Data)
	//此处应该遍历block.Data.Data
	var transaction = make([]*blockdata.ChainTransaction, 0)

	//	 交易详情
	for _, data := range block.Data.Data {
		env, err := futils.GetEnvelopeFromBlock(data)
		chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		chainTransactionup, _ := blockdata.EnvelopeToTrasactionup((*common.Envelope)(env))
		chainTransaction.Height = int64(block.Header.Number) //区块高度
		Pg.Timestamp = chainTransaction.Timestamp
		times := time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")

		Pg.Time = times
		Pg.Height = chainTransaction.Height
		chainTransaction.PreHash = hex.EncodeToString(block.Header.PreviousHash)
		chainTransaction.Hash = hex.EncodeToString(NextBlock.Header.PreviousHash)
		chainTransaction.PreHash = hex.EncodeToString(block.Header.PreviousHash)
		chainTransaction.TxID = chainTransactionup.TxID
		chainTransaction.Time = times
		if chainTransaction.TxID != "" {
			transaction = append(transaction, chainTransaction)
		}
		if err != nil {
			continue
		}
	}

	Pg.Transaction = transaction

	return Pg, nil
}

func (setup *Application) GetBlockHashbf(channelId string, hash string) (*blockdata.Block, error) {
	//
	peer := Peer()
	Pg := &blockdata.Block{}

	// ctx
	ctx := setup.SDK.SDK.ChannelContext(channelId, fabsdk.WithOrg(setup.SDK.OrgName), fabsdk.WithUser(setup.SDK.UserName))

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		return nil, err
	}

	//	查询
	query, _ := hex.DecodeString(hash)
	block, err := cli.QueryBlockByHash(query, ledger.WithTargetEndpoints(peer...))
	if err != nil {
		return nil, err
	}
	//	当前区块Hash 要去另一个区块查询
	//	当前区块交易条数
	Index, _ := strconv.ParseUint("1", 10, 64)

	//	查询区块hash,下一个区块Hash 查询
	NextBlock, err := cli.QueryBlock(block.Header.Number+Index, ledger.WithTargetEndpoints(peer...))

	if err != nil {
		return nil, err
	}

	//	区块Hash
	Pg.BlockHash = hex.EncodeToString(NextBlock.Header.PreviousHash)
	Pg.PreHash = hex.EncodeToString(block.Header.PreviousHash)

	//	交易条数 len
	Pg.TransactionNumber = len(block.Data.Data)
	//此处应该遍历block.Data.Data
	var transaction = make([]*blockdata.ChainTransaction, 0)

	//	 交易详情
	for _, data := range block.Data.Data {
		env, err := futils.GetEnvelopeFromBlock(data)
		chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		chainTransaction.Height = int64(block.Header.Number) //区块高度
		Pg.Timestamp = chainTransaction.Timestamp
		Pg.Height = chainTransaction.Height
		chainTransaction.TxID = chainTransaction.TxID
		if chainTransaction.TxID != "" {
			transaction = append(transaction, chainTransaction)
		}
		if err != nil {
			continue
		}
	}

	Pg.Transaction = transaction

	return Pg, nil
}

/**
@	函数： 区块链浏览器 - 查询通道动态信息
@	时间： 2019年10月21日19:57:15
@	描述:	通道ID  通道ID, cli ， 组织名称 ，用户名称
*/
func (setup *Application) GetBlockDynamic(channelName string, orgName string, userName string) ([]*BlockWordback, error) {
	//
	log.Println("开始根据区块编号查询最新5个区块信息")

	list := []*BlockWordback{}

	// 02. 设置背书节点
	var peer []string
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)

	//ctx
	ctx := setup.SDK.SDK.ChannelContext(
		channelName,
		fabsdk.WithOrg(setup.SDK.OrgName),
		fabsdk.WithUser(setup.SDK.UserName),
	)

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	// QueryInfo
	resp, err := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))
	if err != nil {
		return nil, err
		log.Println("err:", err)
	}

	//resp.BCI.CurrentBlockHash
	// 1. 高度
	blockNum := resp.BCI.Height
	// 区块高度
	log.Println("区块高度是：", blockNum)
	// 常量  1
	indexOne, _ := strconv.ParseUint(strconv.Itoa(1), 10, 64)
	//blockNum := indexOne + indexOne + indexOne

	arr := []uint64{}

	first := true
	// 2. 对应区块高度信息
	for i := 1; i < 20; i++ {
		row := BlockWordback{}
		num, err := strconv.ParseUint(strconv.Itoa(i), 10, 64)
		if err != nil {
			log.Println("err:", err)
			return nil, nil
		}

		if blockNum < num {
			break
			//
			log.Println("跳出")
		} else {
			//
			log.Println("查询区块信息：", int64(blockNum-num))
			block, err := cli.QueryBlock(blockNum-num, ledger.WithTargetEndpoints(peer...))
			//
			if err != nil {
				log.Println("查询区块信息 err:", err)
				break
			}
			//
			log.Println("查询的区块高度是：", block.Header.Number)
			proBlock, err1 := cli.QueryBlock(blockNum-num+indexOne, ledger.WithTargetEndpoints(peer...))
			//
			if err1 != nil {
				log.Println("查询是最高区块，单刀操作")
			}

			//--------------------------------------------------------------------拼接数据
			// 区块高度
			row.Height = block.Header.Number
			//区块hash
			// 判断是否是最新的区块，那么他的hash值就是在最新的查找
			if first {
				row.BlockHash = hex.EncodeToString(resp.BCI.CurrentBlockHash)
				//单刀操作
				first = false
			} else {
				row.BlockHash = hex.EncodeToString(proBlock.Header.PreviousHash)
			}

			//上一个区块hash
			row.PreHash = hex.EncodeToString(block.Header.PreviousHash)
			//区块交易数据量
			row.TransactionNum = strconv.Itoa(len(block.Data.Data))
			//通道ID
			row.ChannelID = channelName

			for _, data := range block.Data.Data {
				env, err := futils.GetEnvelopeFromBlock(data)
				chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
				times := time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")
				//区块时间
				row.BlockTime = times
				chainTransaction.Height = int64(block.Header.Number) //区块高度
				//
				if chainTransaction.TxID != "" {
				}
				if err != nil {
					break
				}
				// tianjai
				log.Println("添加 list:", row.Height)
				//

				if getIsOne(arr, row.Height) {
					log.Println("一个区块存在2笔交易", row.Height)
				} else {
					log.Println("一个新的区块加入数组", row.Height)
					arr = append(arr, row.Height)
					list = append(list, &row)
				}
			}
		}
	}
	return list, nil
}

/**
@	函数： 区块链浏览器 - 一个区块的交易
@	时间： 2019年10月21日20:26:14
@	描述:	区块高度  || 通道 cli 组织名称 用户名称
*/
func (setup *Application) GetBlockTransaction(channelName string, cliName string, orgName string, userName string, blockHeight string) ([]*BlockTransactionBack, error) {
	//
	var peer []string
	row := &BlockTransactionBack{}
	list := []*BlockTransactionBack{}
	// 02. 设置背书节点
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)
	//
	indexNum, _ := strconv.ParseUint("1", 10, 64)
	channelList := strings.Split(Fabric_ChannelList, ",")

	// ctx
	ctx := setup.SDK.SDK.ChannelContext(
		channelList[0],
		fabsdk.WithOrg(setup.SDK.OrgName),
		fabsdk.WithUser(setup.SDK.UserName),
	)

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	num, err := strconv.ParseUint(blockHeight, 10, 64)

	// 2. 对应区块高度信息
	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}

	block, err := cli.QueryBlock(num, ledger.WithTargetEndpoints(peer...))

	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}
	//
	log.Println("查询的区块高度是：", block.Header.Number)

	//最新区块
	NewBlock, err := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))
	log.Println("最新的区块查询的区块高度是：", NewBlock.BCI.Height)
	//
	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}
	indexOne, _ := strconv.ParseUint("1", 10, 64)
	for _, data := range block.Data.Data {
		env, err := futils.GetEnvelopeFromBlock(data)
		//env
		chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		row.Height = int64(block.Header.Number) //区块高度
		times := time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")
		row.BlockTime = times

		//
		if (NewBlock.BCI.Height - indexOne) == block.Header.Number {
			//是最高区块
			//
			log.Println("查询是最高区块,特殊处理")
			row.BlockHash = hex.EncodeToString(NewBlock.BCI.CurrentBlockHash)

		} else {
			log.Println("查询高一级区块")
			lastBlock, err := cli.QueryBlock(num+indexNum, ledger.WithTargetEndpoints(peer...))
			//
			if err != nil {
				log.Println("err:", err)
				return nil, nil
			}
			row.BlockHash = hex.EncodeToString(lastBlock.Header.PreviousHash)
		}

		row.PreHash = hex.EncodeToString(block.Header.PreviousHash)
		row.TransactionID = chainTransaction.TxID
		//

		if chainTransaction.TxID != "" {
			list = append(list, row)
		}
		if err != nil {
			continue
		}
	}
	return list, nil
}

/**
@	函数： 联盟组织	- 添加组织
@	时间： 2019年10月22日10:11:52
@	描述:  对于联盟组织进行注册  ID, USERID  ||  ca  force ca name  ca+name
*/
func (setup *Application) AddAffiliation(caName string, force bool, caPath string, affilName string) (*msp.AffiliationResponse, error) {

	// TODO	获取 CA 的名称
	log.Println("联盟组织	- 添加组织")
	// 拼接请求体
	requset := &msp.AffiliationRequest{
		Name:   caPath + affilName,
		Force:  true,
		CAName: caName,
	}
	//  添加联盟
	affation, err := setup.SDK.MspClient.AddAffiliation(requset)

	if err != nil {
		log.Println("联盟组织	", err)
		return &msp.AffiliationResponse{}, err
	}
	log.Println("联盟组织")
	return affation, nil
}

/**
@	函数： 联盟组织	- 添加用户 || msp
@	时间： 2019年10月22日10:33:58
@	描述:  对于联盟组织进行注册用户  ID, USERID  ||
*/
func (setup *Application) MspAffCreateUser(mspAffClient *msp.Client, affiliationId string, userName string, RegUserName string, RegOrgName string, RegOrgID string, Path string, Pwd string) (string, error) {
	//  msp client
	const psw string = "123456"
	enrollmentSecret, err := mspAffClient.Register(
		&msp.RegistrationRequest{
			Name:        userName,
			Secret:      psw,
			Affiliation: Path + affiliationId,
			Attributes: []msp.Attribute{
				{
					Name:  RegOrgName + RegUserName,
					Value: userName,
					ECert: true,
				}},
		})
	if err != nil {
		fmt.Printf("Register return error %s\n", err)
		return "", errors.WithMessage(err, "register uesr is error")
	}
	err = setup.SDK.MspClient.Enroll(userName, msp.WithSecret(enrollmentSecret))

	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return "", errors.WithMessage(err, "enroll is err")
	}

	_, err3 := setup.SDK.MspClient.GetIdentity(userName)
	if err3 == nil {
		log.Println("验证通过")
	}

	return string(enrollmentSecret), nil
}

//
func (setup *Application) MspAffCreateUserByWei(mspAffClient *msp.Client, affiliationId string, userName string, RegUserName string, RegOrgName string, RegOrgID string) (string, error) {
	//  msp client
	const psw string = "123456"
	enrollmentSecret, err := mspAffClient.Register(
		&msp.RegistrationRequest{
			Name:        userName,
			Secret:      psw,
			Affiliation: Ca_Org1_path + affiliationId,
			Attributes: []msp.Attribute{
				{
					Name:  RegOrgName + RegUserName,
					Value: userName,
					ECert: true,
				}},
		})
	if err != nil {
		fmt.Printf("Register return error %s\n", err)
		return "", errors.WithMessage(err, "register uesr is error")
	}
	err = setup.SDK.MspClient.Enroll(userName, msp.WithSecret(enrollmentSecret))

	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return "", errors.WithMessage(err, "enroll is err")
	}

	_, err3 := setup.SDK.MspClient.GetIdentity(userName)
	if err3 == nil {
		log.Println("验证通过")
	}

	return string(enrollmentSecret), nil
}

/**
@	函数： 联盟组织	-  验证用户信息
@	时间： 2019年10月22日10:42:09
@	描述:  对于联盟组织进行生成 msp cli
*/
func (setup *Application) CreateNewMspClient(affiliationName string, num string) (*msp.Client, error) {

	// 01. 创建资源管理客户端上下文
	clientCTX := setup.SDK.SDK.Context(
		fabsdk.WithOrg(num + affiliationName),
	)

	// cli
	c, err := msp.New(clientCTX)
	//c.Register()
	if err != nil {
		return nil, errors.WithMessage(err, "msp.client is error")
	}
	return c, nil
}

/**
@ 验证user用户信息
@ 2019年10月18日11:50:07
@ No:6001
*/
func (setup *Application) VifUserFull(userHash string) (string, error) {

	//  user
	result, err := setup.SDK.MspClient.GetIdentity(userHash)

	if err != nil {
		fmt.Printf("failed to enroll user: %s\n", err)
		return "", errors.WithMessage(err, "enroll is err")
	} else {
		log.Println("验证成功")
	}

	return string(result.ID), nil
}

// 查询 组织信息
func (setup *Application) GetAllAffiliations(Ca string) (*msp.AffiliationResponse, error) {

	affiliations, err := setup.SDK.MspClient.GetAllAffiliations(
		msp.WithCA(CAfactoring),
	)

	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	return affiliations, nil
}

// 获取 网络中所有peer节点
func ConfigureNodesTheNetwork() (peerNum int) {
	arr := []string{
		"peer0.org1.dinglian.com",
		"peer1.org1.dinglian.com",
		"peer0.org2.dinglian.com",
		"peer1.org2.dinglian.com",
	}
	//
	return len(arr)
}

// 获取 网络中所有Org节点
func ConfigureOrgTheNetwork() (OrgNum int) {
	arr := []string{
		"org1.dinglian.com",
		"org2.dinglian.com",
	}
	//
	return len(arr)
}

/**
get peer
*/

// 调用验证
func getIsOne(arr []uint64, num uint64) bool {
	//
	return isOne(arr, num)

}

//
func isOne(arr []uint64, num uint64) bool {
	//
	state := false
	//
	for k, _ := range arr {
		//
		if arr[k] == num {
			state = true
			break
		}
		//
	}
	//
	return state
}

//	拼接 KYC
func subKyc() {

}

//	拼接 BYC
func subByc() {
	//

}

//  获取对应的组织
func getCorrespondingOrganization(orgIndex string) (caName string) {
	var ca string
	switch orgIndex {
	case Ca01:
		ca = CaMsp01
	case Ca02:
		ca = CaMsp02
	case Ca03:
		ca = CaMsp03
	case Ca04:
		ca = CaMsp04
	case Ca05:
		ca = CaMsp05
	case Ca06:
		ca = CaMsp06
	case Ca07:
		ca = CaMsp07
	case Ca08:
		ca = CaMsp08
	case Ca09:
		ca = CaMsp09
	}

	//
	return ca
}

//  获取对应的组织 Path
func getCorrespondingOrganizationPath(orgIndex string) (caName string) {
	var ca string
	switch orgIndex {
	case Ca01:
		ca = CaMspPath01
	case Ca02:
		ca = CaMspPath02
	case Ca03:
		ca = CaMspPath03
	case Ca04:
		ca = CaMspPath04
	case Ca05:
		ca = CaMspPath05
	case Ca06:
		ca = CaMspPath06
	case Ca07:
		ca = CaMspPath07
	case Ca08:
		ca = CaMspPath08
	case Ca09:
		ca = CaMspPath09
	}
	//
	return ca
}

//	大屏@ 根据区块查询信息 @ 8 BlockData
func (setup *Application) SqlQueryBlock(channelName string, cliName string, orgName string, userName string, blockHeight string) ([]*BlockTransactionBack, error) {
	//
	var peer []string
	row := &BlockTransactionBack{}
	list := []*BlockTransactionBack{}
	AstInfo := blockdata.AstInfo{}

	// 02. 设置背书节点
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)
	//
	indexNum, _ := strconv.ParseUint("1", 10, 64)
	channelList := strings.Split(Fabric_ChannelList, ",")

	// ctx
	ctx := setup.SDK.SDK.ChannelContext(
		channelList[0],
		fabsdk.WithOrg(setup.SDK.OrgName),
		fabsdk.WithUser(setup.SDK.UserName),
	)

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	num, err := strconv.ParseUint(blockHeight, 10, 64)

	// 2. 对应区块高度信息
	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}
	//num = num - 1
	//先对比是否是最高区块
	//最新区块
	//NewBlock, err := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))
	//if NewBlock.BCI.Height == num{
	//	//如果是最高区块
	//
	//}else{
	//
	//
	//}
	//
	block, err := cli.QueryBlock(num, ledger.WithTargetEndpoints(peer...))

	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}
	//
	log.Println("查询的区块高度是：", block.Header.Number)

	//最新区块
	NewBlock, err := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))
	log.Println("最新的区块查询的区块高度是：", NewBlock.BCI.Height)
	//
	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}
	indexOne, _ := strconv.ParseUint("1", 10, 64)
	for _, data := range block.Data.Data {
		env, err := futils.GetEnvelopeFromBlock(data)
		chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		row.Height = int64(block.Header.Number) //区块高度
		times := time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")
		row.BlockTime = times

		//
		if (NewBlock.BCI.Height - indexOne) == block.Header.Number {
			//是最高区块
			//
			log.Println("查询是最高区块,特殊处理")
			row.BlockHash = hex.EncodeToString(NewBlock.BCI.CurrentBlockHash)

		} else {

			//
			log.Println("查询高一级区块")
			lastBlock, err := cli.QueryBlock(num+indexNum, ledger.WithTargetEndpoints(peer...))
			//
			if err != nil {
				log.Println("err:", err)
				return nil, nil
			}
			row.BlockHash = hex.EncodeToString(lastBlock.Header.PreviousHash)
		}

		row.PreHash = hex.EncodeToString(block.Header.PreviousHash)
		row.TransactionID = chainTransaction.TxID
		//
		for _, args := range chainTransaction.TxArgs {
			json.Unmarshal(args, &AstInfo)
			//chainTransaction. = Assets
			if len(AstInfo.AstAssetsInfo.AstAssetsList) == 0 {
			} else {
				row.AssetName = AstInfo.AstAssetsInfo.AstAssetsList[0].AstAssetsIntroduce
				row.AssetUUID = AstInfo.AstAssetsInfo.AstAssetsList[0].AstAssetsuuid
			}
		}
		if chainTransaction.TxID != "" {
			list = append(list, row)
		}
		if err != nil {
			continue
		}
	}
	return list, nil
}

//	SQL --  大屏 展示 TPS  api
/**
@	函数： 区块链浏览器 - 大屏展示
@	时间： 2019年10月21日19:57:15
@	描述:	通道ID  通道ID, cli ， 组织名称 ，用户名称
*/
func (setup *Application) GetBlockDynamicByBig(channelName string, orgName string, userName string) (interface{}, error) {
	//
	log.Println("开始根据区块编号查询最新大屏展示区块信息")
	listByBlock := []*DbBlock{}
	// 02. 设置背书节点
	var peer []string
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)

	//ctx
	ctx := setup.SDK.SDK.ChannelContext(
		channelName,
		fabsdk.WithOrg(setup.SDK.OrgName),
		fabsdk.WithUser(setup.SDK.UserName),
	)

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	// QueryInfo
	resp, err := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))
	if err != nil {
		return nil, err
		log.Println("err:", err)
	}

	//resp.BCI.CurrentBlockHash
	// 1. 高度
	blockNum := resp.BCI.Height
	// 区块高度
	log.Println("区块高度是：", blockNum)
	// 常量  1
	indexOne, _ := strconv.ParseUint(strconv.Itoa(1), 10, 64)
	//blockNum := indexOne + indexOne + indexOne

	first := true
	// 2. 对应区块高度信息
	for i := 1; i < 100; i++ {
		row := BlockWordback{}
		ByBlockOne := DbBlock{}
		ByBlockOne.No = i
		num, err := strconv.ParseUint(strconv.Itoa(i), 10, 64)
		if err != nil {
			log.Println("err:", err)
			return nil, nil
		}

		if blockNum < num {
			break
			//
			log.Println("跳出")
		} else {
			//
			block, err := cli.QueryBlock(blockNum-num, ledger.WithTargetEndpoints(peer...))
			//
			if err != nil {
				log.Println("查询区块信息 err:", err)
				break
			}
			proBlock, err1 := cli.QueryBlock(blockNum-num+indexOne, ledger.WithTargetEndpoints(peer...))
			//
			if err1 != nil {
				log.Println("查询是最高区块，单刀操作")
			}
			//--------------------------------------------------------------------拼接数据
			// 区块高度
			row.Height = block.Header.Number

			// TODO
			ByBlockOne.BlockHeight = strconv.FormatUint(block.Header.Number, 10)

			//区块hash
			// 判断是否是最新的区块，那么他的hash值就是在最新的查找
			if first {
				row.BlockHash = hex.EncodeToString(resp.BCI.CurrentBlockHash)

				//TODO
				ByBlockOne.BlockHash = hex.EncodeToString(resp.BCI.CurrentBlockHash)
				//单刀操作
				first = false
			} else {
				row.BlockHash = hex.EncodeToString(proBlock.Header.PreviousHash)
				//TODO
				ByBlockOne.BlockHash = hex.EncodeToString(proBlock.Header.PreviousHash)
			}

			//上一个区块hash
			row.PreHash = hex.EncodeToString(block.Header.PreviousHash)
			//TODO
			ByBlockOne.Type = hex.EncodeToString(block.Header.PreviousHash)
			//区块交易数据量
			row.TransactionNum = strconv.Itoa(len(block.Data.Data))
			//通道ID
			row.ChannelID = channelName
			//TODO
			ByBlockOne.BlockByChannel = channelName
			listByBlock = append(listByBlock, &ByBlockOne)
		}
	}
	return listByBlock, nil
}

//	区块链大屏 -  点击区块区块上的交易信息
func (setup *Application) GetBlockTransactionByBig(channelName string, cliName string, orgName string, userName string, blockHeight string) (interface{}, error) {

	log.Println("区块链大屏 -查询后五个")

	list := []DbDeal{}
	// 02. 设置背书节点
	peer := Peer()
	// cli
	cli, err := ledger.New(setup.SDK.SDK.ChannelContext(channelName, fabsdk.WithOrg(setup.SDK.OrgName), fabsdk.WithUser(setup.SDK.UserName)))
	if err != nil {
		log.Println("创建 cli err:", err)
		return nil, err
	}

	// 查询的区块高度
	num, err := strconv.ParseUint(blockHeight, 10, 64)
	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}

	block, err := cli.QueryBlock(num, ledger.WithTargetEndpoints(peer...))
	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}

	blockHash, err := blockHash(channelName, peer, blockHeight)
	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}

	// 遍历数据
	for _, v := range block.Data.Data {
		//
		row := DbDeal{}
		env, err := futils.GetEnvelopeFromBlock(v)
		if err != nil {
			log.Println("err:", err)
			return nil, nil
		}
		//env
		chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		if err != nil {
			log.Println("err:", err)
			return nil, nil
		}
		//	TODO
		row.BlockByTXID = chainTransaction.TxID
		row.BlockHash = blockHash
		row.BlockByChannel = channelName
		row.BlockHeight = blockHeight
		row.ChainCodeName = "AssetToChain"
		row.BlockByChainCode = "AssetToChain"
		row.UserTxt = "6a2149a7-a487-4fca-985a-de54aa3d99d4"
		row.TXIDTime = time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")

		//	拼接数组
		list = append(list, row)
	}

	return list, nil
}

//	获取背书节点
func peer() (arr []string) {
	arr = append(arr, Peerfactoring)
	arr = append(arr, Peerproduct)
	return arr
}

// 获取当前区块hash
func blockHash(cName string, anchor []string, height string) (blockHash string, err error) {
	peer := Peer()
	indexOne, _ := strconv.ParseUint("1", 10, 64)

	//	创建 cli 句柄
	cli, err := ledger.New(App.SDK.SDK.ChannelContext(cName, fabsdk.WithOrg(App.SDK.OrgName), fabsdk.WithUser(App.SDK.UserName)))
	if err != nil {
		log.Println("创建 cli err:", err)
		return "", nil
	}

	//	高度转换
	num, err := strconv.ParseUint(height, 10, 64)
	if err != nil {
		log.Println("err:", err)
		return "", nil
	}

	//	查询指定区块
	block, err := cli.QueryBlock(num, ledger.WithTargetEndpoints(peer...))
	if err != nil {
		log.Println("err:", err)
		return "", nil
	}

	//	查询最新区块
	Nblock, err := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))
	if err != nil {
		log.Println("err:", err)
		return "", nil
	}

	// 判断要查询的区块是否是最新区块，是返回最新区块hash,不是 ，返回查询的指定区块的hash

	if (Nblock.BCI.Height - indexOne) == block.Header.Number {
		log.Println("查询是最高区块特殊处理")
		return hex.EncodeToString(Nblock.BCI.CurrentBlockHash), nil
	} else {
		log.Println("查询指定区块直接返回")
		lastBlock, err := cli.QueryBlock(num+indexOne, ledger.WithTargetEndpoints(peer...))
		//
		if err != nil {
			log.Println("err:", err)
			return "", nil
		}
		return hex.EncodeToString(lastBlock.Header.PreviousHash), nil
	}
}

//	获取创世块时间
func (setup *Application) GetsTheCreationBlockTime(channelName string, orgName string, userName string) (timeLable string, err error) {

	// 背书节点
	peer := Peer()

	//ctx
	ctx := setup.SDK.SDK.ChannelContext(channelName, fabsdk.WithOrg(setup.SDK.OrgName), fabsdk.WithUser(setup.SDK.UserName))

	// cli
	cli, err := ledger.New(ctx)

	if err != nil {
		log.Println("err:", err)
		return "", err
	}

	//	默认查第一个块
	indexOne, _ := strconv.ParseUint(strconv.Itoa(1), 10, 64)

	//	查询
	block, err := cli.QueryBlock(indexOne, ledger.WithTargetEndpoints(peer...))

	// 遍历数据
	for _, v := range block.Data.Data {
		//
		env, err := futils.GetEnvelopeFromBlock(v)
		if err != nil {
			log.Println("err:", err)
			return "", err
		}
		//env
		chainTransaction, err := blockdata.EnvelopeToTrasaction((*common.Envelope)(env))
		if err != nil {
			log.Println("err:", err)
			return "", err
		}

		timeLable = time.Unix(chainTransaction.Timestamp, 0).Format("2006-01-02 15:04:05")
	}

	return timeLable, nil
}

// 微分格 数据上链  -- 根据主键查询区块信息 2019年12月24日14:10:33
func getBlockTxIdForBox(channelName string, cliName string, peerName string, orgName string, userName string, txid string) (config string, data []string, err error) {

	var resultRow []string
	peer := Peer()
	// ctx
	ctx := App.SDK.SDK.ChannelContext(
		App.SDK.ChannelID,
		fabsdk.WithOrg(App.SDK.OrgName),
		fabsdk.WithUser(App.SDK.UserName),
	)

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		return "", nil, err
	}
	block, err := cli.QueryBlockByTxID(fab.TransactionID(txid), ledger.WithTargetEndpoints(peer...))

	if err != nil {
		return "", nil, err
	}

	var transaction = make([]*blockdata.ChainTransactionQuery, 0)
	//此处应该遍历block.Data.Data？
	for _, data := range block.Data.Data {

		env, err := futils.GetEnvelopeFromBlock(data)
		chainTransaction, err := blockdata.EnvelopeToTrasactionQuery((*common.Envelope)(env))
		chainTransaction.ChainTransactionConfig.Height = int64(block.Header.Number)              //区块高度
		chainTransaction.ChainTransactionConfig.Hash = hex.EncodeToString(block.Header.DataHash) //hash
		times := time.Unix(chainTransaction.ChainTransactionConfig.Timestamp, 0).Format("2006-01-02 15:04:05")
		chainTransaction.ChainTransactionConfig.Time = times
		chainTransaction.ChainTransactionConfig.PreHash = hex.EncodeToString(block.Header.PreviousHash)
		//
		for _, args := range chainTransaction.ChainTransactionConfig.TxArgs {
			chainTransaction.Data = append(chainTransaction.Data, string(args))
			// TODO 根据反射去获取字段类型
			resultRow = append(resultRow, string(args))
		}
		chainTransaction.ChainTransactionConfig.TxID = txid

		transaction = append(transaction, chainTransaction)
		if err != nil {
			continue
		}
	}

	if err != nil {
		//
	}

	ByteRes, err := json.Marshal(transaction[0].ChainTransactionConfig)

	return string(ByteRes), resultRow, nil

}

// 查询 后五个

func (setup *Application) GetNetBlockDynamicByBig(channelName string, orgName string, userName string, blockNumber string) (interface{}, error) {
	//
	log.Println("开始查询后五个")
	listByBlock := []*DbBlock{}
	// 02. 设置背书节点
	var peer []string
	peer = append(peer, Peerfactoring)
	peer = append(peer, Peerproduct)

	//ctx
	ctx := setup.SDK.SDK.ChannelContext(
		channelName,
		fabsdk.WithOrg(setup.SDK.OrgName),
		fabsdk.WithUser(setup.SDK.UserName),
	)

	// cli
	cli, err := ledger.New(ctx)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	// QueryInfo
	resp, err := cli.QueryInfo(ledger.WithTargetEndpoints(peer...))
	if err != nil {
		return nil, err
		log.Println("err:", err)
	}
	blockNumberUint, err := strconv.ParseUint(blockNumber, 10, 64)
	//resp.BCI.CurrentBlockHash
	// 1. 高度
	blockNum := resp.BCI.Height
	// 区块高度
	log.Println("区块高度是：", blockNum)
	// 常量  1
	indexOne, _ := strconv.ParseUint(strconv.Itoa(1), 10, 64)

	first := true
	// 2. 对应区块高度信息
	for i := 1; i < 30; i++ {
		row := BlockWordback{}
		ByBlockOne := DbBlock{}
		ByBlockOne.No = i
		num, err := strconv.ParseUint(strconv.Itoa(i), 10, 64)
		if err != nil {
			log.Println("err:", err)
			return nil, nil
		}

		if blockNumberUint < num {
			break
			//
			log.Println("跳出")
		} else {
			//
			block, err := cli.QueryBlock(blockNumberUint-num, ledger.WithTargetEndpoints(peer...))
			//
			if err != nil {
				log.Println("查询区块信息 err:", err)
				break
			}
			proBlock, err1 := cli.QueryBlock(blockNumberUint-num+indexOne, ledger.WithTargetEndpoints(peer...))
			//
			if err1 != nil {
				log.Println("查询是最高区块，单刀操作")
			}
			//--------------------------------------------------------------------拼接数据
			// 区块高度
			row.Height = block.Header.Number

			// TODO
			ByBlockOne.BlockHeight = strconv.FormatUint(block.Header.Number, 10)

			//区块hash
			// 判断是否是最新的区块，那么他的hash值就是在最新的查找
			if first {
				row.BlockHash = hex.EncodeToString(resp.BCI.CurrentBlockHash)

				//TODO
				ByBlockOne.BlockHash = hex.EncodeToString(resp.BCI.CurrentBlockHash)
				//单刀操作
				first = false
			} else {
				row.BlockHash = hex.EncodeToString(proBlock.Header.PreviousHash)
				//TODO
				ByBlockOne.BlockHash = hex.EncodeToString(proBlock.Header.PreviousHash)
			}

			//上一个区块hash
			row.PreHash = hex.EncodeToString(block.Header.PreviousHash)
			//TODO
			ByBlockOne.Type = hex.EncodeToString(block.Header.PreviousHash)
			//区块交易数据量
			row.TransactionNum = strconv.Itoa(len(block.Data.Data))
			//通道ID
			row.ChannelID = channelName
			//TODO
			ByBlockOne.BlockByChannel = channelName
			listByBlock = append(listByBlock, &ByBlockOne)
		}
	}
	return listByBlock, nil
}
