package controller

import (
	"fmt"
	"testAPI/service"
)

type liste struct {
	service service.FakeService
}

type Liste interface {
	GetList()
}

func NewController(sFS service.FakeService) Liste {
	return &liste{sFS}
}

func (l liste) GetList() {
	res := l.service.GetListByRepo()

	fmt.Println("API response:", res.Liste)

	for _, post := range res.Liste {
		fmt.Printf("Post %d: %s\n", post.ID, post.Title)
	}
}
