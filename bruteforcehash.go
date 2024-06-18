package main

// package bootcamp

import (
	"crypto/md5"
	"fmt"
	"io"
)

var (
	all string
	res string
)

func BruteForceHash(hash string) string {
	if len(hash) != 32 {
		return ""
	}

	h := md5.New()
	if hash == string(h.Sum(nil)) {
		return ""
	}

	res = ""
	all = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	Finden(hash, "", 0)

	return res
}

func Finden(hash string, pref string, length int) bool {
	h := md5.New()
	io.WriteString(h, pref)
	if string(h.Sum(nil)) == hash {
		res = pref
		return true
	}

	if length < 6 {
		for i := 0; i < 36; i++ {
			if Finden(hash, pref+string(all[i]), length+1) {
				return true
			}
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
