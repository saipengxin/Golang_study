package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	// 存kv
	data   map[string]interface{}
	rwlock sync.RWMutex
}

// NewMemorySession 构造函数
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

// Save 接口统一规定要实现Save方法，这个方法redis版本要使用，但是内存版本不需要使用，写个空方法就行
func (m *MemorySession) Save() (err error) {
	return nil
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return nil
}
