package main

import (
	"log"

	"github.com/genkiroid/vsym"
)

func main() {
	v, err := vsym.NewVSym()
	if err != nil {
		log.SetPrefix("vsym: ")
		log.SetFlags(0)
		log.Fatal(err)
	}

	v.Verify()
}
