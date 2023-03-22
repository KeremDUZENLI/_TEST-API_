package controller

import (
	"fmt"
	"net/http"
	"testAPI/service"
)

type sender struct {
	sendList service.Holder
}

type Sender interface {
	SendList(w http.ResponseWriter, r *http.Request)
}

func NewController(h service.Holder) Sender {
	return &sender{sendList: h}
}

func (s sender) SendList(w http.ResponseWriter, r *http.Request) {
	serviceResult := s.sendList.HoldList()
	fmt.Fprint(w, serviceResult.String())
}
