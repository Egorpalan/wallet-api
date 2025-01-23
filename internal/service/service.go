package service

import (
	"errors"
	"github.com/Egorpalan/wallet/internal/entity"
	"github.com/Egorpalan/wallet/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateWallet(wallet *entity.Wallet) error {
	return s.repo.CreateWallet(wallet)
}

func (s *Service) GetBalance(walletID string) (int64, error) {
	return s.repo.GetBalance(walletID)
}

func (s *Service) Deposit(walletID string, amount int64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	return s.repo.Deposit(walletID, amount)
}

func (s *Service) Withdraw(walletID string, amount int64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	balance, err := s.repo.GetBalance(walletID)
	if err != nil {
		return err
	}

	if balance < amount {
		return errors.New("insufficient funds")
	}

	return s.repo.Withdraw(walletID, amount)
}

func (s *Service) CreateTransaction(transaction *entity.Transaction) error {
	return s.repo.CreateTransaction(transaction)
}
