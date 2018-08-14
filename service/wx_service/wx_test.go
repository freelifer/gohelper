package wx_service_test

import (
	"github.com/freelifer/gohelper/models"
	. "github.com/freelifer/gohelper/service/wx_service"
	"github.com/gin-gonic/gin/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_getWxOpenId(t *testing.T) {
	Convey("getWxOpenId", t, func() {
		wxData, err := GetWxOpenId("appid", "secret", "code")
		So(err, ShouldNotBeNil)
		t.Log(wxData)
	})

}

func Test_WxUserJSON(t *testing.T) {
	Convey("wxuser", t, func() {
		Convey("wxuser_json", func() {
			wxUser := models.WxUser{Id: 1}
			b, err := json.Marshal(wxUser)
			So(err, ShouldBeNil)
			result := string(b)
			t.Log(result)
		})
	})

}
