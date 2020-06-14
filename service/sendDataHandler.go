package service

import (
	"../hexu"
	"strconv"
)

/**
无数据
头字符
开头int类型的个数
是否验证
*/
func SendDataHandler(src []string, head string, intCount int, valid bool) string {
	data := head
	for i := range src {
		s := src[i]
		if i < intCount {
			v, _ := strconv.ParseInt(s, 10, 32)
			data += hexu.Int64ToHex(v)
		} else {
			v, _ := strconv.ParseFloat(s, 2)
			data += hexu.Float32ToHex(float32(v))

		}
	}
	if valid {
		checksum := hexu.CheckSum(data)
		data += checksum
	}
	return data

}
