package router

import (
	"web3-wallet/controller"
	"web3-wallet/middleware"

	"github.com/gin-gonic/gin"
)

func ERC20Route(rg *gin.RouterGroup) {
	router := rg.Group("erc20")
	router.Use(middleware.DeserializeUser())
	router.GET("/balance/:contractAddress", controller.GetERC20Balance)
	router.GET("/supply/:contractAddress", controller.GetERC20TotalSupply)
	router.GET("/allowance/:contractAddress/:spenderAddress", controller.GetERC20Allowance)
}