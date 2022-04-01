package utils

import (
	"regexp"
)

// IsMatch 是否匹配 正则
func IsMatch(src string, regStr string) bool {

	reg := regexp.MustCompile(regStr)
	data := []byte(src)
	return reg.Match(data)
}

// AllMatches AllMatches
func AllMatches(src string, regStr string) []string {
	reg := regexp.MustCompile(regStr)
	return reg.FindAllString(src, -1)
}

// IsChinaMainlandPhone 是否手机号匹配 正则
func IsChinaMainlandPhone(phoneStr string) bool {

	return IsMatch(phoneStr, "^1[0-9]{10}$")
}

// IsEmail 是否电子邮件 正则
func IsEmail(emailStr string) bool {
	match := IsMatch(emailStr, "^[0-9a-zA-Z]+@([0-9a-zA-Z]+\\.)+[a-z]{2,}$")
	return match
}

// IsNumberCode 是否数字验证码 正则
func IsNumberCode(code string) bool {
	match := IsMatch(code, "^[0-9]{4}$")
	return match
}

// IsValidPass 密码格式是否正确 正则
func IsValidPass(password string) bool {
	match := IsMatch(password, "^[a-zA-Z0-9!@#$%^&*]{6,16}$")
	return match
}

// IsValidStrongPass 密码格式是否强密码 正则
func IsValidStrongPass(password string) bool {

	return IsMatch(password, "^(?=.*[A-Za-z])(?=.*\\d)[\\w!@#$%^&*()]{6,16}$")
}
