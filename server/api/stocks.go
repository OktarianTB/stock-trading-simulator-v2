package api

import (
	"errors"
	"net/http"
	"time"

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

	for _, s := range stocks {
		price, err := server.getStockPriceForTicker(s.Ticker)
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
	}

	result.PortfolioBalance = util.RoundFloat(result.PortfolioBalance)

	ctx.JSON(http.StatusOK, result)
}

type tiingoStock struct {
	AdjOpen   float64   `json:"adjOpen"`
	AdjHigh   float64   `json:"adjHigh"`
	AdjLow    float64   `json:"adjLow"`
	AdjClose  float64   `json:"adjClose"`
	AdjVolume int64     `json:"adjVolume"`
	Date      time.Time `json:"date"`
}

func (server *Server) getStockPriceForTicker(ticker string) (float64, error) {
	url := server.config.TiingoAPI + "/daily/" + ticker + "/prices"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	q := req.URL.Query()
	q.Add("token", server.config.TiingoToken)
	req.URL.RawQuery = q.Encode()

	var stocks []tiingoStock
	err = util.MakeGetRequest(req.URL.String(), &stocks)
	if err != nil {
		return 0, err
	}

	return stocks[0].AdjClose, nil
}
