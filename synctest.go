package main

import (
	"fmt"
	"sync"
)

var (
	rwlock sync.RWMutex
	mutex  sync.Mutex
)

func main() {
	mutex.Lock()
	fmt.Println("mutex.locck")
	mutex.Unlock()
	fmt.Println("mutex.unlocck")
	rwlock.RLock()
	fmt.Println("rwlock.rdlock")
	rwlock.RUnlock()
	fmt.Println("rwlock.rwunlock")
	rwlock.Lock()
	fmt.Println("rwlock.unlock")
	rwlock.Unlock()
	m := map[int]int{}
	m[1] = 1
	//delete(m, 1)
	fmt.Println(len(m))
}
