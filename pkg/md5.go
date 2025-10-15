package pkg

import (
	"crypto/md5"
	"fmt"
)

// 双层md5加密
func DoubleMd5(str string) string {
	//第一次
	data := []byte(str)
	has1 := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has1) //将[]byte转成16进制
	//第二次
	data2 := []byte(md5str1)
	has2 := md5.Sum(data2)
	md5str2 := fmt.Sprintf("%x", has2) //将[]byte转成16进制
	return md5str2
}
