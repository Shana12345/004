package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)
var input string

func main() {
	fmt.Println("Please Enter your Message")
	fmt.Scan(&input)

	fmt.Println()

	fmt.Println("Your MD5 hash...")
	h := md5.New()
	io.WriteString(h, input)
	fmt.Printf("%x", h.Sum(nil))

	fmt.Println()
	fmt.Println()

	fmt.Println("converted to byte")
	fmt.Println([]byte(h.Sum(nil)))

	fmt.Println()

	fmt.Println("converted to slice byte")
	hash := []byte(h.Sum(nil))

	for i := 0; i < len(hash); i++ {
		fmt.Printf("%x ", hash[i])
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("base64 Encoding...")
	str := base64.StdEncoding.EncodeToString([]byte(hash))
	fmt.Println(str)

	fmt.Println()

	fmt.Println("base64 Decoding...")
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Printf("%q\n", data)

}
