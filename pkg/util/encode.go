package util

import (
	"admin/pkg/logging"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strconv"
)
//md5加密
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
//sha1加密
func EncodeSHA1(value string) string {
	h := sha1.New()
	h.Write([]byte(value))

	return hex.EncodeToString(h.Sum(nil))
}
//sha256加密
func EncodeSHA256(value string) string {
	h := sha256.New()
	h.Write([]byte(value))

	return hex.EncodeToString(h.Sum(nil))
}
//sha512加密
func EncodeSHA512(value string) string {
	h := sha512.New()
	h.Write([]byte(value))

	return hex.EncodeToString(h.Sum(nil))
}
//base64 编码
func EncodeBase64(value string) string {
	base64Table := "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
	code := base64.NewEncoding(base64Table)
	return code.EncodeToString([]byte(value))
}
//base64 解码
func DecodeBase64(value string) string {
	base64Table := "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
	code := base64.NewEncoding(base64Table)
	res,err := code.DecodeString(value)
	if err != nil{
		logging.Error("DecodeBase64 string:%s error:%v\n",value,err)
		return ""
	}
	return string(res)
}
//url 编码
func EncodeURL(value string) string {
	return url.QueryEscape(value)
}
//url 解码
func DecodeURL(value string) string {
	needToChange := false
	numChars := len(value)
	var i = 0
	var c byte
	result := ""
	var bytes []byte
	for i < numChars {
		c = value[i]
		switch c {
		case '+':
			result += " "
			i++
			needToChange = true
			break
		case '%':
			if bytes == nil {
				bytes = make([]byte, (numChars-i)/3)
			}
			pos := 0
			for (i+2 < numChars) && (c == '%') {
				v, _ := strconv.ParseInt(value[i+1:i+3], 16, 32)
				//                        int v = Integer.parseInt(s.substring(i+1,i+3),16);
				bytes[pos] = byte(v)
				pos++
				i += 3
				if i < numChars {
					c = value[i]
				}
			}
			if (i < numChars) && (c == '%') {
				//                        throw new IllegalArgumentException(
				//                         "URLDecoder: Incomplete trailing escape (%) pattern");
			}
			result += string(bytes[:pos])
			needToChange = true
			break
		default:
			result += string(c)
			i++
			break
		}
	}
	if needToChange {
		return result
	} else {
		return value
	}
}

func EncodeJSON(){

}

func DecodeJSON(){

}
