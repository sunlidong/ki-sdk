package model

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/pkg/errors"
)

// Action
func (swp *SDK) Initialize() error {
	// 01. 读取配置文件config.yaml
	configProvider := config.FromFile(
		swp.ConfigFile,
	)
	// 02. 创建sdk对象
	sdk, err := fabsdk.New(configProvider)
	// 03. 验错
	if err != nil {
		return fmt.Errorf("SDK实例化失败:%v", err)
	}
	// 04. 参数转换至结构体对象 swp
	swp.SDK = sdk
	// 05. 返回
	return nil
}

/**
@ 初始化 msg
@
*/
func (swp *SDK) CreateresMgmtClient() error {
	// 01. 创建资源管理客户端上下文
	resourceManagerClientContext :=
		swp.SDK.Context(fabsdk.WithUser(swp.OrgAdmin),
			fabsdk.WithOrg(swp.OrgName))
	// 02. 创建资源管理客户端
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	// 03. 验错
	if err != nil {
		return fmt.Errorf("创建资源管理客户端失败:%v", err)
	}
	// 04. 参数转换至结构体对象 swp
	swp.Resmgmt = resMgmtClient
	// 05. 返回
	return nil
}

/**
@ 初始化 channel cli
@
*/
func (swp SDK) CreateChannelCli() error {

	// 01. 封装数据Channle cli
	clientContext := swp.SDK.ChannelContext(
		swp.ChannelID,
		fabsdk.WithUser(swp.UserName))

	// 02. 创建Channle cli
	channelCli, err := channel.New(clientContext)

	// 03. 验错
	if err != nil {
		return fmt.Errorf("创建通道管理客户端失败:%v", err)
	}

	// 04. 参数转换至结构体对象 swp
	swp.Client = channelCli
	// 05. 返回
	return nil
}

/**
@ 初始化 msp cli
@
*/
func (swp SDK) CreateMspClient() error {

	// 01. 创建资源管理客户端上下文
	clientCTX := swp.SDK.Context(
		fabsdk.WithUser(swp.OrgAdmin),
		fabsdk.WithOrg(swp.OrgName),
	)

	// 02. 创建资源实例
	c, err := msp.New(clientCTX)

	if err != nil {
		return fmt.Errorf("创建msp管理客户端失败:%v", err)
	}

	if c == nil {
		return fmt.Errorf("创建msp管理客户端为空:%v", c)
	}

	swp.MspClient = c

	return nil
}

//查询peer已经安装的chaincode
func (swp SDK) GetInstalledChaincode(targetPeer string) ([]string, error) {

	// 查询已经安装的CC
	ccInstalledRes, err := swp.Resmgmt.QueryInstalledChaincodes(resmgmt.WithTargetEndpoints(targetPeer))

	if err != nil {
		return nil, errors.Errorf("查询节点%s已安装的chaincode失败，错误为：%s", targetPeer, err)
	}

	list := make([]string, 0)

	if ccInstalledRes != nil {
		for _, cc := range ccInstalledRes.Chaincodes {
			str := cc.Name + cc.Version
			list = append(list, str)
		}
	}
	return list, nil
}

//查询peer已经实例化的chaincode
func (swp SDK) GetInstantiatedChaincode(channelID, targetPeer string) ([]string, error) {

	// 查询已经实例化的CC
	ccInstantiatedRes, err := swp.Resmgmt.QueryInstantiatedChaincodes(channelID, resmgmt.WithTargetEndpoints(targetPeer))
	if err != nil {
		return nil, errors.Errorf("查询目标节点%s在channel:%s已安装的chaincode失败，错误为：%s", targetPeer, channelID, err)
	}

	list := make([]string, 0)

	if ccInstantiatedRes != nil {
		for _, cc := range ccInstantiatedRes.Chaincodes {
			str := cc.Name + cc.Version
			list = append(list, str)
		}
	}
	return list, nil
}

// 向指定peer上安装chaincode
func (swp SDK) InstallChaincode(targetPeer string) error {

	// 查询已经安装的CC
	list, err := swp.GetInstalledChaincode(targetPeer)
	if err != nil {
		return err
	}

	// TODO
	newCC := swp.ChainCodeID + swp.ChainCodeVersion

	for _, v := range list {
		if newCC == v {
			fmt.Println("安装失败，名为：" + newCC + " 的chaincode已经安装在了节点" + targetPeer)
			return nil
		}
	}

	// 打包chaincode

	ccPkg, err := packager.NewCCPackage(swp.ChaincodePath, swp.ChaincodeGoPath)
	if err != nil {
		return errors.Errorf("打包chaincode失败，错误为：%s", err)
	}

	// 安装chaincode
	installCCReq := resmgmt.InstallCCRequest{
		Name:    swp.ChainCodeID,
		Path:    swp.ChaincodePath,
		Version: swp.ChainCodeVersion,
		Package: ccPkg,
	}

	_, err = swp.Resmgmt.InstallCC(
		installCCReq,
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
		resmgmt.WithTargetEndpoints(targetPeer),
	)

	if err != nil {
		return errors.Errorf("向目标节点%s安装名为%s的chaincode失败，错误为：%s", targetPeer, newCC, err)
	}

	fmt.Printf("向目标节点%s安装名为%s的chaincode成功\n", targetPeer, newCC)

	return nil
}

