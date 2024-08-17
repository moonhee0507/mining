package service

import (
	"fmt"
	"mining/types"
)

func (s *Service) newWallet() (string,string, error) {
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