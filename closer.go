package closer

import "sync"

type Closer struct {
	lock        sync.Mutex
	cbs         []func()
	closeOnce   sync.Once
	IsClosing   bool
	WaitClosing chan bool
}

func NewCloser() Closer {
	return Closer{
		WaitClosing: make(chan bool),
	}
}

func (c *Closer) OnClose(f func()) {
	c.lock.Lock()
	c.cbs = append(c.cbs, f)
	c.lock.Unlock()
}

func (c *Closer) Close() {
	c.closeOnce.Do(func() {
		c.IsClosing = true
		close(c.WaitClosing)
		c.lock.Lock()
		for _, f := range c.cbs {
			f()
		}
		c.lock.Unlock()
	})
}
