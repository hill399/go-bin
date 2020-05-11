package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hill399/go-bin/db"
)

type DataPayload struct {
	Data string `json:"Data"`
}

type HashPayload struct {
	Hash string `json:"Hash"`
}

func SubmitData(w http.ResponseWriter, r *http.Request) {
	var newRequest DataPayload

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &newRequest)

	id, err := db.SetRecord(newRequest.Data)

	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusCreated)

	newResponse := HashPayload{id}

	json.NewEncoder(w).Encode(newResponse)

	fmt.Println("Submit:", id)
}

func RequestData(w http.ResponseWriter, r *http.Request) {
	dataId := mux.Vars(r)["id"]

	data, err := db.GetRecord(dataId)

	if err != nil {
		fmt.Println(err)
	}

	newResponse := DataPayload{data}

	json.NewEncoder(w).Encode(newResponse)

	fmt.Println("Request:", dataId)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello Server!")
}
