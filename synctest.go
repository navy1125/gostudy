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
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	m1[1] = 1
	m2 = m1
	l1 := []int{1, 2}
	l2 := l1
	l1 = append(l1, 4)
	if l1 == l2 {
		fmt.Println("xxxxxxxxxxxxxx")
	}
	m1 = make(map[int]int)
	fmt.Println(m2)
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
