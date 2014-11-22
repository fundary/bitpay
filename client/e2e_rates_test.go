package client

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRates(t *testing.T) {
	withContext(func(bitpay *Client) {
		Convey("With the rates endpoint", t, func() {
			Convey("Retrieving all rates should be successful", func() {
				rates, resp, err := bitpay.QueryRates()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(len(rates), ShouldBeGreaterThan, 0)
				So(rates[0].Code, ShouldEqual, "USD")
			})

			Convey("Retrieving exchange rate for USD should be successful", func() {
				rate, resp, err := bitpay.GetRateForCurrency("USD")

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(rate.Code, ShouldEqual, "USD")
			})
		})
	})
}
