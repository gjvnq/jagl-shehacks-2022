package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/oklog/ulid/v2"
)

type Profile struct {
	Ulid           ulid.ULID        `json:"ulid"`
	Name           string           `json:"name"`
	AllowedViewers []ulid.ULID      `json:"allowedViewers"`
	Notices        []string         `json:"notices"`
	Lights         []IndicatorLight `json:"lights"`
	Alters         []Alter          `json:"alters"`
	Timezone       string           `json:"timezone"`
}

// It's the same ULID as the user's ULID
var Profiles map[ulid.ULID]Profile

func ProfilesModuleLoad() {
	fh, err := os.Open("data/Profiles.json")
	if err != nil {
		log.Fatalf("failed to load Profiles (1): %s", err.Error())
	}
	defer fh.Close()
	data, _ := ioutil.ReadAll(fh)
	err = json.Unmarshal(data, &Profiles)
	if err != nil {
		log.Fatalf("failed to load Profiles (2): %s", err.Error())
	}

	for key, profile := range Profiles {
		Profiles[key] = fixProfile(profile)
	}
}

func ProfilesModuleSave() {
	data, err := json.MarshalIndent(Profiles, "", "\t")
	if err != nil {
		log.Printf("failed to save Profiles (1): %s", err.Error())
	}
	err = ioutil.WriteFile("data/Profiles.json", data, 0644)
	if err != nil {
		log.Printf("failed to save Profiles (2): %s", err.Error())
	}
}

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req_ulid, err := ulid.Parse(vars["ulid"])
	if err != nil {
		log.Printf("failed to parse profile ulid: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(Profiles[req_ulid])
	if err != nil {
		log.Printf("failed to encode profile: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SaveProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req_ulid, err := ulid.Parse(vars["ulid"])
	if err != nil {
		log.Printf("failed to parse profile ulid: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	new_profile := Profile{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to read request body: %s", err.Error())
		return
	}
	err = json.Unmarshal(body, &new_profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("failed to load new user: %s", err.Error())
		return
	}

	new_profile.Ulid = req_ulid
	new_profile = fixProfile(new_profile)
	Profiles[req_ulid] = new_profile
	ProfilesModuleSave()

	// show final result
	err = json.NewEncoder(w).Encode(Profiles[req_ulid])
	if err != nil {
		log.Printf("failed to encode profile: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fixProfile(profile Profile) Profile {
	if profile.AllowedViewers == nil {
		profile.AllowedViewers = make([]ulid.ULID, 0)
	}
	if profile.Lights == nil {
		profile.Lights = make([]IndicatorLight, 0)
	}
	if profile.Notices == nil {
		profile.Notices = make([]string, 0)
	}
	if profile.Alters == nil {
		profile.Alters = make([]Alter, 0)
	}
	return profile
}

func newProfile() Profile {
	profile := Profile{}
	return fixProfile(profile)
}
