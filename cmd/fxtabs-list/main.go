package main

import (
	"fmt"
	"os"

	"github.com/zanardo/fxtabs-go"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %v <path to recovery.jsonlz4>\n", os.Args[0])
		os.Exit(1)
	}

	tabs, err := fxtabs.OpenTabs(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, tab := range tabs {
		fmt.Printf("title: %v\nurl: %v\n\n", tab.Title, tab.URL)
	}
}
