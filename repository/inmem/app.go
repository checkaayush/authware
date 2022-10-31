package inmem

import (
	"context"
	"sync"

	"github.com/checkaayush/authware/model"
)

var (
	apps = map[int]*model.App{
		1: {
			ID:          1,
			Name:        "App 1",
			Description: "My App",
		},
	}
	appSeq  = 2
	appLock = sync.Mutex{}
)

func (i *inMemRepo) ListApps(ctx context.Context) ([]model.App, error) {
	appLock.Lock()
	defer appLock.Unlock()

	as := []model.App{}
	for _, app := range apps {
		as = append(as, *app)
	}
	return as, nil
}

func (i *inMemRepo) CreateApp(ctx context.Context, a *model.App) (*model.App, error) {
	appLock.Lock()
	defer appLock.Unlock()

	a.ID = appSeq
	apps[a.ID] = a
	appSeq++
	return a, nil
}

func (i *inMemRepo) DeleteApp(ctx context.Context, id int) error {
	appLock.Lock()
	defer appLock.Unlock()

	delete(apps, id)
	return nil
}
