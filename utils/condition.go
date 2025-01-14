package utils

// If 自定的三元计算
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// Filter 筛选数组
func Filter[T any](slice []T, filter func(T) bool) (ret []T) {
	for _, s := range slice {
		if filter(s) {
			ret = append(ret, s)
		}
	}
	return
}

// ContainInt ContainInt
func ContainInt(slice []int, c int) bool {

	for _, i := range slice {
		if i == c {
			return true
		}
	}
	return false
}

// ErrorCheck 错误检测
func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
