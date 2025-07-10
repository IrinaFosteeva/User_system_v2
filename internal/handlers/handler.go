package handlers

import (
	"User_system_v2/internal/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func MakePersonsHandler(store model.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			persons, err := store.GetAllPersons()
			if err != nil {
				http.Error(w, "Error fetching persons", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(persons)

		case http.MethodPost:
			var p model.Person
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, "Invalid body", http.StatusBadRequest)
				return
			}
			if err := store.CreatePerson(&p); err != nil {
				http.Error(w, "Error creating person", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(p)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func MakePersonHandler(store model.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/persons/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			p, err := store.GetPersonByID(id)
			if err != nil || p == nil {
				http.NotFound(w, r)
				return
			}
			json.NewEncoder(w).Encode(p)

		case http.MethodPatch:
			var upd model.PersonUpdate
			if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
				http.Error(w, "Invalid body", http.StatusBadRequest)
				return
			}
			p, err := store.UpdatePerson(id, &upd)
			if err != nil || p == nil {
				http.NotFound(w, r)
				return
			}
			json.NewEncoder(w).Encode(p)

		case http.MethodDelete:
			if err := store.DeletePerson(id); err != nil {
				http.Error(w, "Error deleting person", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
