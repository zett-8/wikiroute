package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	genType := flag.String("type", "", "generate data type")
	flag.Parse()

	switch *genType {
	case "pages":
		err := GeneratePagesData()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		os.Exit(0)
	case "pagelinks":
		err := GeneratePageLinksData()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		os.Exit(0)
	case "linktargets":
		err := GenerateLinkTargetsData()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		os.Exit(0)
	default:
		fmt.Printf("Error: invalid type %s\nSpecify generate data type with --type flag\nValue must be 'linktargets' | 'pagelinks' | 'pages'\n", *genType)
		os.Exit(1)
	}
}
