package router

import (
	"web3-wallet/controller"
	"web3-wallet/middleware"

	"github.com/gin-gonic/gin"
)

func ERC721Route(rg *gin.RouterGroup) {
	router := rg.Group("erc721")
	router.Use(middleware.DeserializeUser())
	router.GET("/balance/:contractAddress", controller.GetERC721Balance)
	router.GET("/owner/:contractAddress/:tokenId", controller.GetERC721OwnerOf)
	router.GET("/supply/:contractAddress", controller.GetERC721TotalSupply)
}