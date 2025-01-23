package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Egorpalan/wallet/internal/entity"
	"github.com/Egorpalan/wallet/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CreateWallet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO wallets \\(id, balance\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs("wallet1", int64(0)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewRepository(db)

	service := &Service{repo: repo}

	wallet := &entity.Wallet{ID: "wallet1", Balance: 0}
	err = service.CreateWallet(wallet)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_GetBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1").
		WithArgs("wallet1").
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(1000))

	repo := repository.NewRepository(db)

	service := &Service{repo: repo}

	balance, err := service.GetBalance("wallet1")

	assert.NoError(t, err)
	assert.Equal(t, int64(1000), balance)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_Deposit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE wallets SET balance = balance \\+ \\$1 WHERE id = \\$2").
		WithArgs(int64(1000), "wallet1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewRepository(db)

	service := &Service{repo: repo}

	err = service.Deposit("wallet1", 1000)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_Withdraw(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1").
		WithArgs("wallet1").
		WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(2000)) // Баланс 2000

	mock.ExpectExec("UPDATE wallets SET balance = balance - \\$1 WHERE id = \\$2").
		WithArgs(int64(500), "wallet1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewRepository(db)

	service := &Service{repo: repo}

	err = service.Withdraw("wallet1", 500)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestService_CreateTransaction(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO transactions \\(id, wallet_id, operation_type, amount\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)").
		WithArgs("tx1", "wallet1", "DEPOSIT", int64(1000)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewRepository(db)

	service := &Service{repo: repo}

	transaction := &entity.Transaction{
		ID:            "tx1",
		WalletID:      "wallet1",
		OperationType: "DEPOSIT",
		Amount:        1000,
	}
	err = service.CreateTransaction(transaction)

	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
