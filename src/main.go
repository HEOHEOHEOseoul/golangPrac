package main

import (
	"fmt"
	"second/src/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary.Add("hi", "안녕")
	fmt.Println(dictionary.Search("hi"))

}
