package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hiroshijp/hcce-observer/domain"
	"github.com/hiroshijp/hcce-observer/handler"
	"github.com/hiroshijp/hcce-observer/handler/middleware"
	"github.com/hiroshijp/hcce-observer/handler/public"
	postgresRepo "github.com/hiroshijp/hcce-observer/repository/postgres"
	"github.com/hiroshijp/hcce-observer/usecase"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

// set deafult
const (
	defaultAddress = ":8080"
)

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

	// health check　one two
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK \n")
	})

	// prepare repositry
	txRepo := postgresRepo.NewTxRepository(dbConn)
	historyRepo := postgresRepo.NewHistoryRepository(dbConn)
	visitorRepo := postgresRepo.NewVisitorRepository(dbConn)
	userRepo := postgresRepo.NewUserRepository(dbConn)

	// prepare usecase
	historyUsecase := usecase.NewHistoryUsecase(txRepo, historyRepo, visitorRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// prepare middleware and handler
	middleware.NewCORSMiddleware(e, os.Getenv("ALLOWED_ORIGIN"))
	public.NewVisitedHandler(e, historyUsecase)
	public.NewSigninHandler(e, userUsecase)

	api := e.Group("/api")
	middleware.NewJWTMiddleware(api)
	handler.NewHistoryHandler(api, historyUsecase)
	handler.NewUserHandler(api, userUsecase)

	// prepare admin
	adminName := os.Getenv("ADMIN_NAME")
	if adminName == "" {
		log.Fatal("ADMIN_NAME is not set")
	}
	adminPass := os.Getenv("ADMIN_PASS")
	if adminPass == "" {
		log.Fatal("ADMIN_PASS is not set")
	}
	_ = userRepo.Delete(context.TODO(), adminName)
	admin := &domain.User{
		Name:     adminName,
		Password: adminPass,
		Admin:    true,
	}
	err = userRepo.Store(context.TODO(), admin)
	if err != nil {
		log.Fatal(err)
	}

	// start server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}

	log.Fatal(e.Start(address))
}
