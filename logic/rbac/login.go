package rbac

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserInfo struct {
	WxId      string
	NickName  string
	AvatarUrl string
}

func (user *UserInfo) Login() {

}

func (user *UserInfo) CheckLogin() {

}

type WxResponse struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
	ErrCode    int32  `json:"errcode"`
}

func GetOpenId(code string) (string, error) {
	response, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=wxd889ecfff76dac24&secret=5339787bc0bcd8261a3702e569cf795c&js_code=" + code + "&grant_type=authorization_code")
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var wxRes *WxResponse
	err = json.Unmarshal(bytes, &wxRes)
	if err != nil {
		return "", err
	}
	return wxRes.Openid, nil
}
