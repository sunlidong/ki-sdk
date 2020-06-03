package model

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	sdkConfig "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
)

type ResmgmtClient struct {
	Client *resmgmt.Client
}

// 创建ResmgmtClient
func GetResmgmtClient(configFile string, org string) (*resmgmt.Client, error) {
	sdk, err := fabsdk.New(sdkConfig.FromFile(configFile))
	if err != nil {
		return nil, errors.Errorf("创建FabricSDK失败，错误为：%s", err)
	}
	// add new code
	//defer sdk.Close()
	resourceManagerClientContext := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg(org))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return nil, errors.Errorf("创建组织：%s 的资源管理客户端失败，错误为：%s", org, err)
	}
	return resMgmtClient, nil
}

// 查询已经安装的cc
func (c *ResmgmtClient) GetInstalledCC(peerHost string) ([]string, error) {
	// 查询已经安装的CC
	ccInstalledRes, err := c.Client.QueryInstalledChaincodes(resmgmt.WithTargetEndpoints(peerHost))
	if err != nil {
		return nil, errors.Errorf("查询节点%s已安装的chaincode失败，错误为：%s", peerHost, err)
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

// 查询已经实例化的cc
func (c *ResmgmtClient) GetInstantiatedCC(channelName string, peerHost string) ([]string, error) {
	// 查询已经实例化的CC
	ccInstantiatedRes, err := c.Client.QueryInstantiatedChaincodes(channelName, resmgmt.WithTargetEndpoints(peerHost))
	if err != nil {
		return nil, errors.Errorf("查询目标节点%s在channel:%s已安装的chaincode失败，错误为：%s", peerHost, channelName, err)
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
func (c *ResmgmtClient) InstallCCDepend(peerHost string, ccName string, ccVersion string, ccPath string, ccGoPath string) error {
	// 查询已经安装的CC
	list, err := c.GetInstalledCC(peerHost)
	if err != nil {
		return err
	}
	newCC := ccName + ccVersion
	for _, v := range list {
		if newCC == v {
			return errors.New("安装失败，名为：" + newCC + " 的chaincode已经安装在了节点" + peerHost)
		}
	}
	// 打包chaincode
	ccPkg, err := packager.NewCCPackage(ccPath, ccGoPath)
	if err != nil {
		return errors.Errorf("打包chaincode失败，错误为：%s", err)
	}
	// 安装chaincode
	installCCReq := resmgmt.InstallCCRequest{Name: ccName, Path: ccPath, Version: ccVersion, Package: ccPkg}
	_, err = c.Client.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithTargetEndpoints(peerHost))
	if err != nil {
		return errors.Errorf("向目标节点%s安装名为%s的chaincode失败，错误为：%s", peerHost, newCC, err)
	}
	return nil
}

// 实例化chaincode
func (c *ResmgmtClient) InstantiatedCCDepend(peerHost string, ccName string, ccVersion string, ccSigners []string, ccPath string, ccGoPath string, channelName string) error {
	// 查询已经安装的CC
	list, err := c.GetInstalledCC(peerHost)
	if err != nil {
		return err
	}
	installed := false
	newCC := ccName + ccVersion
	for _, v := range list {
		if newCC == v {
			installed = true
		}
	}
	if !installed {
		return errors.New("实例化失败，名为：" + newCC + " 的chaincode还没有安装在节点" + peerHost)
	}
	// 查询已经实例化的CC
	list, err = c.GetInstantiatedCC(channelName, peerHost)
	if err != nil {
		return err
	}
	for _, v := range list {
		if newCC == v {
			return errors.New("实例化失败，名为：" + newCC + " 的chaincode已经实例化在节点" + peerHost)
		}
	}
	// 这里的参数名是msp名称 不是域名
	ccPolicy := cauthdsl.SignedByAnyMember(ccSigners)
	ccInitArgs := [][]byte{[]byte("init"), []byte(" ")}
	request := resmgmt.InstantiateCCRequest{
		Name:    ccName,
		Path:    ccGoPath,
		Version: ccVersion,
		Args:    ccInitArgs,
		Policy:  ccPolicy,
	}
	// opts := requestOptions{Targets: peers}
	resp, err := c.Client.InstantiateCC(channelName, request /*,resmgmt.WithOrdererEndpoint("orderer0.antifake.com") */, resmgmt.WithTargetEndpoints(peerHost))
	if err != nil || resp.TransactionID == "" {
		return errors.Errorf("实例化失败，名为%s的CC实例化到节点%s上失败，错误为：%s", newCC, peerHost, err)
	}
	return nil
}

// 升级chaincode
func (c *ResmgmtClient) UpgradeCCDepend(peerHost string, ccName string, ccVersion string, ccSigners []string, ccPath string, ccGoPath string, channelName string) error {
	// 查询已经安装的CC
	list, err := c.GetInstalledCC(peerHost)
	if err != nil {
		return err
	}
	installed := false
	newCC := ccName + ccVersion
	for _, v := range list {
		if newCC == v {
			installed = true
		}
	}
	if !installed {
		return errors.New("升级CC失败，名为：" + newCC + " 的CC还没有安装在节点" + peerHost)
	}
	// 查询已经实例化的CC
	list, err = c.GetInstantiatedCC(channelName, peerHost)
	if err != nil {
		return err
	}
	for _, v := range list {
		if newCC == v {
			return errors.New("升级CC失败，名为：" + newCC + " 的CC已经实例化在节点" + peerHost)
		}
	}

	// 这里的参数名是msp名称 不是域名
	ccPolicy := cauthdsl.SignedByAnyMember(ccSigners)
	ccInitArgs := [][]byte{[]byte("init"), []byte(" ")}
	ccUpgradeRequest := resmgmt.UpgradeCCRequest{Name: ccName, Path: ccGoPath, Version: ccVersion, Args: ccInitArgs, Policy: ccPolicy}
	resp, err := c.Client.UpgradeCC(channelName, ccUpgradeRequest, resmgmt.WithTargetEndpoints(peerHost))
	if err != nil || resp.TransactionID == "" {
		return errors.Errorf("升级CC失败，名为%s的CC升级到节点%s上失败，错误为：%s", newCC, peerHost, err)
	}
	return nil
}
