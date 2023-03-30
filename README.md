# GoLang-Web3-Wallet
A RESTful Web3 Wallet manager created with GoLang

# Endpoints
## Auth
### Sign In 
Endpoint: **/api/auth/{wallet}**

Returns a JWT Bearer Token in order to use the wallet endpoints.

## ERC20
### Total Token Supply
Endpoint: **/api/erc20/supply/{contractAddress}**

Returns the total supply of a ERC20 contract.

### Wallet balance
Endpoint: **/api/erc20/balance/{contractAddress}**
  - (Bearer token required)  

Returns the balanceOf function from a ERC20 contract for the connected wallet.

### Wallet allowance
Endpoint: **/api/erc20/allowance/{contractAddress}/{spenderAddress}**
  - (Bearer token required)  

Returns the allowance between the connected wallet and the spender address from a ERC20 contract.

## ERC721
### Total Token Supply
Endpoint: **/api/erc721/supply/{contractAddress}**

Returns the total supply of a ERC721 contract.

### Owner of token
Endpoint: **/api/erc721/owner/{contractAddress}/{tokenId}**

Returns the owner address for the specified TokenId from a ERC721 contract.

### Wallet balance
Endpoint: **/api/erc721/balance/{contractAddress}**
  - (Bearer token required)  

Returns the balanceOf function from a ERC721 contract for the connected wallet.
