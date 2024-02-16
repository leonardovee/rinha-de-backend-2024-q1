package api

import (
	"github.com/labstack/echo/v4"
	"leonardovee.com/rinha-de-backend-2024-q1/internal/transactions"
)

type Handlers struct {
	Transactions *transactions.Handler
}

func Setup(e *echo.Echo, h *Handlers) {
	e.POST("/clientes/:id/transacoes", h.Transactions.CreateTransaction)
	e.GET("/clientes/:id/extrato", h.Transactions.GetBalance)
}
