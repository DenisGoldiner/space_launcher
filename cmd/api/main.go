package main

import (
	"log"

	"github.com/DenisGoldiner/space_launcher/platform"
)

func main() {
	if err := platform.RunApp(); err != nil {
		log.Fatal(err)
	}
}
