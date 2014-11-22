package client

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
)

type (
	TokenResp struct {
		Policies          []Policy `json:"policies"`
		Token             string   `json:"token"`
		Facade            *Facade  `json:"facade"`
		Label             string   `json:"label"`
		DateCreated       int64    `json:"dateCreated"`
		PairingExpiration int64    `json:"pairingExpiration"`
		PairingCode       string   `json:"pairingCode"`
	}

	Policy struct {
		Policy string   `json:"policy"`
		Method string   `json:"method"`
		Params []string `json:"params"`
	}
)

// NewToken requests a new token from Bitpay
func (c *Client) NewToken(label, clientID string, facade Facade) (TokenResp, error) {
	data := url.Values{}
	data.Add("label", label)
	data.Add("id", clientID)
	data.Add("facade", string(facade))

	tokenResps := []TokenResp{}
	req, err := http.NewRequest("POST", c.apiBase+"/tokens", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return TokenResp{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	_, err = c.Send(req, &tokenResps)
	if len(tokenResps) == 0 {
		return TokenResp{}, err
	}

	return tokenResps[0], err
}

// ClaimToken claims a generated token by using pairing code
func (c *Client) ClaimToken(label, clientID, pairingCode string) (TokenResp, error) {
	data := url.Values{}
	data.Add("label", label)
	data.Add("id", clientID)
	data.Add("pairingCode", pairingCode)

	tokenResps := []TokenResp{}
	req, err := http.NewRequest("POST", c.apiBase+"/tokens", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return TokenResp{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	_, err = c.Send(req, &tokenResps)
	if len(tokenResps) == 0 {
		return TokenResp{}, err
	}

	return tokenResps[0], err
}
