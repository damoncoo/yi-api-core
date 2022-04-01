package cryp

import (
	"fmt"
	"testing"
)

func Test_RSA(t *testing.T) {

	// // 先生成私钥
	privateKey, privateChiper := RsaPrivateKey("../../keys/id_rsa")
	// publicKey, publicChiper := RsaPublicKey("../../keys/id_rsa.pub")

	// str := "aabab"
	// en, _ := RsaPublicEncrypt(str, publicKey, publicChiper)
	// res, _ := RsaPrivateDecrypt(en, privateKey, privateChiper)
	// fmt.Println(en)
	// fmt.Println(res)

	res2, _ := RsaPrivateDecrypt("ef5uhHLktOT32MjYEX13wHgacvKKSSepdqXHbhABGbfzcSuYterP7Qy+vo4htx0yiyDf+Qtk8ReERPlFdLjhXWIb8tAcor5aUCDO1mN5jbhlmE7wKEcaOYOsMqiaH5jRtIyzv1m/qQSnV1kz7NNV379QP7Pmgs3SVz6jboPYhzI=", privateKey, privateChiper)
	fmt.Println(res2)

}
