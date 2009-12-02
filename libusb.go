package libusb

// #include<usb.h>
import "C"

//import "fmt";


func Init() (int,int)
{
    C.usb_init();

    bn := C.usb_find_busses();
    dn := C.usb_find_devices();

    return int(bn),int(dn);
}

