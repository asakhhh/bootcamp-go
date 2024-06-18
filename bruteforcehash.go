package bootcamp

import (
	"crypto/md5"
)

func hexToByte(c byte) byte {
	if '0' <= c && c <= '9' {
		return c - '0'
	}
	if 'a' <= c && c <= 'f' {
		return c - 'a' + 10
	}
	return c - 'A' + 10
}

// hexStringToByteArray converts a valid hex string to a [16]byte array.
func hexStringToByteArray(hexStr string) [16]byte {
	var result [16]byte
	for i := 0; i < 16; i++ {
		highNibble := hexToByte(hexStr[2*i])
		lowNibble := hexToByte(hexStr[2*i+1])
		result[i] = (highNibble << 4) | lowNibble
	}
	return result
}

func BruteForceHash1(targetHash [16]byte) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	maxLength := 6

	var helper func(prefix string, length int) string
	helper = func(prefix string, length int) string {
		if length > maxLength {
			return ""
		}

		h := md5.Sum([]byte(prefix))
		if h == targetHash {
			return prefix
		}

		for _, c := range chars {
			result := helper(prefix+string(c), length+1)
			if result != "" {
				return result
			}
		}
		return ""
	}

	return helper("", 0)
}

func BruteForceHash(hash string) string {
	if len(hash) != 32 {
		return ""
	}
	for _, c := range hash {
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') {
			return ""
		}
	}

	return BruteForceHash1(hexStringToByteArray(hash))
}

// func main() {
// 	fmt.Println(BruteForceHash("ab6ccd17455d5347c49606d641e0b2af")) // SALEM
// 	fmt.Println(BruteForceHash("3cbfa33db66b830bfcf47ecc956505f8")) // ALEM
// 	fmt.Println(BruteForceHash(""))                                 //
// 	fmt.Println(BruteForceHash("abc"))                              //
// }
