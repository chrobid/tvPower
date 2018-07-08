package main

import (
	"log"

	"github.com/tarm/serial"
)

func main() {
	const deviceName = "/dev/ttyUSB0"

	c := &serial.Config{Name: deviceName, Baud: 9600}
	tv, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n, err := tv.Write([]byte("POWR0   "))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = tv.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%q", buf[:n])
}
