package handler

import (
	"fmt"
	"os"

	"github.com/shplume/cangjie/tools"
)

func handler(name string, offset int64) []byte {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer file.Close()

	buffer := make([]byte, 32)
	if _, err := file.ReadAt(buffer, offset); err != nil {
		fmt.Println(err)
		return nil
	}

	tools.Display(buffer)

	return buffer
}
