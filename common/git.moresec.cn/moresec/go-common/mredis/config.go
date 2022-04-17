package mredis

import "fmt"

// 该文件是为测试用例专用.
var (
	host = "192.168.30.140:6379"
)

var testCli *RedisCli

func InitEnv() {
	txOption := NewDefaultOption()
	txOption.URL = host
	txOption.Password = "moresec@sec"

	testCli = NewRedisCli(txOption)

	fmt.Println("Init test env success.")
}
