package libusb

// #include<usb.h>
import "C"

import "fmt";
import "io";


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

type Device struct
{
    Info;
    handle *C.usb_dev_handle;
    //descriptor C.usb_device_descriptor;
    io.ReadWriteCloser;
}

func Open(info Info) (*Device)
{
    var rdev *Device = nil;

    for bus := C.usb_get_busses() ; bus != nil ; bus=bus.next
    {
        for dev := bus.devices ; dev!=nil ; dev = dev.next
        {
            if int(dev.descriptor.idVendor)  == info.Vid &&
               int(dev.descriptor.idProduct) == info.Pid
            {
                h := C.usb_open(dev);
                rdev = new(Device);
                rdev.Info = info;
                rdev.handle = h;
//                rdev.descriptor = dev.descriptor;
                return rdev;
            }
        }
    }
    return rdev;
}

func (dev *Device) Close() int
{
    r := int(C.usb_close(dev.handle));
    dev.handle = nil;
    return r;
}
//func (dev *Device) Product() string
//{
//    buf := make([]byte,1024);
//
//    C.usb_get_string_simple(dev.handle,dev.descriptor.iProduct,buf,len(buf));
//
//    return string(buf);
//}
