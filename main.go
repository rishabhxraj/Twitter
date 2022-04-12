package main

import (
	"fmt"
	"os"

	"github.com/fogleman/gg"
	"github.com/rishabhxraj/twitter-bot/background"
	"github.com/rishabhxraj/twitter-bot/drawer"
	"github.com/rishabhxraj/twitter-bot/quotable"
)

func main() {
	quote, err := quotable.GetQuotes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if err := background.GetBackground(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	footer := "https://twitter.com/rishabh053"
	img := &drawer.Image{Canvas: gg.NewContext(1200, 628), Quote: quote, Footer: footer}
	if err := img.Create(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
