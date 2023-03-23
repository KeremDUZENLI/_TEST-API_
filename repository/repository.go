package repository

import (
	"fmt"
	"net/http"
	"testAPI/common/env"
)

type database struct{}

type Database interface {
	FindPosts() *http.Response
}

func NewRepository() Database {
	return &database{}
}

func (database) FindPosts() *http.Response {
	response, err := http.Get(env.DB)
	if err != nil {
		fmt.Println(err)
	}
	return response
}
