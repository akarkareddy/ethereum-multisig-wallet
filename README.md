# Ethereum Multisignature Wallet

This project is a basic Ethereum Multisignature Wallet built using **Solidity** for the smart contract and **Go (Golang)** for interacting with the contract on the **Sepolia testnet**. It uses **Remix IDE** for deploying the contract, a **custom wallet implementation** for managing keys and signing transactions, and **Infura** for connecting to the Ethereum network. The Go program uses the **go-ethereum** library to submit and interact with transactions programmatically.

## Features

- **Multisignature Wallet**: Allows multiple owners to approve transactions before execution, ensuring secure approval from multiple parties.
- **Ethereum Blockchain Integration**: Interacts with the Sepolia Ethereum testnet using MetaMask for wallet management.
- **Smart Contract**: Written in Solidity and deployed via Remix IDE, handling wallet creation, transaction submission, and execution.
- **Custom Wallet**: Built-in wallet implementation handles private key management and signing.
- **Programmatic Interaction**: The Go program interacts with the deployed contract, submitting transactions, confirming them, and executing them based on the required number of confirmations.

## Project Structure
/ethereum-multisig-wallet
├── /api                 # REST API implementation for wallet and transaction management
│   ├── handlers.go
│   └── routes.go
├── /blockchain          # Blockchain interaction logic (Go-Ethereum integration)
│   ├── blockchain.go
│   ├── multisig.go
├── /contracts            # Solidity smart contract code and ABI files
│   ├── Multisigwallet.sol
│   └── multisigwallet.abi
|   └── multisig.go
|   └── multisig.bin
├── /interact            # Go code for interacting with the smart contract on Ethereum
│   ├── contract_interaction.go
│   ├── contract_bytecode.txt
│   ├── contract_abi.json
│   └── 
├── /wallet              # Main application entry point to interact with the Ethereum blockchain
│   ├── wallet.go
├── main.go              # Main application entry point to interact with the Ethereum blockchain
├── go.MOD
├── go.sum


Folder Breakdown:
api/: Contains REST API handlers (handlers.go, routes.go) for wallet and transaction management.

blockchain/: Implements Go-Ethereum logic to interact with the Sepolia testnet and the multisig contract.

contracts/: Contains the Solidity smart contract (multisig.sol) for the multi-signature wallet and ABI files (multisig.abi).

interact/: Provides Go code for interacting with the deployed smart contract and handling blockchain interactions.

wallet/: Manages wallet creation and private key management (note: this excludes wallet_key.go).

main.go: The entry point of the Go application that sets up the connection to Ethereum and interacts with the multisig contract.


Requirements
Programming Language: Go (Golang)

Tools:

Docker: For containerizing the application and deployment.

Remix IDE: For deploying the Solidity smart contract.

Infura: For connecting to the Ethereum network (Sepolia testnet).

Go-Ethereum (go-ethereum): For interacting with Ethereum from Go.

Getting Started
1. Install Required Tools
Go: Download and install Go from the official Go website.
Remix IDE: Access the Remix IDE from here.
Infura: Create an account on Infura and obtain your API key for the Sepolia testnet.

2. Clone the Repository
git clone https://github.com/akarkareddy/ethereum-multisig-wallet.git
cd ethereum-multisig-wallet
3. Initialize Go Module
go mod init github.com/akarkareddy/ethereum-multi-sig-wallet
4. Deploy the Smart Contract
Using Remix IDE, deploy the multisig.sol contract to the Sepolia testnet.

Copy the deployed contract's contract address.

Save the ABI of the contract (you will use this in the Go code).

5. Update Configuration
In the main.go file, update the following:

RPC URL: Use the Infura API endpoint for Sepolia.

Private Key: Provide the private key of the wallet you will use to interact with the contract.

Contract Address: Set the deployed contract address in main.go.

go
const (
    rpcURL        = "https://sepolia.infura.io/v3/your-infura-api-key" // Infura RPC URL for Sepolia
    privateKeyHex = "your-wallet-private-key"                           // Private key of the sender
    contractAddr  = "deployed-contract-address"                          // Deployed multisig wallet contract address
)
6. Run the Application
To start the Go application and interact with the Ethereum network:

bash
go run main.go

7. Test the Wallet and Transaction APIs
You can test the wallet and transaction functionality using Postman or curl by sending requests to http://localhost:8080/wallets for wallet creation, and http://localhost:8080/transactions for submitting and confirming transactions.

Smart Contract Details
Contract Overview
The smart contract is written in Solidity and contains the following key features:

Wallet Creation: Allows the creation of multisig wallets with multiple owners.

Transaction Submission: A transaction can be submitted by any wallet owner, requiring confirmation from other owners before execution.

Transaction Execution: Once a transaction has been confirmed by the required number of owners, it is executed on the Ethereum blockchain.

Contract Files
contracts/multisig.sol: Solidity code for the multisig wallet contract.

contracts/multisig.abi: ABI for interacting with the contract from the Go code.

interact/contract_interaction.go: Go code to interact with the deployed smart contract and submit transactions.


## 🛣️ Roadmap

**Step 1**: Set up the project structure and initialize the Go environment.

**Step 2**: Develop the multisig smart contract using Solidity and deploy it to the Sepolia testnet using Remix IDE.

**Step 3**: Build the Go program to interact with the smart contract, including wallet creation, transaction submission, and confirmation.

**Step 4**: Implement secure key generation and management using Go.

**Step 5**: Secure private key handling and harden API endpoints for safe transaction and wallet operations.

**Step 6**: (Optional Future Step) Containerize the application using Docker and set up basic monitoring and logging.

**Step 7**: Extend functionality to support other Ethereum testnets/mainnet, with features like gas optimization and enhanced UI integration.

Contributing
Contributions are welcome! Please follow these steps:

Fork the repository.

Create a new branch (git checkout -b feature-branch-name).

Commit your changes (git commit -m 'Add feature').

Push to the branch (git push origin feature-branch-name).

Open a pull request.

License
This project is licensed under the MIT License. See the LICENSE file for details.

## 🙏 Acknowledgments :
- **Unit 410** – For inspiring the architecture and design approach of this project. Your work in blockchain infrastructure and secure systems engineering served as a key influence.
- **Ethereum Sepolia Testnet** – For providing a reliable testing environment to deploy and interact with smart contracts without spending real ETH.
- **Infura** – For offering seamless and dependable Ethereum node access, which made network integration and contract interaction possible.
- 
Big thanks also to:
- The open-source community for building incredible libraries and documentation.
- The Go-Ethereum team for maintaining `go-ethereum`, making Ethereum development accessible via Go.
- Remix IDE team for enabling easy Solidity development and testing.
