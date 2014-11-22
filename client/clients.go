package client

// https://test.bitpay.com/api#resource-Clients

type (
	// BitpayClient maps to a resource in clients endpoint
	BitpayClient struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	}
)
