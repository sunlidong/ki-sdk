package action

import (
	"ki-sdk/util/api"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"log"
)

/**
@ 数据上链
@ 2019年10月21日09:55:34
@
*/
func UploadAssetTest(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
	return uploadAssettest1(clientName, channelName, chaincodeName, peer, args)
}

// 数据上链
func uploadAssettest1(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {

	var peer111 []string
	// 02. 设置背书节点
	peer111 = append(peer111, Peerproduct)
	peer111 = append(peer111, Peerfactoring)

	request := channel.Request{
		ChaincodeID: ConfChainCodeIDTest,
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

//	根据主键查询关联数据
func QueryAssetTest(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {
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
		ChaincodeID: ConfChainCodeIDTest,
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

//	根据主键查询元数据
func ChainCodeQueryById(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {

	var peer111 []string
	// 02. 设置背书节点
	peer111 = append(peer111, Peerproduct)
	peer111 = append(peer111, Peerfactoring)

	request := channel.Request{
		ChaincodeID: ConfChainCodeIDTest,
		Fcn:         args[0],
		Args: [][]byte{
			[]byte(args[1]),
		}}

	// 请求参数
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

// 微分格
func UploadByBox(clientName string, channelName string, chaincodeName string, funcName string, peer string, args []string, uuid string) (result string, err error) {
	return uploadByBox(clientName, channelName, chaincodeName, funcName, peer, args, uuid)
}

// 微分格 -- 数据查询
func QueryloadByBox(clientName string, channelName string, chaincodeName string, funcName string, peer string, args []string, uuid string) (result string, err error) {
	return queryloadByBox(clientName, channelName, chaincodeName, funcName, peer, args, uuid)
}

// 数据上链
func uploadByBox(clientName string, channelName string, chaincodeName string, funcName string, peer string, args []string, uuid string) (result string, err error) {
	//
	//	01. client TODO
	//	02. channel TODO
	//	03. chaincode TODO
	//	04. peer
	var peerList []string
	// 02. 设置背书节点
	peerList = append(peerList, Peerproduct)
	peerList = append(peerList, Peerfactoring)

	Pg, err := api.ArgsSplicing(args)
	//
	if err != nil {
		log.Println("Args err:", err)
		return "", err
	}

	// 参数拼接
	request := channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         funcName,
		Args:        Pg,
	}

	response, err := App.SDK.client.Execute(
		request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peerList...),
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

//  微分格 -- 数据上链 -- 查询 2019年12月18日10:22:14
func queryLoadByBox(clientName string, channelName string, chaincodeName string, funcName string, peer string, args []string, uuid string) (result string, err error) {

	//	背书节点
	peerList := Peer()

	// 参数拼接
	Pg, err := api.ArgsSplicing(args)
	//
	if err != nil {
		log.Println("err:", err)
		return "", err
	}

	// 参数拼接
	request := channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         funcName,
		Args:        Pg,
	}

	// Query
	response, err := App.SDK.client.Query(
		request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peerList...),
	)

	// 判断错误
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	// 06. 返回调用结果
	return string(response.Payload), nil

}

//	微分格 根据主键查询元数据
func ChainCodeQueryByIdByBox(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {

	var peer111 []string
	// 02. 设置背书节点
	peer111 = append(peer111, Peerproduct)
	peer111 = append(peer111, Peerfactoring)

	request := channel.Request{
		ChaincodeID: chaincodeName,
		Fcn:         args[0],
		Args: [][]byte{
			[]byte(args[1]),
		}}

	// 请求参数
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

//

// 数据上链
func TestAssettest(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {

	peerlist := Peer()

	request := channel.Request{
		ChaincodeID: TestChain,
		Fcn:         args[0],
		Args: [][]byte{
			[]byte(args[1]),
			[]byte(args[2]),
			[]byte(args[3]),
		}}
	response, err := App.SDK.client.Execute(
		request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peerlist...),
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

// 数据上链
func TestSelectChaincode(clientName string, channelName string, chaincodeName string, peer string, args []string) (result string, err error) {

	peerlist := Peer()

	request := channel.Request{
		ChaincodeID: TestChain,
		Fcn:         args[0],
		Args: [][]byte{
			[]byte(args[1]),
		}}
	response, err := App.SDK.client.Query(
		request,
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(peerlist...),
	)
	// 05. 判断错误
	if err != nil {
		//
		log.Println("err:", err)

		return "", err
	}
	log.Println("复合链码查询结果：", string(response.Payload))
	// 06. 返回调用结果
	return string(response.Payload), nil

}
