package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	ioutil.ReadFile(filePath)
	input, err := ioutil.ReadFile(filePath)
	// input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("read data err:", err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, input, "", "    ")
	if err != nil {
		log.Fatal("json format err:", err)
	}
	err = ioutil.WriteFile(filePath, out.Bytes(), 0664)

	if err != nil {
		log.Fatal("WriteFile err:", err)
	}
}
