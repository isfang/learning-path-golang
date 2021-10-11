package mutex

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {

	current := (*int32)(unsafe.Pointer(&m.Mutex))
	if atomic.CompareAndSwapInt32(current, 0, mutexLocked) {
		return true
	}

	old := atomic.LoadInt32(current)

	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	mutexStatus := old | mutexLocked // new 有警告
	return atomic.CompareAndSwapInt32(current, old, mutexStatus)
}

func (m *Mutex) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)
	return int(v)
}

func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	result := state&mutexStarving == mutexStarving
	return result
}
