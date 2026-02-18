package usecase

import (
	"context"
	"errors"
	"playground/03-microservice-foundation/01-clean-arch/internal/domain"
)

type Repo interface {
	Save(context.Context, domain.User) error
	Get(context.Context, string) (domain.User, error)
}
type Service struct{ R Repo }

func (s Service) Create(ctx context.Context, u domain.User) error {
	if u.ID == "" || u.Email == "" {
		return errors.New("id/email required")
	}
	return s.R.Save(ctx, u)
}
func (s Service) Get(ctx context.Context, id string) (domain.User, error) { return s.R.Get(ctx, id) }
