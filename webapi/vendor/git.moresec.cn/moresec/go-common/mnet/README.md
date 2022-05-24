## 网络基本处理库
目录结构:

    -mnet
     ｜-aes.go
     ｜-aes_test.go
     ｜-ipv4.go	
     ｜-ipv4_ip.go	
     ｜-ipv4_lookup.go
     ｜-ipv4_mac.go	
     ｜-nethead.go
     ｜-nethead_test.go

### aes.go
基本的加密算法
使用golang的aes-ecb作为加密算法，分别可采用0填充、PKC5填充以及PKC7填充对数据进行预填充后加密
### `ipv4.go`
* `Ip2long` 将ip地址字符串转换为uint32类型
* `Long2ip` 将uint32类型表示的ip地址转换为ip地址字符串
* `Validate` 校验ip地址字符串是否合法
* `ParseAddress` 解析ip:port格式字符串
### ipv4_ip.go
*  `GetIpArray` 获取当前主机下所有网卡的Ip地址列表
*  `GetIntranetIp` 获取并返回当前机器的第一个内网ip
*  `GetIntranetIpArray` 检索并返回当前主机的内网IP地址列表
*  `IsIntranet` 检查并返回给定的 ip 是否为内网 ip,检测范围如下
    * Local: 127.0.0.1
    * A: 10.0.0.0--10.255.255.255
    * B: 172.16.0.0--172.31.255.255
    * C: 192.168.0.0--192.168.255.255
## ipv4_lookup
* `GetHostByName`返回与给定 Internet 主机名对应的 IPv4 地址
* `GetNameByAddr` 返回与给定 IP 地址对应的 Internet 主机名。
* `GetHostsByName` 返回与给定 Internet 对应的 IPv4 地址列表
## ipv4_mac
* `GetMacArray` 检索并返回当前主机的所有 mac 地址
* `GetMac` 检索并返回当前主机的第一个 mac 地址
## nethead.go
* `SetPacket` 使用字节流和包类型填充数据包头部
* `GetHead` 获取数据包头部
* `GetPacket`解析数据包头部为字节流和包类型