package emas

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/hmac"
	"fmt"
)

func Md5Encrypt(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}


func ShaOneEncrypt(s string) (string) {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func Sha256Encrypt(s string) (string) {
	h := sha256.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func Sha512Encrypt(s string) (string) {
	h := sha512.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func Sha256_HMAC(k string, message string) (string) {
	key := []byte(k)
	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(message))

	hmac := fmt.Sprintf(hex.EncodeToString(sig.Sum(nil)))

	return hmac
}

func Sha1_HMAC(k string, message string) (string) {
	key := []byte(k)
	sig := hmac.New(sha1.New, key)
	sig.Write([]byte(message))

	hmac := fmt.Sprintf(hex.EncodeToString(sig.Sum(nil)))

	return hmac
}
