package router

import (
	"github.com/gorilla/mux"
	"github.com/parking_automation/control"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/parkings", control.GetAllParking).Methods("GET")             // Get all parking slots
	router.HandleFunc("/api/parking", control.CreateNewSlot).Methods("POST")             // Create a new parking slot
	router.HandleFunc("/api/parking/{id}", control.MarkUnavail).Methods("PUT")           // Update or mark a parking slot as unavailable
	router.HandleFunc("/api/parking/{id}", control.DeleteOneSlot).Methods("DELETE")      // Delete a specific parking slot by id
	router.HandleFunc("/api/deleteallparkings", control.DeleteAllSlot).Methods("DELETE") // Delete all parking slots
	return router
}
