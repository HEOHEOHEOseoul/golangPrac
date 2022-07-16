package main

import (
	"fmt"
	"second/src/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary.Add("hi", "안녕")
	fmt.Println(dictionary)
	errr := dictionary.Update("hii", "인사")

	if errr == nil {
		fmt.Println(dictionary)
	} else {
		fmt.Println(errr)
	}
}
