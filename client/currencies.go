package client

// https://bitpay.com/api#resource-Currencies

import (
	"fmt"
	"net/http"
)

type (
	// Currency maps to a resource at the currencies endpoint
	Currency struct {
		Code           string   `json:"code"`
		Symbol         string   `json:"symbol"`
		Precision      int      `json:"precision"`
		ExchangePCTFee int64    `json:"exchangePctFee"`
		PayoutEnabled  bool     `json:"payoutEnabled"`
		Name           string   `json:"name"`
		Plural         string   `json:"plural"`
		Alts           string   `json:"alts"`
		PayoutFields   []string `json:"payoutFields"`
	}
)

// QueryCurrencies returns the list of supported currencies.
func (c *Client) QueryCurrencies() ([]Currency, *http.Response, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/currencies", c.apiBase), nil)
	if err != nil {
		return nil, nil, err
	}

	var currencies []Currency
	resp, err := c.Send(req, &currencies)

	return currencies, resp, err
}
