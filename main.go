package main

import (
	"log"
	"net/http"
	"os"

	"qrgo/platform/authenticator"
	"qrgo/platform/database"
	"qrgo/platform/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

    m2mauth, err := authenticator.NewM2M()
    if err != nil {
        log.Fatalf("Failed to initialize the M2M authenticator: %v", err)
    }

	db, err := database.New(os.Getenv("DB_CONN_STR"))
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	if err := db.Setup(); err != nil {
		log.Fatalf("Failed to setup the database: %v", err)
	}

	rtr := router.New(m2mauth, auth, db)

	log.Print("Server listening on " + os.Getenv("BASE_URL") + os.Getenv("PORT"))
	if err := http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
