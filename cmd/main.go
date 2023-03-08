package main

import (
	"fmt"

	"github.com/nullify-platform/config-file-parser/pkg/parser"
)

func main() {
	config, err := parser.LoadFromFile("data/nullify.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("ignore dirs:\n")
	for _, dir := range config.IgnoreDirs {
		fmt.Printf("  - %s\n", dir)
	}
}
