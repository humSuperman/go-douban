package util

import (
	"encoding/base64"
	"fmt"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"sort"
	"strings"
	"crypto/sha1"
	"crypto"
	"os"
	"admin/pkg/logging"
)
// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDH0BsyCjRcWeuYgqFAMZJyKo716GQRfrOf4I9OIlfeiDTiqs68
uC2/he+u6s64iUkh2XsfFLpadsC+h4r88h27uaibDORWpnxVB8A3feWuImdeC0AC
eo9PFsAt7RLbzAPYHtbWsMHqJq4l1vDUO9aempwk5NgrZZ8VtV/u5FdPgwIDAQAB
AoGANWN2kMVHPlHMcICe408biSOz9SK18jK/ff17bO4iOlR8hQAMo0I2/xCjfUJC
H+6WutOoYSKhtGA8mewPiAyNQs2uOOyF8HvEj4bRo8IPBdZfIcon5UmgImIRcMRB
1a9bfKH87JBQvZ1cJvFvAkIrFWVvB4X1JY3trHcTTAVYzfECQQDcVRq0uxyLzEfd
pPlFp5+wWnsH85h5Z4nsQ1eXLcFGOvNQ0q5zxdsAWmG/5Fs7zCErBCSJfQCovAqF
2/LIJiapAkEA6CikoibYCq3H9yQM00FCdJvzHoY962GEQMFrT7s8CsKXzCqCp9VN
4UiE7P1NX5LZ8JbGJIYNf3tuLVc7xKOcSwJBANghR1wn+32KuqhJ7xeLsVKNvwfy
xQu6LAotmNs8T83zf8A1mlkIqaY0ApT4jSIgQBzxKGIR07axFmVudz4sZlkCQGWk
bbgEk6/RN+xPH0JanxLYuE+T7IYicrm7NRV6XyZC3hzoO73pWKiajIAJwpdmfv6j
tGqHOl+nFazKNYO5MhMCQQC7LLiNNJtHIOhi9FVrclawZHcTZQNvCJJRP9LEO5rF
Bd2h1iS0Xf1p8g/qk3JO7EZ/nwiwOB++Ooh3BgD1hImt
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDH0BsyCjRcWeuYgqFAMZJyKo71
6GQRfrOf4I9OIlfeiDTiqs68uC2/he+u6s64iUkh2XsfFLpadsC+h4r88h27uaib
DORWpnxVB8A3feWuImdeC0ACeo9PFsAt7RLbzAPYHtbWsMHqJq4l1vDUO9aempwk
5NgrZZ8VtV/u5FdPgwIDAQAB
-----END PUBLIC KEY-----
`)

var wxTemplatePubKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDPVRn9ln20yIEk/A08TM6V17FL
9TeRal23j1T1xcH/AoO6wq405IemvyJ8Yk52lim5KwMZ9C1RYfPF8+NOTjmi1g1j
uzBp2QNsKg/ZKdqULvMBfnMkl4E7KJ4RE6hUL7kSAvDL1i74J1PN6nEA+tY3OiIm
jgtM8hICyu3weBPFkQIDAQAB
-----END PUBLIC KEY-----
`)

func ResToBase64(data string) string {
	rsaData, _ := RsaEncryptEx([]byte(data), wxTemplatePubKey)
	base64Data := base64.StdEncoding.EncodeToString(rsaData)
	fmt.Println(base64Data)
	return base64Data
}
//加密
func RSAEncrypt(value string) string {
	return ""
}
//解密
func RSADecrypt(value string) string {
	return ""
}
// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, fmt.Errorf("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 加密
func RsaEncryptEx(origData []byte, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, fmt.Errorf("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, fmt.Errorf("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// 解密
func RsaDecryptEx(ciphertext []byte, priKey []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, fmt.Errorf("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}


func CheckSignature(src_params []string, signedData string) error {
	sort.Strings(src_params)
	sp := strings.Join(src_params, "&")
	signature_src := strings.Replace(strings.Replace(strings.Replace(sp, " ", "", -1), "[", "", -1), "]", "", -1)

	logging.Debug("signedData = ", signedData)
	logging.Debug("signature_src = ", signature_src)

	h := sha1.New() //crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(signature_src))
	SHA1 := h.Sum(nil)
	//s := sha1.New()
	//io.WriteString(s, strings.Join(src_params, "&"))
	//SHA1 := s.Sum(nil)
	logging.Debug(SHA1)

	block_pub, _ := pem.Decode(wxTemplatePubKey)
	if block_pub == nil {
		// 失败情况
		panic("public key error")
	}

	pub, err := x509.ParsePKIXPublicKey(block_pub.Bytes)
	if err != nil {
		panic("ParsePKIXPublicKey error" + err.Error())
	}

	//signedData[0] = byte(0x11)
	sd, _ := base64.StdEncoding.DecodeString(signedData)
	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, SHA1, sd)

	fmt.Println("VerifyPKCS1v15 = ", err)

	return err
}

func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("d:/private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("d:/public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}