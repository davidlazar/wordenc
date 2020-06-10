package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/davidlazar/wordenc"
)

var decodeMode = flag.Bool("d", false, "decode stdin")

func main() {
	flag.Parse()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("error reading stdin: %s", err)
	}

	if *decodeMode {
		// TODO Need better plan for DecodeString length
		panic("NOT IMPLEMENTED")
	}

	s := wordenc.EncodeToString(data)
	fmt.Println(s)
}
