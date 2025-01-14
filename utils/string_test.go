package utils

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	random := RandStringRunes(6)
	fmt.Println(random)
}
