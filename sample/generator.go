package sample

import (
	"github.com/SumanSourav20/pcbook_grpc/pb"
	"github.com/google/uuid"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomKeyboardBacklit(),
	}
}

func NewCPU() *pb.CPU {
	return generateCPU()
}

func NewGPU() *pb.GPU {
	return generateGPU()
}

func NewRAM() *pb.RAM {
	return &pb.RAM{
		RamType: randomRamType(),
		Memory:  randomMemory(),
	}
}

func NewScreen() *pb.Screen {
	return generateScreen()
}

func NewStorage() *pb.Storage {
	return &pb.Storage{
		Driver: randomStorageDriver(),
		Memory: randomMemory(),
	}
}

func NewLaptop() *pb.Laptop {
	return &pb.Laptop{
		Cpu:       NewCPU(),
		Gpu:       []*pb.GPU{NewGPU()},
		Ram:       NewRAM(),
		Storage:   []*pb.Storage{NewStorage()},
		Screen:    NewScreen(),
		Keyboard:  NewKeyboard(),
		ModelName: "Laptop with Random parts",
		Brand:     "My own Brand",
		Id:        uuid.New().String(),
	}
}
