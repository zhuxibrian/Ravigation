package apiserver

import (
	"Ravigation/ravigation/service"
	"Ravigation/ravigation/storage"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

func setCmdGroupRouter() {
	cmdRouter := r.Group("/cmdgroup")
	{
		cmdRouter.POST("/cmdgroup", func(c *gin.Context) {
			body, _ := ioutil.ReadAll(c.Request.Body)

			var cg storage.CmdGroup
			if err := cg.UnmarshalJSON(body); err != nil {
				response(c, Code_Err, "Unmarshal CmdGroup fail")
				return
			}

			err := service.AddCmdGroup(cg)
			if err != nil {
				logrus.Error("create cmd group error:", err)
				response(c, Code_Err, "create cmd group error")
				return
			}
			response(c, Code_OK, "")
			return
		})
	}
}