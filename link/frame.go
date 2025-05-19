package link

import "fmt"

// "fmt"

// startEnd := []byte {0,1,1,1,1,1,1,0}

func Frame(b byte) []byte  {
    startFlag := []byte{0,1,1,1,1,1,1,0}
    endFlag := []byte{0,1,1,1,1,1,1,0}

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
    dataBits = append(dataBits, byte(parity))
    framed := append([]byte{}, startFlag...)
    framed = append(framed, dataBits...)
    framed = append(framed, byte(parity))
    framed = append(framed, endFlag...)
    fmt.Printf("parity: %d\n", parity)
    fmt.Printf("Framed data: %v\n", framed)
    return framed
}


// func main() {
//
//     b := "C"
//     byteValue := b[0]
//     bits := frame(byteValue)
//     fmt.Printf("bits of that byte:%v ",bits)
//     fmt.Println()
// }
