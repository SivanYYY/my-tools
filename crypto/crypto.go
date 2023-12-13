package pkg

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io"
	"my-tools/mlog"
	"os"
)

// Md5 md5加密
// use16=true 返回16位md5值
func Md5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5Sum 校验文件md5sum
func Md5Sum(filePath string) string {
	// 文件全路径名
	pFile, err := os.Open(filePath)
	if err != nil {
		return "打不开文件" + filePath
	}
	defer pFile.Close()
	md5h := md5.New()
	_, _ = io.Copy(md5h, pFile)
	return hex.EncodeToString(md5h.Sum(nil))
}

// GenerateRSAKey  初始化RSA公私钥
func GenerateRSAKey() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		mlog.ZapLogger.Error("generate privateKey err", err)
		return nil, nil, err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	publicKey := &privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		mlog.ZapLogger.Error("generate publicKey err", err)
		return nil, nil, err
	}
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return privateKeyPem, publicKeyPem, nil
}

// EncryptRSA RSA加密
// plainText 要加密的数据
// publicPem 公钥匙文件
func EncryptRSA(plainText, publicPem []byte) []byte {
	//pem解码
	block, _ := pem.Decode(publicPem)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	//返回密文
	return cipherText
}

// DecryptRSA RSA解密
// cipherText 需要解密的byte数据
// privatePem 私钥文件路径
func DecryptRSA(cipherText, privatePem []byte) []byte {
	//pem解码
	block, _ := pem.Decode(privatePem)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText
}
