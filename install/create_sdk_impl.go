package model

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func CreateresMgmtClient() error {
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
