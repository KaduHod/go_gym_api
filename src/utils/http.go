package utils

import (
	"crypto/md5"
	"encoding/hex"
)
func GenerateEtag(content []byte) string {
    hash := md5.Sum(content)
    return "\"" + hex.EncodeToString(hash[:]) + "\""
}
