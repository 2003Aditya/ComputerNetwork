package utils

func ByteToChar(databits []byte)byte {

    var result byte = 0
    for i := 0; i < 8; i++ {
        result <<= 1
        if databits[i] == '1' {
            result |= 1
        }
    }

    return result

}
