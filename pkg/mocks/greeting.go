package mocks

import "github.com/franktore/go-learn/pkg/models"

var Seq_id int64 = 1

var Greetings = []models.Greeting{
	{
		Id:      1,
		Message: "Hi, %v. Welcome!",
		Author:  "Rick",
		Desc:    "A nice greeting",
	},
}
