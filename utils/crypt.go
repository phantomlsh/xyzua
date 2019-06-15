package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5
func MD5(s string, salt string) string {
	h := md5.New()
	h.Write([]byte(s))
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256
func SHA256(s string, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(s + salt))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}

// the HASH function of XYZUA
func HASH(s string, salt string) string {
	return MD5(SHA256(s, salt), salt)
}
