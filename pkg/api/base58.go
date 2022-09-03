package api

import (
	"fmt"
	"strings"

	"github.com/eknkc/basex"
)

type Base58 string

const (
	dict string = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
)

func (s Base58) Validate() error {
	for _, c := range s {
		if !strings.Contains(dict, string(c)) {
			return fmt.Errorf("Illegal character found in base58 string %s", s)
		}
	}
	return nil
}

func Encode(data []byte) (Base58, error) {
	encoding, err := basex.NewEncoding(dict)
	if err != nil {
		return "", err
	}
	return Base58(encoding.Encode(data)), nil
}
