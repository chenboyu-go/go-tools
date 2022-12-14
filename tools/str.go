package tools

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// IsNumber 判断字符串是否为纯数字
func IsNumber(str string) bool {
	for _, s := range str {
		isDigit := unicode.IsDigit(s)
		if !isDigit {
			return false
		}
	}
	return true
}

// RandString 生成指定长度的随机字符串
func RandString(lenNum int) string {
	var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	str := strings.Builder{}
	length := len(chars)
	rand.Seed(time.Now().UnixNano()) //重新播种，否则值不会变
	for i := 0; i < lenNum; i++ {
		str.WriteString(chars[rand.Intn(length)])

	}
	return str.String()
}