// // 实例化chaincode
// func (swp SDK) InstantiateChaincode(targetPeer string) error {

// 	// 查询已经安装的CC
// 	list, err := swp.GetInstalledChaincode(targetPeer)
// 	if err != nil {
// 		return err
// 	}

// 	installed := false

// 	newCC := swp.ChainCodeID + swp.ChainCodeVersion
// 	for _, v := range list {
// 		if newCC == v {
// 			installed = true
// 		}
// 	}
// 	if !installed {
// 		fmt.Println("实例化失败，名为：" + newCC + " 的chaincode还没有安装在节点" + targetPeer)
// 		return nil
// 	}
// 	// 查询已经实例化的CC
// 	list, err = swp.GetInstantiatedChaincode(swp.ChannelID, targetPeer)
// 	if err != nil {
// 		return err
// 	}
// 	newCC = swp.ChainCodeID + swp.ChainCodeVersion
// 	for _, v := range list {
// 		if newCC == v {
// 			fmt.Println("实例化失败，名为：" + newCC + " 的chaincode已经实例化在节点" + targetPeer)
// 			return nil
// 		}
// 	}

// 	// 这里的参数名是msp名称 不是域名  TODO
// 	ccPolicy := cauthdsl.SignedByAnyMember(
// 		[]string{
// 			"eastMSP",
// 			"BoanMSP",
// 			"NorthMSP",
// 		})

// 	// TODO init
// 	ccInitArgs := [][]byte{[]byte("init"), []byte(" ")}

// 	request := resmgmt.InstantiateCCRequest{
// 		Name:    swp.ChainCodeID,
// 		Path:    swp.ChaincodeGoPath,
// 		Version: swp.ChainCodeVersion,
// 		Args:    ccInitArgs,
// 		Policy:  ccPolicy,
// 	}

// 	// opts := requestOptions{Targets: peers}
// 	resp, err := swp.Resmgmt.InstantiateCC(
// 		swp.ChannelID,
// 		request, /*,resmgmt.WithOrdererEndpoint("orderer0.antifake.com") */
// 		resmgmt.WithTargetEndpoints(targetPeer),
// 	)
// 	if err != nil || resp.TransactionID == "" {
// 		return errors.Errorf("实例化失败，名为%s的chaincode实例化到节点%s上失败，错误为：%s", newCC, targetPeer, err)
// 	}
// 	fmt.Printf("实例化成功，实例化名为%s的chaincode成功到节点%s上成功\n", newCC, targetPeer)
// 	return nil
// }

// // 升级chaincode
// func (swp SDK) UpgradeChaincode(targetPeer string) error {
// 	// 查询已经安装的CC
// 	list, err := swp.GetInstalledChaincode(targetPeer)
// 	if err != nil {
// 		return err
// 	}
// 	installed := false
// 	newCC := swp.ChainCodeID + swp.ChainCodeVersion
// 	for _, v := range list {
// 		if newCC == v {
// 			installed = true
// 		}
// 	}
// 	if !installed {
// 		fmt.Println("升级chaincode失败，名为：" + newCC + " 的chaincode还没有安装在节点" + targetPeer)
// 		return nil
// 	}
// 	// 查询已经实例化的CC
// 	list, err = swp.GetInstantiatedChaincode(swp.ChannelID, targetPeer)
// 	if err != nil {
// 		return err
// 	}
// 	newCC = swp.ChainCodeID + swp.ChainCodeVersion
// 	for _, v := range list {
// 		if newCC == v {
// 			fmt.Println("升级chaincode失败，名为：" + newCC + " 的chaincode已经实例化在节点" + targetPeer)
// 			return nil
// 		}
// 	}

// 	// 这里的参数名是msp名称 不是域名
// 	ccPolicy := cauthdsl.SignedByAnyMember(
// 		[]string{
// 			"eastMSP",
// 			"BoanMSP",
// 			"NorthMSP",
// 		})

// 	// TODO
// 	ccInitArgs := [][]byte{[]byte("init"), []byte(" ")}

// 	ccUpgradeRequest := resmgmt.UpgradeCCRequest{
// 		Name:    swp.ChainCodeID,
// 		Path:    swp.ChaincodeGoPath,
// 		Version: swp.ChainCodeVersion,
// 		Args:    ccInitArgs,
// 		Policy:  ccPolicy,
// 	}

// 	resp, err := swp.Resmgmt.UpgradeCC(
// 		swp.ChannelID,
// 		ccUpgradeRequest,
// 		resmgmt.WithTargetEndpoints(targetPeer),
// 	)

// 	if err != nil || resp.TransactionID == "" {
// 		return errors.Errorf("升级chaincode失败，名为%s的chaincode升级到节点%s上失败，错误为：%s", newCC, targetPeer, err)
// 	}
// 	fmt.Printf("升级chaincode成功，实例化名为%s的chaincode成功到节点%s上成功\n", newCC, targetPeer)
// 	return nil
// }
