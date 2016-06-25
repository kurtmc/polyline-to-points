package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("test")
	s, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(s))
}
