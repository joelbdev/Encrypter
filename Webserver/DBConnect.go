//Manages the DB for device enrolment
//http://go-database-sql.org/importing.html
//cd '/usr' ; /usr/bin/mysqld_safe --datadir='/var/lib/mysql' (sudo -s)
package Webserver

import (
	"database/sql"
	"fmt"
	"log"
)

//Connect to DB
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@/devices")
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to db: %v", err)
	}
	//do no close connection, as close called by other functions
	log.Println("Initialised connection to DB")
	return db, nil
}

//Inserts device enrolment data into the DB
func Insert(db *sql.DB, device Enumeration) error {
	insert, err := db.Query("INSERT INTO devices ") //fix insert
	if err != nil {
		return fmt.Errorf("failed inserting data: %v", err)
	}
	defer insert.Close()
	defer db.Close()
	defer log.Printf("Inserted device %s into DB", device.ID)

	return nil
}

//Constructs the struct for display on the /view page
func Query(db *sql.DB) ([]Enumeration, error) {
	var devices []Enumeration
	results, err := db.Query("SELECT ID, Hostname, User, IP, Pwd, OS, Encrypted, Key FROM devices")
	if err != nil {
		return nil, fmt.Errorf("could not query db: %v", err)
	}
	//get all values from db
	for results.Next() {
		var device Enumeration
		err = results.Scan(&device.ID, &device.Hostname, &device.User, &device.IP, &device.Pwd, &device.OS, &device.Encrypted, &device.Key)
		if err != nil {
			return nil, fmt.Errorf("Issue querying data from DB", err)
		}
		devices = append(devices, device)
	}
	defer log.Println("User queried the DB")
	defer db.Close()
	return devices, nil
}

//updates values for devices already enrolled
func Update(device Enumeration) error {

	defer log.Printf("Updated device %s into DB", device.ID) //TODO: Can this be changed to show what was updated

	return nil
}
