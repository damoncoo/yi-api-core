package cryp

import (
	"testing"
)

func Test_AESCrypt(t *testing.T) {

	res, err := AesEncrypt("ababb", []byte("!QAZ2WSX#EDC4RFV"), "!QAZ2WSX#EDC4RFV")
	if err == nil {
		t.Log(string(res))
	} else {
		t.Error(err)
	}
}

func Test_AesDecrypt(t *testing.T) {

	key := "a7fb3283682eef55"
	text := `vpMtVWjfjydQNqtINZZBlWIs8YmaBjB8NMmupe2f8Wvg7VQldWjphk/y+s90g2baqeOeMPsZx+bQd+pbaY9plVl6/uscffJ5tqeSTizF0XNItAlH3lPUx6VixUUxtL51okr7V0MgCxbA+BjLVT3yDSgiJqFjiSV1brK8Cp8xZUu4Ab2QfM5LgTKEYtKkoV/bbAY0xv69zzhHwgcYRXUEIHNsr5qneXCJh1oqFiGs3cxIKWgkZ+dSTHB9c+nazv4rWB/0ml139mZHUCDHNrm+TQ==`
	res, err := AesDecrypt(text, []byte(key), key)
	if err == nil {
		t.Log(string(res))
	} else {
		t.Error(err)
	}
}
