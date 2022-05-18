package user

import "context"

type Repository interface {
	GetUser(ctx context.Context, username string) (User, error)
}
