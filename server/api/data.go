package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type getStockDataRequest struct {
	Ticker    string    `form:"ticker" binding:"required"`
	StartAt   time.Time `form:"start_at" binding:"required"`
	Frequency string    `form:"frequency" binding:"required,frequency"`
}

type stockDatapoint struct {
	Price float64   `json:"price"`
	Date  time.Time `json:"date"`
}

type getStockDataResponse struct {
	Ticker string           `json:"ticker"`
	Stocks []stockDatapoint `json:"data"`
}

func (server *Server) getStockData(ctx *gin.Context) {
	var req getStockDataRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		errResponse := errors.New("invalid input for getting stock data")
		ctx.JSON(http.StatusBadRequest, errorResponse(errResponse))
		return
	}

	data, err := server.getHistoricalStockData(req.Ticker, req.StartAt, req.Frequency)
	if err != nil {
		errResponse := errors.New("unable to get stocks for user")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	var result getStockDataResponse
	result.Ticker = req.Ticker
	result.Stocks = []stockDatapoint{}

	for _, datapoint := range data {
		result.Stocks = append(result.Stocks, stockDatapoint{
			Price: datapoint.AdjClose,
			Date:  datapoint.Date,
		})
	}

	ctx.JSON(http.StatusOK, result)
}

type getStockMetadataRequest struct {
	Ticker    string    `form:"ticker" binding:"required"`
}

func (server *Server) getStockMetadata(ctx *gin.Context) {
	var req getStockMetadataRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		errResponse := errors.New("invalid input for getting ticker metadata")
		ctx.JSON(http.StatusBadRequest, errorResponse(errResponse))
		return
	}

	metadata, err := server.getTickerMetadata(req.Ticker)
	if err != nil {
		errResponse := errors.New("unable to get ticker metadata")
		ctx.JSON(http.StatusInternalServerError, errorResponse(errResponse))
		return
	}

	ctx.JSON(http.StatusOK, metadata)
}
