package store

import (
	"context"

	"github.com/google/uuid"

	"github.com/m0taru/go-app-structure/app/model"
)

// UserRepo is a store for users
type UserRepo interface {
	GetUserByEmail(context.Context, string) (*model.User, error)

	Get(context.Context, uuid.UUID) (*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	Delete(context.Context, uuid.UUID) error
	List(context.Context) (model.Users, error)
}