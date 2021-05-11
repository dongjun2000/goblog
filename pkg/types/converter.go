package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}