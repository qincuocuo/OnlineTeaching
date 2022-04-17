package mredis

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	InitEnv()
}

func TestRedisCli_Append(t *testing.T) {
	testCli.Del("TestRedisCli_Append")
	defer testCli.Del("TestRedisCli_Append")

	_, err := testCli.Set("TestRedisCli_Append", "test")
	if err != nil {
		t.Error(err.Error())
	}
	result, _ := testCli.Get("TestRedisCli_Append")
	if result != "test" {
		t.Error(result)
	}
	result, _ = testCli.Get("TestRedisCli_Append_Nil")
	if result != "" {
		t.Error(result)
	}
}

func TestRedisCli_Decrby(t *testing.T) {
	testCli.Del("TestRedisCli_Decrby")
	defer testCli.Del("TestRedisCli_Decrby")

	result, err := testCli.Decr("TestRedisCli_Decrby")
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
	if result != -1 {
		t.Fail()
		return
	}
	result, _ = testCli.Decr("TestRedisCli_Decrby")
	if result != -2 {
		t.Fail()
		return
	}

	result, _ = testCli.DecrBy("TestRedisCli_Decrby", 1)
	if result != -3 {
		t.Fail()
		return
	}

	result, _ = testCli.Incr("TestRedisCli_Decrby")
	result, _ = testCli.IncrBy("TestRedisCli_Decrby", 2)
	if result != 0 {
		t.Fail()
		return
	}
	result, err = testCli.IncrBy("TestRedisCli_Decrby", 2100000000)
	if result != 2100000000 {
		t.Fatalf(err.Error())
	}
}

func TestRedisCli_Exists(t *testing.T) {
	testCli.Del("TestRedisCli_Exists")
	defer testCli.Del("TestRedisCli_Exists")

	testCli.Set("TestRedisCli_Exists", "fuck")
	rt, err := testCli.Exists("TestRedisCli_Exists")
	if err != nil || !rt {
		t.Error(err.Error())
		return
	}
	testCli.Del("TestRedisCli_Exists")
	rt, _ = testCli.Exists("TestRedisCli_Exists")
	if rt {
		t.Fail()
		return
	}
}

func TestRedisCli_TTL(t *testing.T) {
	testCli.Del("TestRedisCli_TTL")
	defer testCli.Del("TestRedisCli_TTL")

	testCli.SetEx("TestRedisCli_TTL", "test", 100)
	rt, err := testCli.TTL("TestRedisCli_TTL")
	if err != nil || rt != 100 {
		t.Error(err.Error())
		return
	}
	testCli.Set("TestRedisCli_TTL_no", "test")
	rt, err = testCli.TTL("TestRedisCli_TTL_no")
	if err != nil || rt != -1 {
		t.Error(err.Error())
		return
	}

	rt, err = testCli.TTL("TestRedisCli_TTL_err")
	if err != nil || rt != -2 {
		t.Error(err.Error())
		return
	}
	testCli.Del("TestRedisCli_NX_EX")
	rt, err = testCli.SetNxEx("TestRedisCli_NX_EX", "value", 100)
	if err != nil || rt != 1 {
		t.Fail()
	}
	rt, err = testCli.TTL("TestRedisCli_NX_EX")
	if err != nil || rt != 100 {
		t.Fail()
		return
	}
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
