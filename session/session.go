package session

import (
	"time"

	resty "github.com/go-resty/resty/v2"
)

// HTTPManager http管理器
type HTTPManager struct {
	MaxRequest  int
	Client      *resty.Client
	AppId       string
	PushChannel chan map[string]interface{}
}

// NewSessionManager 生成SessionManager
func NewSessionManager(ServerHost string,
	onBeforeRequest func(c *resty.Client, r *resty.Request) error,
	onAfterResponse func(c *resty.Client, r *resty.Response) error,
	appId string) *HTTPManager {

	//管理器
	//SessionManager HttpManager实例
	var SessionManager *HTTPManager = &HTTPManager{
		MaxRequest:  3,
		Client:      resty.New(),
		PushChannel: make(chan map[string]interface{}, 50),
		AppId:       appId,
	}
	SessionManager.Client.SetTimeout(12 * time.Second)
	SessionManager.Client.SetHeader("User-Agent", "imovie")
	SessionManager.Client.HostURL = ServerHost
	SessionManager.Client.OnBeforeRequest(onBeforeRequest)
	SessionManager.Client.DisableWarn = true
	SessionManager.Client.OnAfterResponse(onAfterResponse)

	SessionManager.Client.SetDebug(true)

	return SessionManager
}
