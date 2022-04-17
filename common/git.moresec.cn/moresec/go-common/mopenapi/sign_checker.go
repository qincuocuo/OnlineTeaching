package mopenapi

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
)

type GetSecretKey func(accessKey string) (string, error)

type SignChecker interface {
	Check(req *http.Request, getSecretKey GetSecretKey) error
	Sign(req *http.Request, accessKey string, getSecretKey GetSecretKey) error
}

type defaultSignChecker struct {
}

func (d *defaultSignChecker) Check(req *http.Request, getSecretKey GetSecretKey) error {
	buff, err := getBuff(req)
	if err != nil {
		return err
	}

	bodyHash := hexSha256(buff)

	ts := req.Header.Get(headerTimestamp)
	err = checkTime(ts)
	if err != nil {
		return err
	}

	secretKey, err := getSecretKey(req.Header.Get(headerAccessKey))
	if err != nil {
		return err
	}

	signature := base64HmacSha256([]byte(secretKey), []byte(assemble(ts, bodyHash, req)))
	if signature != req.Header.Get(headerSignature) {
		return ErrSignDiff
	}
	return nil
}

func (d *defaultSignChecker) Sign(req *http.Request, accessKey string, getSecretKey GetSecretKey) error {
	buff, err := getBuff(req)
	if err != nil {
		return err
	}

	bodyHash := hexSha256(buff)

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	req.Header.Set(headerSignatureVersion, versionV1)
	req.Header.Set(headerSignatureMethod, methodHmacSha256)
	req.Header.Set(headerTimestamp, ts)
	req.Header.Set(headerSignatureNonce, uuid.NewV4().String())
	req.Header.Set(headerAccessKey, accessKey)
	secretKey, err := getSecretKey(accessKey)
	if err != nil {
		return err
	}
	req.Header.Set(headerSignature, base64HmacSha256([]byte(secretKey), []byte(assemble(ts, bodyHash, req))))
	return nil
}
