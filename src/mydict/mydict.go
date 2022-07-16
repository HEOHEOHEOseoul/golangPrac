package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var errorNotFount = errors.New("not found")

//search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFount
}
