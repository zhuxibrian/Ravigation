package apiserver

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"Ravigation/ravigation/service"
	"Ravigation/ravigation/storage"
)

func setConnectInfoRouter() {
	agvRouter := r.Group("/device")
	{
		agvRouter.POST("/connectInfo", func(c *gin.Context) {
			body, _ := ioutil.ReadAll(c.Request.Body)

			var info storage.ConnectInfo
			if err := json.Unmarshal(body, &info); err != nil {
				response(c, Code_Err, "Unmarshal connectinfo fail")
				return
			}

			if err := service.AddConnectInfo(info); err != nil {
				response(c, Code_Err, "agv add connectinfo fail")
				return
			}

			response(c, Code_OK, "")
		})

		agvRouter.GET("/device/:device/name/:name", func(c *gin.Context) {
			device := c.Params.ByName("device")
			name := c.Params.ByName("name")

			var info storage.ConnectInfo
			if err := service.GetConnectInfo(&info, device, name); err != nil {
				response(c, Code_Err, "get device connect info error")
				return
			}

			if data, err := json.Marshal(info); err != nil {
				response(c, Code_Err, "Marshal connect info error")
			} else {
				response(c, Code_OK, string(data))
			}

			return
		})
	}
}


