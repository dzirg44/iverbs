package main

import (
	//"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"html/template"
	"bytes"
)


var globalVerbsStore *FileVerbStore
var  templates =  template.Must(template.New("t").ParseGlob("templates/**/*.html"))
var layout = template.Must(template.New("layout.html").Funcs(layoutFunc).ParseFiles("templates/layout.html"),)

var errorTemplate = `
<html>
 <body>
     <h1>Error rendering template %s</h1>
     <p>%s</p>
 </body>
</html>
`
func init() {
	store, err := NewFileVerbStore("./data/verbs.json")
	if err != nil {
		log.Panic(err)
	}
	globalVerbsStore = store
}


func main() {
r := mux.NewRouter()
r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/"))))
r.HandleFunc("/", indexPage).Methods("GET")
r.HandleFunc("/newtest", newTestPage).Methods("GET")
r.HandleFunc("/addverb", addVerbPage).Methods("POST")
server := http.Server{
	Addr: "127.0.0.1:4444",
	Handler: r,
}
fmt.Println("Run http server on :  ", server.Addr)
if err := server.ListenAndServe(); err != nil {
	log.Fatal(err)
}
}

func addVerbPage(w http.ResponseWriter, r *http.Request) {
	infinitive := r.FormValue("infinitive")
	pastsimple  := r.FormValue("pastsimple")
	pastparticiple  := r.FormValue("pastparticiple")
	translate := r.FormValue("translate")
	verb := Verb{
		Infinitive:infinitive,
		PastSimple:pastsimple,
		PastParticiple:pastparticiple,
		Translate:translate,

	}
	err := globalVerbsStore.Save(verb)
	if err != nil {
		log.Panic(err)
	}
//	fmt.Println(infinitive, pastsimple, pastparticiple, translate)
}

func newTestPage(w http.ResponseWriter,r *http.Request) {
	RenderTemplate(w,r, "verbs/new", map[string]interface{}{})
}

func indexPage(w http.ResponseWriter,r *http.Request) {

verb := globalVerbsStore.AllVerbs

RenderTemplate(w,r, "index/index", map[string]interface{}{
	"Verbs": verb,
})
}



var layoutFunc = template.FuncMap{
	"yield" : func() (string, error) {
		return  "", fmt.Errorf("yield called inapropriately")
	},

}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	if data == nil {
		data = map[string]interface{}{}
	}
	//
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err
		},
	}

	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w,data)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf(errorTemplate,name, err),
			http.StatusInternalServerError,
		)
	}
}
