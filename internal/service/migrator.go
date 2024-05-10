package service

import "context"

type Migrator interface {
	Migrate(ctx context.Context) error
}
