package sample

import (
	"github.com/rob-bender/go-test/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewKeyboard() *pb.Keyboard {
	var keyboard *pb.Keyboard = &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCpu() *pb.CPU {
	var brand string = randomCPUBrand()
	var name string = randomCPUName(brand)
	var numberCores int = randomInt(2, 8)
	var numberThreads int = randomInt(numberCores, 12)
	var minGhz float64 = randomFloat64(2.0, 3.5)
	var maxGhz float64 = randomFloat64(minGhz, 5.0)

	var cpu *pb.CPU = &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGpu() *pb.GPU {
	var brand string = randomGPUBrand()
	var name string = randomGPUName(brand)
	var minGhz float64 = randomFloat64(1.0, 1.5)
	var maxGhz float64 = randomFloat64(minGhz, 2.0)
	var memory *pb.Memory = &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	var gpu *pb.GPU = &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

func NewRam() *pb.Memory {
	var memory *pb.Memory = &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return memory
}

func NewSSD() *pb.Storage {
	var ssd *pb.Storage = &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(129, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

func NewHDD() *pb.Storage {
	var ssd *pb.Storage = &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return ssd
}

func NewScreen() *pb.Screen {
	var screen *pb.Screen = &pb.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}

func NewLaptor() *pb.Laptor {
	var brand string = randomLaptorBrand()
	var name string = randomLaptorName(brand)

	var laptor *pb.Laptor = &pb.Laptor{
		Id:       randomId(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCpu(),
		Ram:      NewRam(),
		Grus:     []*pb.GPU{NewGpu()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptor_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2023)),
		UpdatedAt:   timestamppb.Now(),
	}
	return laptor
}
