package transport



func Tcp(SeqNum, AckNum, Payload []byte, SYN, ACK, FIN bool ) []byte{

    var TCP []byte

    TCP = append(TCP, SeqNum...)
    TCP = append(TCP, AckNum...)
    TCP = append(TCP, bootToBit(SYN))
    TCP = append(TCP, bootToBit(ACK))
    TCP = append(TCP, bootToBit(FIN))
    TCP = append(TCP, Payload...)

    return TCP

}

func bootToBit(flag bool) byte {
    if flag {
        return 1
    } else {
        return 0
    }
}

