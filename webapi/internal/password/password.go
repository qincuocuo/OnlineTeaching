package password

import (
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"webapi/utils"

	"git.moresec.cn/moresec/go-common/mnet"

	"golang.org/x/crypto/pbkdf2"
)

const (
	AesKey string = "sun*&znb666*%"
)

var (
	Aes *mnet.Aes
)

func init() {
	Aes = mnet.NewAes([]byte(AesKey), mnet.PKCS7Pad)
}

func MakePassword(passwd string) string {
	salt := utils.String.GetRandomString(12)
	dk := pbkdf2.Key([]byte(passwd), []byte(salt), 36000, 32, sha256.New)
	str := base64.StdEncoding.EncodeToString(dk)
	return "sun_sha256" + "$" + strconv.FormatInt(int64(36000), 10) + "$" + salt + "$" + str
}

func CheckPassword(passwd, encode string) bool {
	t := strings.SplitN(encode, "$", 4)
	algorithm := t[0]
	salt := t[2]
	iterations, _ := strconv.Atoi(t[1])
	digest := sha256.New
	// algorithm most be sun_sha256
	if algorithm != "sun_sha256" {
		return false
	}
	dk := pbkdf2.Key([]byte(passwd), []byte(salt), iterations, 32, digest)
	str := base64.StdEncoding.EncodeToString(dk)
	hashed := "sun_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
	return hashed == encode
}

func decryptPassword(passwd []byte) (password string, err error) {
	var decodePasswd []byte
	if decodePasswd, err = Aes.Decrypt(passwd); err != nil {
		return
	}
	password = string(decodePasswd)
	return
}
