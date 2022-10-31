package models

import (
	"crypto/md5"
	"fmt"
	"time"
)

// 将时间戳转换成日期
func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳 2022-10-25 21:35:00
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetDate() string {
	template := "20060102"
	return time.Now().Format(template)
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x\n", md5.Sum(data))
}
