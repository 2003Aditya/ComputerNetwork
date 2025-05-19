package main

import (
    "fmt"
    "time"
    "os"
    "io"
)

func main() {
    file, err := os.Open("wire.txt")
    if err != nil {
        fmt.Println("Failed to open wire:", err)
        return
    }
    defer file.Close()

    buffer := make([]byte, 1)
    var offset int64 = 0

    for {
        // Move the file pointer to the current offset
        _, seekErr := file.Seek(offset, 0)
        if seekErr != nil {
            fmt.Println("Seek error:", seekErr)
            return
        }

        // Read 1 byte
        n, err := file.Read(buffer)
        if err != nil && err != io.EOF {
            fmt.Println("Error reading wire:", err)
            return
        }

        if n > 0 {
            fmt.Printf("Received bit: %c\n", buffer[0])
            offset++
        }

        time.Sleep(1 * time.Second)
    }

    fmt.Println(buffer)

}

