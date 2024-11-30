package main

import (
	"fmt"

	"github.com/SumanSourav20/pcbook_grpc/pb"
	"github.com/SumanSourav20/pcbook_grpc/sample"
	"github.com/SumanSourav20/pcbook_grpc/serializer"
)

func main() {
	laptop := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop, "new_laptop.bin")

	if err != nil {
		fmt.Println(err.Error())
	}

	var newLaptop pb.Laptop

	serializer.ReadProtobufFromBinaryFile("new_laptop.bin", &newLaptop)

	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(newLaptop)

}
