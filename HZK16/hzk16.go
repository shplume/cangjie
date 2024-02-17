package hzk16

import (
	"fmt"
	"os"
)

var key = [8]byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}

func HZK16(word [2]byte) {
	file, err := os.Open("font/HZK16")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	offset := (94*int64(word[0]-0xa0-1) + int64(word[1]-0xa0-1)) * 32
	buffer := make([]byte, 32)
	if _, err := file.ReadAt(buffer, offset); err != nil {
		fmt.Println(err)
		return
	}

	display(buffer)
}

func display(buffer []byte) {
	for k := 0; k < 16; k++ {
		for j := 0; j < 2; j++ {
			for i := 0; i < 8; i++ {
				if buffer[k*2+j]&key[i] == 0x0 {
					fmt.Printf("○ ")
				} else {
					fmt.Printf("● ")
				}
			}
		}

		fmt.Println()
	}
}
