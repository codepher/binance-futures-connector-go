package binance_futures_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Binance Kline/Candlestick Data endpoint (GET /fapi/v1/premiumIndex)
type PremiumIndex struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *PremiumIndex) Symbol(symbol string) *PremiumIndex {
	s.symbol = &symbol
	return s
}

func (s *PremiumIndex) Do(ctx context.Context, opts ...RequestOption) (res []*PremiumIndexResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/premiumIndex",
		secType:  secTypeNone,
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var PremiumIndexResponseArray PremiumIndexResponseArray
	err = json.Unmarshal(data, &PremiumIndexResponseArray)
	if err != nil {
		return nil, err
	}
	var PremiumIndex []*PremiumIndexResponse
	for _, premium := range PremiumIndexResponseArray {
		Symbol := premium[0].(string)
		MarkPrice := premium[1].(string)
		IndexPrice := premium[2].(string)
		EstimatedSettlePrice := premium[3].(string)
		LastFundingRate := premium[4].(string)
		NextFundingTime := premium[5].(uint64)
		InterestRate := premium[6].(string)
		Time := premium[7].(uint64)

		// create a PremiumIndexResponse struct using the parsed fields
		PremiumIndexResponse := &PremiumIndexResponse{
			Symbol,
			MarkPrice,
			IndexPrice,
			EstimatedSettlePrice,
			LastFundingRate,
			NextFundingTime,
			InterestRate,
			Time,
		}
		PremiumIndex = append(PremiumIndex, PremiumIndexResponse)
	}
	return PremiumIndex, nil
}

type PremiumIndexResponseArray [][]interface{}

// Define PremiumIndex response data
type PremiumIndexResponse struct {
	Symbol               string `json:"symbol"`
	MarkPrice            string `json:"markPrice"`
	IndexPrice           string `json:"indexPrice"`
	EstimatedSettlePrice string `json:"estimatedSettlePrice"`
	LastFundingRate      string `json:"lastFundingRate"`
	NextFundingTime      uint64 `json:"nextFundingTime"`
	InterestRate         string `json:"interestRate"`
	Time                 uint64 `json:"time"`
}
