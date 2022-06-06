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

	services, err := c.ListServices("")
	if err != nil {
		log.Fatal(err)
	}

	for _, service := range services {
		fmt.Println(service.Name)
	}

}
