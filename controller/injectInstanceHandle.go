package controller

import (
	m "ki-sdk/model"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

/*
	------------- 调用  model 层 操作
*/

// chuang jian tong dao
func createChannelConnection(p *SystemByCreateChannelDb) error {

	err := m.CreateChannel(
		p.Path,
		p.ChannelName,
		p.Org,
		p.User,
		p.Orderers,
	)

	return err
}

// jia  ru  tong  dao
func channelPropertyAccess(p *SystemByJoinChannelDb) error {

	err := m.JoinChannel(
		p.Path,
		p.ChannelName,
		p.Org,
		p.User,
		p.Orderers,
	)

	return err
}

// type SystemByXnNodeInfoListFreeDb struct {
// 	PeerHost   string `json:"peerHost"`
// 	ConfigFile string `json:"configFile"`
// 	Org        string `json:"org"`
// }

// 查询 某个 节点 已经 实例化 的 链码
func xnEnumerateExistingNodes(p *SystemByXnNodeInfoListFreeDb) (list []string, err error) {

	resmgmtClient, err := func() (*resmgmt.Client, error) {

		resmgmt, err := m.GetResmgmtClient(
			p.ConfigFile,
			p.Org,
		)

		if err != nil {
			log.Println("实例化 resmgmt_Client 失败： ", err)
		}

		return resmgmt, nil
	}()
	//
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	resmgmtDB := m.ResmgmtClient{
		Client: resmgmtClient,
	}

	arr, err1 := resmgmtDB.GetInstalledCC(p.PeerHost)

	if err1 != nil {
		return nil, err1
	}

	return arr, nil

}

// 查询 某个 节点 已经 实例化 的 链码
func xnEnumerateExistingNodesByInsite(p *SystemByInstantiatedccDb) (list []string, err error) {

	resmgmtClient, err := func() (*resmgmt.Client, error) {

		resmgmt, err := m.GetResmgmtClient(
			p.ConfigFile,
			p.Org,
		)

		if err != nil {
			log.Println("实例化 resmgmt_Client 失败： ", err)
		}

		return resmgmt, nil
	}()
	//
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	resmgmtDB := m.ResmgmtClient{
		Client: resmgmtClient,
	}

	arr, err1 := resmgmtDB.GetInstantiatedCC(p.ChannelName, p.PeerHost)

	if err1 != nil {
		return nil, err1
	}

	return arr, nil

}

// 向某个节点安装链码
func xnEnumerateExistingNodesByInstallCCDepend(p *SystemByInstallCCDependDb) (list []string, err error) {

	resmgmtClient, err := func() (*resmgmt.Client, error) {

		resmgmt, err := m.GetResmgmtClient(
			p.ConfigFile,
			p.Org,
		)

		if err != nil {
			log.Println("实例化 resmgmt_Client 失败： ", err)
		}

		return resmgmt, nil
	}()
	//
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	resmgmtDB := m.ResmgmtClient{
		Client: resmgmtClient,
	}
	// type SystemByInstallCCDependDb struct {
	// 	PeerHost  string `json:"peerHost"`
	// 	CcName    string `json:"ccName"`
	// 	CcVersion string `json:"ccVersion"`
	// 	CcPath    string `json:"ccPath"`
	// 	CcGoPath  string `json:"ccGoPath"`
	// }
	err1 := resmgmtDB.InstallCCDepend(
		p.PeerHost,
		p.CcName,
		p.CcVersion,
		p.CcPath,
		p.CcGoPath,
	)

	return err1

}
