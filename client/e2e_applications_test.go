package client

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestApplications(t *testing.T) {
	withContext(func(bitpay *Client) {
		Convey("With the applications endpoint", t, func() {
			Convey("Creating an application should be successful", func() {
				a := Application{
					Users: []ApplicationUser{
						ApplicationUser{
							Email:            "test@example.com",
							FirstName:        "Foo",
							LastName:         "Bar",
							Phone:            "123456789",
							AgreedToTOSAndPP: false,
						},
					},
					Orgs: []ApplicationOrg{
						ApplicationOrg{
							Name:         "Foo Bar LTD",
							Address1:     "Street 1",
							Address2:     "",
							City:         "City",
							State:        "State",
							Zip:          "99999",
							Country:      "US",
							IsNonProfit:  false,
							USTaxID:      "",
							Industry:     "",
							Website:      "http://example.com",
							CartPOS:      "",
							AffiliateOID: "",
						},
					},
				}

				resp, err := bitpay.CreateApplication(a)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
			})
		})
	})
}
