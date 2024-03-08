package main

import (
	"math/rand"

	"github.com/Lemos1347/inteli-modulo-9-prova-1/internal/repository"
)

func emulateFreezer() int {
	temp := rand.Intn(41)

	return temp * -1
}

func enumlateGeladeira() int {
	return rand.Intn(15)
}

func main() {

	freezer1 := repository.CreateSensor("lj01f01", "freezer", emulateFreezer)
	freezer2 := repository.CreateSensor("lj02f01", "freezer", emulateFreezer)
	geladeira1 := repository.CreateSensor("lj01g01", "geladeira", enumlateGeladeira)
	gelareira2 := repository.CreateSensor("lj02g01", "geladeira", enumlateGeladeira)

	emulator1 := repository.NewEmulator(freezer1)
	emulator2 := repository.NewEmulator(freezer2)
	emulator3 := repository.NewEmulator(geladeira1)
	emulator4 := repository.NewEmulator(gelareira2)

	emulator1.Start()
	go emulator2.Start()
	go emulator3.Start()

	emulator4.Start()

}
