//Manages the DB for device enrolment
//http://go-database-sql.org/importing.html
//cd '/usr' ; /usr/bin/mysqld_safe --datadir='/var/lib/mysql' (sudo -s)
package Webserver

import "log"

//Inserts device enrolment data into the DB
func Insert(device Enumeration) error {

	defer log.Printf("Inserted device %s into DB", device.ID)

	return nil
}

//Constructs the struct for display on the /view page
func Query() ([]Enumeration, error) {
	var devices []Enumeration
	defer log.Println("User queried the DB")

	return devices, nil
}

//updates values for devices already enrolled
func Update(device Enumeration) error {

	defer log.Printf("Updated device %s into DB", device.ID) //TODO: Can this be changed to show what was updated

	return nil
}
