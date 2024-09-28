package main

import (
	"sync"
	"time"
)

type Mutex struct {
	mutex *sync.Mutex
}

func (mutex *Mutex) RLock() {
	mutex.mutex.Lock()
}

func (mutex *Mutex) RUnlock() {
	mutex.mutex.Unlock()
}

func (mutex *Mutex) WLock() {
	mutex.mutex.Lock()
}

func (mutex *Mutex) WUnlock() {
	mutex.mutex.Unlock()
}

type RWMutex struct {
	mutex *sync.RWMutex
}

func (mutex *RWMutex) RLock() {
	mutex.mutex.RLock()
}

func (mutex *RWMutex) RUnlock() {
	mutex.mutex.RUnlock()
}

func (mutex *RWMutex) WLock() {
	mutex.mutex.Lock()
}

func (mutex *RWMutex) WUnlock() {
	mutex.mutex.Unlock()
}

type Config struct {
	Readers    int
	Writers    int
	ReadPause  time.Duration
	WritePause time.Duration
}

type MutexesDuration struct {
	mutexTimeDuration   time.Duration
	rwMutexTimeDuration time.Duration
}

func (config *Config) Run(
	iters int,
) MutexesDuration {
	mutexesDuration := MutexesDuration{}

	wg := sync.WaitGroup{}
	mutexStart := time.Now()
	mutex := Mutex{
		mutex: &sync.Mutex{},
	}

	for i := 0; i < config.Readers; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iters; i += 1 {
				mutex.RLock()
				time.Sleep(config.ReadPause)
				mutex.RUnlock()
			}
		}()
	}

	for i := 0; i < config.Writers; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iters; i += 1 {
				mutex.WLock()
				time.Sleep(config.WritePause)
				mutex.WUnlock()
			}
		}()
	}

	wg.Wait()
	mutexesDuration.mutexTimeDuration = time.Duration(time.Since(mutexStart).Milliseconds())

	rwMutexStart := time.Now()
	rwMutex := RWMutex{
		mutex: &sync.RWMutex{},
	}
	for i := 0; i < config.Readers; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iters; i += 1 {
				rwMutex.RLock()
				time.Sleep(config.ReadPause)
				rwMutex.RUnlock()
			}
		}()
	}

	for i := 0; i < config.Writers; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iters; i += 1 {
				rwMutex.WLock()
				time.Sleep(config.WritePause)
				rwMutex.WUnlock()
			}
		}()
	}

	wg.Wait()
	mutexesDuration.rwMutexTimeDuration = time.Duration(time.Since(rwMutexStart).Milliseconds())
	return mutexesDuration
}
