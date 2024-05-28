package main

import (
	"database/sql"
	"fmt"
	repositoryAuthor "go-cleanv2-riky/author/repository"
	"go-cleanv2-riky/delivery/http/middleware"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("failed to load env file")
	}

	log.Printf("Service Run on %s mode \n", os.Getenv("ENVIRONMENT"))
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbPort, dbUser, dbPass, dbName)
	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect postgre, Error:%v", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("postgre doesn't respond, Error:%v", err)
	}

	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	corsMiddleware := middleware.InitMiddleware()
	e.Use(corsMiddleware.CORS)
	authorRepo := repositoryAuthor.NewAuthorRepo(dbConn)
	articleRepo := 

}
