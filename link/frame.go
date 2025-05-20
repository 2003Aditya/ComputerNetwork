package link

import (
    "fmt"
    "github.com/2003Aditya/ComputerNetwork/network"
)

// "fmt"

// startEnd := []byte {0,1,1,1,1,1,1,0}

func Frame(b byte) []byte  {
    startFlag := []byte{0,1,1,1,1,1,1,0}
    endFlag := []byte{0,1,1,1,1,1,1,0}
    src := []byte{0,1}
    des := []byte{1,0}
    ttl := []byte{0,1,0,1}

    ones := 0
    var parity int

    var dataBits []byte
    fmt.Printf("byte of %c\n", b)

    for i := 7; i >= 0; i-- {
        bits := (b >> i) & 1
        dataBits = append(dataBits, byte(bits))

        if bits == 1 {
            ones++
        }

    }

    if ones % 2 == 0 {
        parity = 0
    } else{
        parity = 1
    }

    packet := network.Packet(src, des, ttl, dataBits)
    framed := append([]byte{}, startFlag...)
    framed = append(framed, packet...)
    framed = append(framed, byte(parity))
    framed = append(framed, endFlag...)
    fmt.Printf("parity: %d\n", parity)
    fmt.Printf("Framed data: %v\n", framed)
    return framed
}

func count(frame []byte) int{
    count := 0
    for range frame {
        count++
    }
    return count
}


// func main() {
//
//     b := "C"
//     byteValue := b[0]
//     bits := Frame(byteValue)
//     fmt.Printf("total no of bits: %d\n" ,count(bits))
//     fmt.Printf("bits of that byte:%v ",bits)
//     fmt.Println()
// }
