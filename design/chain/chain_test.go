package chain

import (
	"testing"
)

func TestChain(t *testing.T) {
	iosHandler := &IOSHandler{}
	androidHandler := &AndroidHandler{}

	iosHandler.handler = androidHandler

	iosHandler.Handle("test content")
}
