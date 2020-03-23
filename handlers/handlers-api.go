package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router, hd *HandlerData) {
	// University
	r.HandleFunc("/universities", hd.AllUniversitiesHandler).Methods("GET")
	r.HandleFunc("/universities", hd.UniversitiesSearchHandler).Methods("GET")
	r.HandleFunc("/university", hd.UniversityByIdGetHandler).Methods("GET")
	r.HandleFunc("/university", hd.AddUniversityHandler).Methods("POST")
	r.HandleFunc("/university", hd.UniversityByIdRemoveHandler).Methods("DELETE")
	r.HandleFunc("/universities", hd.AllUniversitiesRemoveHandler).Methods("DELETE")

	// Subject
	r.HandleFunc("/university/{id}/subject", hd.AddSubjectHandler).Methods("POST")
	r.HandleFunc("/university/{id}/subjects", hd.AllSubjectsHandler).Methods("GET")
	r.HandleFunc("/subject", hd.SubjectByIdGetHandler).Methods("DELETE")
	r.HandleFunc("/university/{id}/subjects", hd.AllSubjectsRemoveHandler).Methods("DELETE")
	r.HandleFunc("/university/{id}/subjects", hd.SubjectsSearchHandler).Methods("GET")
}
