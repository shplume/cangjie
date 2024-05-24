package handler

func SongTi(word []byte) []byte {
	area := (int64(word[0]) - int64(0x81))
	location := (int64(word[1]) - int64(0x40))

	if location <= 16*4 {
		location -= area * 2
	} else {
		location -= area*2 + 1
	}

	offset := (area*192 + location) * 32
	buffer := handler("font/songti", offset)

	return buffer
}
