package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	// "log"
	"net/http"

	"github.com/Flyewzz/group_preparation/errs"
	"github.com/Flyewzz/group_preparation/features"
	"github.com/Flyewzz/group_preparation/models"
)

func (hd *HandlerData) UniversitiesHandler(w http.ResponseWriter, r *http.Request) {
	var page int = 0
	var err error
	strPage := r.URL.Query().Get("page")
	if strPage != "" {
		page, err = strconv.Atoi(strPage)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}
	name := r.URL.Query().Get("name")
	var universities []models.University
	/* If name is empty, then give all universities.
	 * Search universities by name otherwise
	 */
	if name == "" {
		universities, err = hd.UniversityController.Search(name, page)
	} else {
		universities, err = hd.UniversityController.GetAll(page)
	}
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	pagesCount := features.CalculatePageCount(len(universities),
		hd.SubjectController.GetItemsPerPageCount())
	universitiesEncoded, err := json.Marshal(universities)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pagesData := features.PaginatorData{
		Pages:   pagesCount,
		Payload: universitiesEncoded,
	}
	data, err := json.Marshal(pagesData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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

func (hd *HandlerData) AllUniversitiesRemoveHandler(w http.ResponseWriter, r *http.Request) {
	err := hd.UniversityController.RemoveAll()
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("All universities was successfully deleted."))
}

func (hd *HandlerData) UniversityByIdRemoveHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if id < 1 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = hd.UniversityController.RemoveById(id)
	if err != nil {
		if err == errs.UniversityDoesntExist {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		}
		return
	}
	w.Write([]byte(fmt.Sprintf("A university with id %d was deleted\n", id)))
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
