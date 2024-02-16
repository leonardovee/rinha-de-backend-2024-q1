package transactions

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *Repository {
	return &Repository{Conn: conn}
}

func (r Repository) CreateTransaction(ctx context.Context, ctr CreateRequest) (ProcedureResponse, error) {
	var procedureResponse ProcedureResponse

	amount := ctr.Amount
	if ctr.Type == "d" {
		amount = -ctr.Amount
	}

	err := r.Conn.QueryRow(
		context.Background(),
		"CALL create_transaction($1, $2, $3, $4);",
		ctx.Value("clientID"),
		amount,
		ctr.Type,
		ctr.Description,
	).Scan(&procedureResponse.Balance, &procedureResponse.BalanceLimit)

	return procedureResponse, err
}

func (r Repository) GetTransactions(ctx context.Context) ([]TransactionSchema, error) {
	sql := `
		SELECT amount, type, description, created_at 
		FROM transactions
		WHERE client_id = $1
		ORDER BY id DESC
		LIMIT 10;
	`

	rows, err := r.Conn.Query(context.Background(), sql, ctx.Value("clientID"))

	var ts []TransactionSchema
	for rows.Next() {
		var t TransactionSchema
		err = rows.Scan(&t.Amount, &t.Type, &t.Description, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}

	return ts, err
}

func (r Repository) GetBalance(ctx context.Context) (ClientSchema, error) {
	sql := `
		SELECT balance, balance_limit 
		FROM clients 
		WHERE id = $1;
	`

	var c ClientSchema

	err := r.Conn.QueryRow(context.Background(), sql, ctx.Value("clientID")).Scan(&c.Balance, &c.BalanceLimit)

	return c, err
}
