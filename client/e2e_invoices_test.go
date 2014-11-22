package client

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInvoices(t *testing.T) {
	withContext(func(bitpay *Client) {
		Convey("With the invoices endpoint", t, func() {
			Convey("Creating an invoice should be successful", func() {
				i := Invoice{
					Price:             100,
					Currency:          "BTC",
					OrderID:           "100000001",
					ItemDesc:          "Test invoice",
					ItemCode:          "123",
					NotificationEmail: "foo@example.com",
					NotificationURL:   "https://example.com/invoice-notification",
					RedirectURL:       "https://example.com/invoice-redirection",
					POSData:           `{"posData":{"orderId":"100000009"},"hash":"GA.VW7o8puqPA"}`,
					Buyer: Buyer{
						Name:       "Foo Bar",
						Address1:   "Street 1",
						Locality:   "City",
						Region:     "State",
						PostalCode: "99999",
						Country:    "US",
						Email:      "foo@example.com",
						Phone:      "123456789",
					},
				}

				resp, err := bitpay.CreateInvoice(i)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)

				// TODO: To be implemented
				SkipConvey("Retrieving the newly created invoice should be successful", func() {
				})

				// TODO: To be implemented
				SkipConvey("Retrieving events for the invoice", func() {
				})

				// TODO: To be implemented
				SkipConvey("Create a refund for the invoice", func() {
				})

				// TODO: To be implemented
				SkipConvey("Delete a refund for the invoice", func() {
				})

				// TODO: To be implemented
				SkipConvey("Accept adjustment for the invoice", func() {
				})

				// TODO: To be implemented
				SkipConvey("Create notification the invoice", func() {
				})

			})

			Convey("Retrieving all invoices for the merchant", func() {
				invoices, resp, err := bitpay.QueryInvoices()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(len(invoices), ShouldBeGreaterThan, 0)
			})
		})
	})
}
