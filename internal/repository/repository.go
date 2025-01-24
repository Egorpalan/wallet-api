package repository

import (
	"database/sql"
	"github.com/Egorpalan/wallet/internal/entity"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateWallet(wallet *entity.Wallet) error {
	query := `INSERT INTO wallets (id, balance) VALUES ($1, $2)`
	_, err := r.db.Exec(query, wallet.ID, wallet.Balance)
	return err
}

func (r *Repository) GetBalance(walletID string) (int64, error) {
	var balance int64
	query := `SELECT balance FROM wallets WHERE id = $1`
	err := r.db.QueryRow(query, walletID).Scan(&balance)
	return balance, err
}

func (r *Repository) Deposit(walletID string, amount int64) error {
	query := `UPDATE wallets SET balance = balance + $1 WHERE id = $2`
	_, err := r.db.Exec(query, amount, walletID)
	return err
}

func (r *Repository) Withdraw(walletID string, amount int64) error {
	query := `UPDATE wallets SET balance = balance - $1 WHERE id = $2`
	_, err := r.db.Exec(query, amount, walletID)
	return err
}

func (r *Repository) CreateTransaction(transaction *entity.Transaction) error {
	query := `INSERT INTO transactions (id, wallet_id, operation_type, amount) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, transaction.ID, transaction.WalletID, transaction.OperationType, transaction.Amount)
	if err != nil {
		logrus.Errorf("Failed to create transaction: %s", err.Error())
	}
	return err
}
