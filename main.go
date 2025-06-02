package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	protobuffin "protobuffin/gen" // adjust based on your module and directory
)

func main() {
	p := &protobuffin.Person{
		Name: "Alice",
		Age:  30,
	}

	// Marshal to binary
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// Unmarshal from binary
	newPerson := &protobuffin.Person{}
	if err := proto.Unmarshal(data, newPerson); err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Printf("Unmarshaled: %+v\n", newPerson)
}
