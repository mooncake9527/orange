package hashUtil

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// sha256
func SHA256(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

// MD5
func MD5(data []byte) string {
	h := md5.Sum(data)
	return hex.EncodeToString(h[:])
}
