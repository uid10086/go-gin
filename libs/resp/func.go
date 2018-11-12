package resp

import (
	"github.com/gin-gonic/gin"
)

func Error(c * gin.Context,err error)  {
	c.JSON (0, gin.H{
		"msg": err.Error(),
		"code":    0,
		})
}

func Success(c * gin.Context,data interface{}) () {
	c.JSON (200, gin.H{
		"msg": "success",
		"code":    200,
		"data":    data})
}


