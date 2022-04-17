# go-common

## Usage

.gitconfig 配置文件中添加如下变量(文件不存在的话创建)

```
[url "git@git.moresec.cn:"]
	insteadOf = https://git.moresec.cn/
```

然后在项目目录中执行

```sh
go get --insecure git.moresec.cn/moresec/go-common
```

## Arch

```
.
├── README.md
├── mcontainer
├── mconv
├── merrors
├── mbase
├── mdb
├── mip2region
├── mkafka
├── mlicense
├── mlog
├── mnet
├── mpkg
├── mrabbit
├── mredis
├── msingleton
├── mutils
└── mversion
```

## mbase/slice

    提供InSlice函数,判断切片中是否包含某元素

示例

```go
package main

import (
	"git.moresec.cn/moresec/go-common/mbase"
)

func main() {
	s := []string{"a", "b", "c"}
	println(mbase.InSliceAny("a", s)) // true
}
```

## mcontainer/mset

    提供set容器,包含IntSet,StrSet,AnySet

示例

```go
package main

import (
	"git.moresec.cn/moresec/go-common/mcontainer/mset"
)

func main() {
	s1 := mset.NewSet()
	s2 := mset.NewFrom([]string{"a", "b"})

	s1.Add("a")
	s1.Add("b")
	s1.Add("c")

	s1.Remove("c")

	println(s1.Size())    // 2
	println(s1.Slice())   // []interface{}{"a", "b"}
	println(s1.Equal(s2)) // true
}
```

## merrors/merror

    error封装,提供了code,异常栈,以及Cause,Wrap等常用操作

## mconv

    提供常见数据类型转换
    例如转任意类型到指定类型
    any -> int
    any -> map[string]interface{}
    any -> struct

示例

```go
package main

import (
	"fmt"

	"git.moresec.cn/moresec/go-common/mconv"
)

type User struct {
	Uid  int    `c:"uid"`
	Name string `c:"name"`
}

func main() {
	i := 123.456
	fmt.Printf("%10s %v\n", "Int:", mconv.Int(i))
	fmt.Printf("%10s %v\n", "Bool:", mconv.Bool(i))
	fmt.Printf("%10s %v\n", "String:", mconv.String(i))
	fmt.Printf("%10s %v\n", "Ints:", mconv.Ints(i))
	fmt.Printf("%10s %v\n", "Strings:", mconv.Strings(i))
	fmt.Printf("%10s %v\n", "Interfaces:", mconv.Interfaces(i))

	// struct -> map[string]interface{}
	mconv.Map(User{
		Uid:  1,
		Name: "john",
	})

	// pointer -> map[string]interface{}
	mconv.Map(&User{
		Uid:  1,
		Name: "john",
	})

	// any -> map[string]interface{}
	mconv.Map(map[int]int{
		100: 10000,
	})

	// map[string]interface{} -> struct
	params := map[string]interface{}{
		"uid":  1,
		"name": "john",
	}
	var user *User
	mconv.Struct(params, &user)
}
```

## mnet（网络基本处理库）

目录结构:

```
-mnet
  ｜-aes.go
  ｜-aes_test.go
  ｜-ipv4.go
  ｜-ipv4_ip.go
  ｜-ipv4_lookup.go
  ｜-ipv4_mac.go
  ｜-nethead.go
  ｜-nethead_test.go
```

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
