package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

//    MemorySeesionMgr设计：
//    定义MemorySeesionMgr对象（字段：存放所有session的map，读写锁）
//    构造函数
//    Init()
//    CreateSeesion()
//    GetSession()

// MemorySessionMgr 定义对象
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// NewMemorySeesionMgr 构造函数
func NewMemorySeesionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	// go get github.com/satori/go.uuid
	// 我们要使用一个唯一的ID来当作sessionID，这里使用的是go的第三方库uuid来实现的
	// 用uuid作为sessionId
	id := uuid.NewV4()

	// 转string
	sessionId := id.String()
	// 创建个session
	session = NewMemorySession(sessionId)

	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	return
}
