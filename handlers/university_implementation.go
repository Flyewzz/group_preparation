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

func (hd *HandlerData) UniversityByIdGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	university, err := hd.UniversityController.GetById(id)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(university)
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

func (hd *HandlerData) UniversitiesSearchHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	universities, err := hd.UniversityController.SearchByName(name)
	if err != nil {
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	data, _ := json.Marshal(universities)
	w.Write(data)
}
