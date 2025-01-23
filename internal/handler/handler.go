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

func (h *Handler) GetBalance(c *gin.Context) {
	walletID := c.Param("walletId")

	balance, err := h.service.GetBalance(walletID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

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
