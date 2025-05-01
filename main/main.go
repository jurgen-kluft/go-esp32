package main

import (
	"fmt"
	"os"

	"github.com/jurgen-kluft/go-esp32/device"
)

func main() {
	ctrl := device.NewController()
	if err := ctrl.Setup(); err != nil {
		fmt.Fprintf(os.Stderr, "Error on Setup: %s", err.Error())
	}
	for {
		if err := ctrl.Loop(); err != nil {
			fmt.Fprintf(os.Stderr, "Error on Loop: %s", err.Error())
		}
	}
}
