package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	// "log"
	"net/http"

	"github.com/gorilla/mux"
)

func (hd *HandlerData) AllUniversitiesHandler(w http.ResponseWriter, r *http.Request) {
	universities, err := hd.UniversityController.GetAll()
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(universities)
	w.Write(data)
}

func (hd *HandlerData) AllSubjectsHandler(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	subjects, err := hd.UniversityController.GetAllSubjects(id)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(subjects)
	w.Write(data)
}

func (hd *HandlerData) AllUniversitiesRemoveHandler(w http.ResponseWriter, r *http.Request) {
	err := hd.UniversityController.RemoveAll()
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("All universities was successfully deleted."))
}

func (hd *HandlerData) UniversityByIdRemoveHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	strId := r.PostFormValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = hd.UniversityController.RemoveById(id)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
}

func (hd *HandlerData) AddUniversityHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	name := r.PostFormValue("name")
	addedId, err := hd.UniversityController.Add(name)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Added with id %d\n", addedId)))
}
