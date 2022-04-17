/*
@Time : 2020-11-16 14:28
@Author : gaoxl@moresec.cn
@Description:
@Software: GoLand
*/
package mlicense

import (
	"errors"
	"net"

	"git.moresec.cn/moresec/go-common/mnet"
	jsoniter "github.com/json-iterator/go"
)

type licenseV3 struct {
	serviceType string
	vmIp        string
	addr        string
	network     string
}

func (l *licenseV3) socket() (net.Conn, error) {
	network := l.network
	addr := l.addr
	if len(network) == 0 {
		network = "unix"
	}
	if len(addr) == 0 {
		addr = "/var/run/license_unix_sock"
	}

	return net.Dial(network, addr)
}

func (l *licenseV3) init(fun, assetCount int32, data string) (*LicenseResponse, error) {
	conn, err := l.socket()
	defer func() {
		if conn != nil {
			conn.Close()
		}

	}()
	if err != nil {
		return nil, err
	}
	req := LicenseRequest{
		Service:    l.serviceType,
		ReqType:    fun,
		AssetCount: assetCount,
		Ip:         l.vmIp,
		License:    data,
	}

	reqBuf, _ := jsoniter.Marshal(req)
	if err := l.send(conn, reqBuf); err != nil {
		return nil, err
	}

	rspBuf, err := l.recv(conn)
	if err != nil {
		return nil, err
	}

	var rsp LicenseResponse

	err = jsoniter.Unmarshal(rspBuf, &rsp)
	return &rsp, err
}

func (l *licenseV3) send(conn net.Conn, buf []byte) error {
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

func (l *licenseV3) recv(conn net.Conn) ([]byte, error) {
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

func (l *licenseV3) VerifyTime() error {
	head, err := l.init(LICENSE_TYPE_VERIFY_TIME_V3, 0, "")
	if err != nil {
		return err
	}
	if head.Code != RET_CODE_OK {
		return errExpireTime
	}
	return nil
}

func (l *licenseV3) VerifyCount(count int32) error {
	head, err := l.init(LICENSE_TYPE_VERIFY_COUNT_V3, count, "")
	if err != nil {
		return err
	}
	if head.Code != RET_CODE_OK {
		return errors.New("property greater than license data")
	}
	return nil
}

func (l *licenseV3) GetMachineTrait() (string, error) {
	rsp, err := l.init(LICENSE_TYPE_TARIT_V3, 0, "")
	if err != nil {
		return "", err
	}
	return string(rsp.Data), nil
}

func (l *licenseV3) UpdateLicense(data string) (*LicenseData, error) {
	rsp, err := l.init(LICENSE_TYPE_UPDATE_V3, 0, data)
	if err != nil {
		return nil, err
	}

	if rsp.Code != RET_CODE_OK {
		return nil, errors.New(rsp.Msg)
	}
	return &LicenseData{
		StartTime: int64(rsp.StartTime),
		EndTime:   int64(rsp.EndTime),
		Count:     int32(rsp.Count),
	}, nil
}

func (l *licenseV3) GetLicense() (*LicenseData, error) {
	rsp, err := l.init(LICENSE_TYPE_LICENSE_V3, 0, "")
	if err != nil {
		return nil, err
	}

	if rsp.Code != RET_CODE_OK {
		return nil, errors.New(rsp.Msg)
	}
	return &LicenseData{
		StartTime: int64(rsp.StartTime),
		EndTime:   int64(rsp.EndTime),
		Count:     int32(rsp.Count),
	}, nil
}

func (l *licenseV3) GetLicenseContent() (string, error) {
	rsp, err := l.init(LICENSE_TYPE_LICENSE_CONTENT_V3, 0, "")
	if err != nil {
		return "", err
	}

	if rsp.Code != RET_CODE_OK {
		return "", errors.New(rsp.Msg)
	}
	return string(rsp.Data), nil
}

func (l *licenseV3) GetLicenseVersion() (string, error) {
	rsp, err := l.init(LICENSE_TYPE_VERSION_V3, 0, "")
	if err != nil {
		return "", err
	}

	if rsp.Code != RET_CODE_OK {
		return "", errors.New(rsp.Msg)
	}
	return string(rsp.Data), nil
}

func (l *licenseV3) GetLicenseJsonData() (string, error) {
	rsp, err := l.init(LICENSE_TYPE_JSON_DATA_V3, 0, "")
	if err != nil {
		return "", err
	}

	if rsp.Code != RET_CODE_OK {
		return "", errors.New(rsp.Msg)
	}
	return string(rsp.Data), nil
}
