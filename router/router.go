package router

import (
	"net/http"
	"testAPI/controller"
)

type router struct {
	routeList controller.Sender
}

type Router interface {
	Run(url string)
}

func NewRouter(s controller.Sender) Router {
	router := &router{routeList: s}
	router.setup()

	return router
}

func (r *router) Run(url string) {
	http.ListenAndServe(url, nil)
}

func (r *router) setup() {
	http.HandleFunc("/get", r.routeList.SendList)
}
