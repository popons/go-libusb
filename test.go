package main

import "libusb"
import "fmt"

func main()
{
    b,d := libusb.Init();
    fmt.Printf("bus-num=%d\ndev-num=%d\n",b,d);
}
