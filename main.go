package main

import (
	"fmt"
	"log"
)

func main() {
	c, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

}
