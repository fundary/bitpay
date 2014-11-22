package client

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSessions(t *testing.T) {
	withContext(func(bitpay *Client) {
		Convey("With the sessions endpoint", t, func() {
			Convey("Creating a session should be successful", func() {
				session, resp, err := bitpay.CreateSession()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(session, ShouldNotEqual, "")
			})
		})
	})
}
