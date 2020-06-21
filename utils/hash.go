package utils

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
)

const PwdSalt = "abc"

func GenMD5Pwd(pwd string) (string, error) {
	return GenMD5(pwd + PwdSalt)
}

func GenMD5(str string) (string, error) {
	m := md5.New()
	if _, err := m.Write([]byte(str)); err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}

func GenUUID() string {

	return uuid.NewV4().String()
}

