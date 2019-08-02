package wechat_domain_valid_check

import (
	"log"
	"testing"
)

func TestValid(t *testing.T) {
	obj := &WechatDomainValidCheck{}

	b, err := obj.InitWechatApp("wxcbc73f324bfxxxx", "2c054d36f69888cebcfbcfbfd0xxxx").Valid("https://www.baidu.com")

	log.Println(b, err)
}
