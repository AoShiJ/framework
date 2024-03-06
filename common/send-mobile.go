package common

import (
	"fmt"
	"github.com/AoShiJ/framework/redis"
	"github.com/cloopen/go-sms-sdk/cloopen"
	"log"
	"math/rand"
	"time"
)

func SendMobile(mobile string) {
	sprintf := ""
	for i := 0; i < 4; i++ {
		sprintf = fmt.Sprintf("%d", rand.Intn(10000))
	}
	cfg := cloopen.DefaultConfig().
		// 开发者主账号,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIAccount("2c94811c8b1e335b018b5fb396f20c2b").
		// 主账号令牌 TOKEN,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIToken("95ade44c9ed44111acf9cbeb0dee67e0")
	sms := cloopen.NewJsonClient(cfg).SMS()
	// 下发包体参数
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "2c94811c8b1e335b018b5fb398770c32",
		// 手机号码
		To: mobile,
		// 模版ID
		TemplateId: "1",
		// 模版变量内容 非必填
		Datas: []string{sprintf},
	}
	// 下发
	resp, err := sms.Send(input)
	if err != nil {
		log.Fatal(err)
		return
	}

	redis.Red.Set(mobile, sprintf, time.Second*120)

	log.Printf("Response MsgId: %s \n", resp.TemplateSMS.SmsMessageSid)

}
