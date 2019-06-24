package apiserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Ravigation/ravigation/utils"
)


var db = make(map[string]string)

var r *gin.Engine
func Start() {
	setupRouter()
	r.Run(utils.Config().GetString("ApiConfig.Url"))
}

func setupRouter() {
	r = gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	setConnectInfoRouter()
	setPointRouter()
	setCmdGroupRouter()
	setNodeRouter()


	return

}

