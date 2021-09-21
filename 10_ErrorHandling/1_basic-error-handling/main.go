package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Marshalled slice:", string(bs))

	p2 := new(person)
	err = json.Unmarshal(bs, &p2)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Unmarshalled slice: ", p2)
}
