package main

import (
	"fmt"
	"github.com/Sskrill/TestTaskMusic/internal/repository"
	"github.com/Sskrill/TestTaskMusic/internal/service"
	"github.com/Sskrill/TestTaskMusic/internal/transport"
	connDB "github.com/Sskrill/TestTaskMusic/pkg/connectionDB"
	"github.com/Sskrill/TestTaskMusic/pkg/customLogger"
	"log"
	"net/http"
	"os"
)

func main() {

	db, err := connDB.NewDB()
	if err != nil {
		log.Fatalln(err)
	}
	logger := customLogger.NewCSLogger()
	repo := repository.NewRepo(db, logger)
	srvc := service.NewService(repo, logger)
	handler := transport.NewHandler(srvc, logger)
	server := &http.Server{Addr: os.Getenv("PORT"), Handler: handler.InitRouter()}
	fmt.Println("Server started || Сервер запущен")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
