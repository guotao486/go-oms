package util

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 随机字符串
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 使用当前时间创建seed
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// 创建Salt
func CreateSalt() string {
	b := make([]byte, bcrypt.DefaultCost)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// 使用bcrypt 算法加密密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 检查密码是否相等
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
