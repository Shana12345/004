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
	hash := h.Sum(nil)

	fmt.Println()
	fmt.Println()

	fmt.Println("MD5 hash to byte")
	byteString := []byte(hash)
	fmt.Println(byteString)

	fmt.Println()
	fmt.Println()

	fmt.Println("looping through byte")
	//for i := 0; i < len(byteString); i++ {
	for i := range byteString {
		fmt.Printf("%x ", byteString[i])
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("base64 Encoding...")
	str := base64.StdEncoding.EncodeToString([]byte(byteString))
	fmt.Println(str)

	fmt.Println()

	fmt.Println("base64 Decoding...")
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Printf("%q\n", data)
}
