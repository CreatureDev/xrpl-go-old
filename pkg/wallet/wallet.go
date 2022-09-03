package wallet

import (
	"crypto"
	"crypto/ed25519"
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
	pubkey, ok := key.(ed25519.PublicKey)
	if !ok || len(pubkey) != 32 {
		return "", fmt.Errorf("Unable to encode publickey %v", key)
	}
	prefixkey := append([]byte{0xED}, pubkey...)
	if len(prefixkey) != 33 {
		return "", fmt.Errorf("pubkey incorrect size")
	}
	sha := crypto.SHA256.New()
	_, _ = sha.Write(prefixkey)
	hash := sha.Sum(nil)
	ripe := ripemd160.New()
	_, _ = ripe.Write(hash[:])
	acc := ripe.Sum(nil)

	account_base := append([]byte{0x00}, acc...)

	sha = crypto.SHA256.New()
	_, _ = sha.Write(account_base)
	hash = sha.Sum(nil)
	sha = crypto.SHA256.New()
	_, _ = sha.Write(hash[:])
	hash = sha.Sum(nil)
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
