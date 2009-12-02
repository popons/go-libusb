package main

import "libusb"
import "fmt"

func main()
{
    b,d := libusb.Init();
    fmt.Printf("%d,%d\n",b,d);
}
