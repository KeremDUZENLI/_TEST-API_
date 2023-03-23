package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// .env
var (
	eDB  = "https://my-json-server.typicode.com/typicode/demo/db"
	eURL = ":8080"
)

// common/env/variables
var (
	DB  string
	URL string
)

func Load() {
	DB = eDB
	URL = eURL
}

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

func (d DatabaseModelList) String() string {
	var str strings.Builder

	printer(&str, d.Posts...)
	printer(&str, d.Comments...)
	printer(&str, d.Profiles)

	return str.String()
}

func printer[T any](str *strings.Builder, base ...T) {
	for _, model := range base {
		str.WriteString(fmt.Sprintf("%v\n", model))
	}
}

// repository/repository.go
type database struct{}

type Database interface {
	FindPosts() *http.Response
}

func NewRepository() Database {
	return &database{}
}

func (database) FindPosts() *http.Response {
	response, err := http.Get(DB)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

// service/service.go
type holder struct {
	holdList Database
}

type Holder interface {
	HoldList() DatabaseModelList
}

func NewService(d Database) Holder {
	return holder{holdList: d}
}

func (h holder) HoldList() DatabaseModelList {
	databaseResult := h.holdList.FindPosts()

	var dML DatabaseModelList
	err := json.NewDecoder(databaseResult.Body).Decode(&dML)
	if err != nil {
		fmt.Println(err)
	}
	return dML
}

// controller/controller.go
type sender struct {
	sendList Holder
}

type Sender interface {
	SendList(w http.ResponseWriter, r *http.Request)
}

func NewController(h Holder) Sender {
	return &sender{sendList: h}
}

func (s sender) SendList(w http.ResponseWriter, r *http.Request) {
	serviceResult := s.sendList.HoldList()
	fmt.Fprintf(w, serviceResult.String())
}

// router/router.go
type router struct {
	routeList Sender
}

type Router interface {
	Run(url string)
}

func NewRouter(s Sender) Router {
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

// main.go
func main() {
	// common/env/variables (.env)
	Load()
	fmt.Printf("\nDB: %v\nRUL: %v\n", DB, URL)

	// repository/repository.go
	repo := NewRepository()
	fmt.Printf("\nRepository: %v\n", repo)
	resR := repo.FindPosts()
	fmt.Printf("\nRepository Response: %v\n", resR)

	// service/service.go (model/model.go)
	serv := NewService(repo)
	fmt.Printf("\nService: \n%v", serv.HoldList())
	resS := serv.HoldList()
	fmt.Printf("\nService Response: \n%v", resS)

	// controller/controller.go (model/model.go)
	cont := NewController(serv)
	fmt.Printf("\nController: %v\n", cont)

	// router/router.go
	rout := NewRouter(cont)
	fmt.Printf("\nRouter: %v\n", rout)
	rout.Run(URL)
}
