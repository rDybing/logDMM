package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/gousb"
)

type DMM struct {
	Model string
	PID   uint16
	VID   uint16
}

func main() {
	var dmm DMM
	if err := dmm.LoadConfig(); err != nil {
		log.Fatalf("dmm config file error, exiting! %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Model: %s\nVID: %x - PID: %x\n", dmm.Model, dmm.VID, dmm.PID)
	dmm.Connect()
}

func (d *DMM) LoadConfig() error {
	// replace with filename from select at first startup/config menu of available
	// configs in the config folder
	const configFile = "./Config/uni-t_ut161d"
	fmt.Println(configFile)
	f, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("Error, could not load keymap file: %s\n%v", configFile, err)
	}
	if err := json.Unmarshal(f, &d); err != nil {
		return fmt.Errorf("Error, could not unmarshal %s: %v", configFile, err)
	}
	return nil
}

func (d DMM) Connect() {
	// open up usb connection
	ctx := gousb.NewContext()
	defer ctx.Close()
	dev, err := ctx.OpenDeviceWithVIDPID(0x06a3, 0x0d67)
	defer dev.Close()
	if err != nil || dev == nil {
		log.Fatalf("OpenDevice failed - ensure it is connected: %v\n", err)
	}
	fmt.Printf("Device opened: %v\n", dev)
	if err := dev.SetAutoDetach(true); err != nil {
		log.Fatalf("Could not detach device: %v", err)
	}
	fmt.Println("Device auto-detached")
	// grab device, interface and endpoint
	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("Opening %s.Config(1) failed: %v\n", dev, err)
	}
	fmt.Printf("Device config read: %v\n", cfg)
	defer cfg.Close()
	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("Opening %s.Interface(0, 0) failed: %v\n", cfg, err)
	}
	fmt.Println("Interface opened")
	defer intf.Close()
	epIn, err := intf.InEndpoint(1)
	if err != nil {
		log.Fatalf("Opening %s.InEndpoint(1) failed: %v\n", intf, err)
	}
	for {
		buf := make([]byte, epIn.Desc.MaxPacketSize)
		inBytes, err := epIn.Read(buf)
		if err != nil {
			fmt.Printf("Read returned an error: %v\n", err)
		}
		if inBytes == 0 {
			log.Fatalf("IN endpoint 1 returned 0 bytes of data.\n")
		}
		var outBytes [3]uint8
		for i := 0; i < 3; i++ {
			outBytes[i] = uint8(buf[i])
		}
		// display output
	}
}
