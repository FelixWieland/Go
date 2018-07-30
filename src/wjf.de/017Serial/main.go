package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func arduinoReader(port string, baudrate int, readSize int, maskByteStart byte, maskByteEnd byte) string {
	stringBuffer := ""
	started := false
	ended := false
	c := &serial.Config{Name: port, Baud: baudrate}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	s.Flush()
	for !started || !ended {
		buf := make([]byte, readSize)
		_, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		for char := range buf {
			if buf[char] == 10 {
				started = true
			}
			if buf[char] == 9 {
				ended = true
			}
			if started || ended {
				if buf[char] > 32 && buf[char] < 127 {
					stringBuffer = stringBuffer + string(buf[char])
				}
			}
		}
	}
	return stringBuffer
}

func main() {
	output := arduinoReader("COM3", 9600, 12, 10, 9)
	fmt.Println(output)
}
