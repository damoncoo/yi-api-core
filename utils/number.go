package utils

func CeilDivide(a, b uint) uint {

	if b == 0 {
		return 0
	}
	left := a % b
	if left > 0 {
		return uint(a/b + 1)
	}
	return uint(a / b)
}
