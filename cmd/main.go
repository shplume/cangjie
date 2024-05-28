package main

import (
	"fmt"

	"github.com/shplume/cangjie/handler"
	"github.com/shplume/cangjie/tools"
)

func main() {
	var word = tools.UTF8ToGB2312("语文数学")

	for i := 0; i < 8; i++ {
		fmt.Printf("0x%2x, ", word[i])
	}

	fmt.Println()

	buffer := []byte{}
	for i := 0; i < len(word); i += 2 {
		buffer = append(buffer, handler.SongTi(word[i:i+2])...)
		fmt.Println()
	}

	dst := make([]byte, len(buffer))
	tools.Transform(buffer, dst)

	for i := 0; i < len(dst); i++ {
		fmt.Printf("0x%02x, ", dst[i])

		if (i+1)%16 == 0 {
			fmt.Println()
		}
	}
}
