package main

import (
	"log"
	"net/http"

	"github.com/repoleved08/chat-golang-vue/db"
)

func main() {
	db.LoadEnv()
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Unable to connect to the database: %s", err)
	}
	log.Fatal(http.ListenAndServe(":8000", nil))
}
