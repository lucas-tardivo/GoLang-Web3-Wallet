package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ethereum/go-ethereum/common"
)

func GetERC20Balance(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	pubKey := GetUserWalletAddress(c)

	web3 := SetupWeb3()

	contract, err := web3.Eth.NewContract(contractERC20Abi, contractAddress)
	if err != nil {
		panic(err)
	}

	balance, err := contract.Call("balanceOf", pubKey)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": balance})
}

func GetERC20TotalSupply(c *gin.Context) {
	contractAddress := c.Param("contractAddress")

	web3 := SetupWeb3()

	contract, err := web3.Eth.NewContract(contractERC20Abi, contractAddress)
	if err != nil {
		panic(err)
	}

	totalSupply, err := contract.Call("totalSupply")
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": totalSupply})
}

func GetERC20Allowance(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	spenderAddress := c.Param("spenderAddress")
	pubKey := GetUserWalletAddress(c)

	web3 := SetupWeb3()

	contract, err := web3.Eth.NewContract(contractERC20Abi, contractAddress)
	if err != nil {
		panic(err)
	}

	allowance, err := contract.Call("allowance", pubKey, common.HexToAddress(spenderAddress))
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": allowance})
}