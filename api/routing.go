package api

import (
	"encoding/json"
	"net/http"
	"restfulapi/model"
	"strconv"

	"github.com/gorilla/mux"
)

var items []model.Item

func HomePage(w http.ResponseWriter, r *http.Client) {
	json.NewEncoder(w).Encode(items)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		showError(w, err)
		return
	}
	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		showError(w, err)
		return
	}

	item.ID = len(items) + 1
	items = append(items, item)

	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		showError(w, err)
		return
	}

	var updateItem model.Item
	err = json.NewDecoder(r.Body).Decode(&updateItem)
	if err != nil {
		showError(w, err)
		return
	}

	for idx, item := range items {
		if item.ID == id {
			items[idx] = updateItem
			json.NewEncoder(w).Encode(updateItem)
			return
		}
	}
	http.NotFound(w, r)
}

func Deleteitem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		showError(w, err)
		return
	}

	var deleteitem model.Item
	err = json.NewDecoder(r.Body).Decode(&deleteitem)
	if err != nil {
		showError(w, err)
		return
	}
	for idx, item := range items {
		if item.ID == id {
			items = append(items[:idx], items[idx+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func showError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}
