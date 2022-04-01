package utils

import "testing"

func TestBazi(t *testing.T) {

	bazi := GetBazi(1988, 11, 12, 03, 40, 20, 1)
	t.Log(bazi)
}
