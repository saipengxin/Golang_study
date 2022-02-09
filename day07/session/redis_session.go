package session

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
)

// RedisSession redis 版本的session管理
type RedisSession struct {
	sessionId string
	// 设置session，可以先放在内存的map中，然后批量导入redis，提升性能，缺点，不是实时的
	sessionMap map[string]interface{}
	// 读写锁
	rwlock sync.RWMutex
	// 记录内存中的map是否操作
	flag int
	ctx  context.Context
}

// 用常量定义状态
const (
	// 内存数据没有变化
	SessionFlagNone = iota
	// 有变化
	SessionFlagModify
)

// NewRedisSession 构造函数
func NewRedisSession(id string) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		sessionMap: make(map[string]interface{}, 16),
		flag:       SessionFlagNone,
		ctx:        context.Background(),
	}
	return s
}

// Set 将session存储到map中
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 设置值
	r.sessionMap[key] = value
	return
}

// Save 将session存储到redis中
func (r *RedisSession) Save() (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	if r.flag != SessionFlagModify {
		return
	}
	// 将内存中的session序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	// 将数据存储进redis
	err = rdb.Set(r.ctx, r.sessionId, string(data), 0).Err()
	// 改状态
	r.flag = SessionFlagNone
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("session not exists")
	}
	return
}

// LoadFromRedis 从redis中加载session
func (r *RedisSession) LoadFromRedis() (err error) {
	data, err := rdb.Get(r.ctx, r.sessionId).Result()
	if err != nil {
		return
	}

	// 反序列化
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}
