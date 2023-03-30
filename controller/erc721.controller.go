package controller

import (
	"net/http"
	"math/big"

	"github.com/gin-gonic/gin"
)

func GetERC721Balance(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	pubKey := GetUserWalletAddress(c)

	web3 := SetupWeb3()

	contract, err := web3.Eth.NewContract(contractERC721Abi, contractAddress)
	if err != nil {
		panic(err)
	}

	balance, err := contract.Call("balanceOf", pubKey)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": balance})
}

func GetERC721OwnerOf(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	tokenId := c.Param("tokenId")

	web3 := SetupWeb3()

	contract, err := web3.Eth.NewContract(contractERC721Abi, contractAddress)
	if err != nil {
		panic(err)
	}

	tokenBigInteger := new(big.Int)
    tokenBigInteger, ok := tokenBigInteger.SetString(tokenId, 10)
    if !ok {
        panic(err)
    }

	owner, err := contract.Call("ownerOf", tokenBigInteger)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": owner})
}

func GetERC721TotalSupply(c *gin.Context) {
	contractAddress := c.Param("contractAddress")

	web3 := SetupWeb3()

	contract, err := web3.Eth.NewContract(contractERC721Abi, contractAddress)
	if err != nil {
		panic(err)
	}

	totalSupply, err := contract.Call("totalSupply")
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": totalSupply})
}