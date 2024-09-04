package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hiroshijp/try-clean-arch/handler"
	postgresRepo "github.com/hiroshijp/try-clean-arch/repository/postgres"
	"github.com/hiroshijp/try-clean-arch/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

// set deafult
const (
	defaultAddress = ":8080"
)

// set env
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//  prepare database source name
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// prepare echo
	e := echo.New()

	// prepare repositry
	txRepo := postgresRepo.NewTxRepository(dbConn)
	historyRepo := postgresRepo.NewHistoryRepository(dbConn)
	visitorRepo := postgresRepo.NewVisitorRepository(dbConn)

	// prepare service
	historyUsecase := usecase.NewHistoryUsecase(txRepo, historyRepo, visitorRepo)

	// prepare handler
	handler.NewMiddleware(e)
	handler.NewHistoryHandler(e, historyUsecase)

	// start server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}

	log.Fatal(e.Start(address))
}
