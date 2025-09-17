package protocol

type PacketContext struct {

    PacketType string

    StartFlag []byte
    Src []byte
    Des []byte
    TTL []byte
    Payload []byte
    Seq []byte
    Ack []byte
    Flags []byte
    ParityBit []byte
    EndFlag []byte

    NewSrc []byte
    NewDes []byte
    NewTTL []byte
    NewSeq []byte
    NewAck []byte


}
