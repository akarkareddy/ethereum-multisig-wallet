package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	rpcURL        = "https://sepolia.infura.io/v3/0fb5b52beb4d454c965b32ef7dd465a9"    // Ethereum node RPC endpoint
	privateKeyHex = "4b1321d42a47cc5c083a1e6e29e1c4049d43b87de123590d8df55f5fc0ec5ea2" // Private key of the sender
	contractAddr  = "0x842e6C440F574bfc82c27Fa43226B1dC6883A344"                       // Already deployed multisig wallet contract address
)

func main() {
	// Connect to the Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("Failed to connect to Ethereum node:", err)
	}
	defer client.Close()

	// Loading senders  private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("Invalid private key:", err)
	}

	// Retriving ethererum addrss from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("Using address:", fromAddress.Hex())

	// Read and parse contract ABI from file
	abiFile, err := os.ReadFile("interact/contract_abi.json")
	if err != nil {
		log.Fatal("Unable to read ABI file: ", err)
	}
	contractABI, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		log.Fatal("Invalid ABI:  ", err)
	}
	//Bind the contract at the given address
	contractAddress := common.HexToAddress(contractAddr)
	instance := bind.NewBoundContract(contractAddress, contractABI, client, client, client)

	// Create authorized transactor (for signing)
	chainID := big.NewInt(11155111) // Sepolia Chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}

	// Example: Call submitTransaction(to, value, data) Prepare transaction parameters
	to := common.HexToAddress("0xA71125e0210d353Ef642eDdA95419e240f7b3C2a")
	value := big.NewInt(100000000000000000) // 0.1 ETH
	data := []byte{}                        // empty calldata for ETH transfer

	tx, err := instance.Transact(auth, "submitTransaction", to, value, data) // submitting the transaction using a contratc.
	if err != nil {
		log.Fatal("submitTransaction failed:", err)
	}
	fmt.Println("Submitted transaction! Tx hash:", tx.Hash().Hex()) //Transaction hash.
}
