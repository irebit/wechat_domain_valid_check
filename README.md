# wechat_domain_valid_check

利用微信公众号提供的短链接接口判断该域名是否被微信屏蔽

### 接口调用

``` golang
    package main 

    import (
        "github.com/irebit/jingdong_union_go"
        "log"
    )

    obj := &WechatDomainValidCheck{}

	b, err := obj.InitWechatApp("微信公众号APP_ID", "微信公众号APP_SECRET").Valid("https://www.baidu.com")

	log.Println(b, err)
```