package model

// import (
// 	"boan/utils"
// 	"encoding/hex"
// 	"fmt"

// 	"github.com/golang/protobuf/proto"
// 	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
// 	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
// 	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

// 	//cbb "github.com/hyperledger/fabric/protos/common"
// 	"time"

// 	cb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
// 	putils "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/utils"
// 	pb "github.com/hyperledger/fabric/protos/peer"
// 	"github.com/pkg/errors"
// )

// type LedgerInfo struct {
// 	CurrentBlockHash  string
// 	Height            int
// 	PreviousBlockHash string
// }

// func (s *Setup) LedgerInfo() (*LedgerInfo, error) {

// 	// 准备context
// 	boanAdminChannelContext := s.FabSDK.ChannelContext(s.ChannelID, fabsdk.WithUser(s.OrgAdmin), fabsdk.WithOrg(s.OrgName))

// 	//创建ledger client
// 	ledgerClient, err := ledger.New(boanAdminChannelContext)
// 	if err != nil {
// 		return nil, errors.Errorf("创建ledger client失败", err)
// 	}
// 	ledgerInfo, err := ledgerClient.QueryInfo()
// 	if err != nil {
// 		return nil, errors.Errorf("查询ledger info失败", err)
// 	}

// 	result := &LedgerInfo{
// 		CurrentBlockHash:  hex.EncodeToString(ledgerInfo.BCI.CurrentBlockHash),
// 		PreviousBlockHash: hex.EncodeToString(ledgerInfo.BCI.PreviousBlockHash),
// 		Height:            int(ledgerInfo.BCI.Height),
// 	}

// 	//fmt.Println("ledgerClient: ", ledgerClient)
// 	//fmt.Println("ledgerInfo: ", ledgerInfo)
// 	//fmt.Println("ledgerInfo.BCI.CurrentBlockHash: ", utils.Base58Encode(ledgerInfo.BCI.CurrentBlockHash))
// 	//fmt.Println("ledgerInfo.BCI.Height: ", ledgerInfo.BCI.Height)
// 	//fmt.Println("ledgerInfo.BCI.PreviousBlockHash: ", utils.Base58Encode(ledgerInfo.BCI.PreviousBlockHash))

// 	//bb, err := utils.Base64Decode("0jAn18da5DO/2+SthvESrY8kpE1ikxQGc5fnaVBb/Yw=")
// 	//if err != nil {
// 	//	return errors.Errorf("base64 string to []byte failed", err)
// 	//}
// 	//block2, err := ledgerClient.QueryBlockByHash(bb)
// 	//if err != nil {
// 	//	return errors.Errorf("查询block2失败", err)
// 	//}
// 	//fmt.Println("block2: ", block2.Header.Number)

// 	//if err := s.ChannelManager.JoinChannel(targetChannel,
// 	//	resmgmt.WithRetry(retry.DefaultResMgmtOpts),
// 	//	resmgmt.WithOrdererEndpoint(s.Targetorderers[0])); err != nil {
// 	//	return errors.Errorf("Admin加入channel失败，错误为：%s", err)
// 	//}
// 	//fmt.Printf("组织%s的用户%s加入%s成功\n", org, targetPeer, targetChannel)
// 	return result, nil
// }

// func (s *Setup) QueryByTxID(txid string) (*TransactionDetail, error) {
// 	// 准备context
// 	boanAdminChannelContext := s.FabSDK.ChannelContext(s.ChannelID, fabsdk.WithUser(s.OrgAdmin), fabsdk.WithOrg(s.OrgName))

// 	//创建ledger client
// 	ledgerClient, err := ledger.New(boanAdminChannelContext)
// 	if err != nil {
// 		return nil, errors.Errorf("创建ledger client失败", err)
// 	}

// 	input := fab.TransactionID(txid)
// 	block, err := ledgerClient.QueryBlockByTxID(input)
// 	if err != nil {
// 		return nil, errors.Errorf("根据tx id 查询失败", err)
// 	}

// 	fmt.Println("block.Header.Number", block.Header.Number)

// 	ledgerInfo, err := ledgerClient.QueryInfo()
// 	if err != nil {
// 		return nil, errors.Errorf("查询ledger info失败", err)
// 	}
// 	var CurrentDataHash []byte
// 	if block.Header.Number+1 == ledgerInfo.BCI.Height {
// 		CurrentDataHash = ledgerInfo.BCI.CurrentBlockHash
// 	} else {
// 		block2, err := ledgerClient.QueryBlock(block.Header.Number + 1)
// 		if err != nil {
// 			return nil, errors.Errorf("查询下一个块失败", err)
// 		}
// 		CurrentDataHash = block2.Header.PreviousHash
// 	}

// 	fmt.Println("block.Header.CurrentDataHash", utils.Base58Encode(CurrentDataHash))
// 	fmt.Println("block.Header.PreviousHash", utils.Base58Encode(block.Header.PreviousHash))

