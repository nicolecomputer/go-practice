package main

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

// - Types
type Block interface{}

type MarkdownBlock struct {
	Content string
}

type CodeBlock struct {
	Content string
}

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

// - Utils

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
