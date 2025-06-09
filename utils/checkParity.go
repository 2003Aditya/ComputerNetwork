package utils


func CheckParity(databits []byte, parity byte) bool{
    count := 0
    for _, b := range databits{
        if b == '1' {
            count++
        }
    }

    parityBit := 0
    if parity == '1' {
        parityBit = 1
    }

    return (count + parityBit) % 2 == 0

}

// func EqualFlags(a , b []byte) bool {
//     if len(a) != len(b) {
//         return false
//     }
//     for i := range {
//         if a[i] != b[i] {
//             return false
//         }
//     }
//
//     return true
//
// }
