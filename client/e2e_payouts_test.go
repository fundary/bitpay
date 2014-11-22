package client

import (
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPayouts(t *testing.T) {
	withContext(func(bitpay *Client) {
		Convey("With the payouts endpoint", t, func() {
			Convey("Creating a payout should be successful", func() {
				p := Payout{
					Instructions: []Instruction{
						Instruction{
							Amount:  100,
							Address: "Street 1, 99999 City, US",
							Label:   "Test instruction",
						},
					},
					Amount:            100,
					Currency:          "BTC",
					EffectiveDate:     time.Now(),
					Reference:         "Foo Bar",
					PricingMethod:     "Pricing method",
					NotificationEmail: "foo@example.com",
					NotificationURL:   "http://example.com",
				}

				resp, err := bitpay.CreatePayout(p)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)

				// TODO: To be implemented
				SkipConvey("Retrieving the newly created payout should be successful", func() {
				})

				// TODO: To be implemented
				SkipConvey("Updating the payout should be successful", func() {
				})

				// TODO: To be implemented
				SkipConvey("Creating payout transactions should be successful", func() {
				})

				// TODO: To be implemented
				SkipConvey("Deleting the payout should be successful", func() {
				})

				// TODO: To be implemented
				SkipConvey("Creating reports for payouts should be successful", func() {
				})

			})

			Convey("Retrieving all payouts should be successful", func() {
				payouts, resp, err := bitpay.QueryPayouts()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(len(payouts), ShouldBeGreaterThan, 0)
			})
		})
	})
}
