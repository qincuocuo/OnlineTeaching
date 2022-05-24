package mbase

import (
	"io/ioutil"
	"net"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// 得到内网IP列表.
func GetInternalIP() []string {
	result, err := IntranetIP()
	if err != nil || len(result) == 0 {
		return []string{"127.0.0.1"}
	}
	return result
}

func IntranetIP() (ips []string, err error) {
	ips = make([]string, 0)

	ifaces, e := net.Interfaces()
	if e != nil {
		return ips, e
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		// ignore docker and warden bridge
		if strings.HasPrefix(iface.Name, "docker") || strings.HasPrefix(iface.Name, "w-") {
			continue
		}

		addrs, e := iface.Addrs()
		if e != nil {
			return ips, e
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}

			ipStr := ip.String()
			if IsIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}

	return ips, nil
}

func IsIntranet(ipStr string) bool {
	if strings.HasPrefix(ipStr, "10.") || strings.HasPrefix(ipStr, "192.168.") {
		return true
	}
	if strings.HasPrefix(ipStr, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(ipStr, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}

// 从配置文件中读取本机IP信息.
func GetInternalIPFromSysFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	ipStr := string(data)
	// 检验IP是不是在本地IP列表中.
	ipList := GetInternalIP()
	for i := 0; i < len(ipList); i++ {
		if ipStr == ipList[i] {
			return ipStr, nil
		}
	}
	return "", errors.New("system ip config is not in common iplist")
}

// ip合法性校验.
func IsIP(ip string) (b bool) {
	v := net.ParseIP(ip)
	if v == nil {
		return false
	}
	//if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
	//	return false
	//}
	return true
}

// 判断是否在同一网段或同一IP.
// srcIp 被判断的IP， ipDict 可以是合法 Ip 或 合法网段.
func NetContain(ipDict string, srcIp string) bool {
	src := net.ParseIP(srcIp)
	if src == nil {
		return false
	}

	// 如果是IP.
	ip := net.ParseIP(ipDict)
	if ip != nil {
		return ip.Equal(src)
	}

	// 如果是网段.
	_, n, err := net.ParseCIDR(ipDict)
	if err != nil {
		return false
	}
	return n.Contains(src)
}
