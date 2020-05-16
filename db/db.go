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

	fmt.Println("db open...")

	return db
}

func SetRecord(data string) string {

	DB = Open()

	defer DB.Close()

	date := time.Now().AddDate(0, 1, 0).UTC().Format("2006-01-02")

	res, err := DB.Exec("INSERT INTO paste_data(data, expiry_date) VALUES(?,?)", data, date)

	if err != nil {
		log.Panicln(err)
	}

	newId, err := res.LastInsertId()

	id := strconv.Itoa(int(newId))

	if err != nil {
		log.Panicln(err)
	}

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
		pd = PasteData{Id: nId, Data: "Invalid Paste", Expiry: "Invalid Paste", Ts: "Invalid Paste"}
	}

	return pd
}

func DeleteDailyRecords(date time.Time) {

	DB = Open()

	defer DB.Close()

	t := date.UTC().Format("2006-01-02")

	delForm, err := DB.Prepare("DELETE FROM paste_data WHERE expiry_date=?")
	if err != nil {
		log.Panicln(err)
	}

	delForm.Exec(t)
	log.Println("DELETE: " + t)
}
