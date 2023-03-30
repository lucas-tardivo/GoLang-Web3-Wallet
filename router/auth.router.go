package router

import (
	"web3-wallet/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")
	router.POST("/:wallet", controller.SignInUser)
}