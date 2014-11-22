package client

import (
	"fmt"
	"net/http"
)

// https://test.bitpay.com/api#resource-Rates

type (
	// Rate maps to a resource at the rates
	Rate struct {
		Code string  `json:"code"`
		Name string  `json:"name"`
		Rate float64 `json:"rate"`
	}
)

// QueryRates returns a list of exchange rates.
func (c *Client) QueryRates() ([]Rate, *http.Response, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/rates", c.apiBase), nil)
	if err != nil {
		return nil, nil, err
	}

	var rates []Rate
	resp, err := c.Send(req, &rates)

	return rates, resp, err
}

// GetRateForCurrency returns the exchange rate for a given currency.
func (c *Client) GetRateForCurrency(currencyCode string) (*Rate, *http.Response, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/rates/%s", c.apiBase, currencyCode), nil)
	if err != nil {
		return nil, nil, err
	}

	var rate Rate
	resp, err := c.Send(req, &rate)

	return &rate, resp, err
}
