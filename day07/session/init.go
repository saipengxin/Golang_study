package session

import "fmt"

var (
	sessionMgr SessionMgr
)

func Init(provider string, addr string, options ...string) (sessionMgr SessionMgr, err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("不支持%s模式的session", provider)
	}
	err = sessionMgr.Init(addr, options...)
	return
}
