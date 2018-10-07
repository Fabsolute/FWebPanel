package base

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller interface {
	Initialize(defaultPath string, router *mux.Router, controller Controller)
	Handle()
}

type ControllerBase struct {
	defaultPath string
	router      *mux.Router
	writer      http.ResponseWriter
	request     *http.Request
	controller  Controller
}

func (c *ControllerBase) Initialize(defaultPath string, router *mux.Router, controller Controller) {
	c.defaultPath = defaultPath
	c.router = router
	c.controller = controller
	c.controller.Handle()
}

func (c *ControllerBase) Register(path string, handlerFunction func() interface{}) *mux.Route {
	return c.router.HandleFunc(c.defaultPath+path, func(writer http.ResponseWriter, request *http.Request) {
		c.writer = writer
		c.request = request
		response := handlerFunction()
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(response)
	})
}

func (c *ControllerBase) GetWriter() http.ResponseWriter {
	return c.writer
}

func (c *ControllerBase) GetRequest() *http.Request {
	return c.request
}
