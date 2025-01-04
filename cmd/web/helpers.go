package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)
func (app *application)render(w http.ResponseWriter , r *http.Request , files []string  , td *templateData){
    tmpl , err := template.ParseFiles(files...)
    if err != nil {
        http.Error(w , err.Error() , http.StatusInternalServerError)
    }
	if td == nil {
		td = &templateData{}
	}

	td.Flash = app.session.PopString(r , "flash")

    err = tmpl.Execute(w  , nil )
    if err != nil {
        http.Error(w  , err.Error() , http.StatusInternalServerError)
    }
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  return StringWithCharset(length, charset)
}

