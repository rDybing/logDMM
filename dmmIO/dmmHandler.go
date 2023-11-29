package dmmIO

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/google/gousb"

	// project imports
	fileIO "github.com/rDybing/logDMM/fileIO"
)

type DMMT struct {
	Model string
	PID   uint16
	VID   uint16
}

// **************************************************** Init ***********************************************************

// LoadConfig populates the DMMT type with the selected config
func (d *DMMT) LoadConfig() error {
	// replace with filename from select at first startup/config menu of available
	// configs in the config folder
	const configFile = "./ConfigDMM/uni-t_ut161d.json"
	f, err := fileIO.LoadFile(configFile)
	if err != nil {
		return fmt.Errorf("error, could not load dmm config file: %s\n%v", configFile, err)
	}
	type dConvertT struct {
		Model string
		PID   string
		VID   string
	}
	var dcT dConvertT
	if err := json.Unmarshal(f, &dcT); err != nil {
		return fmt.Errorf("error, could not unmarshal %s: %v", configFile, err)
	}
	// convert string hex values from file to uint16
	pid, err := stringToUInt16(dcT.PID)
	if err != nil {
		return err
	}
	vid, err := stringToUInt16(dcT.VID)
	if err != nil {
		return err
	}
	// transfer from temp
	d.Model = dcT.Model
	d.PID = pid
	d.VID = vid
	return nil
}

// **************************************************** Connector ******************************************************

// Connect sets up the USB comms to the given DMM as defined in the DMMT type
func (d DMMT) Connect() {
	// open up usb connection
	ctx := gousb.NewContext()
	defer ctx.Close()
	dev, err := ctx.OpenDeviceWithVIDPID(0x06a3, 0x0d67)
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

// **************************************************** Helpers ********************************************************

func stringToUInt16(s string) (uint16, error) {
	s = "0x" + s
	out, err := strconv.ParseUint(s, 0, 16)
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to hex uint16 value: %v", s, err)
	}
	fmt.Printf("%s:%d\n", s, out)
	return uint16(out), nil
}
