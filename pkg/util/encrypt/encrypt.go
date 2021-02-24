package encrypt

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/bcrypt"
)

const pwHashBytes = 32

// MD5Encode MD5 加密
func MD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// PasswordEncode 密码加密
func PasswordEncode(password string) string {
	pwd := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hash)
}

// PasswordCompare 验证密码
func PasswordCompare(hashedPassword string, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	bytePwd := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

// Salt 随机盐
func Salt() string {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return ""
	}
	return hex.EncodeToString(buf)
}
