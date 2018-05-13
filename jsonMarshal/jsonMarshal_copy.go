package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Book struct{
	Title string
	Price int32
}

var books = []Book{
	{"三生三世",100},
	{"朝阳",92},
	{"科比传",30},
}

func main() {
	data, err := json.MarshalIndent(books, "", "    ")
	if err != nil{
		log.Fatalf("marshal data failed ,%s", err)
	}

	fmt.Printf("%s\n", data)
}
