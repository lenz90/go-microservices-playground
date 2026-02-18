package unit

import "testing"

func TestPrime(t *testing.T) {
	cases := []struct {
		n int
		w bool
	}{{2, true}, {4, false}}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			if IsPrime(c.n) != c.w {
				t.Fatal()
			}
		})
	}
}
func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsPrime(7919)
	}
}
