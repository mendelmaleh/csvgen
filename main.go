package main

import (
	"bufio"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/mitchellh/hashstructure"
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

func groupname(x []int) string {
	r := make([]rune, len(x))
	for i, v := range x {
		r[i] = 'A' + rune(v)
	}

	return string(r)
}

func main() {
	// headers
	scanner := bufio.NewScanner(os.Stdin)

	var headers [][]string
	for scanner.Scan() {
		headers = append(headers, strings.Split(scanner.Text(), ","))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading header from stdin: %s")
	}

	// fields to header indexes
	fields := make(map[string][]int)

	for i, header := range headers {
		for _, field := range header {
			fields[field] = append(fields[field], i)
		}
	}

	// header index group
	var groups [][]int

	// header index group hash to fields
	parts := make(map[uint64][]string)

	for k, v := range fields {
		hash, err := hashstructure.Hash(v, nil)
		if err != nil {
			log.Fatal(err)
		}

		if _, ok := parts[hash]; !ok {
			groups = append(groups, v)
		}

		parts[hash] = append(parts[hash], k)
	}

	// sort groups
	sort.Slice(groups, func(i, j int) bool {
		li := len(groups[i])
		lj := len(groups[j])

		if li == lj {
			for x := 0; x < li; x++ {
				xi := groups[i][x]
				xj := groups[j][x]

				if xi != xj {
					return xi < xj
				}
			}
		}

		// descending, biggest group (most common part) to least
		return li > lj
	})

	// header index to header groups
	backrefs := make([][]int, len(headers))

	for i, group := range groups {
		if len(group) == 1 {
			break
		}

		for _, v := range group {
			backrefs[v] = append(backrefs[v], i)
		}
	}

	// output
	f := jen.NewFile("main")

	for _, v := range groups {
		hash, err := hashstructure.Hash(v, nil)
		if err != nil {
			log.Fatal(err)
		}

		sort.Strings(parts[hash])
		// spew.Dump(parts[hash])

		f.Type().Id(groupname(v)).StructFunc(func(g *jen.Group) {
			// common parts
			if len(v) == 1 {
				for _, v := range backrefs[v[0]] {
					g.Id(groupname(groups[v]))
				}
				g.Line()
			}

			// fields
			for _, v := range parts[hash] {
				id := identifier(v)
				if !token.IsIdentifier(id) {
					log.Fatalf("invalid identifier: %q", id)
				}

				g.Id(id).String().Tag(map[string]string{"csv": v})
			}
		}).Line()

	}

	if err := f.Render(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
