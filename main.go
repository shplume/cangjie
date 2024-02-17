package main

import hzk16 "github.com/shplume/cangjie/HZK16"

func main() {
	var word = [2]byte{0xCE, 0xD2}
	hzk16.HZK16(word)
}
