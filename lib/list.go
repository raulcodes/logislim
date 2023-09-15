package lib

import (
	"C"
	"fmt"
	"os"

	hid "github.com/sstallion/go-hid"
	"github.com/spf13/cobra"
)

func ListDevices(cmd *cobra.Command, args []string) {
	devices := map[uint16]*hid.DeviceInfo{}
	hid.Enumerate(uint16(VENDOR_ID), hid.ProductIDAny, func(info *hid.DeviceInfo) error {
		devices[info.ProductID] = info
		return nil
	})

	for _, d := range devices {
		fmt.Printf("%s: ID %04x:%04x %s %s\n",
		d.Path,
		d.VendorID,
		d.ProductID,
		d.MfrStr,
		d.ProductStr)
	}

	if len(devices) == 0 {
		fmt.Println("No compatible devices found")
		os.Exit(0)
	}
}