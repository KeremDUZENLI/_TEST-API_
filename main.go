package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// .env
var (
	DB  = "https://my-json-server.typicode.com/typicode/demo/db"
	URL = ":8080"
)

// model/model.go
type DatabaseModel1 struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type DatabaseModel2 struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	PostID int    `json:"postId"`
}

type DatabaseModel3 struct {
	Name string `json:"name"`
}

type DatabaseModelList struct {
	Posts    []DatabaseModel1 `json:"posts"`
	Comments []DatabaseModel2 `json:"comments"`
	Profiles DatabaseModel3   `json:"profile"`
}

// repository/repository.go
func FindPosts() *http.Response {
	response, err := http.Get(DB)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

// service/service.go
func HoldList() DatabaseModelList {
	databaseResult := FindPosts()

	var dML DatabaseModelList
	err := json.NewDecoder(databaseResult.Body).Decode(&dML)
	if err != nil {
		fmt.Println(err)
	}
	return dML
}

// controller/controller.go
func SendList(w http.ResponseWriter, r *http.Request) {
	serviceResult := HoldList()
	fmt.Fprintf(w, convertString(serviceResult))
}

func convertString(d DatabaseModelList) string {
	var str strings.Builder

	for _, v := range d.Posts {
		str.WriteString(fmt.Sprintf("%v\n", v))
	}

	for _, v := range d.Comments {
		str.WriteString(fmt.Sprintf("%v\n", v))
	}

	str.WriteString(fmt.Sprintf("%v\n", d.Profiles))

	return str.String()
}

// router/router.go
func Run(url string) {
	http.HandleFunc("/get", SendList)
	http.ListenAndServe(url, nil)
}

// main.go
func main() {
	// common/env/variables (.env)
	fmt.Printf("\nDB: %v\nRUL: %v\n", DB, URL)

	// repository/repository.go
	resR := FindPosts()
	fmt.Printf("\nRepository Response: %v\n", resR)

	// service/service.go (model/model.go)
	resS := HoldList()
	fmt.Printf("\nService Response: \n%v\n", resS)

	// controller/controller.go (model/model.go)

	// router/router.go
	Run(URL)
}
