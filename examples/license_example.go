package main

import (
	"fmt"

	"github.com/alfg/widevine"
)

var (
	key = []byte{
		0x1a, 0xe8, 0xcc, 0xd0, 0xe7, 0x98, 0x5c, 0xc0,
		0xb6, 0x20, 0x3a, 0x55, 0x85, 0x5a, 0x10, 0x34,
		0xaf, 0xc2, 0x52, 0x98, 0x0e, 0x97, 0x0c, 0xa9,
		0x0e, 0x52, 0x02, 0x68, 0x9f, 0x94, 0x7a, 0xb9}

	iv = []byte{
		0xd5, 0x8c, 0xe9, 0x54, 0x20, 0x3b, 0x7c, 0x9a,
		0x9a, 0x9d, 0x46, 0x7f, 0x59, 0x83, 0x92, 0x49}
)

func main() {

	wv := widevine.New(key, iv)

	contentID := "test"
	resp := wv.GetContentKey(contentID)
	fmt.Println(resp.Status)
}