package handlers

import (
	// "log"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Flyewzz/group_preparation/models"
	"github.com/gorilla/mux"
)

func (hd *HandlerData) AddSubjectHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	strId := mux.Vars(r)["id"]
	universityId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	params := r.PostFormValue
	name := params("name")
	semester := params("semester")
	addedId, err := hd.SubjectController.Add(universityId, name, semester)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Added with id %d\n", addedId)))
}

func (hd *HandlerData) SubjectByIdRemoveHandler(w http.ResponseWriter, r *http.Request) {
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
	err = hd.SubjectController.RemoveById(id)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
}

func (hd *HandlerData) SubjectByIdGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	subject, err := hd.SubjectController.GetById(id)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(subject)
	w.Write(data)
}

func (hd *HandlerData) AllSubjectsRemoveHandler(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = hd.SubjectController.RemoveAll(id)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("All subjects was deleted for the university with id %d\n", id)))
}

func (hd *HandlerData) SubjectsSearchHandler(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	universityId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	semester := r.URL.Query().Get("semester")
	var subjects []models.Subject
	if semester == "" {
		subjects, err = hd.SubjectController.SearchByName(universityId, name)
	} else {
		subjects, err = hd.SubjectController.SearchByNameAndSemester(universityId, name, semester)
	}
	if err != nil {
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	data, _ := json.Marshal(subjects)
	w.Write(data)
}
