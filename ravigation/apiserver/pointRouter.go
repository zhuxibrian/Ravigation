package apiserver

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"Ravigation/ravigation/service"
	"Ravigation/ravigation/storage"
	"strconv"
)

func setPointRouter() {
	pointRouter := r.Group("/point")
	{
		pointRouter.POST("/:index", func(c *gin.Context) {
			pointIndex, _ := strconv.Atoi(c.Param("index"))

			body, _ := ioutil.ReadAll(c.Request.Body)

			var point storage.Point
			if err := json.Unmarshal(body, &point); err != nil {
				response(c, Code_Err, "Unmarshal Point fail")
				return
			}

			err := service.AddPoint(pointIndex, point)
			if err!= nil {
				response(c, Code_Err, "add Point fail")
				return
			}

			response(c, Code_OK, "")
		})
	}
}
