package main

import (
	"fmt"
	"log"

	pb "github.com/natemollica-nm/protobuffin/gen"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Name: "Alice",
		Age:  30,
	}

	// Marshal to binary
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// Unmarshal from binary
	newPerson := &pb.Person{}
	if err := proto.Unmarshal(data, newPerson); err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Printf("Unmarshaled: %+v\n", newPerson)
}
