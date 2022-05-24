package mbase

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

type mFile struct{}

var File mFile

// Exists 判断文件是否存在
func (mFile) Exists(path string) bool {
	if _, err := os.Stat(path); (err != nil && os.IsExist(err)) || err == nil {
		return true
	}
	return false
}

// IsDir 判断文件类型是否为目录
func (mFile) IsDir(path string) bool {
	ret, err := os.Stat(path)
	if err != nil {
		return false
	}
	return ret.IsDir()
}

// CreateDir 递归创建目录
func (mFile)CreateDir(filePath string) error {
	if !File.Exists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// Md5 计算文件Md5值
func (mFile) Md5(filePath string) (md5Str string, err error) {
	var f *os.File
	defer f.Close()
	hash := md5.New()
	if f, err = os.Open(filePath); err != nil {
		return
	}
	_, err = io.Copy(hash, f)
	md5Str = hex.EncodeToString(hash.Sum(nil))
	return
}

// GetSize 计算文件大小
func (mFile) GetSize(filePath string) (size int64, err error) {
	var f os.FileInfo
	if f, err = os.Stat(filePath); err != nil {
		return
	}
	size = f.Size()
	return
}

// AppendWrite  文件追加写入,不存在即创建
func (mFile) AppendWrite(filePath string, context string) (err error) {
	fileHandler, _ := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(context + "\n")
	_, err = fileHandler.Write(buf)
	err = fileHandler.Close()
	return
}
