package lazy

import (
	"sync"
	"sync/atomic"
)

var initialized uint32

type singleton struct {
}

var Instance *singleton
var lock = sync.Mutex{}

// GetInstance1 double checks
func GetInstance1() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return Instance
	}
	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&initialized) == 0 {
		Instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}
	return Instance
}

var once = sync.Once{}

// GetInstance2 uses sync.Once
func GetInstance2() *singleton {
	once.Do(func() {
		Instance = &singleton{}
	})
	return Instance
}
