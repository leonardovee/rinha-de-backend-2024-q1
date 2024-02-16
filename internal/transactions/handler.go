package transactions

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Repository *Repository
}

func NewHandler(repository *Repository) *Handler {
	return &Handler{Repository: repository}
}

func (h *Handler) CreateTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	cpr := new(CreateRequest)
	if err := c.Bind(cpr); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	validate := validator.New()
	err := validate.Struct(cpr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	clientID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if clientID > 5 {
		return c.JSON(http.StatusNotFound, nil)
	}

	ctx = context.WithValue(ctx, "clientID", clientID)

	pr, err := h.Repository.CreateTransaction(ctx, *cpr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	if pr.Balance == nil && pr.BalanceLimit == nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	return c.JSON(http.StatusOK, pr)
}

func (h *Handler) GetBalance(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	clientID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if clientID > 5 {
		return c.JSON(http.StatusNotFound, nil)
	}

	ctx = context.WithValue(ctx, "clientID", clientID)

	t, err := h.Repository.GetTransactions(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	b, err := h.Repository.GetBalance(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res := GetBalanceResponse{
		Balance: Balance{
			Total:     b.Balance,
			CreatedAt: time.Now().UTC(),
			Limit:     b.BalanceLimit,
		},
		Transactions: t,
	}

	return c.JSON(http.StatusOK, res)
}
