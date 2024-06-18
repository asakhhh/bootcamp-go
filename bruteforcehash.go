package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func BruteForceHash(hash string) string {
	// Define characters for brute force
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLen := len(characters)

	if len(hash) != 32 {
		return ""
	}
	for _, c := hash {
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') {
			return ""
		}
	}

	// Iterate over all possible combinations up to length 6
	for i := 0; i <= 6; i++ {
		// Generate all strings of length i using characters
		result := make([]byte, i)
		if generateString(hash, result, 0, characters, charsetLen) {
			return string(result)
		}
	}

	return "" // If no match found
}

func generateString(targetHash string, result []byte, pos int, characters string, charsetLen int) bool {
	if pos == len(result) {
		// Convert result to string and compute MD5 hash
		md5Hash := md5.Sum([]byte(result))
		computedHash := hex.EncodeToString(md5Hash[:])

		// Compare computed hash with target hash
		if computedHash == targetHash {
			return true
		}
		return false
	}

	// Try all possible characters for current position
	for i := 0; i < charsetLen; i++ {
		result[pos] = characters[i]
		if generateString(targetHash, result, pos+1, characters, charsetLen) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(BruteForceHash("ab6ccd17455d5347c49606d641e0b2af")) // SALEM
	fmt.Println(BruteForceHash("3cbfa33db66b830bfcf47ecc956505f8")) // ALEM
	fmt.Println(BruteForceHash(""))                                 //
	fmt.Println(BruteForceHash("abc"))                              //
}
