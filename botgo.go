// Package botgo 是一个QQ频道机器人 sdk 的 golang 实现
package bot

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/2mf8/Better-Bot-Go/errs"
	"github.com/2mf8/Better-Bot-Go/log"
	"github.com/2mf8/Better-Bot-Go/openapi"
	v1 "github.com/2mf8/Better-Bot-Go/openapi/v1"
	"github.com/2mf8/Better-Bot-Go/token"
)

var AuthAcess = &AccessConfig{
	Config: make(map[string]*AccessToken, 0),
}

type AccessConfig struct {
	Config map[string]*AccessToken
	rw     sync.RWMutex
}

type AccessToken struct {
	Appid       uint64
	Api         openapi.OpenAPI
	AppSecret   string
	AccessToken string
	ExpiresIn   int64
	IsSandBox   bool
}

func AuthAcessAdd(appid string, accessConfig *AccessToken) *AccessConfig {
	AuthAcess.rw.Lock()
	defer AuthAcess.rw.Unlock()
	AuthAcess.Config[appid] = accessConfig
	return AuthAcess
}

func SendApi(appid string) openapi.OpenAPI {
	at, ok := AuthAcess.Config[appid]
	if ok {
		if at.ExpiresIn-60 < time.Now().Unix() {
			gatr := v1.GetAccessToken(fmt.Sprintf("%v", appid), at.AppSecret)
			if gatr.AccessToken != "" {
				iat, err := strconv.Atoi(gatr.ExpiresIn)
				if err == nil && gatr.AccessToken != "" {
					aei := time.Now().Unix() + int64(iat)
					at.ExpiresIn = aei
					token := token.BotToken(at.Appid, gatr.AccessToken, string(token.TypeQQBot))
					if at.IsSandBox {
						api := NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
						at.Api = api
					} else {
						api := NewOpenAPI(token).WithTimeout(3 * time.Second)
						at.Api = api
					}
					return at.Api
				}
			}
		}
		return at.Api
	}
	return nil
}

func init() {
	v1.Setup() // 注册 v1 接口
}

// SelectOpenAPIVersion 指定使用哪个版本的 api 实现，如果不指定，sdk将默认使用第一个 setup 的 api 实现
func SelectOpenAPIVersion(version openapi.APIVersion) error {
	if _, ok := openapi.VersionMapping[version]; !ok {
		log.Errorf("version %v openapi not found or setup", version)
		return errs.ErrNotFoundOpenAPI
	}
	openapi.DefaultImpl = openapi.VersionMapping[version]
	return nil
}

// NewOpenAPI 创建新的 openapi 实例，会返回当前的 openapi 实现的实例
// 如果需要使用其他版本的实现，需要在调用这个方法之前调用 SelectOpenAPIVersion 方法
func NewOpenAPI(token *token.Token) openapi.OpenAPI {
	return openapi.DefaultImpl.Setup(token, false)
}

// NewSandboxOpenAPI 创建测试环境的 openapi 实例
func NewSandboxOpenAPI(token *token.Token) openapi.OpenAPI {
	return openapi.DefaultImpl.Setup(token, true)
}
