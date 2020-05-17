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
	Data   string `json:"Data"`
	Expiry string `json:"Expiry"`
}

type IdPayload struct {
	Id string `json:"Id"`
}

func SubmitData(w http.ResponseWriter, r *http.Request) {
	var newRequest DataPayload

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(body, &newRequest)

	id := db.SetRecord(newRequest.Data)

	w.WriteHeader(http.StatusCreated)

	newResponse := IdPayload{id}

	json.NewEncoder(w).Encode(newResponse)

	fmt.Println("Submit:", id)
}

func RequestData(w http.ResponseWriter, r *http.Request) {
	dataId := mux.Vars(r)["id"]

	data := db.GetRecord(dataId)

	newResponse := DataPayload{Data: data.Data, Expiry: data.Expiry}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(newResponse)

	fmt.Println("Request:", dataId)
}
