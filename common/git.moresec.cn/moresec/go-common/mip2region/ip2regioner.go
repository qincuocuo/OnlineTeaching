package mip2region

import (
	"os"
	"strings"
	"sync"

	"git.moresec.cn/moresec/go-common/mbase"

	"github.com/labstack/gommon/log"
)

var (
	//ip 库.
	ip2RegionHandle *Ip2Region

	mtxIpregion sync.Mutex
)

func InitIp2Region(ipDbPath string) bool {
	_, err := os.Stat(ipDbPath)
	if os.IsNotExist(err) {
		log.Error("ip2region.db not exist.")
		return false
	}

	ip2RegionHandle, err = New(ipDbPath)
	if err != nil {
		log.Error("init ip2region err.")
		return false
	}
	return true
}

func GetIpLocation(ipAddr string) (string, error) {
	if !mbase.IsIP(ipAddr) || mbase.IsIntranet(ipAddr) {
		return "局域网", nil
	}

	mtxIpregion.Lock()
	ipInfo, err := ip2RegionHandle.MemorySearch(ipAddr)
	mtxIpregion.Unlock()

	if err != nil {
		return "", err
	}

	var result []string
	if len(ipInfo.Country) > 0 && ipInfo.Country != "0" {
		result = append(result, ipInfo.Country)
	}
	if len(ipInfo.Province) > 0 && ipInfo.Province != "0" {
		result = append(result, ipInfo.Province)
	}
	if len(ipInfo.City) > 0 && ipInfo.City != "0" {
		result = append(result, ipInfo.City)
	}
	return strings.Join(result, "-"), nil
}

func Close() {
	if ip2RegionHandle != nil {
		ip2RegionHandle.Close()
	}
}
