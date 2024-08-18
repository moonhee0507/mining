package service

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"mining/types"

	"github.com/hacpy/go-ethereum/common/hexutil"
	"github.com/hacpy/go-ethereum/crypto"
)

func (s *Service) newWallet() (string,string, error) {
	p256 := elliptic.P256()
	// y^2 = x^3 + ax + b
	if private, err := ecdsa.GenerateKey(p256, rand.Reader); err != nil {
		return "", "", err
	} else if private == nil {
		return "", "", errors.New("Pk is Nil")
	} else {
		privateKeyToBytes := crypto.FromECDSA(private)
		fmt.Println(hexutil.Encode(privateKeyToBytes))
		fmt.Println(privateKeyToBytes)
	}

	return "", "", nil
}

func (s *Service) MakeWallet() *types.Wallet {
	fmt.Println("들어옴")
	var wallet types.Wallet
	var err error

	wallet.PrivateKey, wallet.PublicKey, err = s.newWallet()
	if  err != nil {
		panic(err)
	}

	// TODO: connect repository
	// s.repository

	return &wallet
}