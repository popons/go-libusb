package main

// #include<libusb-1.0/libusb.h>
import "C"

import "fmt";

var ctx *C.libusb_context;

func main()
{
    C.libusb_init(&ctx);
    fmt.Printf("アホの坂田\n");
}

