package nodeManager

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"Ravigation/ravigation/storage"
	"Ravigation/ravigation/utils"
	"strings"
	"time"
)

func SubcsribeNodeMsg() {
	utils.Subcsribe(TOPIC_RAVI_NODE_CONNECT + "/#" , 2, nodeCommandHandler)
}


func nodeCommandHandler(client mqtt.Client, message mqtt.Message) {
	conn := storage.GetConn()
	defer conn.Close()

	subTopic := strings.Split(string(message.Topic()), "/")
	nodeName := subTopic[2]
	btnName := subTopic[3]

	cmdGroupName, err := storage.GetBtnCmdMap(conn, nodeName, btnName)
	if  err != nil {
		logrus.Error("node manager GetBtnCmdMap error:", err)
		return
	}

	ncg := storage.NodeCmdGroup{
		NodeName:nodeName,
		BtnName:btnName,
		Timestamp:time.Now().Unix(),
		CmdGroupName:cmdGroupName,
	}
	storage.RPushCmdGroupToWaitList(conn, &ncg)
}
