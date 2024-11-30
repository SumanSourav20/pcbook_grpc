package sample

import (
	"math/rand"

	"github.com/SumanSourav20/pcbook_grpc/pb"
)

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(8) {
	case 0:
		return pb.Keyboard_QWERTZ
	case 1:
		return pb.Keyboard_AZERTY
	case 2:
		return pb.Keyboard_DVORAK
	case 3:
		return pb.Keyboard_COLEMAK
	default:
		return pb.Keyboard_QWERTY
	}
}

func randomKeyboardBacklit() bool {
	return rand.Intn(2) == 1
}

func randomRamType() pb.RAM_RAMType {
	switch rand.Intn(3) {
	case 0:
		return pb.RAM_DDR3
	case 1:
		return pb.RAM_DDR4
	default:
		return pb.RAM_DDR5
	}
}

func randomMemoryUnit() (pb.Memory_Unit, int, int) {
	switch rand.Intn(4) {
	case 0:
		return pb.Memory_KILOBYTE, 100_000, 1_000_000_000
	case 1:
		return pb.Memory_MEGABYTE, 1_000, 10_000_000
	case 2:
		return pb.Memory_GIGABYTE, 1, 10_000
	default:
		return pb.Memory_TERABYTE, 1, 10
	}
}

func randomMemory() *pb.Memory {
	unit, min, max := randomMemoryUnit()
	return &pb.Memory{
		Size: uint64(min + rand.Intn(max-min+1)),
		Unit: unit,
	}
}

func generateCPU() *pb.CPU {
	brand := randomStringfromSet("AMD", "Intel")

	var model_name string

	switch brand {
	case "AMD":
		model_name = randomStringfromSet("Ryzen 5 5500", "Ryzen 7 5700", "Ryzen 9 5800", "Ryzen 5 6500", "Ryzen 7 6700", "Ryzen 9 6800")
	default:
		model_name = randomStringfromSet("i5 11gen", "i7 11gen", "i9 11gen", "i5 12gen", "i7 12gen", "i9 12gen")
	}

	var cores uint32 = uint32(1 + rand.Intn(8))

	var l1cache uint32 = 100000
	var l2cache uint32 = 1000000
	var l3cache uint32 = 10000000
	treads := 2 * cores
	minGhz := 2 + rand.Intn(3)
	maxGhz := 4 + rand.Intn(3)

	return &pb.CPU{
		Brand:     brand,
		ModelName: model_name,
		Cores:     cores,
		L1Cache:   l1cache,
		L2Cache:   l2cache,
		L3Cache:   l3cache,
		Treads:    treads,
		MinGhz:    uint32(minGhz),
		MaxGhz:    uint32(maxGhz),
	}
}

func randomStringfromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	r := rand.Intn(n)
	return a[r]
}

func generateGPU() *pb.GPU {
	brand := randomStringfromSet("AMD", "NVDIA")

	var model_name string

	switch brand {
	case "AMD":
		model_name = randomStringfromSet("RX 5000", "RX 3000")
	default:
		model_name = randomStringfromSet("A100", "RTX 3070", "RTX 4090", "RTX 3060")
	}

	minGhz := 2 + rand.Intn(2)
	maxGhz := 4 + rand.Intn(2)

	return &pb.GPU{
		Brand:     brand,
		ModelName: model_name,
		MinGhz:    uint32(minGhz),
		MaxGhz:    uint32(maxGhz),
		Memory:    randomMemory(),
	}
}

func randomScreenPanel() pb.Screen_Panel {
	switch rand.Intn(2) {
	case 0:
		return pb.Screen_IPS
	default:
		return pb.Screen_OLED
	}
}
func generateScreen() *pb.Screen {
	return &pb.Screen{
		Resolution: &pb.Screen_Resolution{Width: 1600, Height: uint32(randomIntFromSet(1080, 1400))},
		Panel:      randomScreenPanel(),
		Size:       15,
		Multitouch: false,
	}
}

func randomIntFromSet(a ...int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	r := rand.Intn(n)
	return a[r]
}

func randomStorageDriver() pb.Storage_Driver {
	switch rand.Intn(2) {
	case 0:
		return pb.Storage_HDD
	default:
		return pb.Storage_SSD
	}
}
