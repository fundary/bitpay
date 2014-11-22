package client

import (
	"fmt"
	"net/http"
	"time"
)

// https://test.bitpay.com/api#resource-Payouts

type (
	// Payout maps to a resource at the payouts endpoint
	Payout struct {
		Instructions      []Instruction `json:"instructions"`
		Amount            int64         `json:"amount"`
		Currency          string        `json:"currency"`
		EffectiveDate     time.Time     `json:"effectiveDate"`
		Reference         string        `json:"reference,omitemtpy"`
		PricingMethod     string        `json:"pricingMethod,omitempty"`
		NotificationEmail string        `json:"notificationEmail,omitempty"`
		NotificationURL   string        `json:"notificationURL,omitempty"`
	}

	// Instruction maps to an item in Instructions field of Payout
	Instruction struct {
		Amount  int64  `json:"amount"`
		Address string `json:"address"`
		Label   string `json:"label"`
	}
)

// CreatePayout creates a payout batch request.
func (c *Client) CreatePayout(p Payout) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("POST", fmt.Sprintf("%s/payouts", c.apiBase), p)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}

// QueryPayouts returns all of the caller's payout requests by status
func (c *Client) QueryPayouts() ([]Payout, *http.Response, error) {
	req, err := c.NewRequestWithAuth("GET", fmt.Sprintf("%s/payouts", c.apiBase), nil)
	if err != nil {
		return nil, nil, err
	}

	var payouts []Payout
	resp, err := c.Send(req, &payouts)

	return payouts, resp, err
}

// DeletePayout cancels the given payout request if status is still new.
func (c *Client) DeletePayout(payoutID string) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("DELETE", fmt.Sprintf("%s/payouts/%s", c.apiBase, payoutID), nil)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}

// UpdatePayout sets the rate for a payout request and/or mark as funded.
// TODO: Reimplement
func (c *Client) UpdatePayout(p Payout) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("PUT", fmt.Sprintf("%s/payouts", c.apiBase), p)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}

// CreatePayoutTransaction
// TODO: Implement

// GetPayout return the specified payout request
func (c *Client) GetPayout(ID string) (*Payout, *http.Response, error) {
	req, err := c.NewRequestWithAuth("GET", fmt.Sprintf("%s/payouts/%s", c.apiBase, ID), nil)
	if err != nil {
		return nil, nil, err
	}

	var payout Payout
	resp, err := c.Send(req, &payout)

	return &payout, resp, err
}

// CreatePayoutsReports creates and returns a payout request report
// TODO: implement
func (c *Client) CreatePayoutsReports() (*http.Response, error) {
	req, err := c.NewRequestWithAuth("POST", fmt.Sprintf("%s/reports/payouts", c.apiBase), nil)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}
