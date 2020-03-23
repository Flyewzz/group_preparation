package handlers

import (
	// "log"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Flyewzz/group_preparation/features"
	"github.com/Flyewzz/group_preparation/models"
	"github.com/gorilla/mux"
)

func (hd *HandlerData) SubjectsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	strId := mux.Vars(r)["id"]
	universityId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	var page int = 0
	strPage := r.URL.Query().Get("page")
	if strPage != "" {
		page, err = strconv.Atoi(strPage)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}
	name := r.URL.Query().Get("name")
	semester := r.URL.Query().Get("semester")
	var subjects []models.Subject
	if len(name) == 0 && len(semester) == 0 {
		subjects, err = hd.SubjectController.GetAllSubjects(universityId, page)
		return
	} else {
		subjects, err = hd.SubjectController.Search(universityId, name, semester, page)
	}

	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	pagesCount := features.CalculatePageCount(len(subjects),
		hd.SubjectController.GetItemsPerPageCount())
	subjectsEncoded, err := json.Marshal(subjects)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pagesData := features.PaginatorData{
		Pages:   pagesCount,
		Payload: subjectsEncoded,
	}
	data, err := json.Marshal(pagesData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

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
	if id < 1 {
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
