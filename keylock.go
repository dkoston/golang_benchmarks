package golang_unlock_benchmark

import "sync"

type KeyLock struct {
	GlobalLock sync.RWMutex
	KeyLocks   map[uint32]*sync.Mutex
}

func NewKeyLock() *KeyLock {
	return &KeyLock{
		GlobalLock: sync.RWMutex{},
		KeyLocks:   map[uint32]*sync.Mutex{},
	}
}

func (self *KeyLock) getLock(key uint32) *sync.Mutex {
	self.GlobalLock.RLock()
	if lock, ok := self.KeyLocks[key]; ok {
		self.GlobalLock.RUnlock()
		return lock
	}

	self.GlobalLock.RUnlock()
	self.GlobalLock.Lock()

	if lock, ok := self.KeyLocks[key]; ok {
		self.GlobalLock.Unlock()
		return lock
	}

	lock := &sync.Mutex{}
	self.KeyLocks[key] = lock
	self.GlobalLock.Unlock()
	return lock
}

func (self *KeyLock) Lock(key uint32) {
	self.getLock(key).Lock()
}

func (self *KeyLock) Unlock(key uint32) {
	self.getLock(key).Unlock()
}

func (self *KeyLock) KeyLocker(key uint32) sync.Locker {
	return self.getLock(key)
}

type KeyRWLock struct {
	GlobalLock sync.RWMutex
	KeyLocks   map[uint32]*sync.RWMutex
}

func NewKeyRWLock() *KeyRWLock {
	return &KeyRWLock{
		GlobalLock: sync.RWMutex{},
		KeyLocks:   map[uint32]*sync.RWMutex{},
	}
}

func (self *KeyRWLock) getLock(key uint32) *sync.RWMutex {
	self.GlobalLock.RLock()
	if lock, ok := self.KeyLocks[key]; ok {
		self.GlobalLock.RUnlock()
		return lock
	}

	self.GlobalLock.RUnlock()
	self.GlobalLock.Lock()

	if lock, ok := self.KeyLocks[key]; ok {
		self.GlobalLock.Unlock()
		return lock
	}

	lock := &sync.RWMutex{}
	self.KeyLocks[key] = lock
	self.GlobalLock.Unlock()
	return lock
}

func (self *KeyRWLock) Lock(key uint32) {
	self.getLock(key).Lock()
}

func (self *KeyRWLock) Unlock(key uint32) {
	self.getLock(key).Unlock()
}

func (self *KeyRWLock) RLock(key uint32) {
	self.getLock(key).RLock()
}

func (self *KeyRWLock) RUnlock(key uint32) {
	self.getLock(key).RUnlock()
}

func (self *KeyRWLock) SymbolLocker(key uint32) sync.Locker {
	return self.getLock(key)
}

func (self *KeyRWLock) SymbolRLocker(key uint32) sync.Locker {
	return self.getLock(key).RLocker()
}
