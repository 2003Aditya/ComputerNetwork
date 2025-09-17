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

func HandleSYN(ctx *PacketContext) {

	msg := "d"
	msgB := []byte(msg)

    msgByte, parity := utils.MsgToByte(msgB)
    newMsg := utils.AsciiBytesToDigitBytes(msgByte)
    fmt.Printf("newMsg %v\n", newMsg)
    fmt.Printf("parity %v\n", parity)




	fmt.Printf("Syn Received from %c to %c\n", ctx.Src, ctx.Des)
	fmt.Printf("Initial sequence number %v\n", ctx.Seq)

	ctx.NewTTL = utils.AsciiBytesToDigitBytes(ctx.TTL)
	fmt.Printf("New TTL %v\n", ctx.NewTTL)

	wuhuuSeq := utils.AsciiBytesToDigitBytes(ctx.Seq)
	fmt.Printf("wuhuuSeq: %v\n", wuhuuSeq)

	newAck := utils.AsciiBytesToDigitBytes(ctx.Ack)
	fmt.Printf("newAck %v\n", newAck)

    newSrc := utils.AsciiBytesToDigitBytes(ctx.Des)
	fmt.Printf("newSrc: %v\n", newSrc)

    newDes := utils.AsciiBytesToDigitBytes(ctx.Src)
	fmt.Printf("newdes: %v\n", newDes)

	newSeq := utils.Increment(wuhuuSeq)
	fmt.Printf("NewSeq: %v\n", newSeq)

	startFlag := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	endFlag := []byte{0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Println("creating ACK-SYN Packet")
	fmt.Printf("Source : %s, Destination: %s\n",ctx.Src, ctx.Des)

	newTcpSegment := transport.Tcp(newSeq, newAck, newMsg, true, true, false)
	fmt.Printf("newTcpSegment %v\n", newTcpSegment)

	newPacket := network.Packet(newSrc, newDes,ctx.NewTTL, newTcpSegment)


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

func HandleSYNACK(ctx *PacketContext) {
    StartFlag := ctx.StartFlag
    EndFlag := ctx.EndFlag
    newSrc := ctx.Des
    newDes := ctx.Src
    newTTL := ctx.TTL
    newSeq := ctx.Seq
    newAck := ctx.Ack

    //msg formulation
    msg := "b"
    msgb := []byte(msg)

    msgByte, parity := utils.MsgToByte(msgb)
    newMsg := utils.AsciiBytesToDigitBytes(msgByte)
    fmt.Printf("NewMsg %v\n", newMsg)
    fmt.Printf("parity %v\n", parity)



    // tcp segment creation
    newTcpSegment := transport.Tcp(newSeq, newAck, newMsg,false, false, true )
    fmt.Printf("NewTCPSegment: %v", newTcpSegment)

    newPacket := network.Packet(newSrc, newDes, ctx.TTL, )


    fmt.Printf("ctx.NewTTL: %c", newTTL)

    fmt.Printf("Src1: %c, Des1: %c", newSrc, newDes)

}

func BuildPacket() {
}
