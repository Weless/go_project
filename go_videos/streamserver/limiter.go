package main

type ConnLimiter struct {
	concurrentConn int
	bucket         chan struct{}
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan struct{}, cc),
	}
}

func (c *ConnLimiter) GetConn() bool {
	if len(c.bucket) >= c.concurrentConn {
		return false
	}
	c.bucket <- struct{}{}
	return true
}

func (c *ConnLimiter) ReleaseConn() {
	<-c.bucket
}
