package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// programmatically set swagger info
	//docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	//docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	//docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Static("/dist", "./dist")
	return router
}
