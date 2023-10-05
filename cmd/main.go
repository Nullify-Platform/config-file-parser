package main

import (
	"bytes"

	"github.com/nullify-platform/config-file-parser/pkg/parser"
	"gopkg.in/yaml.v3"
)

func main() {
	config, err := parser.LoadFromFile("data/nullify.yaml")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)

	err = enc.Encode(config)
	if err != nil {
		panic(err)
	}

	println(buf.String())
}
