package main

import (
	"fmt"
	"io"
	"os"
	// "time"

	"github.com/2003Aditya/ComputerNetwork/link"
	"github.com/2003Aditya/ComputerNetwork/protocol"
	"github.com/2003Aditya/ComputerNetwork/utils"
)

const (
    StartFlagSize = 8
    SrcSize =2
    DesSize = 2
    TTLSize = 4
    SeqNumSize = 8
    AckNumSize = 8
    FlagSize = 3

    PayLoadSize = 8
    ParitySize = 1

    PacketSize = SrcSize + DesSize + TTLSize + SeqNumSize + AckNumSize + FlagSize + PayLoadSize
    FrameSize = StartFlagSize + PacketSize + ParitySize + EndFlagSize
    EndFlagSize = 8
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
            fmt.Printf("representing 52 bit slice %c\n", frame)
            FrameSize := link.Count(frame)
            fmt.Println("Size of the frame ", FrameSize)

            // startFlag := frame[:StartFlagSize]
            //          src := frame[StartFlagSize : StartFlagSize + SrcSize]
            //          des := frame[StartFlagSize + SrcSize : StartFlagSize + SrcSize + DesSize]
            //          ttl := frame[StartFlagSize + SrcSize + DesSize : StartFlagSize + SrcSize + DesSize + TTLSize]
            //          seq := frame[StartFlagSize + SrcSize + DesSize : StartFlagSize + SrcSize + DesSize + TTLSize + SeqNum :
            //          ack := frame[]
            //          flags := frame[]
            //
            //
            //          payload := frame[StartFlagSize + SrcSize + DesSize + TTLSize: StartFlagSize + SrcSize + DesSize + TTLSize + PayLoadSize]
            //          parityBit := byte(frame[StartFlagSize + PacketSize ])
            //          endFlag := frame[StartFlagSize + PacketSize + ParitySize : StartFlagSize + PacketSize + ParitySize + EndFlagSize]
            //
            //          packet := frame[StartFlagSize : StartFlagSize + PacketSize]


            start := 0
            startFlag := frame[start : start+ StartFlagSize]
            start += StartFlagSize

            src := frame[start : start + SrcSize]
            start += SrcSize

            des := frame[start : start + DesSize]
            start += DesSize

            ttl := frame[start : start + TTLSize]
            start += TTLSize

            seq := frame[start : start + SeqNumSize]
            start += SeqNumSize

            ack := frame[start : start + AckNumSize]
            start += AckNumSize

            flags := frame[start : start + FlagSize]
            start += FlagSize

            payload := frame[start : start + PayLoadSize]
            start += PayLoadSize

            parityBit := frame[start]
            start ++

            endFlag := frame[start : start + EndFlagSize]




            fmt.Printf("startbit: %c\n", startFlag)
            // fmt.Printf("packet: %c\n", packet)
            fmt.Printf("src: %c\n", src)
            fmt.Printf("des: %c\n", des)
            fmt.Printf("TTL: %c\n", ttl)
            fmt.Printf("payload: %c\n", payload)

            fmt.Printf("seq: %c\n", seq)
            fmt.Printf("ack: %c\n", ack)
            fmt.Printf("flags: %c (SYN: %c, ACK: %c, FIN: %c)\n", flags, flags[0], flags[1], flags[2])
            fmt.Printf("parityBit: %c\n", parityBit)
            fmt.Printf("endFlag: %c\n", endFlag)

            checkParity := utils.CheckParity(payload, parityBit)
            if checkParity {

                packetType := utils.GetPacketTypeSimple(flags)
                fmt.Println("PAcket TYpe:", packetType)
                ctx := &protocol.PacketContext {
                    PacketType: utils.GetPacketTypeSimple(flags),
                    StartFlag: startFlag,
                    EndFlag: endFlag,
                    Src: src,
                    Des: des,
                    TTL: ttl,
                    Payload: payload,
                    Seq: seq,
                    Ack: ack,
                    Flags: flags,
                    // ParityBit: parityBit,
                }

                switch ctx.PacketType {
                case "SYN":
                    fmt.Println(seq)
                    fmt.Printf("seq: %c\n",seq)
                    protocol.HandleSYN(ctx)
                case "SYN-ACK":
                    fmt.Println("Handling SYN-ACK Packet")
                    fmt.Println("NewSrc23:", ctx.Src)
                    protocol.HandleSYNACK(ctx)
                case "ACK":
                    fmt.Println("Handling ACK Packet")
                case "FIN":
                    fmt.Println("HAndling FIN Packet")
                case "DATA":
                    dataBitsToAscii := utils.ByteToChar(payload)
                    fmt.Printf("Data received %c\n", dataBitsToAscii)
                default :
                    fmt.Println("Unknown or Malformed flags")
                }


                if err != nil {
                    fmt.Println("error in data receiving", err)
                    return
                }

            } else {
                fmt.Println("Parity check Failed - discarding frame")
            }

            fmt.Println("correct parity", checkParity)

            //checking parity if valid or not
            bitStream = bitStream[FrameSize:]
            continue


            // break
        }



        // time.Sleep(1 * time.Second)
    }

}
