package model

import (
	"fmt"
	"strings"
)

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
