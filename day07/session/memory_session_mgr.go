package session

import (
	"errors"
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

// NewMemorySessionMgr 构造函数
func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

// redis版本需要初始化，内存版本不需要，写一个空方法就行，接口规定了要实现Init方法
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
	// 将session对象添加进管理的map中
	s.sessionMap[sessionId] = session

	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
