package transactions

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(conn *pgxpool.Pool) *Handler {
	repository := NewRepository(conn)
	return NewHandler(repository)
}
