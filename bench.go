package main

import "libusb"
import "fmt"

func main()
{
    libusb.Init();
    for _,d := range libusb.Enum()
    {
        fmt.Printf("BUS:%s DEVICE:%s VID:%04x PID:%04x\n",d.Bus,d.Device,d.Vid,d.Pid);
    }
}
