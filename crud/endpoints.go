package crud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var event = []productIntentory{}

// POST  create new
func Create(w http.ResponseWriter, r *http.Request) {
	var newEvent productIntentory
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to create")
	}

	fmt.Println("Method used:", r.Method)

	json.Unmarshal(reqBody, &newEvent)
	event = append(event, newEvent)
	w.WriteHeader(http.StatusCreated)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(event)
	fmt.Println("Method used:", r.Method)
}

func GetByid(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]

	for _, singleEvent := range event {
		if singleEvent.Product.ID == eventId {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
	fmt.Println("Method used:", r.Method)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updateEvent productIntentory

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updateEvent)

	for i, singleEvent := range event {
		if singleEvent.Product.ID == eventID {

			singleEvent.Product.ID = updateEvent.Product.ID
			singleEvent.Product.Code = updateEvent.Product.Name
			singleEvent.Product.Name = updateEvent.Product.Name
			singleEvent.Product.Price = updateEvent.Product.Price

			event = append(event[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
	fmt.Println("Method used:", r.Method)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range event {
		if singleEvent.Product.ID == eventID {
			event = append(event[:i], event[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
	fmt.Println("Method used:", r.Method)
}
