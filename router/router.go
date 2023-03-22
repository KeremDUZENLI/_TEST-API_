package router

import (
	"testAPI/controller"
)

type router struct {
	controller controller.Liste
}

type Router interface {
	Run(string)
}

func NewRouter(cL controller.Liste) Router {
	router := &router{controller: cL}
	router.setup()

	return router
}

func (r *router) Run(url string) {}

func (r *router) setup() {
	r.controller.GetList()
}
