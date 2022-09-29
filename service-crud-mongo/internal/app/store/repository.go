package store

import (
	"context"
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Open() error
	Disconnect() error
	GetCollection() *mongo.Collection
	CreateRecord(ctx context.Context, data model.Data) ([]model.Data, error)
	DeleteRecord(ctx context.Context, data model.Data) ([]model.Data, error)
	UpdateRecord(ctx context.Context, data model.Data) ([]model.Data, error)
	GetRecords(ctx context.Context) ([]model.Data, error)
}
