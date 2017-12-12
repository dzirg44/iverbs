package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"os"
//	"fmt"
)
type VerbStore interface {
	Save(Verb) error
}
type Verb struct {
	Id             int    `json:"Id"`
	Infinitive     string `json:"Infinitive"`
	PastSimple     string `json:"PastSimple"`
	PastParticiple string `json:"PastParticiple"`
	Translate      string `json:"Translate"`
}

type FileVerbStore struct {
	filename string
	AllVerbs []Verb
}

func NewFileVerbStore(filename string) (*FileVerbStore, error) {
	store := &FileVerbStore{
		filename: filename,
		AllVerbs: []Verb{},
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			panic(err)
		}
		return nil, err
	}

	err = json.Unmarshal(contents, &store)
	if err != nil {
		return nil, err
	}

	return store, nil

}

func (store FileVerbStore) Save(verb Verb) error {
	store.AllVerbs[verb.Id] = verb
	contents, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(store.filename, contents, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (store FileVerbStore) Find(verbs int) (*Verb, error) {
	if infinitiveName := store.AllVerbs[verbs]; verbs !=  0  {
		return &infinitiveName, nil
	}
	return nil, nil
}

func (store FileVerbStore) FindByPastSimple(pastsimple string) (*Verb, error) {
	if pastsimple == "" {
		return nil, nil
	}
	for _, psimple := range store.AllVerbs {
		if strings.ToLower(pastsimple) == strings.ToLower(psimple.PastSimple) {
			return &psimple, nil
		}
	}
	return nil, nil
}

func (store FileVerbStore) FindByPastParticiple(pastparticiple string) (*Verb, error) {
	if pastparticiple == "" {
		return nil, nil
	}
	for _, pparticiple := range store.AllVerbs {
		if strings.ToLower(pastparticiple) == strings.ToLower(pparticiple.PastParticiple) {
			return &pparticiple, nil
		}
	}
	return nil, nil
}
