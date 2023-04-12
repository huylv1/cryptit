package main

import (
	"fmt"

	"github.com/huylv1/cryptit/decrypt"
	"github.com/huylv1/cryptit/encrypt"
)

func main() {
	encryptStr := encrypt.Nimbus("abc")
	fmt.Println(encryptStr)

	fmt.Println(decrypt.Nimbus(encryptStr))
}
