package main

import (
	"bufio"
	"fmt"
	"go/token"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
)

func identifier(id string) string {
	words := strings.FieldsFunc(id, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	for i, v := range words {
		words[i] = strings.Title(strings.ToLower(v))
	}

	return strings.Join(words, "")
}

func main() {
	// TODO: help flag

	scanner := bufio.NewScanner(os.Stdin)

	var header string
	if scanner.Scan() {
		header = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading header from stdin: %s")
	}

	// TODO: sep flag, unquote fields
	fields := strings.Split(header, ",")

	/*
		fmt.Println(strings.Join(fields, "\n"))
		os.Exit(0)
	*/

	// TODO: struct name flag
	t := jen.Type().Id("Raw").StructFunc(func(g *jen.Group) {
		for _, v := range fields {
			id := identifier(v)
			if !token.IsIdentifier(id) {
				log.Fatalf("invalid identifier: %q", id)
			}

			g.Id(id).String().Tag(map[string]string{"csv": v})
		}
	})

	if err := t.Render(os.Stdout); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}
