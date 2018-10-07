package base

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type API struct {
	router *mux.Router
	addr   string
}

func NewAPI(addr string) *API {
	return &API{router: mux.NewRouter(), addr: addr}
}

func (a *API) Register(defaultPath string, controller Controller) *API {
	controller.Initialize(defaultPath, a.router, controller)
	return a
}

func (a *API) Run() {
	log.Fatal(http.ListenAndServe(a.addr, a.router))
}
