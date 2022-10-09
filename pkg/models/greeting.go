package models

type Greeting struct {
	Id      int64  `json:"id"`
	Message string `json:"greeting"`
	Author  string `json:"author"`
	Desc    string `json:"desc"`
}
