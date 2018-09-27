package locktest

import "testing"

func BenchmarkDeferFunction(b *testing.B) {
	for n := 0; n < b.N; n++ {
		check1, check2 := getTestValues(n)
		DeferFunction(check1, check2)
	}
}

func BenchmarkExplicitUnlockFunction(b *testing.B) {
	for n := 0; n < b.N; n++ {
		check1, check2 := getTestValues(n)
		ExplicitUnlockFunction(check1, check2)
	}
}
