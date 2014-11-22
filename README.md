#Bitpay

Bitpay command line tool and REST API client for Go
[https://bitpay.com/api](https://bitpay.com/api)

[![GoDoc](https://godoc.org/github.com/fundary/bitpay?status.svg)](https://godoc.org/github.com/fundary/bitpay)

## Status

Currently, not all endpoints are working yet and the request signing seems to be rejected by the API. The new Bitpay API version 2.0.0 is lacking in documentation and the example libraries (PHP, Ruby, Nodejs) also don't work. Hopefully it will be fixed soon by the Bitpay team.

## Installation

```sh
go get github.com/fundary/bitpay
```

## Usage

### Command line tool

The command line tool enable generating and managing client IDs and tokens. To see all the commands and option, run:
```sh
bitpay help
```

For example, to create a new token, run:
```sh
bitpay new-token "Label for the token" TfLgB8tzxxwsefunU3Ec8cjt81bJuvYxX1P merchant --env=test
```

### Go package

The Go client package can be imported and used directly. First generate keys and token using the command line tool. Then pass it to your application.

```go
package main

import(
	"github.com/fundary/bitpay/client"
)

func main() {
	privateKey := os.Getenv("BITPAY_PRIVATE_KEY")
	if privateKey == "" {
		panic("Bitpay private key is missing")
	}

	token := os.Getenv("BITPAY_TOKEN")
	if token == "" {
		panic("Bitpay token is missing")
	}

	bitpay = NewClientWithAuth(
		privateKey,
		token,
		APIBaseTest,
	)

	// Get rates
	rates, resp, err := bitpay.QueryRates()

	// Create invoice
	invoice, _, err := bitpay.CreateInvoice(Invoice{
		Price:             100,
		Currency:          "USD",
		NotificationURL: 'http://your-ipn-server'
	})
	if err != nil {
		panic(err)
	}
	log.Println("Invoice ID", invoice.ID)
}
```

## TODO
- [ ] Make all tests pass
- [ ] Use sessions
- [ ] Allow persisting generated keys and folders to encrypted files
