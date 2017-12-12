package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	jsonStr, err := ioutil.ReadFile("verbs.json")
	if err != nil {
		panic(err)
	}
	//jsonStr := `{"Name":"Adam","Age":36,"Job":"CEO"}`

	personMap := make(map[string]interface{})

	err = json.Unmarshal([]byte(jsonStr), &personMap)

	if err != nil {
		panic(err)
	}

	for key, value := range personMap {
		fmt.Println("index : ", key, " value : ", value)
	}

}
