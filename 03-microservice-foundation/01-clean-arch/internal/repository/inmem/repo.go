package inmem

import (
	"context"
	"errors"
	"playground/03-microservice-foundation/01-clean-arch/internal/domain"
)

type Repo struct{ M map[string]domain.User }

func New() *Repo                                            { return &Repo{M: map[string]domain.User{}} }
func (r *Repo) Save(_ context.Context, u domain.User) error { r.M[u.ID] = u; return nil }
func (r *Repo) Get(_ context.Context, id string) (domain.User, error) {
	u, ok := r.M[id]
	if !ok {
		return domain.User{}, errors.New("not found")
	}
	return u, nil
}
