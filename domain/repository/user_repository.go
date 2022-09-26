package repository

import (
	"context"
)

type (
	UserRepository interface {
		Fetch(ctx context.Context) error
	}
)
