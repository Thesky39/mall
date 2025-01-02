package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

// AES对称加密
type Encryption struct {
	Key string
}

var Encrypt *Encryption

func init() {
	Encrypt = NewEncryption()
}
func NewEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// AesEncoding加密
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	block, err := aes.NewCipher([]byte(k.Key))
	if err != nil {
		return ""
	}
	NewSrcByte := PadPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(NewSrcByte))
	block.Encrypt(dst, NewSrcByte)
	//base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}

// UnPadPwd去掉填充部分
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) < 0 {
		return nil, nil
	}
	unpadNum := int(dst[len(dst)-1])
	strErr := "error"
	op := []byte(strErr)
	if len(dst) < unpadNum {
		return op, nil
	}
	str := dst[:len(dst)-unpadNum]
	return str, nil
}

// 解密
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte := []byte(pwd)
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return ""
	}
	block, errBlock := aes.NewCipher([]byte(k.Key))
	if errBlock != nil {
		return ""
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	dst, err = UnPadPwd(dst)
	if err != nil {
		return "0"
	}
	return string(dst)
}
func (k *Encryption) SetKey(key string) {
	k.Key = key
}
