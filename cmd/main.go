package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"leonardovee.com/rinha-de-backend-2024-q1/internal/api"
	"leonardovee.com/rinha-de-backend-2024-q1/internal/transactions"
)

func main() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	e := echo.New()
	api.Setup(e, &api.Handlers{Transactions: transactions.Setup(pool)})
	e.Logger.Fatal(e.Start(":80"))
}
