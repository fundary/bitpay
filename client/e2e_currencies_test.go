package client

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCurrencies(t *testing.T) {
	withContext(func(bitpay *Client) {
		Convey("With the currencies endpoint", t, func() {
			Convey("Retrieving all currencies should be successful", func() {
				currencies, resp, err := bitpay.QueryCurrencies()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(len(currencies), ShouldBeGreaterThan, 0)
				So(currencies[0].Code, ShouldEqual, "BTC")
				So(currencies[0].Symbol, ShouldEqual, "฿")
				So(currencies[0].Name, ShouldEqual, "Bitcoin")
				So(currencies[0].Plural, ShouldEqual, "Bitcoin")
				So(currencies[0].Alts, ShouldEqual, "btc")
				So(currencies[0].PayoutFields[0], ShouldEqual, "bitcoinAddress")
			})
		})
	})
}
