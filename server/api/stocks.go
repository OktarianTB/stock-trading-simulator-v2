package api

import (
	"errors"
	"net/http"
	"sync"

	db "github.com/OktarianTB/stock-trading-simulator-golang/db/sqlc"
	"github.com/OktarianTB/stock-trading-simulator-golang/token"
	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	"github.com/gin-gonic/gin"
)

type stock struct {
	Ticker         string  `json:"ticker"`
	Quantity       int64   `json:"quantity"`
	CurrentPrice   float64 `json:"current_price"`
	CurrentBalance float64 `json:"current_balance"`
	PurchaseTotal  float64 `json:"purchase_total"`
}

type listUserStocksResponse struct {
	PortfolioBalance float64 `json:"portfolio_balance"`
	Stocks           []stock `json:"stocks"`
}

func (server *Server) listUserStocks(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	stocks, err := server.store.ListStockQuantitiesForUser(ctx, authPayload.Username)
	if err != nil {
		errResponse := errors.New("unable to get stocks for user")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	var result listUserStocksResponse
	result.Stocks = []stock{}

	wg := sync.WaitGroup{}

	for _, s := range stocks {
		wg.Add(1)
		go func(s db.ListStockQuantitiesForUserRow) {
			if s.Quantity <= 0 {
				wg.Done()
				return
			}

			price, err := server.getStockPrice(s.Ticker)
			if err != nil {
				errResponse := errors.New("unable to get stocks for user")
				ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
				return
			}
			balance := price * float64(s.Quantity)

			purchaseTotal, err := server.store.GetPurchasePriceForTicker(ctx, db.GetPurchasePriceForTickerParams{
				Username: authPayload.Username,
				Ticker:   s.Ticker,
			})
			if err != nil {
				errResponse := errors.New("unable to get stocks for user")
				ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
				return
			}

			result.Stocks = append(result.Stocks, stock{
				Ticker:         s.Ticker,
				Quantity:       s.Quantity,
				CurrentPrice:   price,
				CurrentBalance: util.RoundFloat(balance),
				PurchaseTotal:  util.RoundFloat(purchaseTotal),
			})
			result.PortfolioBalance += balance

			wg.Done()
		}(s)
	}

	wg.Wait()

	result.PortfolioBalance = util.RoundFloat(result.PortfolioBalance)

	ctx.JSON(http.StatusOK, result)
}
