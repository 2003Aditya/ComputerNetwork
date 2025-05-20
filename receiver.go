package main

import (
	"fmt"
	"io"
	"os"
	"time"
    "github.com/2003Aditya/ComputerNetwork/utils"
)

const (
    StartFlagSize = 8
    EndFlagSize = 8
    SrcSize =2
    DesSize = 2
    TTLSize = 4
    PayLoadSize = 8
    ParitySize = 1

    PacketSize = SrcSize + DesSize + TTLSize + PayLoadSize
    FrameSize = StartFlagSize + PacketSize + ParitySize + EndFlagSize
)

func main() {
	file, err := os.Open("wire.txt")
	if err != nil {
		fmt.Println("Failed to open wire:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1)
	var offset int64 = 0
	var bitStream []byte

	for {
		// Move the file pointer to the current offset
		_, seekErr := file.Seek(offset, 0)
		if seekErr != nil {
			fmt.Println("Seek error:", seekErr)
			return
		}

		// Read 1 byte
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading wire:", err)
			return
		}

		if n > 0 {
			bitStream = append(bitStream, buffer[0])
			fmt.Printf("Received bit: %c\n", buffer[0])
			offset++
		}

		if len(bitStream) >= FrameSize {
			frame := bitStream[:FrameSize]
			fmt.Printf("respresenting 33 bit slice %c\n", frame)

			startFlag := frame[:StartFlagSize]
            packet := frame[StartFlagSize : StartFlagSize + PacketSize]
            src := frame[StartFlagSize : StartFlagSize + SrcSize]
            des := frame[StartFlagSize + SrcSize : StartFlagSize + SrcSize + DesSize]
            ttl := frame[StartFlagSize + SrcSize + DesSize : StartFlagSize + SrcSize + DesSize + TTLSize]
            payload := frame[StartFlagSize + SrcSize + DesSize + TTLSize: StartFlagSize + SrcSize + DesSize + TTLSize + PayLoadSize]
            parityBit := byte(frame[StartFlagSize + PacketSize + ParitySize])
            endFlag := frame[StartFlagSize + PacketSize + ParitySize : StartFlagSize + PacketSize + ParitySize + EndFlagSize]

            fmt.Printf("startbit: %c\n", startFlag)
            fmt.Printf("packet: %c\n", packet)
            fmt.Printf("src: %c\n", src)
            fmt.Printf("des: %c\n", des)
            fmt.Printf("TTL: %c\n", ttl)
            fmt.Printf("payload: %c\n", payload)

            fmt.Printf("parityBit: %c\n", parityBit)
            fmt.Printf("endFlag: %c\n", endFlag)

            checkParity := utils.CheckParity(packet, parityBit)
            if checkParity {
                dataBitsToAscii := utils.ByteToChar(payload)
                fmt.Printf("Data received %c\n", dataBitsToAscii)
            } else {
                fmt.Println("Parity check Failed - discarding frame")
            }

            fmt.Println("correct parity", checkParity)

            //checking parity if valid or not
            bitStream = bitStream[FrameSize:]
            continue


            // break
		}



		time.Sleep(1 * time.Second)
	}

}
