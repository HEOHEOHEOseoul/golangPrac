package mydict

import (
	"errors"
)

//Dictionary type
type Dictionary map[string]string

var (
	errorNotFount   = errors.New("not found")
	errorAlreadyHas = errors.New("this word is already has")
)

//search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFount
}

//Add a word to the dictionary
func (d Dictionary) Add(word, des string) error {
	_, err := d.Search(word)

	switch err {
	case errorNotFount:
		d[word] = des
	case nil:
		return errorAlreadyHas
	}
	return nil

}

//Update a description of a word
func (d Dictionary) Update(word, des string) error {
	_, err := d.Search(word)

	// if err == errorNotFount {
	// 	return err
	// }
	// d[word] = des
	// return nil
	switch err {
	case errorNotFount:
		return errorNotFount
	case nil:
		d[word] = des
	}
	return nil
}

//Delete a word from Dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	if err == errorNotFount {
		return errorNotFount
	}
	delete(d, word)
	return nil
}
