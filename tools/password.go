package tools

import "golang.org/x/crypto/bcrypt"

// 密码加密
func EncodePassword(rawPassword string) string  {
	hash, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hash)
}

// 验证密码
func ValidatePassword(encodePassword, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(rawPassword))
	return err == nil
}
