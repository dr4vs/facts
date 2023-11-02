package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dr4vs/facts/internal/config"
	"github.com/dr4vs/facts/internal/facts"
	"github.com/dr4vs/facts/internal/storage"
)

func main() {
	log.Println("Starting random facts app...")
	onShutdown, srv := createServer()
	defer onShutdown()
	if srv == nil {
		log.Fatal("Creating server failed. Closing it!")
	}

	go start(srv)

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM)
	defer close(shutdownChannel)
	<-shutdownChannel

	shutdown(srv)
}

func createServer() (func(), *http.Server) {
	// InMemoryRepository
	// repository := storage.InitInMemoryFactsRepository()
	//onShutdown := func(){}

	//TmpDirRepository
	// repository := storage.InitTmpDirFactsRepository()
	//onShutdown := func() {
	// repository.CloseTmpDirFactsRepository()
	// }

	//TmpFileRepository
	// repository := storage.InitTmpFileFactsRepository()
	// onShutdown := func() {
	// 	repository.CloseTmpFileFactsRepository()
	// }

	//PostgresqlRepository
	postgresql, err := config.NewPostgresqlDB("postgres://root:pass@localhost/facts?sslmode=disable", 10*time.Second)
	onShutdown := func() { config.ClosePostgresqlDB(postgresql) }
	if err != nil {
		log.Println(err)
		return onShutdown, nil
	}
	repository := storage.InitPostgresqlFactsRepository(postgresql)
	queryService := storage.InitPostgresqlFactsQueryService(postgresql)

	// Mongo Repository
	// mongodb, err := config.NewMongoDatabse("mongodb://root:pass@localhost:27017", "facts", 10*time.Second)
	//  onShutdown := func() { config.CloseMongoDatabase(mongodb) }
	// if err != nil {
	// 	log.Println(err)
	//    return onShutdown, nil
	// }
	// repository := storage.InitMongoFactsRepository(mongodb)
	// queryService := storage.InitMongoFactsQueryService(mongodb)
	facts.InitializeFacts(&storage.HttpFactsQueryService{}, repository, 100)
	service := facts.InitFactService(queryService, repository)
	handler := facts.NewRouter(service)

	return onShutdown, &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}

func start(srv *http.Server) {
	log.Printf("Server listening on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func shutdown(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Bye bye!")
}
