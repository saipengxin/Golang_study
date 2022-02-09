package session

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	rwlock     sync.RWMutex
	sessionMap map[string]Session
}

// 定义redis连接
var (
	rdb *redis.Client
)

func NewRedisSessionMgr() *RedisSessionMgr {
	s := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return s
}

// Init 初始化连接
func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	passwd := ""
	if len(options) > 0 {
		passwd = options[0]
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd, // no password set
		DB:       0,      // use default DB
		PoolSize: 100,    // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	// 使用UUID作为sessionid
	id := uuid.NewV4()
	sessionId := id.String()
	// 创建session
	session = NewRedisSession(sessionId)
	// 存入管理session的map中
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
