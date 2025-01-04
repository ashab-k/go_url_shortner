package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)
func (app *application)routes() http.Handler {
	mux := pat.New()
	mux.Get("/" , app.session.Enable(http.HandlerFunc(app.home)))
	mux.Get("/shorten" , app.session.Enable(http.HandlerFunc(app.showShortenLink)))
	mux.Post("/shorten" , app.session.Enable(http.HandlerFunc(app.shorten)))
	mux.Get("/:code" , app.session.Enable(http.HandlerFunc(app.redirect)))
	
	return mux
}