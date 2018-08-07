package wx_service

import (
	"errors"
	"fmt"
	"gin/json"
	"github.com"
	"github.com/bitly/go-simplejson"
	"github.com/freelifer/gohelper/models"
	"github.com/freelifer/gohelper/pkg/settings"
	"github.com/freelifer/gohelper/pkg/utils"
	"io/ioutil"
	"net/http"
)

const (
	wxUrl = "https://api.weixin.qq.com/sns/jscode2session?grant_type=authorization_code"
)

var (
	WX_LOGIN_UNKNOW error = errors.New("wx error unknow")
)

type WxService struct {
	Code      string
	SessionId string
}

type WeiXinData struct {
	SessionKey string
	Openid     string
}

func (s *WxService) Login() error {
	data, err := getWxOpenId(settings.WxCfg.Appid, settings.WxCfg.Secret, s.Code)
	if err != nil {
		return err
	}

	wxUser, err := models.CreateWxUserWhenNoExist(data.Openid)
	if err != nil {
		return err
	}

	s.SessionId = utils.NewSessionID()

	// save [key, value] to radis

	return nil
}

func Certificate(sessionKey string) error {

}

// From WeiXin Service, Get User's openid and sessionKey
func getWxOpenId(appid, secret, code string) (*WeiXinData, error) {
	url := fmt.Sprintf("%s&appid=%s&secret=%s&js_code=%s", wxUrl, appid, secret, code)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	js, err := simplejson.NewJson(body)

	if err != nil {
		return nil, err
	}

	fmt.Println(js)
	openid := js.Get("openid").MustString()
	if len(openid) == 0 {
		errmsg := js.Get("errmsg").MustString()
		if len(errmsg) == 0 {
			return nil, WX_LOGIN_UNKNOW
		} else {
			return nil, errors.New(errmsg)
		}
	}

	var data = &WeiXinData{}
	data.SessionKey = js.Get("session_key").MustString()
	data.Openid = js.Get("unionid").MustString()
	return data, nil
}
