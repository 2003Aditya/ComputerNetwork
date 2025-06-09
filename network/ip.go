package network


func Packet(src, des, ttl, TCPSegment []byte) []byte{

    var packet []byte

    packet = append(packet, src...)
    packet = append(packet, des...)
    packet = append(packet, ttl...)
    packet = append(packet, TCPSegment...)

    return packet


}
