package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/OktarianTB/stock-trading-simulator-golang/token"
	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	"github.com/gin-gonic/gin"
)

func (server *Server) listUserStocks(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	stocks, err := server.store.ListStockQuantitiesForUser(ctx, authPayload.Username)
	if err != nil {
		errResponse := errors.New("unable to get stocks for user")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	ctx.JSON(http.StatusOK, stocks)
}

type stock struct {
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

	var stocks []stock
	err = util.MakeGetRequest(req.URL.String(), &stocks)
	if err != nil {
		return 0, err
	}

	return stocks[0].AdjClose, nil
}
