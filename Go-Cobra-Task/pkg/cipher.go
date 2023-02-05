package strman

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/ascii85"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func EncodeString(input string, format string) (result string) {
	if format == "binary" {
		for _, c := range input {
			result = fmt.Sprintf("%s%.8b", result, c)
		}
	}
	if format == "ascii" {
		ascii85.Encode([]byte(result), []byte(input))
	}
	if format == "hex" {
		result = hex.EncodeToString([]byte(input))
	}
	if format == "base64" {
		result = base64.StdEncoding.EncodeToString([]byte(input))
	}
	return result
}

func DecodeString(input string, format string) (result string) {
	if format == "ascii" {
		ascii85.Decode([]byte(result), []byte(input), true)
	}
	if format == "hex" {
		res, _ := hex.DecodeString(input)
		result = string(res)
	}
	if format == "base64" {
		res, _ := base64.StdEncoding.DecodeString(input)
		result = string(res)
	}
	return result
}

func HashString(input string, format string) (result string) {
	if format == "md5" {
		hash := md5.Sum([]byte(input))
		result = hex.EncodeToString(hash[:])
	}
	if format == "sha1" {
		hasher := sha1.New()
		hasher.Write([]byte(input))
		result = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	}
	if format == "sha256" {
		hasher := sha256.New()
		hasher.Write([]byte(input))
		result = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	}
	if format == "sha512" {
		hasher := sha512.New()
		hasher.Write([]byte(input))
		result = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	}
	return result
}
