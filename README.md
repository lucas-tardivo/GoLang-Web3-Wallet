# GoLang-Web3-Wallet
A RESTful Web3 Wallet manager created with GoLang

# Endpoints
## Auth
- Sign In: **/api/auth/{wallet}**

Returns a JWT Bearer Token in order to use the wallet endpoints.

## ERC20
- Wallet balance: **/api/erc20/{contractAddress}**
  - (Bearer token required)  

Returns the balanceOf function from a ERC20 contract for the connected wallet.
