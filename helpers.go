package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

func renderTempate(w *http.ResponseWriter, tmpl string, tmp *Store) {

	t := template.New("json")

	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, *tmp); err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := *w
	res.Write(buf.Bytes())
	res.(http.Flusher).Flush()
}
