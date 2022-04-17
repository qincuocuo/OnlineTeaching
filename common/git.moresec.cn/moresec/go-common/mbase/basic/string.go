package mbase

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	randomId "github.com/satori/go.uuid"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type str struct{}

var String str

// ToInt 字符串转换为整型
func (str) ToInt(src string) (dst int) {
	dst, _ = strconv.Atoi(src)
	return
}

// ToInt32 字符串转换为32位整型
func (str) ToInt32(src string) (dst int32) {
	tmp, _ := strconv.ParseInt(src, 10, 32)
	dst = int32(tmp)
	return
}

// ToInt64 字符串转换为64位整型
func (str) ToInt64(src string) (dst int64) {
	dst, _ = strconv.ParseInt(src, 10, 64)
	return
}

// ToFloat64 字符串转换为64位浮点型
func (str) ToFloat64(src string) (dst float64) {
	var err error
	if dst, err = strconv.ParseFloat(src, 64); err != nil {
		dst = 0.0
	}
	return
}

// IntToString 整型转化为字符串
func (str) IntToString(src int) (dst string) {
	dst = strconv.Itoa(src)
	return dst
}

// Int32ToString 32位整型转化为字符串
func (str) Int32ToString(src int32) (dst string) {
	dst = strconv.Itoa(int(src))
	return dst
}

// Float32ToString 32为位浮点型转化为字符串
func (str) Float32ToString(src float32) (dst string) {
	dst = fmt.Sprintf("%f", src)
	return dst
}

// Float64ToString 64位整型转化位字符串
func (str) Float64ToString(src float64) (dst string) {
	dst = fmt.Sprintf("%f", src)
	return dst
}

// UpperFirst 字符串首字母大写
func (str) UpperFirst(src string) (result string) {
	result = strings.ToUpper(src[:1]) + src[1:]
	return
}

// LowerFirst 字符串首字母小写
func (str) LowerFirst(src string) (result string) {
	result = strings.ToLower(src[:1]) + src[1:]
	return
}

// GetRandomString 生成随机字符串
func (str) GetRandomString(len int) string {
	return String.MD5String(randomId.NewV1().String(), len)
}

// MD5String 生成32位MD5值
func (str) MD5String(key string, length int) string {
	h := md5.New()
	h.Write([]byte(key))
	result := hex.EncodeToString(h.Sum(nil))
	if length > len(result) {
		length = len(result)
	}
	return result[:length]
}

// Md5 计算字符串md5值
func (str) Md5(str ...string) (md5Str string) {
	var srcStr string
	srcStr = strings.Join(str, "")
	data := []byte(srcStr)
	md5Str = fmt.Sprintf("%x", md5.Sum(data))
	return
}

// Base64Decrypt 计算字符串Base64
func (str) Base64Decrypt(str string) (decrypt string) {
	var decryptByte []byte
	decryptByte, _ = base64.StdEncoding.DecodeString(str)
	decrypt = string(decryptByte)
	return decrypt
}

// GetPasswordStrength 判断密码字符串密码强度
func (str) GetPasswordStrength(passwd string) (passwdLevel int) {
	counter := 0
	numberMatch := "[0-9]+"
	lowLetter := "[a-z]+"
	upLetter := "[A-Z]+"
	specialSymbol := `[*()~!@#$%^&*-+=_|:;'<>,.?/\[\]\{\}<>]+`
	if match, _ := regexp.MatchString(numberMatch, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(lowLetter, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(upLetter, passwd); match {
		counter++
	}
	if match, _ := regexp.MatchString(specialSymbol, passwd); match {
		counter++
	}
	if len(passwd) < 8 || counter <= 1 {
		passwdLevel = 0
	} else if counter <= 2 {
		passwdLevel = 1
	} else if counter <= 3 {
		passwdLevel = 2
	} else {
		passwdLevel = 3
	}
	return
}

// KeyInMap 判断字符串是否存在于map的key中
func (str) KeyInMap(key string, src map[string]string) bool {
	if _, ok := src[key]; ok {
		return true
	}
	return false
}

// Compare 判断两个字符串是否相同
func (str) Compare(str1, str2 string) bool {
	if strings.Compare(str1, str2) == 0 {
		return true
	}
	return false
}

func (str) IsIp(ipStr string) bool {
	v := net.ParseIP(ipStr)
	if v == nil {
		return false
	}
	return true
}
