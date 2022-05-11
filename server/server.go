package server

import (
	"gocode/router"
	"log"
	"net/http"
)

var App = &Myserver{}

type Myserver struct {
}

func (*Myserver) Start(ip, port string) {
	router.Router()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}
}
