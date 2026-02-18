package mocking

import (
	"testing"
	"time"
)

type fc struct{}

func (fc) Now() time.Time { return time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) }

type fs struct{ m map[string]string }

func (s *fs) Save(n, m string) error { s.m[n] = m; return nil }
func TestSvc(t *testing.T) {
	st := &fs{m: map[string]string{}}
	sv := Service{C: fc{}, S: st}
	m, _ := sv.Greet("Ada")
	if st.m["Ada"] != m {
		t.Fatal()
	}
}
