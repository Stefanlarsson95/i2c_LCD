package main

import (
	"fmt"
	"github.com/i2c_LCD"
	"log"
	"net"
)
/* TODO:
	 Solve i2c-com error resulting in unknown display-error. Suspected to be HW related.
*/

// 	Adapted from http://golang-examples.tumblr.com/post/99458329439/get-local-ip-addresses

func main() {
	// 16x2 LCD-screen @ /dev/i2c-1 #0x27
	i, err := i2c.New(0x27, 1)
	check(err)

	// Releases i2c-buss when done.
	defer i.Close()
	lcd, err := i2c.NewLcd(i, 2, 1, 0, 4, 5, 6, 7, 3)
	check(err)
	lcd.BacklightOn()
	lcd.Clear()
	// Local IP-address(es)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				lcd.Home()
				ip := ipnet.IP.String()
				lcd.SetPosition(1, 0)
				fmt.Fprint(lcd, "---=: MyIP :=---")
				lcd.SetPosition(2, (16 - byte(len(ip)))/2)
				fmt.Fprint(lcd, ip)
				fmt.Println("My ip: " + ip)
			}
		}
	}
}


func check(err error) {
	if err != nil { log.Fatal(err) }
}