package controller

import (
	"github.com/gin-gonic/gin"
	"lcq/services"
	"lcq/libs/resp"
	log "github.com/sirupsen/logrus"
)

func GetUser (c * gin.Context)  {
	name := c.Query("name")
	m := make(map[string]interface{})
	m["Name"] = name
    log.Debug("asdasd")
    result,err := services.Only().Service.Getuser(m)
    if err != nil {
    	resp.Error(c,err)
    	return
	}
	resp.Success(c,result)
}

