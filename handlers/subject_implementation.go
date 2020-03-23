package handlers

import (
	// "log"
	"fmt"
	"net/http"
	"strconv"

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
