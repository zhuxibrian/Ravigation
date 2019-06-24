package apiserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Code_OK = iota
	Code_Err

)

func response(c *gin.Context, code int32, msg string)  {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}
