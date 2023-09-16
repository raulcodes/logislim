package lib

import (
	"C"
	"fmt"
	"math"
	"os"

	hid "github.com/sstallion/go-hid"
	"github.com/spf13/cobra"
)

func AllDevicesON(cmd *cobra.Command, args []string, brightness int) {
	AllDevicesAction(turnON, brightness)
}

func AllDevicesOFF(cmd *cobra.Command, args []string, brightness int) {
	AllDevicesAction(turnOFF, -1)
}

func AllDevicesAction(fn func(d *hid.Device, brightness int) error, brightness int) {
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
	
		err := fn(d, brightness)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func turnON(device *hid.Device, brightness int) error {
	_, err := device.Write([]byte{
		0x11, 0xff, 0x04, 0x1c, TURN_ON,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 
		0x00, 0x00, 0x00, 0x00, 0x00,
	})

	b := adjustBrightness(brightness)
	_, err = device.Write([]byte{
		0x11, 0xff, 0x04, 0x4c, 0x00,
		b, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 
		0x00, 0x00, 0x00, 0x00, 0x00,
	})
	return err
}

func turnOFF(device *hid.Device, b int) error {
	_, err := device.Write([]byte{
		0x11, 0xff, 0x04, 0x1c, TURN_OFF,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 
		0x00, 0x00, 0x00, 0x00, 0x00,
	})
	return err
}

func adjustBrightness(brightness int) byte {
	minBrightness := float64(0x14)
	maxBrightness := float64(0xfa)

	if brightness < 0 {
		brightness = 0
	}
	if brightness > 100 {
		brightness = 100
	}

	value := minBrightness + ((float64(brightness) / 100) * (maxBrightness - minBrightness))
	return byte(math.Floor(float64(value)))
}