package repository

import (
	"context"

	"github.com/checkaayush/authware/model"
)

// TODO: Add Update endpoints
type Repository interface {
	UserRepository
	MetricRepository
	AppRepository
	BlockRepository
}

type UserRepository interface {
	ListUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, u *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type MetricRepository interface {
	ListMetrics(ctx context.Context) ([]model.Metric, error)
	CreateMetric(ctx context.Context, m *model.Metric) (*model.Metric, error)
	DeleteMetric(ctx context.Context, id int) error
}

type AppRepository interface {
	ListApps(ctx context.Context) ([]model.App, error)
	CreateApp(ctx context.Context, u *model.App) (*model.App, error)
	DeleteApp(ctx context.Context, id int) error
}

type BlockRepository interface {
	ListBlocks(ctx context.Context) ([]model.Block, error)
	CreateBlock(ctx context.Context, u *model.Block) (*model.Block, error)
	DeleteBlock(ctx context.Context, id int) error
}
