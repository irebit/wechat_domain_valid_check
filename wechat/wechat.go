package wechat

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/irebit/wechat_domain_valid_check/request"
)

//
const AccessTokenRequestURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

const BanHost = "weixin110.qq.com"

//
const ShortURLRequestURL = "https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s"

type App struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

//
type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpireTime  int64  `json:"expire_time"`
}

//
type ShortURLResp struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	ShortURL string `json:"short_url"`
}

//
func (app *App) GetAccessToken() (accessTokenResp *AccessTokenResp, err error) {
	now := time.Now().Unix()
	body, _, err := request.Get(fmt.Sprintf(AccessTokenRequestURL, app.AppID, app.AppSecret))
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &accessTokenResp); err != nil {
		return
	}
	accessTokenResp.ExpireTime = now + accessTokenResp.ExpiresIn

	return
}

//
func (app *App) GetShortURL(accessToken string, longURL string) (shortURLResp *ShortURLResp, err error) {
	params := map[string]string{
		"action":   "long2short",
		"long_url": longURL,
	}

	requestData, err := json.Marshal(params)
	if err != nil {
		return
	}

	body, _, err := request.Post(fmt.Sprintf(ShortURLRequestURL, accessToken), "application/text", string(requestData))
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &shortURLResp); err != nil {
		return
	}

	return
}

func (app *App) CheckBanHost(url string) bool {
	return strings.Contains(url, BanHost)
}
