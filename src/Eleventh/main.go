package main

import (
	"fmt"
	"sync"
	"time"
)

func write(d map[string]int) {
	d["test"] = 2
}

func read(d map[string]int) {
	fmt.Println(d["test"])
}

type safeDict struct {
	data  map[string]int
	mutex *sync.Mutex
}

func (s *safeDict) put(key string, value int) {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	s.data[key] = value
}

func (s safeDict) get(key string) int {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	v := s.data[key]
	return v
}

func (s *safeDict) delete(key string) bool {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	_, ok := s.data[key]
	if ok {
		delete(s.data, key)
	}
	return ok
}

func safeWrite(s *safeDict) {
	s.put("apple", 1)
	s.put("peach", 2)
}

type safeCount struct {
	cnt   int
	mutex *sync.RWMutex
}

func (s *safeCount) Write() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i := 5000000; i > 0; i-- {
		s.cnt++
	}
}

func (s safeCount) Read(no int) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	fmt.Println("Now is ", s.cnt, " from repoter ", no)
	if s.cnt == 10000000 {
		return true
	}
	return false
}

func CountWrite(s *safeCount) {
	s.Write()
}
func CountRead(s *safeCount, no int) {
	for {
		v := s.Read(no)
		time.Sleep(time.Second / 1000)
		if v {
			break
		}
	}

}
func main() {
	dict := make(map[string]int)
	go write(dict)
	go read(dict)

	time.Sleep(time.Second / 10)

	var s safeDict = safeDict{data: make(map[string]int), mutex: &sync.Mutex{}}

	go safeWrite(&s)
	go func() {
		fmt.Println(s.get("apple"))
	}()

	time.Sleep(time.Second / 10)
	sa := safeCount{cnt: 0, mutex: &sync.RWMutex{}}
	for i := 0; i < 2; i++ {
		go CountWrite(&sa)
		go CountRead(&sa, i)
	}
	time.Sleep(time.Second)
	fmt.Println(sa.cnt)
}
