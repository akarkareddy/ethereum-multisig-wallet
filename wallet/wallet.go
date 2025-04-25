package wallet

import (
	"crypto/ecdsa" // Elliptic Curve Digital Signature Algorithm (used in Ethereum)
	"encoding/hex" // For converting byte slices to hex strings
	"fmt"

	"github.com/ethereum/go-ethereum/crypto" // Ethereum's cryptographic utility package
)

// WalletKey represents the structure of wallet with public and private keys of a wallet
type WalletKey struct {
	Address    string `json:"address"`
	PublicKey  string `json:"publicKey"` // Hex-encoded public key
	PrivateKey string `json:"privateKey"`
}

// GenerateKey creates a new Ethereum wallet key pair
func GenerateKey() (*WalletKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)      // Convert private key to bytes
	privateKeyHex := hex.EncodeToString(privateKeyBytes) // Encode bytes as hex string

	publicKey := privateKey.Public().(*ecdsa.PublicKey) // Extract public key from private key
	publicKeyBytes := crypto.FromECDSAPub(publicKey)    // Convert public key to bytes
	publicKeyHex := hex.EncodeToString(publicKeyBytes)  // Encode public key as hex string

	address := crypto.PubkeyToAddress(*publicKey).Hex()

	// Return the WalletKey struct with address, public key, and private key

	return &WalletKey{
		Address:    address,
		PublicKey:  publicKeyHex,
		PrivateKey: privateKeyHex,
	}, nil
}
