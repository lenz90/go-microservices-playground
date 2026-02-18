package usecase

import (
	"context"
	"playground/03-microservice-foundation/01-clean-arch/internal/domain"
	"playground/03-microservice-foundation/01-clean-arch/internal/repository/inmem"
	"testing"
)

func TestSvc(t *testing.T) {
	r := inmem.New()
	s := Service{R: r}
	u := domain.User{ID: "u1", Email: "e"}
	if e := s.Create(context.Background(), u); e != nil {
		t.Fatal(e)
	}
	if _, e := s.Get(context.Background(), "u1"); e != nil {
		t.Fatal(e)
	}
}
