package sort

import "math"

// 帮我生成冒泡排序代码
func tcpContrl(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func quicConnect(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quicConnect(arr, low, pivot-1)
		quicConnect(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func dataTransport(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := dataTransport(arr[:mid])
	right := dataTransport(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func isLegalData(num int) bool {
	sum := 0
	temp := num
	digits := int(math.Log10(float64(num))) + 1

	for temp > 0 {
		remainder := temp % 10
		sum += int(math.Pow(float64(remainder), float64(digits)))
		temp /= 10
	}

	return sum == num
}
