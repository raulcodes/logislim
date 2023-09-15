package lib

import (
	"C"
	"fmt"
	"os"

	hid "github.com/sstallion/go-hid"
	"github.com/spf13/cobra"
)

func AllDevicesON(cmd *cobra.Command, args []string) {
	AllDevicesAction(turnON)
}

func AllDevicesOFF(cmd *cobra.Command, args []string) {
	AllDevicesAction(turnOFF)
}

func AllDevicesAction(fn func(*hid.Device) error) {
	devices := map[uint16]*hid.DeviceInfo{}
	hid.Enumerate(uint16(VENDOR_ID), hid.ProductIDAny, func(info *hid.DeviceInfo) error {
		devices[info.ProductID] = info
		return nil
	})

	if len(devices) == 0 {
		fmt.Println("No compatible devices found")
		os.Exit(0)
	}

	for _, device := range devices {
		d, e := hid.OpenFirst(device.VendorID, device.ProductID)
		if e != nil {
			fmt.Println(e.Error())
		}
	
		err := fn(d)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func turnON(device *hid.Device) error {
	_, err := device.Write([]byte{
		0x11, 0xff, 0x04, 0x1c, TURN_ON,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 
		0x00, 0x00, 0x00, 0x00, 0x00,
	})
	return err
}

func turnOFF(device *hid.Device) error {
	_, err := device.Write([]byte{
		0x11, 0xff, 0x04, 0x1c, TURN_OFF,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 
		0x00, 0x00, 0x00, 0x00, 0x00,
	})
	return err
}