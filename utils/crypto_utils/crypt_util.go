package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(str string) string {
	hash := md5.New()
	defer hash.Reset()

	_, _ = hash.Write([]byte(str))           // Convert str to byte array
	return hex.EncodeToString(hash.Sum(nil)) // Convert byte array to string for return
}
