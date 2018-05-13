package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct {
	Title string
	Year int `json:"release"`
	Color bool `json:"color,omitempty"`
	Actors []string
}
var movies = []Movie{
	{Title:"Casablanca",Year:1942,Color:false,Actors:[]string{"Humphrey Bogart","Ingrid Bergman"}},
	{Title:"Casablanca_2",Year:1943,Color:true,Actors:[]string{"Humphrey Bogart2","Ingrid Bergman2"}},
	{Title:"Casablanca_3",Year:1944,Color:false,Actors:[]string{"Humphrey Bogart3","Ingrid Bergman3"}},
	{Title:"Casablanca_4",Year:1945,Color:true,Actors:[]string{"Humphrey Bogart4","Ingrid Bergman4"}},
}

func main() {
	data, err := json.MarshalIndent(movies,"","    ")
	//data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json marshaling falled:%s", err)
	}

	fmt.Printf("%s\n", data)
}
