package client

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBills(t *testing.T) {
	withContext(func(bitpay *Client) {
		Convey("With the bills endpoint", t, func() {
			Convey("Creating a bill should be successful", func() {
				b := Bill{
					Items: []BillItem{
						BillItem{
							Description: "Item 1",
							Price:       100,
							Quantity:    1,
						},
					},
					Currency: "USD",
					ShowRate: "",
					Archived: "",
					Name:     "Foo bar",
					Address1: "Street 1",
					Address2: "",
					City:     "City",
					State:    "State",
					Zip:      "99999",
					Country:  "US",
					Email:    "foo@example.com",
					Phone:    "123456789",
				}

				resp, err := bitpay.CreateBill(b)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
			})

			Convey("Retrieving all bills for the merchant", func() {
				bills, resp, err := bitpay.QueryBills()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(len(bills), ShouldBeGreaterThan, 0)
			})
		})
	})
}
