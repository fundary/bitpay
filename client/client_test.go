package client

import "os"

var testClient *Client
var done = make(chan bool)

func init() {
	getTestClient()
}

func getTestClient() *Client {
	if testClient == nil {

		privateKey := os.Getenv("BITPAY_PRIVATE_KEY")
		if privateKey == "" {
			panic("Bitpay test private key is missing")
		}

		token := os.Getenv("BITPAY_TOKEN")
		if token == "" {
			panic("Bitpay test token is missing")
		}

		testClient = NewClientWithAuth(
			privateKey,
			token,
			APIBaseTest,
		)

		close(done)
	}

	return testClient
}

func withContext(fn func(c *Client)) {
	for {
		_, ok := <-done
		if !ok {
			break
		}
	}
	fn(getTestClient())
}
