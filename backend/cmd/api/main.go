package main

import (
	"log"
	"net/http"
	"os"
	"shohinsan/MeetMinder/src/dbrepo"
	"shohinsan/MeetMinder/src/driver"
	"shohinsan/MeetMinder/src/http/handlers"
	"shohinsan/MeetMinder/src/http/router"
	"shohinsan/MeetMinder/src/services/hashrepo"
	"shohinsan/MeetMinder/src/services/tokenrepo"
	"time"

	"github.com/joho/godotenv"
)

const PORT = ":8080"

func main() {
	godotenv.Load()

	dsn := os.Getenv("DB_CONNECTIOn")
	if dsn == "" {
		log.Fatal("DB_CONNECTION environment variable not set")
	}

	_, err := driver.ConnectSQL(dsn)
	if err != nil {
		panic(err)
	}

	// Set up Database Repository
	db := dbrepo.NewTestDBRepo()

	// Set up Hash Repository
	hr := hashrepo.NewBcryptRepo(10)

	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		log.Fatal("JWT_KEY environment variable not set")
	}

	// Set up Token Repository
	tr := tokenrepo.NewJWTTokenRepo(
		"HS256",
		[]byte(jwtKey),
		time.Hour*24*7,
	)

	// Set up HTTP handlers
	handlers.NewHandlers(
		handlers.NewRepository(db, hr, tr),
	)

	// Set up HTTP router
	r := router.NewRouter()

	// Set up HTTP server
	server := &http.Server{
		Addr:    PORT,
		Handler: r,
	}

	log.Printf("Starting server on port %s", PORT)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
