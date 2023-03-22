package service

import (
	"encoding/json"
	"fmt"
	"testAPI/model"
	"testAPI/repository"
)

type fakeService struct {
	fakeRepository repository.FakeDatabase
}

type FakeService interface {
	GetListByRepo() model.ApiResponse
}

func NewService(repository repository.FakeDatabase) FakeService {
	return fakeService{fakeRepository: repository}
}

func (fS fakeService) GetListByRepo() model.ApiResponse {
	response := fS.fakeRepository.GetListFromFakeDatabase()

	var apiResponse model.ApiResponse
	err := json.NewDecoder(response.Body).Decode(&apiResponse)

	if err != nil {
		fmt.Println("API isteği gönderilemedi:", err)
	}

	return apiResponse
}
