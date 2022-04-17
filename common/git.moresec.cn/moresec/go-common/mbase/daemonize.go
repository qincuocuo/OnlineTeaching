package mbase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sevlyar/go-daemon"
)

func Daemonize(args []string) (*daemon.Context, error) {

	cntxt := &daemon.Context{
		PidFileName: "",
		PidFilePerm: 0644,
		LogFileName: "",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        args,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		fmt.Println("Unable to run: ", err)
		return cntxt, err
	}
	if d != nil {
		err = errors.New("child")
		return cntxt, err
	}
	//defer cntxt.Release()

	return cntxt, err
}

func Exit(code int) {
	time.Sleep(time.Second)
	os.Exit(code)
}
