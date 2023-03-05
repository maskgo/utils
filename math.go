package utils

import "math"

// round()
func Round(value float64) float64 {
	return math.Floor(value + 0.5)
}

// floor()
func Floor(value float64) float64 {
	return math.Floor(value)
}

// ceil()
func Ceil(value float64) float64 {
	return math.Ceil(value)
}

// max()
func Max(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}

	max := nums[0]

	for i := 1; i < len(nums); i++ {
		max = math.Max(max, nums[i])
	}

	return max
}

// min()
func Min(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("nums: the nums length is less than 2")
	}

	min := nums[0]

	for i := 1; i < len(nums); i++ {
		min = math.Min(min, nums[i])
	}

	return min
}
