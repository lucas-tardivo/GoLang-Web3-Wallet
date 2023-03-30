package controller

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/common"
)

var (
	chainRpc	string
	chainId		int64
	contractERC20Abi	string
	contractERC721Abi	string
)

func init() {
	godotenv.Load("app.env")

	intChainId, err := strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
	if err != nil {
		panic(err)
	}else{
		chainId = intChainId
	}

	chainRpc = os.Getenv("CHAIN_RPC")

	erc20Abi, err := os.ReadFile("contract-erc20-abi.txt")
	if err != nil {
		panic(err)
	}

	contractERC20Abi = string(erc20Abi)

	erc721Abi, err := os.ReadFile("contract-erc721-abi.txt")
	if err != nil {
		panic(err)
	}

	contractERC721Abi = string(erc721Abi)
}

func GetUserWalletAddress(c *gin.Context) (common.Address) {
	wallet := c.MustGet("currentWallet").(string)

	return common.HexToAddress(wallet)
}

func SetupWeb3() (*web3.Web3) {
	web3, err := web3.NewWeb3(chainRpc)
	if err != nil {
		panic(err)
	}

	web3.Eth.SetChainId(chainId)

	return web3
}