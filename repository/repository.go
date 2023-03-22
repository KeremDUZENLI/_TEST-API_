package repository

import (
	"fmt"
	"net/http"
	"testAPI/common/env"
)

type fakeDatabase struct{}

type FakeDatabase interface {
	GetListFromFakeDatabase() *http.Response
}

func NewRepository() FakeDatabase {
	return &fakeDatabase{}
}

func (fD fakeDatabase) GetListFromFakeDatabase() *http.Response {
	response, err := http.Get(env.URL)

	if err != nil {
		fmt.Println("API isteği gönderilemedi:", err)
	}

	return response
	// defer response.Body.Close()
}
