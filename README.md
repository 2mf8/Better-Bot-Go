# Better-Bot-Go WebSocket 方式使用
# [QQ交流群 677742758](https://qm.qq.com/q/okWktIaAqk)
```golang
package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/2mf8/Better-Bot-Go/dto"
	"github.com/2mf8/Better-Bot-Go/event"
	v1 "github.com/2mf8/Better-Bot-Go/openapi/v1"
	"github.com/2mf8/Better-Bot-Go/token"
	"github.com/2mf8/Better-Bot-Go/webhook"
	"github.com/2mf8/Better-Bot-Go/websocket2"
)

// 仅供参考，可自由修改成适合你自己的

var botId = "AppId" // 你的AppId

func main() {
	webhook.InitLog()
	ctx := context.Background()
	as := webhook.ReadSetting()
	for _, v := range as.Apps {
		atr := v1.GetAccessToken(fmt.Sprintf("%v", v.AppId), v.AppSecret)
		iat, err := strconv.Atoi(atr.ExpiresIn)
		if err == nil && atr.AccessToken != "" {
			aei := time.Now().Unix() + int64(iat)
			token := token.BotToken(v.AppId, atr.AccessToken, string(token.TypeQQBot))
			if v.IsSandBox {
				api := NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
				go AuthAcessAdd(fmt.Sprintf("%v", v.AppId), &AccessToken{AccessToken: atr.AccessToken, ExpiresIn: aei, Api: api, AppSecret: v.AppSecret, IsSandBox: v.IsSandBox, Appid: v.AppId})
			} else {
				api := NewOpenAPI(token).WithTimeout(3 * time.Second)
				go AuthAcessAdd(fmt.Sprintf("%v", v.AppId), &AccessToken{AccessToken: atr.AccessToken, ExpiresIn: aei, Api: api, AppSecret: v.AppSecret, IsSandBox: v.IsSandBox, Appid: v.AppId})
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
	at, ok := AuthAcess.Config[botId]
	if ok {
		botId = fmt.Sprintf("%v", at.Appid)
		itoken := token.BotToken(at.Appid, at.AccessToken, string(token.TypeQQBot))
		api := NewOpenAPI(itoken).WithTimeout(3 * time.Second)
		wsInfo, err := api.WS(ctx, nil, "")
		if err != nil {
			log.Fatalln(err)
		}
		intent := websocket2.RegisterHandlers(ATMessageEventHandler())
		if err = NewSessionManager().Start(wsInfo, itoken, &intent); err != nil {
			log.Fatalln(err)
		}
	}
}

func ATMessageEventHandler() event.GroupMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSGroupMessageData) error {
		newMsg := &dto.GroupMessageToCreate{
			Content: "测试成功",
			MsgType: dto.C2CMsgTypeText,
			MsgID:   data.MsgId,
			MsgSeq:  4,
		}
		SendApi(botId).PostGroupMessage(context.Background(), data.GroupId, newMsg)
		return nil
	}
}

```

# Better-Bot-Go WebHook 方式使用
## 由原来的 [Go-QQ-SDK](https://github.com/2mf8/Go-QQ-SDK) 迁移过来。

## 可以基于快速开发包快速开发机器人。

