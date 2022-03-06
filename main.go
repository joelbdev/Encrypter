package main

import (
	"Encrypter/Webserver"
	"log"
	"net/http"
)

func main() {
	//Initialise webserver
	http.HandleFunc("/", Webserver.HomeHandler)
	http.HandleFunc("/infected", Webserver.InfectedHandler)
	http.HandleFunc("/commands", Webserver.CommandHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatalf("Webserver runtime issue: %s", err.Error())

}
