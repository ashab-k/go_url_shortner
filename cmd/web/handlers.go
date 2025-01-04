package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)



func (app *application)home(w http.ResponseWriter , r *http.Request) {

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	app.render(w , r , files , &templateData{})
}

func (app *application)shorten(w http.ResponseWriter , r *http.Request) {

	ctx := context.Background()
	err := godotenv.Load()
    if err != nil {
        log.Panic("Error in loading .env file")
    }
    pass := os.Getenv("PASS")
	fmt.Println(pass)
	rdb , err := redisConnect(pass)

	if err != nil {
		log.Fatalf("Error in connecting to redis")
	}

	fmt.Println("Connected to redis")
	err = r.ParseForm()

	if err != nil {
		log.Println("Error in parsing form")
		return
	}

	url := r.Form.Get("url")

	if url == "" {
		log.Println("Empty url")
		return
	}

	log.Println(url)

	randString := String(5)

    for {
        exists , err := rdb.Exists(ctx , randString).Result()

        if err != nil {
        log.Printf("Database Error 1")
        }

        if exists > 0 {
            randString = String(5)
            continue
        }else{
             err := rdb.Set(ctx , randString , url , 0).Err()
            if err != nil {
                log.Println("Database Error 2")
                return
            }
            break
        }

    }
	log.Println(randString)
    app.session.Put(r , "flash" , fmt.Sprintf("http://localhost:5050/%v" , randString))
    http.Redirect(w , r , "/shorten" , http.StatusSeeOther)
}


func (app *application)showShortenLink(w http.ResponseWriter , r *http.Request){
    files := []string{
        "./ui/html/shorten.page.tmpl", 
        "./ui/html/base.layout.tmpl",
    }



    app.render(w , r , files , nil)
}


func (app *application)redirect(w http.ResponseWriter , r *http.Request){
    code := r.URL.Query().Get(":code")

    ctx := context.Background()
    err := godotenv.Load()
    if err != nil {
        log.Panic("Error in loading .env file")
    }
    pass:= os.Getenv("PASS")
    rdb , err  := redisConnect(pass)
    if err != nil {
        log.Panic("Error connecting to redis")
    }

    val , err := rdb.Get(ctx , code).Result()

    if err != nil  {
        log.Panic("Error in getting data from redis")
    }

    if val == ""{
        http.Error(w , err.Error() , http.StatusNotFound)
    }else {
        log.Println(val)
        http.Redirect(w ,r , val , http.StatusSeeOther)
    }
}