package main

import (
	"fmt"
	"os"

	// Load necessary providers.
	_ "github.com/zquestz/s/providers/amazon"
	_ "github.com/zquestz/s/providers/bing"
	_ "github.com/zquestz/s/providers/dockerhub"
	_ "github.com/zquestz/s/providers/duckduckgo"
	_ "github.com/zquestz/s/providers/github"
	_ "github.com/zquestz/s/providers/google"
	_ "github.com/zquestz/s/providers/reddit"
	_ "github.com/zquestz/s/providers/stackoverflow"
	_ "github.com/zquestz/s/providers/twitter"
	_ "github.com/zquestz/s/providers/wikipedia"
	_ "github.com/zquestz/s/providers/yahoo"

	"github.com/zquestz/s/cmd"
)

func main() {
	setupSignalHandlers()

	if err := cmd.SearchCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}