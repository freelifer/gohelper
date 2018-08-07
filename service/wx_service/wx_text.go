package wx_service

import (
	"testing"
)

func Test_getWxOpenId(t *testing.T) {
	wxData, err := getWxOpenId("appid", "secret", "code")

	if err != nil {
		t.Error(err)
	}
	t.Log(wxData)
}
