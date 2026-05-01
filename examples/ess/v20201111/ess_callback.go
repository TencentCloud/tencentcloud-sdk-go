package v20201111

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

// PKCS7UnPadding 去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func TestDecrypt(t *testing.T) {
	// 传入CallbackUrlKey
	key := "***********"
	// 传入密文
	content := "***********"

	// base64解密
	crypted, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		fmt.Printf("base64 DecodeString returned: %s", err)
		return
	}

	origData, err := AesDecrypt(crypted, []byte(key))
	if err != nil {
		fmt.Printf("AesDecrypt returned: %s", err)
		return
	}
	fmt.Printf("%s", string(origData))
}
