package main

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
)
var jsons = `
[
    {
        "Id": 0,
        "Infinitive": "be",
        "PastSimple": "was, were",
        "PastParticiple": "been",
        "Translate": "быть, являться"
    },
    {
        "Id": 1,
        "Infinitive": "beat",
        "PastSimple": "beat",
        "PastParticiple": "beaten",
        "Translate": "бить, колотить"
    }
]
`
type Verb struct {
	Id             int    `json:"Id"`
	Infinitive     string `json:"Infinitive"`
	PastSimple     string `json:"PastSimple"`
	PastParticiple string `json:"PastParticiple"`
	Translate      string `json:"Translate"`
}
type VerbStore struct {
	filename string
	AllVerbs []Verb
}

func main() {
	filename := "verbstore"
	store := &VerbStore{
		filename: filename,
		//		AllVerbs: map[int]Verb,
		AllVerbs: []Verb{},
	}

	//var store []Verb

	err := json.Unmarshal([]byte(jsons), &store.AllVerbs)

	if err != nil {
		panic(err)
	}

	fmt.Println(store)

}
