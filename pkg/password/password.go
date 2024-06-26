package password

import (
    "github.com/Cynthia/goblog/pkg/logger"

    "golang.org/x/crypto/bcrypt"
)


func Hash(password string) string {
    // GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    logger.LogError(err)

    return string(bytes)
}


func CheckHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    logger.LogError(err)
    return err == nil
}


func IsHashed(str string) bool {
    return len(str) == 60
}