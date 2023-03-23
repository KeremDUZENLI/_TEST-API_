package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DatabaseModel1 struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type DatabaseModelList struct {
	Liste []DatabaseModel1 `json:"posts"`
}

func main() {
	url := "https://my-json-server.typicode.com/typicode/demo/db"

	// repository
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Database not connected:", err)
		return
	}
	defer response.Body.Close()

	// service
	var dML DatabaseModelList
	err = json.NewDecoder(response.Body).Decode(&dML)
	if err != nil {
		fmt.Println("Result is not parsed:", err)
		return
	}

	// controller
	fmt.Printf("List: %v\n\n", dML.Liste)
	for _, post := range dML.Liste {
		fmt.Printf("ID: %d --- Title: %s\n", post.ID, post.Title)
	}
}
