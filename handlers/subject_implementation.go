package handlers

import (
	// "fmt"
	// "log"
	"fmt"
	"net/http"
	"strconv"
	// . "github.com/Flyewzz/golang-itv/features"
	// "github.com/Flyewzz/golang-itv/models"
)

func (hd *HandlerData) AddSubjectHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	universityId, err := strconv.Atoi(r.FormValue("university_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	semester := r.FormValue("semester")
	addedId, err := hd.SubjectController.Add(universityId, name, semester)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Added with id %d\n", addedId)))
}
