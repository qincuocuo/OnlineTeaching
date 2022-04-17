package mopenapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"hash"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func base64HmacSha256(secret, data []byte) string {
	return Encode(&hashEncoder{
		h:       hmac.New(sha256.New, secret),
		encodeF: base64.StdEncoding.EncodeToString,
	}, data)
}

func hexSha256(data []byte) string {
	return Encode(&hashEncoder{
		h:       sha256.New(),
		encodeF: hex.EncodeToString,
	}, data)
}

type Encoder interface {
	Encode(data []byte) string
}

type hashEncoder struct {
	h       hash.Hash
	encodeF func([]byte) string
}

func (he *hashEncoder) Encode(data []byte) string {
	defer he.h.Reset()
	he.h.Write(data)
	return he.encodeF(he.h.Sum(nil))
}

func Encode(encoder Encoder, data []byte) string {
	return encoder.Encode(data)
}

func assemble(ts, bodyHash string, req *http.Request) string {
	return ts + "\n" + req.Method + "\n" + req.URL.Query().Encode() + "\n" + bodyHash
}

func UUID() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func randStr() string {
	dict := "abcdefghijklmdopqestuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*(){}[]-_=+,.<>?/"
	rand.Seed(time.Now().UnixNano())

	data := ""
	for i := 0; i < SecretKeyLength; i++ {
		data += string(dict[rand.Intn(len(dict))])
	}
	return data
}
