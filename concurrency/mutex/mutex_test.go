package mutex

import "testing"

func TestBadCounter(t *testing.T) {
	badCounter()
}

func TestMutex(t *testing.T) {
	mutex()
}
