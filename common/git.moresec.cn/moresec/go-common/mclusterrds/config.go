package mclusterrds

import (
	"context"
	"fmt"
)

// 该文件是为测试用例专用.
var (
	host        = "192.168.30.140:6379"
	clusterHost = []string{"192.168.30.140:7000", "192.168.30.140:7001", "192.168.30.140:7002", "192.168.30.140:7003"}
)

var testCli *RedisCli

func initEnv() {
	initClusterEnv()
	//initSentinelEnv()
}

func initNormalEnv() {
	txOption := NewDefaultOption()
	txOption.Addr = host

	testCli = NewRedisCli(txOption)
	fmt.Println("Init test env success.")
}

func initClusterEnv() {
	txOption := NewDefaultOption()
	txOption.AddrList = clusterHost
	testCli = NewClusterCli(txOption)
	testCli.redisCli.Ping(context.TODO())
	fmt.Println("Init test env success.")
}

func initSentinelEnv() {
	txOption := NewDefaultOption()
	txOption.AddrList = []string{"192.168.120.160:26379", "192.168.120.160:26380"}
	txOption.Password = "moresec#sec"
	txOption.SentinelPassword = "moresec#sec"
	txOption.MasterName = "master"
	testCli = NewSentinelCli(txOption)
	testCli.redisCli.Ping(context.TODO())
	fmt.Println("Init test env success.")
}
