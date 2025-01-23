package repository

import (
	"github.com/Egorpalan/wallet/internal/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateWallet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO wallets \\(id, balance\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs("wallet1", int64(0)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &Repository{db: db}

	wallet := &entity.Wallet{ID: "wallet1", Balance: 0}
	err = repo.CreateWallet(wallet)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1").
		WithArgs("wallet1").
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(1000))

	repo := &Repository{db: db}

	balance, err := repo.GetBalance("wallet1")

	assert.NoError(t, err)
	assert.Equal(t, int64(1000), balance)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Deposit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE wallets SET balance = balance \\+ \\$1 WHERE id = \\$2").
		WithArgs(int64(1000), "wallet1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &Repository{db: db}

	err = repo.Deposit("wallet1", 1000)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Withdraw(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE wallets SET balance = balance - \\$1 WHERE id = \\$2").
		WithArgs(int64(500), "wallet1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &Repository{db: db}

	err = repo.Withdraw("wallet1", 500)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_CreateTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO transactions \\(id, wallet_id, operation_type, amount\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)").
		WithArgs("tx1", "wallet1", "DEPOSIT", int64(1000)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &Repository{db: db}

	transaction := &entity.Transaction{
		ID:            "tx1",
		WalletID:      "wallet1",
		OperationType: "DEPOSIT",
		Amount:        1000,
	}
	err = repo.CreateTransaction(transaction)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
