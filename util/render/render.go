package render

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func HTMLRender(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {
	t := template.Must(template.ParseGlob("templates/*"))
	err := t.ExecuteTemplate(w, name, data)
	if err != nil{
		log.Println(err.Error())
	}
}

func JSONRender(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
