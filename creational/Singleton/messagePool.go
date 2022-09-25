package singleton

import "sync"

type Message struct {
	Count int
}
type messagePool struct {
	pool *sync.Pool
}

// 消息池单例
var msgPool = &messagePool{
	pool: &sync.Pool{
		New: func() interface{} {
			return &Message{Count: 0}
		},
	},
}

// 访问消息池单例的唯一方法
func Instance() *messagePool {
	return msgPool
}

// 往消息池里添加消息
func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

// 从消息池里获取消息
func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}
