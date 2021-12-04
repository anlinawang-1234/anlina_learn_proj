package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

func (c *Config) T() {}
func BenchmarkAtomic(b *testing.B) {
	var v atomic.Value
	v.Store(&Config{})

	go func() {
		i := 0
		for {
			i++
			cfg := Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			v.Store(&cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				cfg := v.Load().(*Config)
				cfg.T()
				//fmt.Println(cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkRWMutex(b *testing.B) {
	var cfg Config
	var rwMutex sync.RWMutex

	go func() {
		i := 0
		for {
			i++
			rwMutex.Lock()
			cfg = Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			rwMutex.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				rwMutex.RLock()
				//fmt.Println(cfg)
				cfg.T()
				rwMutex.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkMutex(b *testing.B) {
	var cfg Config
	var mutex sync.Mutex

	go func() {
		i := 0
		for {
			i++
			mutex.Lock()
			cfg = Config{a: []int{i, i + 1, i + 2, i + 3, i + 4}}
			mutex.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				mutex.Lock()
				cfg.T()
				mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 测试命令  go test -bench=. atomic_test.go
func main() {
	b := testing.B{}
	BenchmarkAtomic(&b)
	BenchmarkRWMutex(&b)
	BenchmarkMutex(&b)
}
