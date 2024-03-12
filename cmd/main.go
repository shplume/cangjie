package main

import (
	"fmt"

	hzk16 "github.com/shplume/cangjie/HZK16"
	tools "github.com/shplume/cangjie/tools"
)

func main() {
	var word = tools.UTF8ToGB2312("选课成功")
	if len(word) != 8 {
		fmt.Println("word is invalid")

		return
	}

	buffer := []byte{}
	for i := 0; i < len(word); i += 2 {
		buffer = append(buffer, hzk16.HZK16(word[i:i+2])...)
		fmt.Println()
	}

	dst := make([]byte, len(buffer))
	tools.Transform(buffer, dst)

	for i := 0; i < len(dst); i++ {
		fmt.Printf("0x%02x, ", dst[i])

		if i%32 == 0 && i != 0 {
			fmt.Println()
		}
	}
}
