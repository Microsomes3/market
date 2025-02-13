package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("load genesis block")

	f, err := os.Open("genesis_block2")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	bytesi := bytes.Buffer{}

	io.Copy(&bytesi, f)

	fmt.Println(bytesi.Bytes())

	os.Exit(0)

	// json.Unmarshal(bytesi.Bytes(), &genesisBlock)

	// isValidBlock := genesisBlock.VerifyBlock()

	// fmt.Println("is valid block:", isValidBlock)

}
