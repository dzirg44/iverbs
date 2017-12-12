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

	//jsonStr, err := ioutil.ReadFile("verbs.json")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(jsonStr)
	//jsonStr := `{"Name":"Adam","Age":36,"Job":"CEO"}`


	var people []Verb

	//var personMap []map[string]interface{}

	err := json.Unmarshal([]byte(jsons), &people)

	if err != nil {
		panic(err)
	}
	//fmt.Printf("Birds : %+v", people)
   //fmt.Println(personMap)
   /*
	for _, personData := range personMap {
        fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
		// convert map to array of Person struct
		var p Verb
		p.Id, _ = strconv.Atoi(fmt.Sprintf("%v", personData["Id"]))
		//p.Infinitive = fmt.Sprintf("%s", personData["Infinitive"])
		//p.PastParticiple = fmt.Sprintf("%s", personData["PastSimple"])
		people = append(people, p)

	}
   */
	fmt.Println(people)

}
