package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/asahi-zip/api/models"
	"github.com/asahi-zip/api/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Media{}, &models.User{}, &models.Org{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	port := flag.String("port", "8080", "Port to run the server on")
	dbConnection := flag.String("db", "host=localhost user=postgres dbname=asahicdn port=5432 sslmode=disable", "Postgres database connection string")

	flag.Parse()

	db, err := InitDatabase(*dbConnection)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	models.DB = db

	r := routes.SetupRouter()

	serverAddress := fmt.Sprintf(":%s", *port)
	log.Printf("Server running on port %s", *port)
	err = http.ListenAndServe(serverAddress, r)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
