package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	filePath = flag.String("file", "", "json file")
)

func init() {
	flag.StringVar(filePath, "f", "", "json file")
}

func main() {
	var err error
	var input []byte
	flag.Parse()
	if len(*filePath) > 0 {
		input, err = ioutil.ReadFile(*filePath)
	} else {
		input, err = ioutil.ReadAll(os.Stdin)
	}
	// input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("read data err:", err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, input, "", "    ")
	if err != nil {
		log.Fatal("json format err:", err)
	}
	io.Copy(os.Stdout, bytes.NewReader(out.Bytes()))
	// if len(filePath) > 0 {
	// 	err = ioutil.WriteFile(filePath, out.Bytes(), 0664)
	// }

	// if err != nil {
	// 	log.Fatal("WriteFile err:", err)
	// }
}
