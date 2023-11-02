package storage

import (
	"context"

  "github.com/dr4vs/facts/internal/config"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFactsRepository struct {
	db *mongo.Database
}

type Fact struct {
	Id   string `bson:"_id" json:"id"`
	Fact string `bson:"fact" json:"fact"`
}

func InitMongoFactsRepository(db *mongo.Database) *MongoFactsRepository {
	return &MongoFactsRepository{
		db: db,
	}
}

func (r *MongoFactsRepository) SaveFact(fact string) error {
	coll := r.db.Collection(config.CollectionName)
	doc := Fact{uuid.New().String(), fact}
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	return nil
}
