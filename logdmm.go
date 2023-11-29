package main

import (
	"fmt"
	"log"
	"os"

	// project imports
	dmmIO "github.com/rDybing/logDMM/dmmIO"
)

func main() {
	var dmm dmmIO.DMMT
	if err := dmm.LoadConfig(); err != nil {
		log.Fatalf("dmm config file error, exiting! %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Model: %s\nVID: %x - PID: %x\n", dmm.Model, dmm.VID, dmm.PID)
	//dmm.Connect()
}
