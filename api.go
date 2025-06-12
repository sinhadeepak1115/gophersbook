package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

var users = []User{}

func (a *api) getUserHandeler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// encode user slice to JSON
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) creatUserHandeler(w http.ResponseWriter, r *http.Request) {
	// decode request body to User struct
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}
	if err = insertUser(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	// input Validation
	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}
	// storage validation
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User already exists")
		}
	}
	users = append(users, u)
	return nil
}
