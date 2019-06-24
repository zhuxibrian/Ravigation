package main

import (
	"Ravigation/ravigation/apiserver"
	"Ravigation/ravigation/controllerManager"
	"Ravigation/ravigation/scheduler"
	"Ravigation/ravigation/utils"
)


func main() {
	utils.Connect()

	controllerManager.Start()
	scheduler.Start()
	apiserver.Start()

	utils.Disconnect()
}

