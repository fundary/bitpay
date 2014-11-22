package client

// https://bitpay.com/api#resource-Applications

import (
	"fmt"
	"net/http"
)

type (
	// Application maps to a resource at the applications endpoint
	Application struct {
		Users []ApplicationUser `json:"users"`
		Orgs  []ApplicationOrg  `json:"orgs"`
	}

	ApplicationUser struct {
		Email            string `json:"email"`
		FirstName        string `json:"firstName"`
		LastName         string `json:"lastName"`
		Phone            string `json:"phone"`
		AgreedToTOSAndPP bool   `json:"agreedToTOSandPP"`
	}

	ApplicationOrg struct {
		Name         string `json:"name"`
		Address1     string `json:"address1"`
		Address2     string `json:"address2"`
		City         string `json:"city"`
		State        string `json:"state"`
		Zip          string `json:"zip"`
		Country      string `json:"country"`
		IsNonProfit  bool   `json:"isNonProfit,omitempty"`
		USTaxID      string `json:"usTaxId,omitempty"`
		Industry     string `json:"industry"`
		Website      string `json:"website"`
		CartPOS      string `json:"cartPos,omitempty"`
		AffiliateOID string `json:"affiliateOid,omitempty"`
	}
)

// CreateApplication creates an application for a new merchant account
func (c *Client) CreateApplication(a Application) (*http.Response, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s/applications", c.apiBase), a)
	if err != nil {
		return nil, err
	}

	return c.Send(req, nil)
}
