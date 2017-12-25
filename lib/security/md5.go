package security

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encrypt(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}