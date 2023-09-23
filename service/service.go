package service

import (
	"encoding/json"
	"fmt"
	"testAPI/model"
	"testAPI/repository"
)

type holder struct {
	holdList repository.Database
}

type Holder interface {
	HoldList() model.DatabaseModelList
}

func NewService(d repository.Database) Holder {
	return holder{holdList: d}
}

func (h holder) HoldList() model.DatabaseModelList {
	databaseResult := h.holdList.FindPosts()

	var dML model.DatabaseModelList

	err := json.NewDecoder(databaseResult.Body).Decode(&dML)
	if err != nil {
		fmt.Println(err)
	}

	return dML
}
