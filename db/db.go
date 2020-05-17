package db

/*

MySQL Database - testdb

CREATE TABLE IF NOT EXISTS paste_data (
paste_id INT AUTO_INCREMENT PRIMARY KEY,
data VARCHAR(3000),
expiry_date DATE NOT NULL,
ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;

*/

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type PasteData struct {
	Id     int
	Data   string
	Expiry string
	Ts     string
}

var DB *sql.DB

func Open() *sql.DB {
	dbType := os.Getenv("DBTYPE")
	user := os.Getenv("DBUSER")
	pw := os.Getenv("DBPASS")
	table := os.Getenv("DBTABLE")
	address := os.Getenv("DBADDRESS")

	db, err := sql.Open(dbType, user+":"+pw+"@tcp("+address+")/"+table)

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	return db
}

func SetRecord(data string) string {

	DB = Open()
	defer DB.Close()

	res, err := DB.Exec("INSERT INTO paste_data(data, expiry_date) VALUES(?, NOW() + INTERVAL 30 DAY)", data)
	if err != nil {
		log.Panicln(err)
	}

	newId, err := res.LastInsertId()
	if err != nil {
		log.Panicln(err)
	}

	id := strconv.Itoa(int(newId))

	date := time.Now().AddDate(0, 1, 0).UTC().Format("2006-01-02")
	log.Println("INSERT: ID: " + id + " |  Data: " + data + " | Expiry: " + date)

	return id
}

func GetRecord(Id string) PasteData {

	DB = Open()
	defer DB.Close()

	nId, _ := strconv.Atoi(Id)
	selDB, err := DB.Query("SELECT * FROM paste_data WHERE paste_id=?", nId)
	if err != nil {
		log.Panicln(err)
	}

	pd := PasteData{}

	for selDB.Next() {
		var id int
		var data, exp, ts string

		err = selDB.Scan(&id, &data, &exp, &ts)
		if err != nil {
			log.Panicln(err)
		}

		pd.Id = id
		pd.Data = data
		pd.Expiry = exp
		pd.Ts = ts
	}

	if pd.Ts == "" {
		pd = PasteData{Id: nId, Data: "Invalid Paste", Expiry: "-", Ts: "-"}
	}

	return pd
}

func DeleteExpiredRecords() {

	DB = Open()
	defer DB.Close()

	res, err := DB.Exec("DELETE FROM paste_data WHERE expiry_date < NOW()")
	if err != nil {
		log.Panicln(err)
	}

	nDel, err := res.RowsAffected()
	if err != nil {
		log.Panicln(err)
	}

	date := time.Now().AddDate(0, 1, 0).UTC().Format("2006-01-02")

	log.Println("DELETE: " + strconv.Itoa(int(nDel)) + " records older than " + date)
}
