package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/oklog/ulid/v2"
)

type User struct {
	Ulid           ulid.ULID   `json:"ulid"`
	Name           string      `json:"name"`
	CanSeeProfiles []ulid.ULID `json:"canSeeProfile"`
}

var Users []User

func UsersModuleLoad() {
	fh, err := os.Open("data/users.json")
	if err != nil {
		log.Fatalf("failed to load users (1): %s", err.Error())
	}
	defer fh.Close()
	data, _ := ioutil.ReadAll(fh)
	err = json.Unmarshal(data, &Users)
	if err != nil {
		log.Fatalf("failed to load users (2): %s", err.Error())
	}
}

func UsersModuleSave() {
	data, err := json.MarshalIndent(Users, "", "\t")
	if err != nil {
		log.Printf("failed to save users (1): %s", err.Error())
	}
	err = ioutil.WriteFile("data/users.json", data, 0644)
	if err != nil {
		log.Printf("failed to save users (2): %s", err.Error())
	}
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(Users)
	if err != nil {
		log.Printf("failed to encode users: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to read request body: %s", err.Error())
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to load new user: %s", err.Error())
		return
	}
	user.Ulid = ulid.Make()
	user.CanSeeProfiles = make([]ulid.ULID, 0)

	if len(strings.TrimSpace(user.Name)) == 0 {
		err = errors.New("name cannot be empty")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to encode user: %s", err.Error())
		return
	}

	Users = append(Users, user)
	profile := newProfile()
	profile.Ulid = user.Ulid
	profile.Name = user.Name
	Profiles[user.Ulid] = profile

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to encode user: %s", err.Error())
		return
	}

	UsersModuleSave()
	ProfilesModuleSave()
}
