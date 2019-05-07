package main

import (
	"fmt"
	"log"

	"github.com/shenwei356/bwt"
)

func main() {
	fmt.Println("vim-go")
	s := "abracadabra"
	tr, _, err := bwt.Transform([]byte(s), '$')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(tr))
}
