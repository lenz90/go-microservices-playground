package mocking

import "time"

type Clock interface{ Now() time.Time }
type Store interface{ Save(string, string) error }
type Service struct {
	C Clock
	S Store
}

func (s Service) Greet(name string) (string, error) {
	m := "Hello " + name + " @ " + s.C.Now().Format(time.RFC3339)
	return m, s.S.Save(name, m)
}
