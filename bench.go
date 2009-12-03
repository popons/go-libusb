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
    fmt.Printf("input number and enter key:");

    rd := bufio.NewReader(os.Stdin);
    moji,_ := rd.ReadString('\n');
    moji = moji[0:len(moji)-1];
    sel,_ := Atoi(moji);

    dev := libusb.Open(libusb.Enum()[sel]);

    fmt.Printf("Vendor     : %s\n",dev.Vendor());
    fmt.Printf("Product    : %s\n",dev.Product());

    dev.Close();

    fmt.Printf("Last Error : %s\n",dev.LastError());
}

