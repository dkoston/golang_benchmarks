package golang_unlock_benchmark

import (
  "testing"
)

func getRandomKey(i int) uint32 {
  if i % 3 == 0 {
    var key uint32 = 1
    return key
  } else if i % 2 == 0 {
    var key uint32 = 9999
    return key
  }
  var key uint32 = 56
  return key
}


func BenchmarkKeyLock(b *testing.B) {
	for n := 0; n < b.N; n++ {
		cache := NewKeyLock()
    key := getRandomKey(n)
    cache.Lock(key)
    cache.Unlock(key)
	}
}

func BenchmarkKeyRWLock(b *testing.B) {
	for n := 0; n < b.N; n++ {
    cache := NewKeyRWLock()
    key := getRandomKey(n)
    cache.Lock(key)
    cache.Unlock(key)
	}
}
