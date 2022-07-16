package mydict

import (
	"errors"
)

//Dictionary type
type Dictionary map[string]string

var errorNotFount = errors.New("not found")
var errorAlreadyHas = errors.New("this word is already has")

//search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFount
}

func (d Dictionary) Add(word, des string) error {
	_, err := d.Search(word)
	if err == errorNotFount {
		d[word] = des
		return nil
	} else {
		return errorAlreadyHas
	}

}
