package client

// https://test.bitpay.com/api#resource-Bills

import (
	"fmt"
	"net/http"
)

type (
	// Bill maps to a resource at the bills endpoint
	Bill struct {
		ID       string     `json:"id,omitempty"`
		Items    []BillItem `json:"items"`
		Currency string     `json:"currency,omitempty"`
		ShowRate string     `json:"showRate,omitempty"`
		Archived string     `json:"archived,omitempty"`
		Name     string     `json:"name,omitempty"`
		Address1 string     `json:"address1,omitempty"`
		Address2 string     `json:"address2,omitempty"`
		City     string     `json:"city,omitempty"`
		State    string     `json:"state,omitempty"`
		Zip      string     `json:"zip,omitempty"`
		Country  string     `json:"country,omitempty"`
		Email    string     `json:"email,omitempty"`
		Phone    string     `json:"phone,omitempty"`
	}

	// BillItem maps to an entry in the items array of Bill
	BillItem struct {
		Description string `json:"description"`
		Price       int64  `json:"price"`
		Quantity    int64  `json:"quantity"`
	}
)

// CreateBill creates a bill for the calling merchant
func (c *Client) CreateBill(b Bill) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("POST", fmt.Sprintf("%s/bills", c.apiBase), b)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}

// QueryBills returns all of the caller's bills.
func (c *Client) QueryBills() ([]Bill, *http.Response, error) {
	req, err := c.NewRequestWithAuth("GET", fmt.Sprintf("%s/bills", c.apiBase), nil)
	if err != nil {
		return nil, nil, err
	}

	var bills []Bill
	resp, err := c.Send(req, &bills)

	return bills, resp, err
}

// GetBill returns the specified bill by ID
func (c *Client) GetBill(ID string) (*Bill, *http.Response, error) {
	req, err := c.NewRequestWithAuth("GET", fmt.Sprintf("%s/bills/%s", c.apiBase, ID), nil)
	if err != nil {
		return nil, nil, err
	}

	var bill Bill
	resp, err := c.Send(req, &bill)

	return &bill, resp, err
}

// UpdateBill updates a specified bill by ID
func (c *Client) UpdateBill(b Bill) (*http.Response, error) {
	// Copy ID and unset it so it doesn't get included in the signing of request
	id := b.ID
	b.ID = ""

	req, err := c.NewRequestWithAuth("PUT", fmt.Sprintf("%s/bills/%s", c.apiBase, id), b)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}
