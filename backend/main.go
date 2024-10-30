package main

import (
	"log"
	"net/http"

	"github.com/repoleved08/chat-golang-vue/db"
)

func main() {
	db.LoadEnv()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
