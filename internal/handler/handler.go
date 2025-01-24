package handler

import (
	"github.com/Egorpalan/wallet/internal/entity"
	"github.com/Egorpalan/wallet/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// HandleWalletOperation handles wallet deposit and withdrawal operations
// @Summary Perform wallet operation
// @Description Deposit or withdraw funds from a wallet
// @Tags wallet
// @Accept json
// @Produce json
// @Param request body entity.TransactionRequest true "Transaction data"
// @Success 200 {object} handler.StandardResponse "Operation successful"
// @Failure 400 {object} handler.ErrorResponse "Invalid request payload"
// @Failure 500 {object} handler.ErrorResponse "Failed to process transaction"
// @Router /api/v1/wallet [post]
func (h *Handler) HandleWalletOperation(c *gin.Context) {
	var req entity.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	transaction := &entity.Transaction{
		ID:            generateUUID(),
		WalletID:      req.WalletID,
		OperationType: req.OperationType,
		Amount:        req.Amount,
	}

	switch req.OperationType {
	case "DEPOSIT":
		if err := h.service.Deposit(req.WalletID, req.Amount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	case "WITHDRAW":
		if err := h.service.Withdraw(req.WalletID, req.Amount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation type"})
		return
	}

	if err := h.service.CreateTransaction(transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Operation successful"})
}

// GetBalance retrieves wallet balance
// @Summary Get wallet balance
// @Description Get the current balance of a wallet by walletId
// @Tags wallet
// @Produce json
// @Param walletId path string true "Wallet ID"
// @Success 200 {object} handler.BalanceResponse "Wallet balance"
// @Failure 500 {object} handler.ErrorResponse "Failed to get balance"
// @Router /api/v1/wallets/{walletId} [get]
func (h *Handler) GetBalance(c *gin.Context) {
	walletID := c.Param("walletId")

	balance, err := h.service.GetBalance(walletID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// CreateWallet creates a new wallet
// @Summary Create a wallet
// @Description Create a new wallet with initial balance
// @Tags wallet
// @Accept json
// @Produce json
// @Param request body entity.Wallet true "Wallet ID"
// @Success 200 {object} handler.StandardResponse "Wallet created successfully"
// @Failure 400 {object} handler.ErrorResponse "Invalid request payload"
// @Failure 500 {object} handler.ErrorResponse "Failed to create wallet"
// @Router /api/v1/wallets [post]
func (h *Handler) CreateWallet(c *gin.Context) {
	var req struct {
		WalletID string `json:"walletId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	wallet := &entity.Wallet{
		ID:      req.WalletID,
		Balance: 0,
	}

	if err := h.service.CreateWallet(wallet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wallet created successfully"})
}

func generateUUID() string {
	return uuid.New().String()
}
