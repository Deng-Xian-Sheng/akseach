package model

import "sync"

////接口
//type Pool interface {
//	// NewConcurrencyLimiter 创建一个并发限制器，limit为并发限制数量，可通过 ResetLimit() 动态调整limit。
//	// 每次调用 Get() 来获取一个资源，然后创建一个协程，完成任务后通过 Release() 释放资源。
//	newConcurrencyLimiter(limit int32) *concurrencyLimiter
//}

type concurrencyLimiter struct {
	resource    int32
	limit       int32
	blockingNum int32
	cond        *sync.Cond
	mu          *sync.Mutex
}

func NewConcurrencyLimiter(limit int32) *concurrencyLimiter {
	l := new(sync.Mutex)
	return &concurrencyLimiter{
		limit: limit,
		cond:  sync.NewCond(l),
		mu:    l,
	}
}

// ResetLimit 可更新limit，需要保证limit > 0
func (c *concurrencyLimiter) resetLimit(limit int32) {
	c.mu.Lock()
	defer c.mu.Unlock()

	tmp := c.limit
	c.limit = limit
	if limit-tmp > 0 {
		for i := int32(0); i < limit-tmp; i++ {
			c.cond.Signal()
		}
	}
}

// Get 当 concurrencyLimiter 没有资源时，会阻塞。
func (c *concurrencyLimiter) get() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.resource < c.limit {
		c.resource++
		return
	}
	c.blockingNum++
	for !(c.resource < c.limit) {
		c.cond.Wait()
	}
	c.resource++
	c.blockingNum--
}

// Release 释放一个资源
func (c *concurrencyLimiter) release() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.blockingNum > 0 {
		c.resource--
		c.cond.Signal()
		return
	}

	c.resource--
}