快速开发包下载 [GoSDK](https://2mf8.cn/GoSDK.zip)

## 已支持正向 WebSocket ,正向 WebSocket 地址为 `wss://你的域名:端口/websocket` , 程序启动默认启用。

示例 `wss://fw1009zb5979.vicp.fun:443/websocket`

## 对应的 WebSocket 客户端地址 [Bot-Client-Go](https://github.com/2mf8/Bot-Client-Go) ,欢迎大家使用。

QQ机器人，官方 GOLANG SDK。

[![Go Reference](https://pkg.go.dev/badge/github.com/2mf8/Better-Bot-Go.svg)](https://pkg.go.dev/github.com/2mf8/Better-Bot-Go)

# [QQ交流群 677742758](https://qm.qq.com/q/okWktIaAqk)

<details>

<summary><font size="4">已完成功能/开发计划列表</font></summary>

### **登录**

- [x] 登录

### **消息类型**
- [x] 文本
- [x] 图片
- [x] 语音
- [x] MarkDown
- [ ] 表情
- [ ] At
- [ ] 回复
- [ ] 长消息(仅群聊/私聊)
- [ ] 链接分享
- [ ] 小程序(暂只支持RAW)
- [x] 短视频
- [ ] 合并转发
- [ ] 群文件(上传与接收信息)

### **群聊**

- [x] 收发群消息
- [x] 机器人加群通知
- [x] 机器人离群通知
- [x] 群接收机器人主动消息通知
- [x] 群拒绝机器人主动消息通知
- [x] 机器人撤回自己在2分钟内的消息
- [x] 机器人获取群成员列表【需要申请权限】

### **C2C**

- [x] 收发C2C消息
- [x] 机器人加好友通知
- [x] 机器人删好友通知
- [x] 接收机器人消息通知
- [x] 拒绝机器人消息通知
- [x] 机器人撤回自己在2分钟内的消息

</details>

## 一、如何使用

### 1.回调地址配置

#### is_open 为 true 时, 服务端只需要关注 port, cert_file 和 cert_key
https://你的域名:端口/qqbot/你的应用appid/你的应用app_secret

示例 `https://fw1009zb5979.vicp.fun:443/qqbot/101981675/hjksdfhi3jkslfjlksdfjksejkdjk`

#### is_open 为 false 时
https://你的域名:端口/qqbot/你的应用appid

示例 `https://fw1009zb5979.vicp.fun:443/qqbot/101981675`

### 2.配置文件填写（支持多账号）

默认配置文件为
```
{
	"apps": {
		"123456": {
			"qq": 123456,
			"app_id": 123456,
			"token": "你的AppToken",
			"app_secret": "你的AppSecret"
			"is_sandbox": true,
			"wss_addr": "你的wss地址，服务端不用管"
		}
	},
	"port": 8443,
	"cert_file": "ssl证书文件路径",
	"cert_key": "ssl证书密钥"
	"is_open": true
}
```
多账号

```
{
	"apps": {
		"5123456": {
			"qq": 123456,
			"app_id": 5123456,
			"token": "你的AppToken",
			"app_secret": "你的AppSecret"
			"is_sandbox": true,
			"wss_addr": "你的wss地址，服务端不用管"
		},
		"7234567": {
			"qq": 234567,
			"app_id": 7234567,
			"token": "你的AppToken",
			"app_secret": "你的AppSecret"
			"is_sandbox": true,
			"wss_addr": "你的wss地址，服务端不用管"
		}
	},
	"port": 8443,
	"cert_file": "ssl证书文件路径",
	"cert_key": "ssl证书密钥"
	"is_open": true
}
```

### 3.请求 openapi 接口，操作资源

```golang
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	bot1 "github.com/2mf8/Better-Bot-Go"
	"github.com/2mf8/Better-Bot-Go/dto"
	v1 "github.com/2mf8/Better-Bot-Go/openapi/v1"
	"github.com/2mf8/Better-Bot-Go/token"
	"github.com/2mf8/Better-Bot-Go/webhook"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.WithValue(context.Background(), "key", "value")
	webhook.InitLog()
	as := webhook.ReadSetting()
	for _, v := range as.Apps {
		atr := v1.GetAccessToken(fmt.Sprintf("%v", v.AppId), v.AppSecret)
		iat, err := strconv.Atoi(atr.ExpiresIn)
		if err == nil && atr.AccessToken != "" {
			aei := time.Now().Unix() + int64(iat)
			token := token.BotToken(v.AppId, atr.AccessToken, string(token.TypeQQBot))
			if v.IsSandBox {
				api := bot1.NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
				go bot1.AuthAcessAdd(fmt.Sprintf("%v", v.AppId), &bot1.AccessToken{AccessToken: atr.AccessToken, ExpiresIn: aei, Api: api, AppSecret: v.AppSecret, IsSandBox: v.IsSandBox, Appid: v.AppId})
			} else {
				api := bot1.NewOpenAPI(token).WithTimeout(3 * time.Second)
				go bot1.AuthAcessAdd(fmt.Sprintf("%v", v.AppId), &bot1.AccessToken{AccessToken: atr.AccessToken, ExpiresIn: aei, Api: api, AppSecret: v.AppSecret, IsSandBox: v.IsSandBox, Appid: v.AppId})
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
	b, _ := json.Marshal(as)
	fmt.Println("配置", string(b))
	webhook.GroupAtMessageEventHandler = func(bot *webhook.BotHeaderInfo, event *dto.WSPayload, data *dto.WSGroupATMessageData) error {
		fmt.Println(bot.XBotAppid, data.GroupId, data.Content)
		if len(data.Attachments) > 0 {
			log.Infof(`BotId(%s) GroupId(%s) UserId(%s) <- %s <image id="%s">`, bot.XBotAppid[0], data.GroupId, data.Author.UserId, data.Content, data.Attachments[0].URL)
		} else {
			log.Infof("BotId(%s) GroupId(%s) UserId(%s) <- %s", bot.XBotAppid[0], data.GroupId, data.Author.UserId, data.Content)
		}
		if strings.TrimSpace(data.Content) == "测试" {
			bot1.SendApi(bot.XBotAppid[0]).PostGroupMessage(ctx, data.GroupId, &dto.GroupMessageToCreate{
				Content: "成功",
				MsgID:   data.MsgId,
				MsgType: 0,
			})
		}
		return nil
	}
	webhook.C2CMessageEventHandler = func(bot *webhook.BotHeaderInfo, event *dto.WSPayload, data *dto.WSC2CMessageData) error {
		b, _ := json.Marshal(event)
		fmt.Println(bot.XBotAppid, string(b), data.Content)
		return nil
	}
	webhook.MessageEventHandler = func(bot *webhook.BotHeaderInfo, event *dto.WSPayload, data *dto.WSMessageData) error {
		b, _ := json.Marshal(event)
		fmt.Println(bot.XBotAppid, string(b), data.Content)
		return nil
	}
	webhook.InitGin(as.IsOpen)
	select {}
}
```