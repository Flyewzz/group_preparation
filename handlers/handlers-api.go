package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router, hd *HandlerData) {
	r.HandleFunc("/universities", hd.AllUniversitiesHandler).Methods("GET")
	r.HandleFunc("/university/{id}/subjects", hd.AllSubjectsHandler).Methods("GET")
	r.HandleFunc("/universities", hd.AllUniversitiesRemoveHandler).Methods("DELETE")
	// r.HandleFunc("/requests", uh.AllRequestsHandler).Methods("GET")
	// r.HandleFunc("/requests/page/{number}", uh.PageHandler).Methods("GET")
	// r.HandleFunc("/requests/{id}", uh.RequestIdHandler).Methods("GET")
	// r.HandleFunc("/requests/{id}", uh.RemoveRequestIdHandler).Methods("DELETE")
}
