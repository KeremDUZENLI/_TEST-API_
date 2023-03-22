package model

type ApiResponseParse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ApiResponse struct {
	Liste []ApiResponseParse `json:"posts"`
}
