package hexu

import (
	"math"
	"strconv"
)

/*
 float to hex
*/
func Float32ToHex(src float32) string {
	bits := math.Float32bits(src)
	return Int64ToHex(int64(bits))
}

/*
 int to hex
*/
func Int64ToHex(src int64) string {
	s := strconv.FormatInt(src, 16)
	length := len(s)
	if length < 8 { //如果小于8位
		for i := 1; i <= 8-length; i++ {
			s = "0" + s
		}
	}
	return s
}

/*
 hex to byte[]
*/
func HexToByte(str string) []byte {
	slen := len(str)
	bHex := make([]byte, len(str)/2)
	ii := 0
	for i := 0; i < len(str); i = i + 2 {
		if slen != 1 {
			ss := string(str[i]) + string(str[i+1])
			bt, _ := strconv.ParseInt(ss, 16, 32)
			bHex[ii] = byte(bt)
			ii = ii + 1
			slen = slen - 2
		}
	}
	return bHex
}
