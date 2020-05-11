package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/google/uuid"
)

const (
	dbPath = "./db/data"
)

var Database *badger.DB

func InitDatabase() *badger.DB {
	opts := badger.DefaultOptions(dbPath)

	database, err := badger.Open(opts)
	if err != nil {
		fmt.Println("db cannot be opened", err)
	}

	bytesNow := time.Now().String()

	err = database.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("last_mod")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing database found, creating...")

			err = txn.Set([]byte("last_mod"), []byte(bytesNow))

			return err
		}

		fmt.Println("Existing database found...")

		return nil
	})

	if err != nil {
		fmt.Println("Error in opening database:", err)
	}

	return database
}

func SetRecord(data string) (string, error) {

	id := uuid.New().String()
	id = strings.ReplaceAll(id, "-", "")

	bytesNow := time.Now().String()

	err := Database.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte(id)); err == badger.ErrKeyNotFound {

			err = txn.Set([]byte(id), []byte(data))

			err = txn.Set([]byte("last_mod"), []byte(bytesNow))

			return err
		} else {
			SetRecord(data)
		}

		return nil
	})

	return id, err
}

func GetRecord(id string) (string, error) {

	var response []byte

	err := Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))

		err = item.Value(func(val []byte) error {
			response = val
			return err
		})

		return err
	})

	return string(response), err
}
