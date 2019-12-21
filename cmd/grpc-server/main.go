package main

import (
	"fmt"
	"log"
	"os"

	database "github.com/vskut/twigo/pkg/common/db/postgresql"
	"github.com/vskut/twigo/pkg/server"
)

func main() {
	connStr := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err := database.ConnectDB(connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Start server...")

	srv := server.NewServer(db)

	sAddr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	if err := srv.Run(sAddr); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
