package main

import (
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/db"
	"github.com/HarrisonWAffel/dbTrain/handlers/post"
	"github.com/HarrisonWAffel/dbTrain/handlers/user"
	"github.com/HarrisonWAffel/dbTrain/util"
	"log"
	"net/http"
	"time"
)

func main() {
	config.Read()
	time.Sleep(time.Second * 10)
	serviceCtx, err := util.NewServiceContext()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(config.Dsn); err != nil {
		if err.Error() != "no change" {
			panic(err)
		}
	}

	mux := http.DefaultServeMux
	mux.Handle("/user/", user.Handler{SrvCtx: serviceCtx})
	mux.Handle("/post/", post.Handler{SrvCtx: serviceCtx})

	log.Println("listening on " + config.GetString("service.port"))
	log.Fatal(http.ListenAndServe(":"+config.GetString("service.port"), mux))
}
