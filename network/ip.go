package network


func Packet(src, des, ttl, payload []byte) []byte{

    var packet []byte

    packet = append(packet, src...)
    packet = append(packet, des...)
    packet = append(packet, ttl...)
    packet = append(packet, payload...)

    return packet


}
