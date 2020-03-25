package handlers

import (
	// "log"

	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Flyewzz/group_preparation/errs"
	"github.com/Flyewzz/group_preparation/features"
	"github.com/Flyewzz/group_preparation/models"
	"github.com/gorilla/mux"
)

func (hd *HandlerData) MaterialsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	strId := mux.Vars(r)["id"]
	subjectId, err := strconv.Atoi(strId)
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
		if page < 1 {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}
	var materials []models.MaterialData
	name := r.URL.Query().Get("name")
	strTypeId := r.URL.Query().Get("type_id")
	typeId, err := strconv.Atoi(strTypeId)

	var elementsCount int = 0
	if len(name) == 0 && (typeId == 0 || err != nil) {
		materials, err = hd.MaterialController.GetAllMaterials(subjectId, page)
		elementsCount, err = hd.UniversityController.GetElementsCount()
	} else {
		materials, err = hd.MaterialController.Search(subjectId, name, typeId, page)
		elementsCount = len(materials)
	}

	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	pagesCount := features.CalculatePageCount(elementsCount,
		hd.MaterialController.GetItemsPerPageCount())
	materialsEncoded, err := json.Marshal(materials)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pagesData := features.PaginatorData{
		Pages:   pagesCount,
		Payload: materialsEncoded,
	}
	data, err := json.Marshal(pagesData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) AddMaterialHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	strId := mux.Vars(r)["id"]
	subjectId, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	params := r.PostFormValue
	name := params("name")
	strTypeId := params("type_id")
	typeId, err := strconv.Atoi(strTypeId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// AuthorId always will be 1 in debug (the first user makes everything temporary)
	hd.MaterialController
	addedId, err := hd.MaterialController.Add(subjectId, name, typeId, 1)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Added with id %d\n", addedId)))
}

func (hd *HandlerData) MaterialByIdRemoveHandler(w http.ResponseWriter, r *http.Request) {
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
	err = hd.MaterialController.RemoveById(id)
	if err != nil {
		if err == errs.MaterialDoesntExist {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		}
		return
	}
	w.Write([]byte(fmt.Sprintf("A material with id %d was deleted\n", id)))
}

func (hd *HandlerData) MaterialByIdGetHandler(w http.ResponseWriter, r *http.Request) {
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
	material, err := hd.MaterialController.GetById(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			http.Error(w, "Not found", http.StatusNotFound)
		default:
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		}
		return
	}
	data, _ := json.Marshal(material)
	w.Write(data)
}

func (hd *HandlerData) AllMaterialsRemoveHandler(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = hd.MaterialController.RemoveAll(id)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("All materials was deleted for the subject with id %d\n", id)))
}
