package libusb

// #include<usb.h>
import "C"

import "fmt";


func Init() (int,int)
{
    C.usb_init();

    bn := C.usb_find_busses();
    dn := C.usb_find_devices();

    return int(bn),int(dn);
}

type Info struct
{
    Bus     string;
    Device  string;
    Vid     int;
    Pid     int;
}

func Enum() []Info
{
    fmt.Printf("");

    bus := C.usb_get_busses();
    n := 0;
    for ;bus != nil; bus=bus.next
    {
        for dev := bus.devices ; dev!=nil ; dev = dev.next
        {
            n += 1;
        }
    }
    infos := make([]Info,n);

    bus = C.usb_get_busses();
    n =0;

    for ;bus != nil; bus=bus.next
    {
        busname := C.GoString(&bus.dirname[0]);

        for dev := bus.devices ; dev!=nil ; dev = dev.next
        {
            devname := C.GoString(&dev.filename[0]);

            var info Info;
            info.Bus = busname;
            info.Device = devname;
            info.Vid = int(dev.descriptor.idVendor);
            info.Pid = int(dev.descriptor.idProduct);

            infos[n] = info;
            n += 1;
        }
    }
    return infos;
}
