package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes registers all API routes
func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // welcome message to make sure its working
		fmt.Fprintln(w, "Welcome to the Ethereum Multisig Wallet Service!")
	}).Methods("GET")

	router.HandleFunc("/wallet/create", CreateWalletHandler).Methods("POST")         // POST endpoint to generate a new Ethereum wallet (private/public key + address)
	router.HandleFunc("/wallet/balance/{address}", GetBalanceHandler).Methods("GET") // GET endpoint to retrieve ETH balance
	router.HandleFunc("/wallet/transfer", TransferHandler).Methods("POST")           // POST endpoint to transfer
	router.HandleFunc("/wallet/multisig/deploy", DeployMultisigHandler).Methods("POST")
	router.HandleFunc("/wallet/multisig/submit", SubmitMultisigTxHandler).Methods("POST") // POST endpoint to submit a transaction
}
