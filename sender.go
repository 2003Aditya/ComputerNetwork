package main

import (
	"fmt"
	"os"
	"time"

	"github.com/2003Aditya/ComputerNetwork/link"
)

func main() {

	message := "C"

	file, err := os.OpenFile("wire.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open wire: ", err)
		return
	}
	defer file.Close()

	for _, ch := range message {
		framedBits := link.Frame(byte(ch))

		for _, bit := range framedBits {
			strBit := fmt.Sprintf("%d", bit)
			_, err := file.WriteString(strBit)

			if err != nil {
				fmt.Println("Write error: ", err)
				return
			}

			fmt.Printf("Send bit : %v\n", strBit)
			time.Sleep(1 * time.Second)

		}

	}

	fmt.Println("Transmission complete")
}
