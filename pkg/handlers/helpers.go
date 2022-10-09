package handlers

import (
	"errors"

	"github.com/franktore/go-learn/pkg/mocks"
	"github.com/franktore/go-learn/pkg/models"
)

func getGreetingById(id int) (int, models.Greeting, error) {
	var greeting models.Greeting
	var idx int

	for i := range mocks.Greetings {
		if mocks.Greetings[i].Id == int64(id) {
			greeting = mocks.Greetings[i]
			idx = i
			break
		}
	}

	if greeting == (models.Greeting{}) {
		return -1, (models.Greeting{}), errors.New("No greeting found for id")
	}

	return idx, greeting, nil
}

func removeByIdx(s []models.Greeting, i int) ([]models.Greeting, error) {
	itemLength := len(s)
	if i >= itemLength || i < 0 {
		return s, errors.New("Index out of range")
	}

	// replace idx with last item
	// then remove last item
	if i != itemLength-1 {
		s[i] = s[itemLength-1]
	}
	s = s[:itemLength-1]
	return s, nil
}
