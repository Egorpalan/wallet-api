package entity

import (
	"time"
)

type Wallet struct {
	ID      string `json:"id"`
	Balance int64  `json:"balance"`
}

type Transaction struct {
	ID            string    `json:"id"`
	WalletID      string    `json:"wallet_id"`
	OperationType string    `json:"operation_type"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type TransactionRequest struct {
	WalletID      string `json:"walletId"`
	OperationType string `json:"operationType"`
	Amount        int64  `json:"amount"`
}
