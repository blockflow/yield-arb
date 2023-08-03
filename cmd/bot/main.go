package main

import (
	"log"
	"yield-arb/pkg/protocols"
)

func main() {
	log.Println("Starting bot...")
	p, err := protocols.GetProtocol("aavev2")
	if err != nil {
		panic(err)
	}
	p.Connect("ethereum")
}
