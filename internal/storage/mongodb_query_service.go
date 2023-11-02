package storage

import (
	"context"

	"github.com/dr4vs/facts/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFactsQueryService struct{
  db *mongo.Database
}

type FactReadModel struct{
  Id   string `bson:"_id" json:"id"`
  Fact string `bson:"fact" json:"fact"`
}

func InitMongoFactsQueryService(db *mongo.Database) *MongoFactsQueryService {
  return &MongoFactsQueryService{
    db: db,
  }
}

func (qs *MongoFactsQueryService) GetFact() (string, error){
  ctx, cancel := context.WithCancel(context.Background())
  defer cancel()
  pipeline := []bson.M{{"$sample": bson.M{"size": 1}}}
  cursor, err := qs.db.Collection(config.CollectionName).Aggregate(ctx, pipeline)
  defer cursor.Close(ctx)
  if err != nil{
    return "", err
  }
  var facts []FactReadModel
  if err = cursor.All(ctx, &facts); err != nil || len(facts) == 0{
    return "", nil
  }
  return facts[0].Fact, nil
  }
