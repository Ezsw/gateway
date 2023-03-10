package controller

import "github.com/gin-gonic/gin"

type LoginController struct{}

func LoginRegister(group *gin.RouterGroup) {
	Login := &LoginController{}
	group.POST("/login", Login.Login)
	group.GET("/logout", Login.Logout)
}

func (*LoginController) Login(c *gin.Context) {

}

func (*LoginController) Logout(c *gin.Context) {

}
