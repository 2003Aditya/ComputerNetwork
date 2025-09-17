package utils

import (
	"fmt"
	"os"

)

func ByteToChar(databits []byte)byte {

    var result byte = 0
    for i := 0; i < len(databits)-1; i++ {
        result <<= 1
        if databits[i] == '1' {
            result |= 1
        }
    }

    return result

}

func MsgToByte(value []byte)( []byte, int ){

    var msgByte []byte
    var parity int
    ones := 0

    for _, b := range value{
        for i := 7; i >= 0; i-- {
            bit := ( b >> i) & 1
            msgByte = append(msgByte, byte('0'+bit))

            if bit == 1 {
                ones++
            }
        }
        if ones % 2 == 0 {
            parity = 0
        } else {
            parity = 1
        }
    }
    return msgByte, parity
}

func AsciiBytesToDigitBytes(value []byte) []byte {
    result := make([]byte, len(value))
    for i, v := range value {
        result[i] = v - '0'
    }
    return result
}


func WriteToWire(frame []byte) {
    file, err := os.OpenFile("wire.txt", os.O_APPEND|os.O_CREATE| os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Opening failed", err)
        return
    }
    defer file.Close()

    _, err = file.Write(frame)
    if err != nil {
        fmt.Println("error while writing", err)
    }
    return

}



func GetPacketTypeSimple(flags []byte) string {
    if len(flags) < 3 {
        return "UNKOWN"
    }

    syn := flags[0]
    ack := flags[1]
    fin := flags[2]

    if syn == '1' && ack == '0' && fin == '0' {
        return "SYN"
    } else if syn == '1' && ack == '1' && fin == '0' {
        return "SYN-ACK"
    } else if syn == '0' && ack == '1' && fin == '0' {
        return "ACK"
    } else if syn == '0' && ack == '0' && fin == '1' {
        return "FIN"
    } else if syn == '0' && ack == '0' && fin == '0' {
        return "DATA"
    }

    return "UNKNOWN"



}


func Increment(value []byte) []byte {
    carry := byte(1)

    for i := len(value)-1; i >= 0; i-- {
        sum := value[i] + carry

        value[i] = sum % 2
        carry = sum / 2

        if carry == 0 {
            break
        }
    }
    if carry == 1 {
        value = append([]byte{1}, value...)

    }
    fmt.Printf("from inside Incremented:%v\n", value)

    return value

}

// this function will delete the wire bits after using them in a 52 bit of representation
func DeleteWire() {

    file, err := os.OpenFile("wire.txt", os.O_TRUNC|os.O_CREATE| os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Opening failed", err)
        return
    }
    defer file.Close()


}

// func ByteToBinary(value []byte ) string {
// var binaryString string
// for _, i := range value {
// binaryString += byte
// }
//
// }
