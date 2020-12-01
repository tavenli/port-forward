package Utils

import (
	"math"
)

func AbsInt(num float64) int {
	//result := math.Abs(float64(num))
	result := math.Abs(num)
	return int(result)
}

func AbsInt64(num float64) int64 {
	result := math.Abs(num)
	return int64(result)
}

func CeilInt(num float64) int {
	result := math.Ceil(num)
	return int(result)
}

func CeilInt64(num float64) int64 {
	//CeilInt64(5.9) = 6
	//CeilInt64(5.3) = 6
	//CeilInt64(5) = 5
	result := math.Ceil(num)
	return int64(result)
}

func Float64ToInt64(num float64) int64 {
	return int64(num)
}

func Float64TryToInt64(num interface{}) int64 {
	return int64(num.(float64))
}

//	返回最大值
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//	返回最小值
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Pages(total, psize int64) int64 {

	pages := float64(total) / float64(psize)
	result := math.Ceil(pages)
	return int64(result)
}
