package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash/crc32"
	"math/rand"
	"os"
	"runtime"
	"time"
)

const (
	RandChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// 获取正在运行的函数名
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// 获取随机字符串
func RandStr(l int) string {
	le := len(RandChars)
	data := make([]byte, l, l)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		data[i] = byte(RandChars[rand.Intn(le)])
	}
	return string(data)
}

// md5
func Md5(s string) string {
	c := md5.New()
	c.Write([]byte(s))
	return hex.EncodeToString(c.Sum(nil))
}

// sha1
func SHA1(str string) string {
	c := sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

// crc32
func CRC32(str string) uint32  {
	return crc32.ChecksumIEEE([]byte(str))
}

// 获取主机名
func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	return name
}