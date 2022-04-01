package utils

import (
	"encoding/json"
	"math/rand"
	"regexp"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// ObjecToString ObjecToString
func ObjecToString(obj interface{}) string {

	jsonByte, _ := json.MarshalIndent(obj, "", " ")
	jsonStr := string(jsonByte)
	return jsonStr
}

// RandStringRunes RandStringRunes
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// string match pattern
func Match(str, pattern string) bool {
	reg := regexp.MustCompile(pattern)
	return reg.Match([]byte(str))
}
