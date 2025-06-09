package protocol

import (
	"fmt"

	"github.com/2003Aditya/ComputerNetwork/link"
	"github.com/2003Aditya/ComputerNetwork/network"
	"github.com/2003Aditya/ComputerNetwork/transport"
	"github.com/2003Aditya/ComputerNetwork/utils"
)

func ParseFrame() {
}

func HandleSYN(src, des, ttl, seq, ack []byte) {

	msg := "d"
	msgB := []byte(msg)

    msgByte, parity := utils.MsgToByte(msgB)
    newMsg := utils.AsciiBytesToDigitBytes(msgByte)
    fmt.Printf("newMsg %v\n", newMsg)
    fmt.Printf("parity %v\n", parity)




	fmt.Printf("Syn Received from %c to %c\n", src, des)
	fmt.Printf("Initial sequence number %v\n", seq)

	newTTL := utils.AsciiBytesToDigitBytes(ttl)
	fmt.Printf("New TTL %v\n", newTTL)

	wuhuuSeq := utils.AsciiBytesToDigitBytes(seq)
	fmt.Printf("wuhuuSeq: %v\n", wuhuuSeq)

	newAck := utils.AsciiBytesToDigitBytes(ack)
	fmt.Printf("newAck %v\n", newAck)

	newSrc := utils.AsciiBytesToDigitBytes(des)
	fmt.Printf("newSrc: %v\n", newSrc)

	newDes := utils.AsciiBytesToDigitBytes(src)
	fmt.Printf("newdes: %v\n", newDes)

	newSeq := utils.Increment(wuhuuSeq)
	fmt.Printf("NewSeq: %v\n", newSeq)

	startFlag := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	endFlag := []byte{0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println("creating ACK-SYN Packet")
	fmt.Printf("Source : %s, Destination: %s\n", src, des)

	newTcpSegment := transport.Tcp(newSeq, newAck, newMsg, true, true, false)
	fmt.Printf("newTcpSegment %v\n", newTcpSegment)

	newPacket := network.Packet(newSrc, newDes, newTTL, newTcpSegment)

	var frame []byte

	frame = append(frame, startFlag...)
	fmt.Println("StartFlag:", startFlag)
	fmt.Println("frame after startFlag", frame)

	frame = append(frame, newPacket...)
	fmt.Printf("Frame after adding Packet %v\n", frame)
    frame = append(frame, byte(parity))
    fmt.Printf("Frame after adding Parity %v\n", frame)


	// frame = append(frame,newTcpSegment...)
	// fmt.Printf("newTcpSegment : %c\n", newTcpSegment)
	// fmt.Println("frame after newPayload", frame)
	frame = append(frame, endFlag...)
	fmt.Println("EndFlag:", endFlag)
	fmt.Printf("Frame 1after endflag %v\n", frame)

    cound := link.Count(frame)
    fmt.Println("Cound", cound)

	// fmt.Println("Frame", utils.ByteToBinary(frame))

	for _, bit := range frame {
		strBit := fmt.Sprintf("%d", bit)
		utils.WriteToWire([]byte(strBit))

	}
}


func BuildPacket() {
}
