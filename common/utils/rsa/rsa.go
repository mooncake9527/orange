package rsaUtil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/mooncake9527/x/xerrors/xerror"
)

func ParsePublicKey(publicKeyStr string) (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode([]byte(publicKeyStr))
	if publicKeyBlock == nil || publicKeyBlock.Type != "PUBLIC KEY" {
		return nil, xerror.New("无效的公钥")
	}
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, xerror.New("无效的RSA公钥")
	}
	return rsaPublicKey, nil
}

// 将字符串转换为RSA私钥
func ParsePrivateKey(privateKeyStr string) (*rsa.PrivateKey, error) {
	privateKeyBlock, _ := pem.Decode([]byte(privateKeyStr))
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		return nil, xerror.New("无效的私钥")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func Encrypt(message, pubKey string) (string, error) {
	publicKey, err := ParsePublicKey(pubKey)
	if err != nil {
		return "", xerror.New(err.Error())
	}
	encryptedMessage, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(message))
	if err != nil {
		return "", xerror.New(err.Error())
	}

	return base64.StdEncoding.EncodeToString(encryptedMessage), nil
}

func Decrypt(encryptedMsg, priKey string) (string, error) {
	privateKey, err := ParsePrivateKey(priKey)
	if err != nil {
		return "", xerror.New(err.Error())
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedMsg)
	if err != nil {
		return "", xerror.New(err.Error())
	}
	decryptedMessage, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return "", xerror.New(err.Error())
	}
	return string(decryptedMessage[:]), nil

}

func GenerateRsaKey(len int) (string, string, error) {
	if len == 0 {
		len = 2048
	}
	key, err := rsa.GenerateKey(rand.Reader, len)
	if err != nil {
		return "", "", xerror.Newf("无法生成RSA密钥对:%s", err.Error())
	}
	publicKeyStr, err := PublicKeyToString(&key.PublicKey)
	if err != nil {
		return "", "", err
	}
	privateKeyStr, err := PrivateKeyToString(key)
	if err != nil {
		return "", "", err
	}
	return publicKeyStr, privateKeyStr, nil

}

func PrivateKeyToString(privateKey *rsa.PrivateKey) (string, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privateKeyStr := string(pem.EncodeToMemory(privateKeyBlock))
	return privateKeyStr, nil
}

func PublicKeyToString(publicKey *rsa.PublicKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", xerror.New(err.Error())
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicKeyStr := string(pem.EncodeToMemory(publicKeyBlock))
	return publicKeyStr, nil
}
