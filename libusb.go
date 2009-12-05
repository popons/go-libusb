package libusb

// #include<usb.h>
import "C"
import "unsafe"

import "fmt";
//import "container/list";


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
    for ;bus != nil; bus=bus.next {
        for dev := bus.devices ; dev!=nil ; dev = dev.next {
            n += 1;
        }
    }
    infos := make([]Info,n);

    bus = C.usb_get_busses();
    n =0;

    for ;bus != nil; bus=bus.next {
        busname := C.GoString(&bus.dirname[0]);

        for dev := bus.devices ; dev!=nil ; dev = dev.next {
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
    *Info;
    handle *C.usb_dev_handle;
    descriptor _Cstruct_usb_device_descriptor;
    timeout int;
}

/// open usb device with info
//func Open(info Info) (*Device)
//{
//    var rdev *Device = nil;
//
//    for bus := C.usb_get_busses() ; bus != nil ; bus=bus.next
//    {
//        for dev := bus.devices ; dev!=nil ; dev = dev.next
//        {
//            if int(dev.descriptor.idVendor)  == info.Vid &&
//               int(dev.descriptor.idProduct) == info.Pid
//            {
//                h := C.usb_open(dev);
//                rdev = &Device{&info,h,dev.descriptor,10000};
//                return rdev;
//            }
//        }
//    }
//    return rdev;
//}
/// open usb device with info
func Open(vid , pid int) (*Device)
{
    for bus := C.usb_get_busses() ; bus != nil ; bus=bus.next
    {
        for dev := bus.devices ; dev!=nil ; dev = dev.next
        {
            if int(dev.descriptor.idVendor)  == vid &&
               int(dev.descriptor.idProduct) == pid
            {
                h := C.usb_open(dev);
                rdev := &Device{
                    &Info{
                        C.GoString(&bus.dirname[0]),
                        C.GoString(&dev.filename[0]),vid,pid},
                    h, dev.descriptor,10000};
                return rdev;
            }
        }
    }
    return nil;
}

func (dev *Device) Close() int
{
    r := int(C.usb_close(dev.handle));
    dev.handle = nil;
    return r;
}
func (dev *Device) String(key int) string
{
    buf := make([]C.char,256);

    C.usb_get_string_simple(
            dev.handle,
            C.int(key),
            &buf[0],
            C.size_t(len(buf)));

    return C.GoString(&buf[0]);

}
func (self *Device) Vendor() string
{
    return self.String(int(self.descriptor.iManufacturer));
}
func (self *Device) Product() string
{
    return self.String(int(self.descriptor.iProduct));
}
func LastError() string
{
    return C.GoString(C.usb_strerror());
}
func (*Device) LastError() string
{
    return LastError();
}

func (self *Device) BulkWrite(ep int ,dat []byte) int
{
    return int( C.usb_bulk_write(self.handle,
                                    C.int(ep),
                                  (*C.char)(unsafe.Pointer(&dat[0])),
                                    C.int(len(dat)),
                                    C.int(self.timeout)) );
}
func (self *Device) BulkRead(ep int ,dat []byte) int
{
    return int( C.usb_bulk_read( self.handle,
                                    C.int(ep),
                                  (*C.char)(unsafe.Pointer(&dat[0])),
                                    C.int(len(dat)),
                                    C.int(self.timeout)) );
}
func (self *Device) Configuration(conf int) int
{
    return int( C.usb_set_configuration(self.handle, C.int(conf)) );
    //return int( C.usb_set_configuration( (*C.uint)(123), C.int(conf)) );
}
func (self *Device) Interface (ifc int) int
{
    return int( C.usb_claim_interface(self.handle, C.int(ifc)));
}

const (
    USB_TYPE_STANDARD   =(0x00 << 5);
    USB_TYPE_CLASS      =(0x01 << 5);
    USB_TYPE_VENDOR     =(0x02 << 5);
    USB_TYPE_RESERVED   =(0x03 << 5);
);

func (self *Device) ControlMsg (reqtype int, req int, value int, index int, dat []byte) int
{
    return int( C.usb_control_msg(    self.handle,
                                    C.int(reqtype),
                                    C.int(req),
                                    C.int(value),
                                    C.int(index),
                                  (*C.char)(unsafe.Pointer(&dat[0])),
                                    C.int(len(dat)),
                                    C.int(self.timeout)));
}


