// package main

package bootcamp

import (
	"crypto/md5"
	"fmt"
	"io"
)

var (
	rrange []string
	all    string
	res    string
)

func BruteForceHash(hash string) string {
	if len(hash) != 32 {
		return ""
	}

	h := md5.New()
	if hash == string(h.Sum(nil)) {
		return ""
	}

	c := make(chan string, 1)
	res = ""
	rrange = rrange[:0]
	rrange = append(rrange, "ABC")
	rrange = append(rrange, "DEF")
	rrange = append(rrange, "GHI")
	rrange = append(rrange, "JKL")
	rrange = append(rrange, "MNO")
	rrange = append(rrange, "PQR")
	rrange = append(rrange, "STU")
	rrange = append(rrange, "VWX")
	rrange = append(rrange, "YZ0")
	rrange = append(rrange, "123")
	rrange = append(rrange, "456")
	rrange = append(rrange, "789")
	all = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	go Find(hash, 0, &c)
	go Find(hash, 1, &c)
	go Find(hash, 2, &c)
	go Find(hash, 3, &c)
	go Find(hash, 4, &c)
	go Find(hash, 5, &c)
	go Find(hash, 6, &c)
	go Find(hash, 7, &c)
	go Find(hash, 8, &c)
	go Find(hash, 9, &c)
	go Find(hash, 10, &c)
	go Find(hash, 11, &c)

	return <-c
}

func Find(hash string, ind int, c *chan string) {
	for i := 0; i < 3; i++ {
		if Finden(hash, string(rrange[ind][i]), 1, c) {
			return
		}
	}
}

func Finden(hash string, pref string, length int, c *chan string) bool {
	h := md5.New()
	io.WriteString(h, pref)
	if string(h.Sum(nil)) == hash {
		*c <- pref
		close(*c)
		return true
	}

	if length < 6 {
		for i := 0; i < 36; i++ {
			if Finden(hash, pref+string(all[i]), length+1, c) {
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
