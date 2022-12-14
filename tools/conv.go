package tools

import (
	"fmt"
	"strconv"
)

/*
	Float64ToString float64转成string
*/
func Float64ToString(target float64) string {
	res := strconv.FormatFloat(target, 'f', -1, 64)

	return res
}

/*
	StringToFloat64 string转成float64
*/
func StringToFloat64(target string) float64 {
	res, err := strconv.ParseFloat(target, 64)
	if err != nil {
		fmt.Println("parse float failed, err:", err)
		return 0
	}
	return res
}

/*
	StringToInt64 string 转 int64
*/
func StringToInt64(target string) int64 {
	res, _ := strconv.ParseInt(target, 10, 64)

	return res
}

/*
	StringToInt string 转 int
*/
func StringToInt(target string) int {
	res, _ := strconv.Atoi(target)

	return res
}

/*
	Int64ToString int64 转 string
*/
func Int64ToString(target int64) string {
	return strconv.FormatInt(target, 10)
}
