package api

import (
	"encoding/json" // For encoding and decoding JSON requests/responses
	"net/http"

	"github.com/ethereum/go-ethereum/common" // Ethereum address conversion utility

	"github.com/akarkareddy/ethereum-multisig-wallet/blockchain" // Internal blockchain interaction logic
	"github.com/akarkareddy/ethereum-multisig-wallet/wallet"     // Internal wallet generation logic

	"github.com/gorilla/mux"
)

const (
	rpcURL        = "https://sepolia.infura.io/v3/0fb5b52beb4d454c965b32ef7dd465a9"    // Sepolia testnet RPC endpoint (via Infura)
	privateKeyHex = "4b1321d42a47cc5c083a1e6e29e1c4049d43b87de123590d8df55f5fc0ec5ea2" //  wallet private key used for transactions
)

// CreateWalletHandler creates a new Ethereum wallet
func CreateWalletHandler(w http.ResponseWriter, r *http.Request) { // This generates wallets address/ public key and private key.
	key, err := wallet.GenerateKey()
	if err != nil {
		http.Error(w, "Failed to generate wallet", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(key) // Respond with wallet info in JSON
}

// GetBalanceHandler returns the balance of an address
func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"] // Extract address from URL path
	balance, err := blockchain.GetBalance(address)
	if err != nil {
		http.Error(w, "Failed to get balance", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"address": address,
		"balance": balance,
	})
}

// TransferHandler sends ETH from one address to another
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	var txReq blockchain.TransferRequest                           // Struct to hold request body
	if err := json.NewDecoder(r.Body).Decode(&txReq); err != nil { // Decode JSON body into txReq
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	txHash, err := blockchain.SendTransaction(txReq)
	if err != nil {
		http.Error(w, "Transaction failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"txHash": txHash}) // Return tx hash
}

// DeployMultisigHandler deploys a new multisig contract
func DeployMultisigHandler(w http.ResponseWriter, r *http.Request) {
	var req blockchain.DeployMultisigRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var owners []common.Address // Convert string addresses to Ethereum address type
	for _, addr := range req.Owners {
		owners = append(owners, common.HexToAddress(addr))
	}

	contractAddress, err := blockchain.DeployMultisigWallet(rpcURL, privateKeyHex, owners, req.RequiredConfirmations)
	if err != nil {
		http.Error(w, "Deployment failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"contractAddress": contractAddress.Hex()}) // Return contract address
}

// SubmitMultisigTxHandler submits a multisig transaction to the contract

func SubmitMultisigTxHandler(w http.ResponseWriter, r *http.Request) {
	var txReq blockchain.SubmitMultisigTxRequest
	if err := json.NewDecoder(r.Body).Decode(&txReq); err != nil { // Decode request body
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	txHash, err := blockchain.SubmitMultisigTransaction(txReq)
	if err != nil {
		http.Error(w, "Multisig tx failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"txHash": txHash}) // Respond with tx hash
}

/*func SubmitMultisigTxHandler(w http.ResponseWriter, r *http.Request) {
	var txReq blockchain.SubmitMultisigTxRequest
	if err := json.NewDecoder(r.Body).Decode(&txReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Pass the correct arguments to DeployMultisigWallet
	contractAddress, err := blockchain.DeployMultisigWallet(
		"https://sepolia.infura.io/v3/0fb5b52beb4d454c965b32ef7dd465a9",    // rpc URL
		"7791d222a29027e9b88f96f3cb5393061c4de7e964109a59a301e44e1e22fd18", // private key
		req.Owners, // owners
		req.RequiredConfirmations,
	)
	//txHash, err := blockchain.SubmitMultisigTransaction(txReq)
	if err != nil {
		http.Error(w, "Multisig tx failed: "+err.Error(), http.StatusInternalServerError)
		return
	}
	//json.NewEncoder(w).Encode(map[string]string{"txHash": txHash})
	json.NewEncoder(w).Encode(map[string]string{"contractAddress": contractAddress.Hex()})
}
*/
