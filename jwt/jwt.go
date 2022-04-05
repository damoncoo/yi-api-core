package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Session interface {
	Token() string
	LoginRequired() bool
	AdminRequired() bool
	Set(key string, value interface{})
	Path() string
	RSA() string
}

type AuthInterface interface {
	Auth(session Session) error
	Signing(uid int, living int64) (string, error)
	VerifySigning(tokenString string) (map[string]interface{}, error)
}

// AuthImpl AuthImpl
type AuthImpl struct {
	Key string
}

func (authImp AuthImpl) Auth(session Session) error {

	// 是否需要鉴权
	auth := session.LoginRequired()
	// 是否是admin
	admin := session.AdminRequired()

	var errMsg = "无权访问"
	var bearerToken = session.Token()

	if bearerToken == "" {
		if !auth {
			return nil
		}
		return errors.New(errMsg)
	}

	// 解析失败
	c, err := authImp.VerifySigning(bearerToken)
	if err != nil {
		errMsg += ("," + err.Error())
		return errors.New(errMsg)
	}

	uidS := c["jti"].(string)
	uid, _ := strconv.Atoi(uidS)
	session.Set("uid", uid)

	// 需要admin 而且就是admin
	if admin && uid != 1 {
		return errors.New("需要管理权限")
	}

	return nil
}

// Signing 签名
func (auth AuthImpl) Signing(uid int, living int64) (string, error) {

	// claims := jwt.MapClaims(payload)
	claims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + living,
		Id:        strconv.Itoa(uid),
		Issuer:    "IMOVIE",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(auth.Key)
	return tokenString, err
}

// VerifySigning token 验证
func (authImp AuthImpl) VerifySigning(tokenString string) (map[string]interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(authImp.Key), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return map[string]interface{}(claims), nil
	}

	return nil, errors.New("登录过期")
}