// 	result, err := GetTransactionInfoFromData(block.Data.Data[0], utils.Base58Encode(CurrentDataHash))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// 	//for _, v := range block.Data.Data {
// 	//	result, err := GetTransactionInfoFromData(v)
// 	//	if err!=nil {
// 	//		return nil ,err
// 	//	}
// 	//
// 	//}
// 	//
// 	//return nil

// }

// type TransactionDetail struct {
// 	TransactionId string
// 	CreateTime    string
// 	Args          []string
// 	BlockHash     string
// }

// func GetTransactionInfoFromData(data []byte, blockHash string) (*TransactionDetail, error) {

// 	env, err := putils.GetEnvelopeFromBlock(data)
// 	if err != nil {
// 		return nil, errors.Errorf("error extracting Envelope from block", err)
// 	}
// 	if env == nil {
// 		return nil, errors.Errorf("empty envelope")
// 	}

// 	payload, err := putils.GetPayload(env)
// 	if err != nil {
// 		return nil, errors.Errorf("error extracting Payload from envelope", err)
// 	}

// 	channelHeaderBytes := payload.Header.ChannelHeader
// 	channelHeader := &cb.ChannelHeader{}

// 	if err := proto.Unmarshal(channelHeaderBytes, channelHeader); err != nil {
// 		return nil, errors.Errorf("error extracting ChannelHeader from payload", err)
// 	}

// 	args := []string{}

// 	//获取tx
// 	tx, err := putils.GetTransaction(payload.Data)
// 	if err != nil {
// 		return nil, errors.Errorf("error unmarshalling transaction payload", err)
// 	}

// 	chaincodeActionPayload, err := putils.GetChaincodeActionPayload(tx.Actions[0].Payload)
// 	if err != nil {
// 		return nil, errors.Errorf("error unmarshalling chaincode action payload", err)
// 	}
// 	propPayload := &pb.ChaincodeProposalPayload{}
// 	if err := proto.Unmarshal(chaincodeActionPayload.ChaincodeProposalPayload, propPayload); err != nil {
// 		return nil, errors.Wrap(err, "error extracting ChannelHeader from payload")
// 	}
// 	invokeSpec := &pb.ChaincodeInvocationSpec{}
// 	err = proto.Unmarshal(propPayload.Input, invokeSpec)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "error extracting ChannelHeader from payload")
// 	}
// 	for _, v := range invokeSpec.ChaincodeSpec.Input.Args {
// 		args = append(args, string(v))
// 	}

// 	res := &TransactionDetail{
// 		TransactionId: channelHeader.TxId,
// 		Args:          args,
// 		CreateTime:    time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05"),
// 		BlockHash:     blockHash,
// 	}
// 	return res, nil

// }

// type BlockInfo struct {
// 	BlockNum          int
// 	PreviousBlockHash string
// 	BlockHash         string
// 	SizeOfBlockHeader string
// 	SizeOfBlock       string
// 	NumsOfTx          int
// 	TxIDs             []string
// }

// func (s *Setup) QueryBlockByBlockHash(blockHash string) (*BlockInfo, error) {
// 	// 准备context
// 	boanAdminChannelContext := s.FabSDK.ChannelContext(s.ChannelID, fabsdk.WithUser(s.OrgAdmin), fabsdk.WithOrg(s.OrgName))

// 	//创建ledger client
// 	ledgerClient, err := ledger.New(boanAdminChannelContext)
// 	if err != nil {
// 		return nil, errors.Errorf("创建ledger client失败", err)
// 	}
// 	bb := utils.Base58Decode(blockHash)
// 	block, err := ledgerClient.QueryBlockByHash(bb)
// 	if err != nil {
// 		return nil, errors.Errorf("查询block失败", err)
// 	}

// 	blockSize := len(block.String())
// 	blockHeaderSize := len(block.Header.String())
// 	result := &BlockInfo{
// 		BlockNum:          int(block.Header.Number),
// 		BlockHash:         blockHash,
// 		PreviousBlockHash: utils.Base58Encode(block.Header.PreviousHash),
// 		SizeOfBlock:       fmt.Sprintf("%d", blockSize),
// 		SizeOfBlockHeader: fmt.Sprintf("%d", blockHeaderSize),
// 		NumsOfTx:          len(block.Data.Data),
// 	}
// 	return result, nil

// 	//fmt.Println("block.Header.Number", block.Header.Number)

// 	//blockHashBytes, err := utils.Base64Decode(blockHash)
// 	//if err != nil {
// 	//	return err
// 	//}
// 	//block, err := ledgerClient.QueryBlockByHash(blockHashBytes)
// 	//if err != nil {
// 	//	return err
// 	//}
// 	//fmt.Println("block.Header.Number", block.Header.Number)
// 	//return nil
// }

// func (s *Setup) QueryBlock(blockHash string, blockNum uint64, txId string) error {
// 	// 准备context
// 	boanAdminChannelContext := s.FabSDK.ChannelContext(s.ChannelID, fabsdk.WithUser(s.OrgAdmin), fabsdk.WithOrg(s.OrgName))

