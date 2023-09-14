package encrypts

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
)

func Md5(str string) string {
	hash := md5.New()
	_, _ = io.WriteString(hash, str)
	return hex.EncodeToString(hash.Sum(nil))
}

func GetUserID() int64 {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间的纳秒部分作为随机数种子
	min := 100000                    // 最小的6位数
	max := 999999                    // 最大的6位数
	return int64(rand.Intn(max-min+1) + min)
}
