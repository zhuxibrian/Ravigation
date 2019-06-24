package agvManager

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"Ravigation/ravigation/storage"
	"Ravigation/ravigation/utils"
	"strings"
)

func SubcsribeAgvMsg() {
	utils.Subcsribe(TOPIC_RAVI_AGV_STATE+ "/#" , 1, AgvStateHandler)
	utils.Subcsribe(TOPIC_RAVI_AGV_FINISH+ "/#", 1, AgvFinishHandler)
}


func AgvStateHandler(client mqtt.Client, message mqtt.Message) {
	conn := storage.GetConn()
	defer conn.Close()

	var agv storage.AgvState
	json.Unmarshal(message.Payload(), &agv)

	if err := storage.SetAgvState(conn, agv); err != nil {
		logrus.Info("AgvStateHandler error : ", err)
	}
}

func AgvFinishHandler(client mqtt.Client, message mqtt.Message)  {
	conn := storage.GetConn()
	defer conn.Close()

	subTopic := strings.Split(string(message.Topic()), "/")
	agvName := subTopic[2]

	storage.DeleteCmdFromRunList(conn, agvName)
}