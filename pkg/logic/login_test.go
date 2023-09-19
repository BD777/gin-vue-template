package logic

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_passwordHash(t *testing.T) {
	convey.Convey("Test_passwordHash", t, func() {
		convey.Convey("ok", func() {
			resp := passwordHash("lhu8396")
			t.Logf("resp: %v", resp)
		})
	})
}
