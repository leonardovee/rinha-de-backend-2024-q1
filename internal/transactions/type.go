package transactions

import "time"

type CreateRequest struct {
	Amount      int    `json:"valor" validate:"required"`
	Type        string `json:"tipo" validate:"required,oneof=c d"`
	Description string `json:"descricao" validate:"required,max=10"`
}

type CreateResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type Balance struct {
	Total     int       `json:"total"`
	CreatedAt time.Time `json:"data_extrato"`
	Limit     int       `json:"limite"`
}

type GetBalanceResponse struct {
	Balance      Balance             `json:"saldo"`
	Transactions []TransactionSchema `json:"ultimas_transacoes"`
}

type TransactionSchema struct {
	ID          int       `json:"id,omitempty"`
	Amount      int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

type ClientSchema struct {
	ID           int
	Name         string
	Balance      int
	BalanceLimit int
	CreatedAt    time.Time
}

type ProcedureResponse struct {
	Balance      *int `json:"saldo"`
	BalanceLimit *int `json:"limite"`
}
