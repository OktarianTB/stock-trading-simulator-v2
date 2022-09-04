package api

import (
	"fmt"
	"net/http"
	"time"

	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
)

type tiingoStock struct {
	AdjClose float64   `json:"adjClose"`
	Date     time.Time `json:"date"`
}

func (server *Server) getStockPrice(ticker string) (float64, error) {
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

func (server *Server) getHistoricalStockData(ticker string, startAt time.Time, frequency string) ([]tiingoStock, error) {
	url := server.config.TiingoAPI + "/daily/" + ticker + "/prices"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("startDate", fmt.Sprintf("%v-%v-%v", startAt.Year(), int(startAt.Month()), startAt.Day()))
	q.Add("resampleFreq", frequency)
	q.Add("token", server.config.TiingoToken)
	req.URL.RawQuery = q.Encode()
	
	var stocks []tiingoStock
	err = util.MakeGetRequest(req.URL.String(), &stocks)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

type tiingoMetadata struct {
	Ticker       string `json:"ticker"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	ExchangeCode string `json:"exchangeCode"`
}

func (server *Server) getTickerMetadata(ticker string) (tiingoMetadata, error) {
	url := server.config.TiingoAPI + "/daily/" + ticker
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return tiingoMetadata{}, err
	}

	q := req.URL.Query()
	q.Add("token", server.config.TiingoToken)
	req.URL.RawQuery = q.Encode()

	var metadata tiingoMetadata
	err = util.MakeGetRequest(req.URL.String(), &metadata)
	if err != nil {
		return tiingoMetadata{}, err
	}

	return metadata, nil
}
