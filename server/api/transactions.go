package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/OktarianTB/stock-trading-simulator-golang/db/sqlc"
	"github.com/OktarianTB/stock-trading-simulator-golang/token"
	"github.com/gin-gonic/gin"
)

type listTransactionsRequest struct {
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=1,max=20"`
	Ticker   string `form:"ticker"`
}

func (server *Server) listTransactions(ctx *gin.Context) {
	var req listTransactionsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if req.Ticker != "" {
		arg := db.ListTransactionsForUserForTickerParams{
			Username: authPayload.Username,
			Limit:    req.PageSize,
			Offset:   (req.PageID - 1) * req.PageSize,
			Ticker:   req.Ticker,
		}

		transactions, err := server.store.ListTransactionsForUserForTicker(ctx, arg)
		if err != nil {
			errResponse := errors.New("unable to list transactions for user for ticker")
			ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
			return
		}

		ctx.JSON(http.StatusOK, transactions)
	} else {
		arg := db.ListTransactionsForUserParams{
			Username: authPayload.Username,
			Limit:    req.PageSize,
			Offset:   (req.PageID - 1) * req.PageSize,
		}

		transactions, err := server.store.ListTransactionsForUser(ctx, arg)
		if err != nil {
			errResponse := errors.New("unable to list transactions for user")
			ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
			return
		}

		ctx.JSON(http.StatusOK, transactions)
	}
}

type purchaseStockRequest struct {
	Ticker   string `json:"ticker" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required,min=1"`
}

type purchaseStockResponse struct {
	UserBalance float64   `json:"user_balance"`
	Ticker      string    `json:"ticker"`
	Quantity    int64     `json:"quantity"`
	Price       float64   `json:"price"`
	Total       float64   `json:"total"`
	PurchasedAt time.Time `json:"purchased_at"`
}

func (server *Server) purchaseStock(ctx *gin.Context) {
	var req purchaseStockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errResponse := errors.New("invalid input for purchasing stock")
		ctx.JSON(http.StatusBadRequest, errorResponse(errResponse))
		return
	}

	tickerPrice, err := server.getStockPriceForTicker(req.Ticker)
	if err != nil {
		errResponse := errors.New("unable to get stock price for ticker")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	result, err := server.store.PurchaseTx(ctx, db.CreateTransactionParams{
		Username: authPayload.Username,
		Ticker:   req.Ticker,
		Quantity: req.Quantity,
		Price:    tickerPrice,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp := purchaseStockResponse{
		UserBalance: result.User.Balance,
		Ticker:      result.Transaction.Ticker,
		Quantity:    result.Transaction.Quantity,
		Price:       result.Transaction.Price,
		Total:       float64(result.Transaction.Quantity) * result.Transaction.Price,
		PurchasedAt: result.Transaction.CreatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}

type sellStockRequest struct {
	Ticker   string `json:"ticker" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required,min=1"`
}

type sellStockResponse struct {
	UserBalance float64   `json:"user_balance"`
	Ticker      string    `json:"ticker"`
	Quantity    int64     `json:"quantity"`
	Price       float64   `json:"price"`
	Total       float64   `json:"total"`
	SoldAt      time.Time `json:"sold_at"`
}

func (server *Server) sellStock(ctx *gin.Context) {
	var req sellStockRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errResponse := errors.New("invalid input for selling stock")
		ctx.JSON(http.StatusBadRequest, errorResponse(errResponse))
		return
	}

	tickerPrice, err := server.getStockPriceForTicker(req.Ticker)
	if err != nil {
		errResponse := errors.New("unable to get stock price for ticker")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	result, err := server.store.SellTx(ctx, db.CreateTransactionParams{
		Username: authPayload.Username,
		Ticker:   req.Ticker,
		Quantity: req.Quantity,
		Price:    tickerPrice,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	resp := sellStockResponse{
		UserBalance: result.User.Balance,
		Ticker:      result.Transaction.Ticker,
		Quantity:    result.Transaction.Quantity,
		Price:       result.Transaction.Price,
		Total:       float64(result.Transaction.Quantity) * result.Transaction.Price,
		SoldAt:      result.Transaction.CreatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}
