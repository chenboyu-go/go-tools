package tools

import "strconv"

// KeepOneDecimal 保留一位小数
func KeepOneDecimal(num float64) (ret float64) {
	retStr := strconv.FormatFloat(num, 'f', 1, 64)
	ret = StringToFloat64(retStr)
	return
}
