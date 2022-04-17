package mopenapi

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type bodyReadCloser struct {
	io.Reader
	io.Closer
}

var (
	// formDataCheckLength read MAX length multipart/form-data body
	// if formDataCheckLength lt 1,skip read body
	formDataCheckLength int64 = 1 << 12
	//timeRange use for check timestamp header
	//if timeRange lt 1,skip
	timeRange int64 = 15 * 60
	//SecretKeyLength secret key length
	SecretKeyLength = 32
)

func UpdateFormDataCheckLength(length int64) {
	if length < 0 {
		return
	}
	formDataCheckLength = length
}

func UpdateTimeRange(tr int64) {
	timeRange = tr
}

func UpdateSecretKeyLength(length int) {
	if length < 1 {
		return
	}
	SecretKeyLength = length
}

const (
	versionV1        = "1.0"
	methodHmacSha256 = "HMAC-SHA256"

	//header name
	headerSignatureMethod  = "SignatureMethod"
	headerSignatureVersion = "SignatureVersion"
	headerSignatureNonce   = "SignatureNonce"
	headerTimestamp        = "Timestamp"
	headerSignature        = "Signature"
	headerAccessKey        = "AccessKey"
)

func getBuff(req *http.Request) (buff []byte, err error) {
	if strings.Split(req.Header.Get("Content-Type"), ";")[0] == "multipart/form-data" {
		buff, err = ioutil.ReadAll(io.LimitReader(req.Body, formDataCheckLength))
		if err != nil {
			return
		}

		req.Body = bodyReadCloser{
			Reader: io.MultiReader(bytes.NewReader(buff), req.Body),
			Closer: req.Body,
		}
		return
	}

	buff, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	err = req.Body.Close()
	if err != nil {
		return
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(buff))
	return
}

// Sign sign req body && set header info
func Sign(req *http.Request, accessKey string, getSecretKey GetSecretKey, options ...SignCheckerOption) error {
	getter := &signCheckGetter{
		signVersion: versionV1,
		signMethod:  methodHmacSha256,
	}
	for _, option := range options {
		option(getter)
	}

	checker, ok := signCheckM[getter.Key()]
	if !ok {
		return ErrSignNotSupported
	}

	return checker.Sign(req, accessKey, getSecretKey)
}

var (
	ErrSignDiff          = errors.New("signature diff error")
	ErrTimestampEmpty    = errors.New("empty timestamp error")
	ErrMethodEmpty       = errors.New("empty method error")
	ErrVersionEmpty      = errors.New("empty version error")
	ErrCheckerEmpty      = errors.New("empty checker func error")
	ErrTimestampOutRange = errors.New("timestamp out of range error")
	ErrSignNotSupported  = errors.New("signature not supported")
)

// CheckSign check client request signature && timestamp
func CheckSign(req *http.Request, getSecretKey func(accessKey string) (string, error)) error {
	methodVersion := req.Header.Get(headerSignatureMethod) + ":" + req.Header.Get(headerSignatureVersion)
	checker, ok := signCheckM[methodVersion]
	if !ok {
		return ErrSignNotSupported
	}

	return checker.Check(req, getSecretKey)
}

func checkTime(ts string) error {
	if ts == "" {
		return ErrTimestampEmpty
	}

	tsI, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}

	// skip check
	if timeRange < 1 {
		return nil
	}

	if time.Now().Unix()-tsI > timeRange {
		return ErrTimestampOutRange
	}
	return nil
}

func GenerateKey() (string, string) {
	return UUID(), randStr()
}

var signCheckM map[string]SignChecker

func init() {
	signCheckM = make(map[string]SignChecker)
	signCheckM[methodHmacSha256+":"+versionV1] = &defaultSignChecker{}
}

func NewSignature(method, version string, signChecker SignChecker) error {
	if method == "" {
		return ErrMethodEmpty
	}
	if version == "" {
		return ErrVersionEmpty
	}

	if signChecker == nil {
		return ErrCheckerEmpty
	}

	signCheckM[method+":"+version] = signChecker
	return nil
}
