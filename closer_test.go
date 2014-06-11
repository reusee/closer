package closer

import "testing"

type foo struct {
	Closer
}

func TestCloser(t *testing.T) {
	// create
	f := foo{
		Closer: NewCloser(),
	}

	// OnClose test
	b := false
	f.OnClose(func() {
		b = true
	})

	// IsClosing test
	f.OnClose(func() {
		if !f.IsClosing {
			t.Fatalf("not IsClosing")
		}
	})

	f.Close()

	if !b {
		t.Fail()
	}

	// WaitClosing test
	select {
	case <-f.WaitClosing:
	default:
		t.Fatalf("WaitClosing not closed")
	}
}
