package wechat_domain_valid_check

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/irebit/wechat_domain_valid_check/request"
	"github.com/irebit/wechat_domain_valid_check/wechat"
)

//
type WechatDomainValidCheck struct {
	WechatApp       *wechat.App
	AccessTokenInfo *wechat.AccessTokenResp
	ShortURLInfo    *wechat.ShortURLResp
	Error           error
}

//
func (w *WechatDomainValidCheck) InitWechatApp(appId, appSecret string) *WechatDomainValidCheck {
	w.WechatApp = &wechat.App{appId, appSecret}
	w.AccessTokenInfo = &wechat.AccessTokenResp{}
	w.ShortURLInfo = &wechat.ShortURLResp{}
	w.Error = nil
	return w
}

func (w *WechatDomainValidCheck) GetAccessToken() *WechatDomainValidCheck {
	if w.AccessTokenInfo.AccessToken == "" || time.Now().Unix() >= w.AccessTokenInfo.ExpireTime {
		w.AccessTokenInfo, w.Error = w.WechatApp.GetAccessToken()
	}
	return w
}

//
func (w *WechatDomainValidCheck) Valid(url string) (b bool, err error) {
	w.GetAccessToken()
	if w.Error != nil {
		return false, w.Error
	}
	log.Println(w.AccessTokenInfo.AccessToken)
	w.ShortURLInfo, w.Error = w.WechatApp.GetShortURL(w.AccessTokenInfo.AccessToken, url)
	log.Println(w.ShortURLInfo)
	if w.Error != nil {
		return false, w.Error
	}

	_, resp, err := request.Get(w.ShortURLInfo.ShortURL)
	redirectURL := resp.Header["Location"][0]
	log.Println(resp.Header["Location"][0])

	if redirectURL == url {
		return true, nil
	} else if w.WechatApp.CheckBanHost(redirectURL) {
		return false, nil
	} else {
		return false, errors.New(fmt.Sprintf("跳转链接[%s]和原链接[%s]不一致", redirectURL, url))
	}
}
