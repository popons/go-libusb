package main

import "libusb"
import "fmt"
import "os"
import . "strconv"
import "bufio"

func main()
{
    libusb.Init();
    for i,d := range libusb.Enum()
    {
        fmt.Printf("%3d :  BUS:%s DEVICE:%s VID:%04x PID:%04x\n",i,d.Bus,d.Device,d.Vid,d.Pid);
    }
    fmt.Printf("どれを見るか\n");
    fmt.Printf(":");

    rd := bufio.NewReader(os.Stdin);
    moji,_ := rd.ReadString('\n');
    moji = moji[0:len(moji)-1];
    fmt.Printf("sel=%s\n",moji);
    sel,_ := Atoi(moji);
    fmt.Printf("sel=%d\n",sel);

    dev := libusb.Open(libusb.Enum()[sel]);

    fmt.Printf("dev %04x %04x\n",dev.Vid,dev.Pid);

    dev.Close();
}
