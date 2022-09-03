package wallet

import (
	"crypto"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/pkg/api"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PublicKey  crypto.PublicKey
	PrivateKey crypto.PrivateKey
	Account    api.Account
}

func EncodeAccount(key crypto.PublicKey) (api.Account, error) {
	pubkey, ok := key.([]byte)
	if !ok || len(pubkey) != 32 {
		return "", fmt.Errorf("Unable to encode publickey %v", key)
	}
	prefixkey := append([]byte{0xED}, pubkey...)
	hash := sha256.Sum256(prefixkey)
	acc := ripemd160.New().Sum(hash[:])

	account_base := append([]byte{0x00}, acc...)

	hash = sha256.Sum256(account_base)
	hash = sha256.Sum256(hash[:])
	checksum := hash[:4]

	account := append(account_base, checksum...)
	encoded, err := api.Encode(account)
	return api.Account(encoded), err
}

func GenerateWallet(seed []byte) (*Wallet, error) {
	var ret Wallet
	var err error
	if len(seed) == 0 {
		ret.PublicKey, ret.PrivateKey, err = ed25519.GenerateKey(nil)
		if err != nil {
			return nil, err
		}
	} else {
		priv := ed25519.NewKeyFromSeed(seed)
		ret.PublicKey = priv.Public()
		ret.PrivateKey = priv
	}

	ret.Account, err = EncodeAccount(ret.PublicKey)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (w *Wallet) Sign(tx api.Tx) ([]byte, error) {
	msg, _ := json.Marshal(tx)
	key, ok := w.PrivateKey.(ed25519.PrivateKey)
	if !ok {
		return []byte{}, fmt.Errorf("Failed to sign: Invalid private key")
	}
	return ed25519.Sign(key, msg), nil
}
