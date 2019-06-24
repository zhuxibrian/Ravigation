package controllerManager

import (
	"github.com/sirupsen/logrus"
	"Ravigation/ravigation/controllerManager/agvManager"
	"Ravigation/ravigation/controllerManager/nodeManager"
	"Ravigation/ravigation/storage"
	"time"
)



func Start() {
	agvManager.SubcsribeAgvMsg()
	nodeManager.SubcsribeNodeMsg()

	go ConnectAllDevice()
}


func ConnectAllDevice() {
	conn := storage.GetConn()
	defer conn.Close()
	for {
		var infoList []storage.ConnectInfo
		if err := storage.GetAllConnectInfo(conn, &infoList); err != nil {
			logrus.Error("get all agv keys error")
			continue
		}

		for _, info := range infoList {

			switch info.Device {
			case "agv":
				agvManager.ConnectAgv(conn, info)
			case "node":
			default:
				logrus.Error("no define the device type")
			}

		}
		time.Sleep(time.Duration(20) * time.Second)
	}
}


