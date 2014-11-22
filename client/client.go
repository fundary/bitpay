/*
Package bitpay provides a Go client for the Bitpay REST API

Docs: https://bitpay.com/api
*/
package client

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/conformal/btcec"
	"github.com/fundary/bitauth"

	"code.google.com/p/go-uuid/uuid"
)

const (
	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseTest = "https://test.bitpay.com"

	// APIBaseLive points to the live version of the API
	APIBaseProd = "https://bitpay.com"
)

var (
	// Debug is a flag to enable/disable debug messages
	Debug bool

	FacadePublic   Facade = "public"
	FacadePOS      Facade = "pos"
	FacadeMerchant Facade = "merchant"
)

type (

	// Facade is a named collection of capabilitites
	// https://test.bitpay.com/api#facades
	Facade string

	// Client represents a Bitpay REST API Client
	Client struct {
		client     *http.Client
		privateKey string
		publicKey  string
		sin        string
		token      string
		apiBase    string
	}

	// Response represents a response from Bitpay API, it contains either an error
	// or data
	Response struct {
		Error string          `json:"error"`
		Data  json.RawMessage `json:"data"`
	}
)

func init() {
	var err error
	debug := os.Getenv("BITPAY_DEBUG")
	if debug != "" {
		Debug, err = strconv.ParseBool(debug)
		if err != nil {
			panic("Invalid value for BITPAY_DEBUG")
		}
	}
}

// NewClient returns a new Client struct without keys and SIN
func NewClient(APIBase string) *Client {
	return &Client{
		client:  &http.Client{},
		apiBase: APIBase,
	}
}

// NewClientWithAuth returns a new client with keys and SIN
func NewClientWithAuth(privateKey, token, APIBase string) *Client {
	decoded, err := hex.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}

	_, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), decoded)
	if pubKey == nil {
		panic("Invalid private key")
	}
	publicKey := hex.EncodeToString(pubKey.SerializeCompressed())
	if err != nil {
		panic(err)
	}

	sin, err := bitauth.GetSINFromPublicKeyString(publicKey)
	if err != nil {
		panic(err)
	}
	clientID := string(sin)

	client := NewClient(APIBase)
	client.privateKey = privateKey
	client.publicKey = publicKey
	client.sin = clientID
	client.token = token

	return client
}

// NewRequest constructs a request. If payload is not empty, it will be
// marshalled into JSON
func (c *Client) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		var b []byte
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, url, buf)
}

// NewRequestWithAuth constructs a request. Will do the following things:
// 1. If payload is not empty, it will be marshalled into JSON.
// 2. Applies signing and auth headers.
// 3. Add token and guid to the body
func (c *Client) NewRequestWithAuth(method, endpoint string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	var b []byte
	var err error

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	if payload != nil {
		b, err = json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
	}

	if method == "POST" || method == "PUT" {
		// Add token as field in request body
		var intermediate map[string]interface{}
		err = json.Unmarshal(b, &intermediate)
		if err != nil {
			return nil, err
		}

		intermediate["token"] = c.token

		// If we are creating a new resource, then generate a guid and pass it along
		if method == "POST" {
			intermediate["guid"] = uuid.New()
		}

		b, err = json.Marshal(&intermediate)
		if err != nil {
			return nil, err
		}
	} else {
		// Add token as query param
		if u.RawQuery != "" {
			u.RawQuery += "&"
		}
		u.RawQuery += "token=" + c.token
	}

	if len(b) > 0 {
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, u.String(), buf)

	// Sign the request
	signed, err := bitauth.Sign(u.String()+string(b), c.privateKey)

	req.Header.Set("X-Identity", c.publicKey)
	req.Header.Set("X-Signature", signed)

	return req, err
}

// Send makes a request to the API, the response body will be
// unmarshaled into v, or if v is an io.Writer, the response will
// be written to it without decoding
func (c *Client) Send(req *http.Request, v interface{}) (*http.Response, error) {
	// Set default headers
	req.Header.Set("Accept", "application/json")

	// Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	if req.Header.Get("X-Accept-Version") == "" {
		req.Header.Set("X-Accept-Version", "2.0.0")
	}

	if Debug {
		log.Println(req.Method, ":", req.URL)
		log.Println(req.Header)
		log.Println("Request body:", req.Body)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	if Debug {
		log.Println(resp.Status)
		log.Println(resp.Header)
		log.Println("Response body:", string(data))
	}

	r := Response{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		return resp, err
	}

	if r.Error != "" {
		return resp, errors.New(r.Error)
	}

	if c := resp.StatusCode; c < 200 || c > 299 {
		return resp, errors.New("unknown error occurred")
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.Unmarshal(r.Data, v)
			if err != nil {
				return resp, err
			}
		}
	}

	return resp, nil
}
