package mutex

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBadCounter(t *testing.T) {
	badCounter()
}

func TestMutex(t *testing.T) {
	mutex()
}

func TestTryLock(t *testing.T) {
	var mu Mutex
	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)
	if mu.TryLock() {
		fmt.Println("got lock")
		mu.Unlock()
		return
	}
	fmt.Println("can not get lock")
}

func TestState(t *testing.T) {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	fmt.Printf("waitings: %d, loaked: %t, woken: %t starving: %t", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
