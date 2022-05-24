package mnet

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var (
	errPkgHead = errors.New("invalid packet head")
	errpkg     = errors.New("invalid packet")
)

const (
	NetSartSign         uint32 = 0x11111111
	NetHeadTypeJson     uint8  = 0x02
	NetHeadTypeProtoBuf uint8  = 0x03
	NetHeadLen          int    = 9
)

type NetHead struct {
	Start  uint32
	Type   uint8
	PkgLen uint32
}

func SetPacket(pkg []byte, pkgtype uint8) (out []byte, err error) {
	var netHead NetHead
	netHead.Start = NetSartSign
	netHead.Type = pkgtype
	netHead.PkgLen = uint32(binary.Size(netHead)) + uint32(len(pkg))

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, netHead); err != nil {
		return nil, err
	}
	if _, err := buf.Write(pkg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GetHead(pkg []byte) (head *NetHead, err error) {
	var netHead NetHead
	headLen := binary.Size(netHead)
	if len(pkg) < headLen {
		return nil, nil
	}

	buf := new(bytes.Buffer)
	buf.Write(pkg[0:headLen])
	if err := binary.Read(buf, binary.BigEndian, &netHead); err != nil {
		return nil, err
	}

	if netHead.Start == NetSartSign {
		return &netHead, nil
	}

	return nil, errPkgHead
}

func GetPacket(pkg []byte) (data []byte, pkgType uint8, err error) {
	netHead, err := GetHead(pkg)
	if err != nil {
		return nil, 0, err
	}
	headLen := uint32(binary.Size(netHead))
	dataLen := netHead.PkgLen - headLen

	packet := pkg[headLen:netHead.PkgLen]

	if uint32(len(packet)) != dataLen {
		return nil, netHead.Type, errpkg
	}

	return packet, netHead.Type, nil
}
