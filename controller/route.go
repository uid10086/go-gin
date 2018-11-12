package controller

import (
	"github.com/gin-gonic/gin"
)

func init()  {
	//gin.DisableConsoleColor()
	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//os.Setenv("Authorization","lichuanqi")
}
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(Logger())

	r.Use(gin.Recovery())



	video := r.Group("/video",Middleware)//视频
	{
		video.GET("",GetUser)

	}

	return r
}