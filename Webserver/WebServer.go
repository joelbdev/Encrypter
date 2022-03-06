package Webserver

import (
	"log"
	"net/http"
)

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	defaultMessage := []byte("Listening for incoming connections ")
	_, err := writer.Write(defaultMessage)
	if err != nil {
		log.Fatal(err)
	}

	//It will need to take incoming strings
}

func InfectedHandler(writer http.ResponseWriter, request *http.Request) {
	//this is where registered devices are kept - receives the JSON
	//eventually this should append to a database
	//sends a command response - tied to Agent.RegisterHost()
}

func CommandHandler(writer http.ResponseWriter, request *http.Request) {
	//Reads the header of the request and issue commands
	//Tied to Agent.KeepAlive()
}
