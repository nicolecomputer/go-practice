/*:
# Explain
:*/

package main

/*:
Explain is a tool for taking source-code formatted with special comment
identifiers and turning it into pretty documentation.

For a well documented example see [Cracklepop!](https://new-rc-2024.neocities.org/crackle)

This version is *sketch*. It does not contain tests or rigid constraints. I
often work like this in code where I will sketch on a problem and wait to see
how useful the result is before committing to a full scale implemenation.
:*/

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

/*:
## Code types

These are the types that are processed by the parse function
:*/
// - Types
type Block interface{}

type MarkdownBlock struct {
	Content string
}

type CodeBlock struct {
	Content string
}

/*:
## Core parsing
This parses a string to look for code blocks and markdown blocks.

In a future implementation it would be neat if this could also process aside
blocks that would be show as margin notes with code. If that happened I think I
would use the identifier `/*|``.
:*/
// - Parse
func parse(str string) []Block {
	blocks := []Block{}
	currentContent := ""

	for _, line := range strings.Split(str, "\n") {
		if strings.HasPrefix(line, "/*:") {
			if len(currentContent) > 0 {
				blocks = append(blocks, CodeBlock{currentContent})
			}
			currentContent = ""
		} else if strings.HasPrefix(line, ":*/") {
			if len(currentContent) > 0 {
				blocks = append(blocks, MarkdownBlock{currentContent})
			}
			currentContent = ""
		} else {
			currentContent = currentContent + line + "\n"
		}
	}

	if len(currentContent) > 0 {
		blocks = append(blocks, CodeBlock{currentContent})
	}

	return blocks
}

/*:
## Util functions
:*/

func SanitizedInput(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	str := string(bytes)
	str = strings.TrimSpace(str)

	return str
}

/*:
## Output

These functions take the syntax tree build up by parse and convert it to HTML
:*/

type OutputData struct {
	Filename string
	Blocks   []OutputBlock
}

type OutputBlock struct {
	Type    string // code or words
	Content string
}

func OutputCodeExplainer(inputFile string) {
	input := SanitizedInput(inputFile)

	blocks := []OutputBlock{}

	for _, block := range parse(input) {
		switch b := block.(type) {
		case MarkdownBlock:
			var buf bytes.Buffer
			if err := goldmark.Convert([]byte(b.Content), &buf); err != nil {
				panic(err)
			}

			blocks = append(blocks, OutputBlock{Type: "words", Content: buf.String()})
		case CodeBlock:
			// By trimming space we no-longer have accurate line numbers
			// I did some work to preserve line numbers but there is often space
			// between markdown and code blocks that made it awkward
			cleaned := strings.TrimSpace(b.Content)
			blocks = append(blocks, OutputBlock{Type: "code", Content: cleaned})
		}
	}

	filename := filepath.Base(inputFile)
	data := OutputData{
		Filename: filename,
		Blocks:   blocks,
	}

	templateFile := "output.tmpl"
	template, err := template.New(templateFile).ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	err = template.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

// This might not make sense as part of this program- just markdown might want
// to be its own thing
func OutputJustMarkdown(inputFile string) {
	input := SanitizedInput(inputFile)

	filename := filepath.Base(inputFile)

	var buf bytes.Buffer
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	if err := markdown.Convert([]byte(input), &buf); err != nil {
		panic(err)
	}

	data := OutputData{
		Filename: filename,
		Blocks: []OutputBlock{
			{Type: "words", Content: buf.String()},
		},
	}

	templateFile := "output.tmpl"
	template, err := template.New(templateFile).ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}
	err = template.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func main() {
	inputFile := os.Args[1]
	fileType := filepath.Ext(inputFile)
	if fileType == ".md" {
		OutputJustMarkdown(inputFile)
	} else {
		OutputCodeExplainer(inputFile)
	}
}
