package controller

import (
	m "ki-sdk/model"
	"log"
)

/*
	------------- 调用  model 层 操作
*/

func createChannelConnection(p *SystemByCreateChannelDb) error {

	//1 .调用  model 层  操作
	// func CreateChannel(path string, channelName string, org string, user string, orderers []string) error {
	log.Println("调用  model 层  操作")
	err := m.CreateChannel(
		p.Path,
		p.ChannelName,
		p.Org,
		p.User,
		p.Orderers,
	)
	log.Println("调用  model 层  操作 err=>", err)
	return err
}
