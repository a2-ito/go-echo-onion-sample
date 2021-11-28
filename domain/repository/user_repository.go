package repository

import ()

type (
	UserRepository interface {
		Fetch(ctx context.Context) error
	}
)
