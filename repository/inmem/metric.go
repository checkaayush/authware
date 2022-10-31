package inmem

import (
	"context"
	"sync"

	"github.com/checkaayush/authware/model"
)

var (
	metrics = map[int]*model.Metric{
		1: {
			ID:          1,
			Name:        "ARR",
			Description: "Annual Recurring Revenue",
		},
	}
	metricSeq  = 2
	metricLock = sync.Mutex{}
)

func (i *inMemRepo) ListMetrics(ctx context.Context) ([]model.Metric, error) {
	metricLock.Lock()
	defer metricLock.Unlock()

	us := []model.Metric{}
	for _, metric := range metrics {
		us = append(us, *metric)
	}
	return us, nil
}

func (i *inMemRepo) CreateMetric(ctx context.Context, u *model.Metric) (*model.Metric, error) {
	metricLock.Lock()
	defer metricLock.Unlock()

	u.ID = metricSeq
	metrics[u.ID] = u
	metricSeq++
	return u, nil
}

func (i *inMemRepo) DeleteMetric(ctx context.Context, id int) error {
	metricLock.Lock()
	defer metricLock.Unlock()

	delete(metrics, id)
	return nil
}
