package handler

type StandardResponse struct {
	Message string `json:"message"`
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