// 	//创建ledger client
// 	ledgerClient, err := ledger.New(boanAdminChannelContext)
// 	if err != nil {
// 		return errors.Errorf("创建ledger client失败", err)
// 	}

// 	block1, err := ledgerClient.QueryBlockByHash(utils.Base58Decode(blockHash))
// 	if err != nil {
// 		return errors.Errorf("QueryBlockByHash failed", err)
// 	}

// 	fmt.Println("block1.Header.DataHash", block1.Header.Number)
// 	fmt.Println("block1.Header.DataHash", utils.Base58Encode(block1.Header.DataHash))

// 	block2, err := ledgerClient.QueryBlock(blockNum)
// 	if err != nil {
// 		return errors.Errorf("QueryBlockByNum failed", err)
// 	}
// 	fmt.Println("block2.Header.DataHash", block2.Header.Number)
// 	fmt.Println("block2.Header.DataHash", utils.Base58Encode(block2.Header.DataHash))

// 	block3, err := ledgerClient.QueryBlockByTxID(fab.TransactionID(txId))
// 	if err != nil {
// 		return errors.Errorf("QueryBlockByTxID failed", err)
// 	}
// 	fmt.Println("block3.Header.DataHash", block3.Header.Number)
// 	fmt.Println("block3.Header.DataHash", utils.Base58Encode(block3.Header.DataHash))
// 	return nil
// }

// func (s *Setup) QueryTxTime(txid string) (string, error) {
// 	// 准备context
// 	boanAdminChannelContext := s.FabSDK.ChannelContext(s.ChannelID, fabsdk.WithUser(s.OrgAdmin), fabsdk.WithOrg(s.OrgName))

// 	//创建ledger client
// 	ledgerClient, err := ledger.New(boanAdminChannelContext)
// 	if err != nil {
// 		return "", errors.Errorf("创建ledger client失败", err)
// 	}

// 	input := fab.TransactionID(txid)
// 	block, err := ledgerClient.QueryBlockByTxID(input)
// 	if err != nil {
// 		return "", errors.Errorf("根据tx id 查询失败", err)
// 	}

// 	for _, txData := range block.Data.Data {

// 		env, err := putils.GetEnvelopeFromBlock(txData)
// 		if err != nil {
// 			return "", errors.Errorf("error extracting Envelope from block", err)
// 		}
// 		if env == nil {
// 			return "", errors.Errorf("empty envelope")
// 		}

// 		payload, err := putils.GetPayload(env)
// 		if err != nil {
// 			return "", errors.Errorf("error extracting Payload from envelope", err)
// 		}
// 		channelHeaderBytes := payload.Header.ChannelHeader
// 		channelHeader := &cb.ChannelHeader{}
// 		if err := proto.Unmarshal(channelHeaderBytes, channelHeader); err != nil {
// 			return "", errors.Errorf("error extracting ChannelHeader from payload", err)
// 		}
// 		txidInner := channelHeader.TxId
// 		if txidInner != txid {
// 			continue
// 		}
// 		return time.Unix(channelHeader.Timestamp.Seconds, 0).Format("2006-01-02 15:04:05"), nil
// 	}
// 	return "", errors.New("failed to get create time")
// }

// func (s *Setup) AllTxTime(blockNum int64) (string, error) {
// 	// 准备context
// 	boanAdminChannelContext := s.FabSDK.ChannelContext(s.ChannelID, fabsdk.WithUser(s.OrgAdmin), fabsdk.WithOrg(s.OrgName))

// 	//创建ledger client
// 	ledgerClient, err := ledger.New(boanAdminChannelContext)
// 	if err != nil {
// 		return "", errors.Errorf("创建ledger client失败", err)
// 	}

// 	block, err := ledgerClient.QueryBlock(uint64(blockNum))
// 	if err != nil {
// 		return "", errors.Errorf("根据block num 查询失败", err)
// 	}
// 	mres := make(map[string]int64)
// 	for _, txData := range block.Data.Data {

// 		env, err := putils.GetEnvelopeFromBlock(txData)
// 		if err != nil {
// 			return "", errors.Errorf("error extracting Envelope from block", err)
// 		}
// 		if env == nil {
// 			return "", errors.Errorf("empty envelope")
// 		}

// 		payload, err := putils.GetPayload(env)
// 		if err != nil {
// 			return "", errors.Errorf("error extracting Payload from envelope", err)
// 		}
// 		channelHeaderBytes := payload.Header.ChannelHeader
// 		channelHeader := &cb.ChannelHeader{}
// 		if err := proto.Unmarshal(channelHeaderBytes, channelHeader); err != nil {
// 			return "", errors.Errorf("error extracting ChannelHeader from payload", err)
// 		}
// 		mres[channelHeader.TxId] = channelHeader.Timestamp.Seconds
// 	}
// 	fmt.Println("~~~~~~~~")
// 	fmt.Printf("%+v\n", mres)
// 	fmt.Println("~~~~~~~~")
// 	return "", errors.New("failed to get create time")
// }
