package cryp

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
)

// RsaPublicEncrypt 公钥加密
func RsaPublicEncrypt(payload string, key *rsa.PublicKey, chiper []byte) (string, error) {

	return encrypt(payload, key, chiper)
}

// RsaPrivateDecrypt 私钥解密
func RsaPrivateDecrypt(encryptStr string, key *rsa.PrivateKey, chiper []byte) (string, error) {

	return decrypt(encryptStr, key, chiper)
}

// RsaPublicKey RsaPublicKey
func RsaPublicKey(path string) (*rsa.PublicKey, []byte) {

	keyData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("ERROR: fail get idrsapub, %s", err.Error())
		return nil, []byte{}
	}

	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil {
		log.Printf("ERROR: fail get idrsapub, invalid key")
		return nil, []byte{}
	}

	publicKey, err := x509.ParsePKIXPublicKey(keyBlock.Bytes)
	if err != nil {
		log.Printf("ERROR: fail get idrsapub, %s", err.Error())
		return nil, []byte{}
	}
	switch publicKey := publicKey.(type) {
	case *rsa.PublicKey:
		return publicKey, keyData
	default:
		return nil, []byte{}
	}
}

// RsaPrivateKey RsaPrivateKey
func RsaPrivateKey(path string) (*rsa.PrivateKey, []byte) {

	keyData, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("ERROR: fail get idrsa, %s", err.Error())
		os.Exit(1)
	}

	keyBlock, _ := pem.Decode(keyData)
	if keyBlock == nil {
		log.Printf("ERROR: fail get idrsa, invalid key")
		os.Exit(1)
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		log.Printf("ERROR: fail get idrsa, %s", err.Error())
		panic(err)
	}

	return privateKey, keyData
}

func encrypt(payload string, key *rsa.PublicKey, chiper []byte) (string, error) {
	// params
	msg := []byte(payload)

	// encrypt with PKCS1v15
	ciperText, err := rsa.EncryptPKCS1v15(rand.Reader, key, msg)
	if err != nil {
		log.Printf("ERROR: fail to encrypt, %s", err.Error())
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciperText), nil
}

func decrypt(payload string, key *rsa.PrivateKey, chiper []byte) (string, error) {
	// decode base64 encoded signature
	msg, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		log.Printf("ERROR: fail to base64 decode, %s", err.Error())
		return "", err
	}

	// decrypt with DecryptPKCS1v15
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, key, msg)
	if err != nil {
		log.Printf("ERROR: fail to decrypt, %s", err.Error())
		return "", err
	}

	return string(plainText), nil
}
