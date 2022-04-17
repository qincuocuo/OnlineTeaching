package mlicense

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"net"
	"time"

	"git.moresec.cn/moresec/go-common/mnet"
)

var errIp = errors.New("vm ip is empty")
var errSrvType = errors.New("service type is empty")
var errExpireTime = errors.New("license expiration time")

const (
	UNIX_PATH                   = "/var/run/license_unix_sock"
	AES_KEY                     = "0123456789abcdef"
	LICENSE_TYPE_AUTH           = 1
	LICENSE_TYPE_TARIT          = 2
	LICENSE_TYPE_UPDATE         = 3
	LICENSE_TYPE_LICENT_CONTENT = 4
    LICENSE_TYPE_JSON_DATA      = 5
    LICENSE_TYPE_GET_VERSION    = 6
)

var aesUtil *mnet.Aes = mnet.NewAes([]byte(AES_KEY), mnet.PKCS5Pad)

type reqHead struct {
	STime uint64
	ETime uint64
	Code  int32
	Count int32
	Buf   [32]byte
}

type reqMsg struct {
	Fun     int    `json:"fun"`
	SrvType string `json:"service"`
	Ip      string `json:"ip"`
	License string `json:"license"`
}

type licenseContent struct {
	Code    int    `json:"code"`
	License string `json:"license"`
	JsonData string `json:"jsonData"`
}

type license struct {
	strSrvType string
	strVmIp    string
	isVm       bool
}

func (l *license) sock() (*net.UnixConn, error) {
	addr := net.UnixAddr{UNIX_PATH, "unix"}

	conn, err := net.DialUnix("unix", nil, &addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (l *license) send(conn *net.UnixConn, buf []byte) error {
	var outBuf []byte
	outBuf, err := aesUtil.Encrypt(string(buf))
	if err != nil {
		return err
	}
	msg, err := mnet.SetPacket(outBuf, mnet.NetHeadTypeJson)
	if err != nil {
		return err
	}

	if _, err := conn.Write(msg); err != nil {
		return err
	}
	return nil
}

func (l *license) recv(conn *net.UnixConn) ([]byte, error) {
	buf := make([]byte, 0, 1024)
	pack := make([]byte, 2048)
	for {
		n, err := conn.Read(pack)
		if n == 0 || err != nil {
			return nil, err
		}
		tmpData := pack[0:n]
		buf = append(buf, tmpData...)
		netHead, err := mnet.GetHead(buf)
		if netHead != nil {
			if netHead.PkgLen <= uint32(len(buf)) {
				msg1 := buf[9:netHead.PkgLen]
				msg, err := aesUtil.Decrypt(msg1)
				if err != nil {
					return nil, err
				}
				return msg, nil
			}
		} else {
			return nil, err
		}
	}
	return nil, nil
}

func (l *license) init(fun int, data string) (*reqHead, error) {
	conn, err := l.sock()
	defer func() {
		if conn != nil {
			conn.Close()
		}

	}()
	if err != nil {
		return nil, err
	}
	reqmsg := reqMsg{
		SrvType: l.strSrvType,
		Fun:     fun,
		Ip:      l.strVmIp,
		License: data,
	}

	reqBuf, _ := json.Marshal(reqmsg)
	if err := l.send(conn, reqBuf); err != nil {
		return nil, err
	}

	rspBuf, err := l.recv(conn)
	if err != nil {
		return nil, err
	}

	var req reqHead
	buf := new(bytes.Buffer)
	buf.Write(rspBuf)
	if err := binary.Read(buf, binary.LittleEndian, &req); err != nil {
		return nil, err
	}

	if req.Code != 1 {
		if req.Code == -2 {
			return nil, errors.New("system time modified")
		} else {
			return nil, errors.New("get data from license_srv failed")
		}

	}

	return &req, nil
}

func (l *license) VerifyTime() error {
	head, err := l.init(LICENSE_TYPE_AUTH, "")
	if err != nil {
		return err
	}
	curTime := time.Now().Unix()
	if head.STime > uint64(curTime) || head.ETime < uint64(curTime) {
		return errExpireTime
	}
	return nil
}

func (l *license) VerifyCount(count int32) error {
	head, err := l.init(LICENSE_TYPE_AUTH, "")
	if err != nil {
		return err
	}
	if head.Count < count {
		return errors.New("property greater than license data")
	}
	return nil
}

func (l *license) GetMachineTrait() (string, error) {
	head, err := l.init(int(LICENSE_TYPE_TARIT), "")
	if err != nil {
		return "", err
	}

	return string(head.Buf[:]), nil
}

func (l *license) UpdateLicense(data string) (*LicenseData, error) {
	head, err := l.init(LICENSE_TYPE_UPDATE, data)
	if err != nil {
		return nil, err
	}

	return &LicenseData{
		StartTime: int64(head.STime),
		EndTime:   int64(head.ETime),
		Count:     int32(head.Count),
	}, nil
}

func (l *license) GetLicense() (*LicenseData, error) {
	head, err := l.init(LICENSE_TYPE_AUTH, "")
	if err != nil {
		return nil, err
	}

	return &LicenseData{
		StartTime: int64(head.STime),
		EndTime:   int64(head.ETime),
		Count:     int32(head.Count),
	}, nil
}

func (l *license) GetLicenseContent() (string, error) {
	conn, err := l.sock()
	defer func() {
		if conn != nil {
			conn.Close()
		}

	}()
	if err != nil {
		return "", err
	}
	reqmsg := reqMsg{
		SrvType: l.strSrvType,
		Fun:     LICENSE_TYPE_LICENT_CONTENT,
		Ip:      l.strVmIp,
		License: "",
	}
	reqBuf, _ := json.Marshal(reqmsg)
	if err := l.send(conn, reqBuf); err != nil {
		return "", err
	}

	rspBuf, err := l.recv(conn)
	if err != nil {
		return "", err
	}
	var lContent licenseContent
	json.Unmarshal(rspBuf, &lContent)
	if lContent.Code != 1 {
		return "", errors.New("license content is empty")
	}
	return lContent.License, nil
}

func (l *license) GetLicenseVersion() (string, error) {
	conn, err := l.sock()
	defer func() {
		if conn != nil {
			conn.Close()
		}

	}()
	if err != nil {
		return "", err
	}
	reqmsg := reqMsg{
		SrvType: l.strSrvType,
		Fun:     LICENSE_TYPE_GET_VERSION,
		Ip:      l.strVmIp,
		License: "",
	}
	reqBuf, _ := json.Marshal(reqmsg)
	if err := l.send(conn, reqBuf); err != nil {
		return "", err
	}

	rspBuf, err := l.recv(conn)
	if err != nil {
		return "", err
	}

	return string(rspBuf), nil
}

func (l *license) GetLicenseJsonData() (string, error) {
	conn, err := l.sock()
	defer func() {
		if conn != nil {
			conn.Close()
		}

	}()
	if err != nil {
		return "", err
	}
	reqmsg := reqMsg{
		SrvType: l.strSrvType,
		Fun:     LICENSE_TYPE_JSON_DATA,
		Ip:      l.strVmIp,
		License: "",
	}

	reqBuf, _ := json.Marshal(reqmsg)
	if err := l.send(conn, reqBuf); err != nil {
		return "", err
	}

	rspBuf, err := l.recv(conn)
	if err != nil {
		return "", err
	}

	var lContent licenseContent
	json.Unmarshal(rspBuf, &lContent)
	if lContent.Code != 1 {
		return "", errors.New("get license jsondata faield")
	}
	return lContent.JsonData, nil
}