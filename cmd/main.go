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

	fmt.Println("ignore dirs:")
	for _, dir := range config.IgnoreDirs {
		fmt.Printf("  - %s\n", dir)
	}

	fmt.Println("notifications:")
	for name, notification := range config.Notifications {
		fmt.Printf("- %s\n", name)
		fmt.Println("\tevents:")
		for _, event := range notification.Events {
			fmt.Printf("\t- %s\n", event.Type)
		}

		fmt.Println("\ttargets:")
		for _, event := range notification.Targets {
			fmt.Printf("\t- %s\n", event.Type)
		}
	}
}
