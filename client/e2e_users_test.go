package client

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUsers(t *testing.T) {
	withContext(func(bitpay *Client) {
		Convey("With the users endpoint", t, func() {
			Convey("Retrieving the caller's user information should be successful", func() {
				user, resp, err := bitpay.GetUser()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(user.Name, ShouldNotEqual, "")

				Convey("Updating caller's user information should be successful", func() {
					user.Name += " Test"
					updatedUser, resp, err := bitpay.UpdateUser(*user)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(updatedUser.Name, ShouldNotEqual, user.Name)
				})
			})
		})
	})
}
