package client

// https://test.bitpay.com/api#resource-Sessions

import (
	"fmt"
	"net/http"
)

// Session is a unique session ID to protect against replay attacks
type Session string

// CreateSession creates an API session to protect against replay attacks
// and ensure requests are received in the same order they are sent.
func (c *Client) CreateSession() (Session, *http.Response, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s/sessions", c.apiBase), nil)
	if err != nil {
		return "", nil, err
	}

	var session Session
	resp, err := c.Send(req, &session)

	return session, resp, err
}
