package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router, hd *HandlerData) {
	r.HandleFunc("/universities", hd.AllUniversitiesHandler).Methods("GET")
	r.HandleFunc("/university/{id}/subjects", hd.AllSubjectsHandler).Methods("GET")
	r.HandleFunc("/university", hd.AddUniversityHandler).Methods("POST")
	r.HandleFunc("/university", hd.UniversityByIdRemoveHandler).Methods("DELETE")
	r.HandleFunc("/universities", hd.AllUniversitiesRemoveHandler).Methods("DELETE")
	r.HandleFunc("/university/{id}/subject", hd.AddSubjectHandler).Methods("POST")
}
