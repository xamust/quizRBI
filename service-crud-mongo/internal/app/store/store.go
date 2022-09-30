package store

import (
	"context"
	"github.com/xamust/quizRBI/service-crud-mongo/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client *mongo.Client
	config *Config
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {

	client, err := mongo.NewClient(options.Client().ApplyURI(s.config.ConnString))
	if err != nil {
		return err
	}
	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

func (s *Store) Disconnect() error {
	return s.client.Disconnect(context.TODO())
}

func (s *Store) GetCollection() *mongo.Collection {
	return s.client.Database(s.config.DBName).Collection(s.config.CollectionName)
}

func (s *Store) CreateRecord(ctx context.Context, data model.Data) ([]model.Data, error) {

	req, err := s.GetCollection().InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	result := model.Data{
		ID:         data.ID,
		SimpleData: data.SimpleData,
		InsertedID: req.InsertedID.(primitive.ObjectID).Hex(),
	}
	return s.UpdateRecord(ctx, result)
}

func (s *Store) DeleteRecord(ctx context.Context, data model.Data) ([]model.Data, error) {
	filter := bson.D{{"id", data.ID}}
	_, err := s.GetCollection().DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}

	return []model.Data{
		{
			ID:         data.ID,
			SimpleData: data.SimpleData,
			InsertedID: data.InsertedID,
		},
	}, nil
}

func (s *Store) UpdateRecord(ctx context.Context, data model.Data) ([]model.Data, error) {
	filter := bson.D{{"id", data.ID}}
	fields := bson.D{
		{"$set", bson.D{
			{"simple_data", data.SimpleData},
			{"inserted_id", data.InsertedID},
		}},
	}
	_, err := s.GetCollection().UpdateOne(ctx, filter, fields)
	if err != nil {
		return nil, err
	}
	return []model.Data{
		{
			ID:         data.ID,
			SimpleData: data.SimpleData,
			InsertedID: data.InsertedID},
	}, nil
}

func (s *Store) GetRecords(ctx context.Context) ([]model.Data, error) {

	opt := options.Find()
	var results []model.Data
	cur, err := s.client.Database(s.config.DBName).Collection(s.config.CollectionName).Find(ctx, bson.D{}, opt)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var elem model.Data

		if err = cur.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}

	if err = cur.Close(ctx); err != nil {
		return nil, err
	}

	return results, nil
}
