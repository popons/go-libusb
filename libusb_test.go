package libusb

import "fmt"
import "os"
import "testing";

type method func();

func enum()
{

    Init();

    //buf := make([]byte,1024);

    for i,info := range Enum()
    {
        fmt.Printf("======================================================\n");
        fmt.Printf(" %10d : BUS:%s DEVICE:%s VID:%04x PID:%04x\n",i,info.Bus,info.Device,info.Vid,info.Pid);
        dev := Open(info.Vid,info.Pid);
        if dev != nil
        {
            fmt.Printf(" Vendor     : %s\n",dev.Vendor());
            fmt.Printf(" Product    : %s\n",dev.Product());
            fmt.Printf(" Last Error : %s\n",dev.LastError());
//            l := dev.BulkRead(1,buf);
//            fmt.Printf(" Last Error : %s\n",dev.LastError());
//            fmt.Printf("%d read\n",l);
//            if l>=0{
//                for _,d := range buf[0:l]{
//                    fmt.Printf("%08X ",d);
//                }
//            }

            dev.Close();
        }else{
            os.Exit(1);
        }
    }
}

func conf()
{
    vid , pid := 0x04b4,0x8613;

    Init();

    device := Open(vid,pid);
    println("dev=",device);
    println("dev.bus=",device.Bus);
    println("dev.dev=",device.Device);
    println("dev.handle=",device.handle);
    fmt.Printf(" Last Error : %s\n",device.LastError());
    var r int;
    r = device.Configuration(1);
    println("Configuration=" , r);
    fmt.Printf(" Last Error : %s\n",device.LastError());
    device.Interface(0);
    println("Interface=" , r);
    fmt.Printf(" Last Error : %s\n",device.LastError());
    device.Close();
}


var methods = []method {
//enum,
conf,
};


func TestAll(t *testing.T)
{
    for i,m := range methods
    {
        println("==============================================");
        println("========= test ",i);
        println("==============================================");
        m();
    }
}
