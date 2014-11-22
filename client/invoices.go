package client

import (
	"fmt"
	"net/http"
)

// https://test.bitpay.com/api#resource-Invoices

var (
	InvoiceAdjustmentAcceptUnderpayment InvoiceAdjustment = "acceptUnderpayment"
	InvoiceAdjustmentAcceptOverpayment  InvoiceAdjustment = "acceptOverpayment"
)

type (
	// InvoiceAdjustment is used when accepting the overpayment or underpayment for an invoice.
	InvoiceAdjustment string

	// Invoice maps to a resource at the invoices endpoint
	Invoice struct {
		ID                string `json:"id,omitempty"`
		Price             int64  `json:"price"`
		Currency          string `json:"currency"`
		OrderID           string `json:"orderID,omitempty"`
		ItemDesc          string `json:"itemDesc,omitempty"`
		ItemCode          string `json:"itemCode,omitempty"`
		NotificationEmail string `json:"notificationEmail,omitempty"`
		NotificationURL   string `json:"notificationURL,omitempty"`
		RedirectURL       string `json:"redirectURL,omitempty"`
		POSData           string `json:"posData,omitempty"`
		TransactionSpeed  string `json:"transactionSpeed,omitempty"`
		FullNotifications string `json:"fullNotifications,omitempty"`
		Physical          string `json:"physical,omitempty"`
		Buyer             Buyer  `json:"buyer,omitempty"`
	}

	// Buyer maps to the buyer object in an Invoice
	Buyer struct {
		Name       string `json:"name,omitempty"`
		Address1   string `json:"address1,omitempty"`
		Address2   string `json:"address2,omitempty"`
		Locality   string `json:"locality,omitempty"`
		Region     string `json:"region,omitempty"`
		PostalCode string `json:"postalCode,omitempty"`
		Country    string `json:"country,omitempty"`
		Email      string `json:"email,omitempty"`
		Phone      string `json:"phone,omitempty"`
	}

	// EventResp maps to the response from the GET /invoices/:invoiceId/events
	EventResp struct {
		URL     string   `json:"url"`
		Token   string   `json:"token"`
		Events  []string `json:"events"`
		Actions []string `json:"actions"`
	}

	// InvoiceRefund maps to a resource at the invoice refunds endpoint
	InvoiceRefund struct {
		RequestID      string `json:"requestID,omitempty"`
		BitcoinAddress string `json:"bitcoinAddress,omitempty"`
		Amount         int64  `json:"amount,omitempty"`
		Currency       string `json:"currency,omitempty"`
	}
)

// CreateInvoice creates an invoice for the calling merchant
func (c *Client) CreateInvoice(i Invoice) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("POST", fmt.Sprintf("%s/invoices", c.apiBase), i)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}

// QueryInvoices returns invoices for the calling merchant filtered by query.
func (c *Client) QueryInvoices() ([]Invoice, *http.Response, error) {
	req, err := c.NewRequestWithAuth("GET", fmt.Sprintf("%s/invoices", c.apiBase), nil)
	if err != nil {
		return nil, nil, err
	}

	var invoices []Invoice
	resp, err := c.Send(req, &invoices)

	return invoices, resp, err
}

// GetInvoice returns the specified invoice by ID for the calling merchant
func (c *Client) GetInvoice(ID string) (*Invoice, *http.Response, error) {
	req, err := c.NewRequestWithAuth("GET", fmt.Sprintf("%s/invoices/%s", c.apiBase, ID), nil)
	if err != nil {
		return nil, nil, err
	}

	var invoice Invoice
	resp, err := c.Send(req, &invoice)

	return &invoice, resp, err
}

// GetInvoiceEvents returns a bus token which can be used to subscribe to invoice events
func (c *Client) GetInvoiceEvents(ID string) (*EventResp, *http.Response, error) {
	req, err := c.NewRequestWithAuth("GET", fmt.Sprintf("%s/invoices/%s/events", c.apiBase, ID), nil)
	if err != nil {
		return nil, nil, err
	}

	var eventResp EventResp
	resp, err := c.Send(req, &eventResp)

	return &eventResp, resp, err
}

// CreateInvoiceRefund creates a refund request for a given invoice
func (c *Client) CreateInvoiceRefund(invoiceID string, r InvoiceRefund) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("POST", fmt.Sprintf("%s/invoices/%s/refunds", c.apiBase, invoiceID), r)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}

// GetInvoiceRefund returns the status of a refund
// TODO: to be implemented

// DeleteInvoiceRefund cancels a pending refund request
func (c *Client) DeleteInvoiceRefund(invoiceID, refundID string) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("DELETE", fmt.Sprintf("%s/invoices/%s/refunds/%s", c.apiBase, invoiceID, refundID), nil)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}

// AcceptInvoiceAdjustment accepts the overpayment or underpayment for the invoice.
func (c *Client) AcceptInvoiceAdjustment(invoiceID string, adjustment InvoiceAdjustment) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("POST", fmt.Sprintf("%s/invoices/%s/refunds", c.apiBase, invoiceID), struct {
		Type InvoiceAdjustment `json:"type"`
	}{
		Type: adjustment,
	})
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}

// CreateInvoiceNotification resends the IPN for the specified invoice
func (c *Client) CreateInvoiceNotification(invoiceID string) (*http.Response, error) {
	req, err := c.NewRequestWithAuth("POST", fmt.Sprintf("%s/invoices/%s/notifications", c.apiBase, invoiceID), nil)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}
