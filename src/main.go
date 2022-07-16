package main

import (
	"fmt"
	"second/src/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "firstWord"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
