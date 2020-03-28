package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	// "log"
	"net/http"

	"github.com/Flyewzz/group_preparation/errs"
	"github.com/Flyewzz/group_preparation/features"
	"github.com/Flyewzz/group_preparation/models"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
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
	var elementsCount int = 0
	if name == "" {
		universities, err = hd.UniversityController.GetAll(page)
		elementsCount, err = hd.UniversityController.GetElementsCount()
	} else {
		universities, err = hd.UniversityController.Search(name, page)
		elementsCount = len(universities)
	}
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	pagesCount := features.CalculatePageCount(elementsCount,
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

func (hd *HandlerData) AvatarByIdGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	iconPath, err := hd.UniversityController.GetAvatar(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		}
		return
	}
	file, err := os.Open(iconPath)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
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
	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	name := r.PostFormValue("name")
	fullName := r.PostFormValue("full_name")
	file, header, err := r.FormFile("icon")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// An icon must not be larger than 1 MB
	if header.Size > 1*1024*1024 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()
	log.Printf("%v", header.Header)
	uuid := uuid.NewV4().String()
	extension := features.GetExtension(header.Filename)
	if !features.IsExtensionPicture(extension) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	dir := viper.GetString("icons.directory")
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	path := filepath.Join(
		dir,
		uuid+"."+extension, // Create a name for the file
	)
	f, err := os.OpenFile(path,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	addedId, err := hd.UniversityController.Add(name, fullName, path)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Added with id %d\n", addedId)))
}
