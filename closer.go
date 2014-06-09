package closer

import "sync"

type Closer struct {
	lock      sync.Mutex
	cbs       []func()
	closeOnce sync.Once
}

func (c *Closer) OnClose(f func()) {
	c.lock.Lock()
	c.cbs = append(c.cbs, f)
	c.lock.Unlock()
}

func (c *Closer) Close() {
	c.closeOnce.Do(func() {
		for _, f := range c.cbs {
			f()
		}
	})
}
