package cash

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/model"
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/store"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"time"
)

type Cash struct {
	config  *Config
	ctx     context.Context
	mu      *sync.Mutex
	logger  *logrus.Logger
	cashMap map[string]model.Data
	store   store.Repository
}

func New(config *Config, ctx context.Context, store store.Repository, mu *sync.Mutex, logger *logrus.Logger) *Cash {
	filter := bson.D{}
	count, err := store.GetCollection().CountDocuments(ctx, filter)
	if err != nil {
		return nil
	}
	return &Cash{
		config:  config,
		mu:      mu,
		logger:  logger,
		ctx:     ctx,
		cashMap: make(map[string]model.Data, count),
		store:   store,
	}
}

func (c *Cash) UpdateCash() error {

	for range time.Tick(time.Second * time.Duration(c.config.UpdateInterval)) {
		c.mu.Lock()
		records, err := c.store.GetRecords(c.ctx)
		for _, record := range records {
			c.cashMap[record.InsertedID] = record
		}
		if err != nil {
			return err
		}
		c.mu.Unlock()
		c.logger.Infof("cash update successfully, count of record: %v", len(c.cashMap))
	}
	return nil
}
