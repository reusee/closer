package closer

import "testing"

type foo struct {
	Closer
}

func TestCloser(t *testing.T) {
	f := new(foo)
	b := false
	f.OnClose(func() {
		b = true
	})
	f.Close()
	if !b {
		t.Fail()
	}
}
