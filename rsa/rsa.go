package rsa

import (
	"crypto/rsa"
	"errors"

	"github.com/damoncoo/yi-api-core/jwt"
	cryp "github.com/damoncoo/yi-api-core/utils/crypto"
)

// PublicKey PublicKey
type PublicKey struct {
	RSAPublicKey *rsa.PublicKey
	Data         []byte
}

// PrivateKey PrivateKey
type PrivateKey struct {
	RSAPrivateKey *rsa.PrivateKey
	Data          []byte
}

type RSA struct {
	RSAPrivatePath string
	RSAPublicPath  string
	RSAPrivateKey  *PrivateKey
	RSAPublicKey   *PublicKey
}

func (r RSA) InitRsaPrivateBlock() {

	privateKey, privateChiper := cryp.RsaPrivateKey(r.RSAPrivatePath)
	r.RSAPrivateKey = &PrivateKey{
		RSAPrivateKey: privateKey,
		Data:          privateChiper,
	}

	publicKey, publicChiper := cryp.RsaPublicKey(r.RSAPublicPath)
	r.RSAPublicKey = &PublicKey{
		RSAPublicKey: publicKey,
		Data:         publicChiper,
	}
	if r.RSAPrivateKey == nil ||
		r.RSAPublicKey == nil {
		panic("无效的RSA秘钥")
	}
}

type RSAInterface interface {
	VerifySigning(session jwt.Session) error
}

func (r RSA) VerifySigning(session jwt.Session) error {

	snoc := session.RSA()
	if snoc == "" {
		return errors.New("无效请求")
	}

	AesKey, err := cryp.RsaPrivateDecrypt(snoc, r.RSAPrivateKey.RSAPrivateKey, r.RSAPrivateKey.Data)
	if err != nil || len(AesKey) != 16 {
		return errors.New("无效请求")
	}

	session.Set("AesKey", AesKey)

	return nil
}
