package utils

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"io"
	"strings"
	"encoding/json"
	"io/ioutil"
)

type iFields struct {
	Id int `json:"Id"`
	Infinitive string `json:"Infinitive"`
	PastSimple string  `json:"PastSimple"`
	PastParticiple string `json:"PastParticiple"`
	Translate string `json:"Translate"`
}
type aFieldsSlice struct {
	arrayFields []iFields `json:"verbs"`
//	arrayFields map[string]iFields `json:"verbs"`

}
func main() {
fileName := commandline()
fmt.Println(fileName)
lines := fileReader(fileName)
json := jsonner(lines)
fileWriter(json, "output")
}

func fileWriter(b []byte, filename string)  {
	//out := bufio.NewWriter(os.Stdout)
	fmt.Println(string(b))
	var err error
	file := os.Stdout

	if file, err = os.Create(filename); err != nil {
		log.Panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()
	for _, bs  := range b {
		if err := writer.WriteByte(bs); err != nil {
			log.Panic(err)
		}
	}
	ioutil.WriteFile("aaaa", b, 0644)

}
func fileReader(filename string) []string {
	var file *os.File
	var err error
	var lines []string
	if file, err = os.Open(filename); err != nil  {
		log.Panic(err)
	}
	defer file.Close()
	read := bufio.NewReader(file)
	for {
		line, err := read.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic(err)
		}
		lines = append(lines, line)
	}

	return lines

}

func jsonner(lines []string) []byte {
	var afields aFieldsSlice
	var fields iFields

for id, vars := range lines {
	mama := strings.Split(vars,"\t")
	fields.Id = id
	for  i, k := range mama {
		if k != "" && len(mama) == 4 {
			//i++
			//fmt.Println(i, k)
			switch i {
			case  0:
				fields.Infinitive = strings.TrimSpace(k)
			case 1:
				fields.PastSimple = strings.TrimSpace(k)
			case 2:
				fields.PastParticiple = strings.TrimSpace(k)
			case 3:
				fields.Translate = strings.TrimSpace(k)
			}


		}

	}
	afields.arrayFields = append(afields.arrayFields, fields)
}

///"", "    "
//	b, err := json.Marshal(afields.arrayFields)
	b, err := json.MarshalIndent(afields.arrayFields, "", "    ")

	if err != nil {
		fmt.Println("error:", err)
	}
	return b

}
func commandline() string {
	var file string
	if len(os.Args) == 1 || (os.Args[1] == "-h" || os.Args[1] == "-help") {
		fmt.Println("you must  use  file  as argument ")
		os.Exit(1)
	} else if len(os.Args) >= 1 && os.Args[1] != "" {

		 file = os.Args[1]
	}
return file
}