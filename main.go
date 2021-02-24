package main

import (
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/handlers"
	"github.com/HarrisonWAffel/dbTrain/util"
	"log"
	"net/http"
	"time"
)

func main() {
	time.Sleep(time.Second*10)
	serviceCtx, err := util.NewServiceContext()
	if err != nil {
		panic(err)
	}

	mux := http.DefaultServeMux
	mux.Handle("/user", handlers.Handler{H: handlers.User, AppCtx: serviceCtx})
	mux.Handle("/post", handlers.Handler{H: handlers.Post, AppCtx: serviceCtx})

	log.Println("listening on " + config.GetString("service.port"))
	log.Fatal(http.ListenAndServe(":"+config.GetString("service.port"), mux))
}
