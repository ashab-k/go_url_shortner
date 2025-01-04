package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golangcollege/sessions"
	"github.com/redis/go-redis/v9"
)

type application struct {
	session *sessions.Session

}

func main(){
	session := sessions.New([]byte("s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	app := &application{
		session: session,
	}

	fmt.Println("Starting server on port 5050")

	err := http.ListenAndServe(":5050" , app.routes())
	if err != nil {
		log.Fatal(err)
	}
}

func redisConnect(pass string) (*redis.Client , error){
	rdb  := redis.NewClient(&redis.Options{
		Addr:     "redis-19192.crce179.ap-south-1-1.ec2.redns.redis-cloud.com:19192",
		Username: "default",
		Password: pass,
		DB:       0,
	})
	
	if rdb == nil {
		return nil , fmt.Errorf("error in connecting to redis")
	}
	return rdb , nil;
}