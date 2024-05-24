package handler

func HZK16(word []byte) []byte {
	offset := (94*int64(word[0]-0xa0-1) + int64(word[1]-0xa0-1)) * 32

	buffer := handler("font/HZK16", offset)

	return buffer
}
