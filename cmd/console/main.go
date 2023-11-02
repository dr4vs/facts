package main

import (
	"fmt"
	"time"

	"github.com/dr4vs/facts/internal/config"
	"github.com/dr4vs/facts/internal/storage"
)

func main() {

	mongoDb, err := config.NewMongoDatabse("mongodb://root:pass@localhost:27017", "facts", 10*time.Second)
	if err != nil {
		panic(err)
	}

	queryService := storage.InitMongoFactsQueryService(mongoDb)
	fact, err := queryService.GetFact()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fact)
}
