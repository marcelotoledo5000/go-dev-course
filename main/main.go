package main

import (
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

func main() {
	u, err := uuid.NewV4()
	fmt.Printf("My uuid is %s\n", u)

	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	log.Printf("generated Version 4 UUID %v", u)
}
