package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/akarkareddy/ethereum-multisig-wallet/contracts" // Auto-generated Go bindings from the smart contract ABI
	"github.com/ethereum/go-ethereum/accounts/abi/bind"         // For smart contract interaction
	"github.com/ethereum/go-ethereum/common"                    // For Ethereum address conversion
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"    // For key and address generation
	"github.com/ethereum/go-ethereum/ethclient" // Ethereum client for RPC communication
)

const RPC_URL = "https://sepolia.infura.io/v3/0fb5b52beb4d454c965b32ef7dd465a9" // with Infura Id

// DeployMultisigRequest for deploying multisig wallet
type DeployMultisigRequest struct {
	Owners                []string `json:"owners"`                // List of wallet owners
	RequiredConfirmations uint8    `json:"requiredConfirmations"` // Minimum confirmations required to execute a transaction
}

// SubmitMultisigTxRequest placeholder structucture for submitting the transaction
type SubmitMultisigTxRequest struct {
	// Add fields like contractAddress, to, value, data, senderPrivateKey .
	ContractAddress string `json:"contractAddress"`
	To              string `json:"to"`
	Value           string `json:"value"`
	SenderKey       string `json:"senderPrivateKey"`
}

// GetBalance returns the ETH balance of an address
func GetBalance(address string) (string, error) {
	client, err := ethclient.Dial(RPC_URL) //establishes a connection to an Ethereum node
	if err != nil {
		return "", err
	}
	defer client.Close()

	account := common.HexToAddress(address)                              //Converts the hexadecimal string representation of the Ethereum address
	balance, err := client.BalanceAt(context.Background(), account, nil) //fetches the balance of the specified account
	if err != nil {
		return "", err
	}

	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	return ethValue.String(), nil
}

// TransferRequest struct for transfer
type TransferRequest struct {
	PrivateKey string `json:"privateKey"`
	ToAddress  string `json:"toAddress"`
	Amount     string `json:"amount"` // in ETH
}

// SendTransaction sends ETH from one wallet to another
func SendTransaction(req TransferRequest) (string, error) {
	client, err := ethclient.Dial(RPC_URL) //establishes a connection to an Ethereum node
	if err != nil {
		return "", err
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(req.PrivateKey, "0x")) //ecdsa key to hexa and removes the 0x prefix coomon in key formt
	if err != nil {
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)            //sender address is derived from the public key associated with the private key.
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) //Each transaction must have a unique nonce to prevent replay attacks
	if err != nil {
		return "", err
	}

	value, _ := new(big.Float).SetString(req.Amount) // Rrquired amount is convert into big.float=10^18  to wei
	valueWei := new(big.Int)
	value.Mul(value, big.NewFloat(1e18)).Int(valueWei)

	gasLimit := uint64(21000)                                   //standard gas limit for simple ETH transfers
	gasPrice, _ := client.SuggestGasPrice(context.Background()) // can retrive current gas prize for the eth blockchain network.

	toAddress := common.HexToAddress(req.ToAddress)
	tx := types.NewTransaction(nonce, toAddress, valueWei, gasLimit, gasPrice, nil) //A new transaction is created with noce ,toaddress,etc..

	chainID, _ := client.NetworkID(context.Background())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey) //transaction is signed with sender's private key and the EIP-155 signer
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx) //signed transaction is broadcast to the Ethereum network.
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

func SubmitMultisigTransaction(req SubmitMultisigTxRequest) (string, error) {
	client, err := ethclient.Dial(RPC_URL) //establishes a connection to an Ethereum node
	if err != nil {
		return "", err
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(req.SenderKey, "0x")) //private key to ECDSA
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) //Creates (auth) object for signing and submitting transactions.
	if err != nil {                                                                   //chain ID ensures the transaction is valid
		return "", err
	}

	contractAddress := common.HexToAddress(req.ContractAddress)      //Converts the contract address from a string to an Ethereum address
	instance, err := contracts.NewContracts(contractAddress, client) // contract binding for interacting with the deployed multisig contract.
	if err != nil {
		return "", err
	}

	to := common.HexToAddress(req.To) // value converted to Eth adess(common)

	ethFloat, ok := new(big.Float).SetString(req.Value) //transaction of ethereum to float
	if !ok {
		return "", fmt.Errorf("invalid value: %s", req.Value)
	}
	weiFloat := new(big.Float).Mul(ethFloat, big.NewFloat(1e18)) //converting Wei to big.converting to int float
	value := new(big.Int)                                        //
	weiFloat.Int(value)

	tx, err := instance.SubmitTransaction(auth, to, value, []byte{}) //submitting the transaction with autha and chain iD ,to and value.
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil //Tx hash convetted to Hex.
}
