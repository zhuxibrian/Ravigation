package apiserver

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"Ravigation/ravigation/service"
	"Ravigation/ravigation/storage"
)

func setNodeRouter() {
	nodeRouter := r.Group("/node")
	{
		nodeRouter.POST("/btncmdmap", func(c *gin.Context) {
			body, _ := ioutil.ReadAll(c.Request.Body)
			var btnCmdMap storage.ButtonCmdGroup
			if err := json.Unmarshal(body, &btnCmdMap); err != nil {
				response(c, Code_Err, "Unmarshal Button_Cmd map fail")
				return
			}

			if err := service.AddBtnCmdMap(btnCmdMap); err != nil {
				response(c, Code_Err, "add Button_Cmd map fail")
				return
			}

			response(c, Code_OK, "")
		})
	}
}