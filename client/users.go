package client

import (
	"fmt"
	"net/http"
)

// https://test.bitpay.com/api#resource-Users

type (
	// User maps to a resource in the users endpoint
	User struct {
		Name  string `json:"name,omitempty"`
		Phone string `json:"phone,omitempty"`
	}
)

// GetUser returns caller's user information.
func (c *Client) GetUser() (*User, *http.Response, error) {
	req, err := c.NewRequestWithAuth("GET", fmt.Sprintf("%s/user", c.apiBase), nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := c.Send(req, &user)

	return &user, resp, err
}

// UpdateUser updates caller's user information.
func (c *Client) UpdateUser(u User) (*User, *http.Response, error) {
	req, err := c.NewRequestWithAuth("PUT", fmt.Sprintf("%s/user", c.apiBase), &u)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := c.Send(req, &user)

	return &user, resp, err
}
