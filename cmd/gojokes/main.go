// Command Line interface (CLI) Tool.
//
// Usage:
//
//	gojokes                     Print a random joke from any category.
//	gojokes -c <category>       Print a random joke from the named category.
//	gojokes --category <name>   Same as -c.
//	gojokes --list              List all available categories.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yuvrajmodz/gojokes"
)

func main() {
	var (
		category string
		list     bool
	)

	flag.StringVar(&category, "c", "", "Category to pick a joke from.")
	flag.StringVar(&category, "category", "", "Category to pick a joke from.")
	flag.BoolVar(&list, "list", false, "List all available categories.")
	flag.Usage = usage
	flag.Parse()

	switch {
	case list:
		printCategories()
	case category != "":
		printCategoryJoke(category)
	default:
		printRandomJoke()
	}
}

func printRandomJoke() {
	joke, err := gojokes.Random()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gojokes: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(joke)
}

func printCategoryJoke(name string) {
	joke, err := gojokes.Category(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gojokes: %v\n", err)
		fmt.Fprintf(os.Stderr, "Available categories: %s\n",
			strings.Join(gojokes.Categories(), ", "))
		os.Exit(1)
	}
	fmt.Println(joke)
}

func printCategories() {
	for _, cat := range gojokes.Categories() {
		fmt.Println(cat)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `gojokes - Random Jokes from the Command Line

Usage:
  gojokes                     Print a random joke
  gojokes -c <category>       Print a joke from a specific category
  gojokes --category <name>   Same as -c
  gojokes --list              List all available categories

Examples:
  gojokes
  gojokes -c programming
  gojokes --category golang
  gojokes --list
`)
}