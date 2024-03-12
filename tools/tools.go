package tools

import (
	"github.com/gogf/gf/v2/encoding/gcharset"
)

func Transform(src []byte, dst []byte) {
	for i := 0; i < 16; i++ {
		dst[i*8] = src[2*i]
		dst[i*8+1] = src[2*i+1]
		dst[i*8+2] = src[32+2*i]
		dst[i*8+3] = src[32+2*i+1]
		dst[i*8+4] = src[64+2*i]
		dst[i*8+5] = src[64+2*i+1]
		dst[i*8+6] = src[96+2*i]
		dst[i*8+7] = src[96+2*i+1]
	}
}

func UTF8ToGB2312(text string) []byte {
	str, err := gcharset.UTF8To("GBK", text)
	if err != nil {
		panic(err)
	}

	return []byte(str)
}
