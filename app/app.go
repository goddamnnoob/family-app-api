package app

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Start() {
	router := httprouter.New()
	router.GET("/ping", Ping)
	log.Fatal(http.ListenAndServe(":8000", router))
}
