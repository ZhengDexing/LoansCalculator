package util

import "encoding/base64"

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(encodeString string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodeString)
}
