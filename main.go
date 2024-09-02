package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func main() {
	sourceDir := "markdown" // Current directory
	destDir := "docs"

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			relPath, err := filepath.Rel(sourceDir, path)
			if err != nil {
				return err
			}

			destPath := filepath.Join(destDir, strings.TrimSuffix(relPath, ".md")+".html")

			err = convertMarkdownToHTML(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
		return
	}

	fmt.Println("Conversion completed successfully.")
}

func convertMarkdownToHTML(sourcePath, destPath string) error {
	mdContent, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("error reading markdown file: %v", err)
	}

	// Check if the markdown file has front matter.
	_, mdContentString, hasFrontmatter := strings.Cut(string(mdContent), "---")
	if hasFrontmatter {
		mdContent = []byte(mdContentString)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	html := markdown.ToHTML(mdContent, p, nil)

	html = []byte(HTMLTemplate(html))

	err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating destination directory: %v", err)
	}

	err = ioutil.WriteFile(destPath, html, 0644)
	if err != nil {
		return fmt.Errorf("error writing HTML file: %v", err)
	}

	fmt.Printf("Converted %s to %s\n", sourcePath, destPath)
	return nil
}

func HTMLTemplate(html []byte) string {
	return fmt.Sprintf(`<html lang="en">
<head>
  <style>
  html {
    max-width: 70ch;
    padding: 3em 1em;
    margin: auto;
    line-height: 1.75;
    font-size: 1.25em;
  }
  </style>
</head>
<body>
<article>
%s
</article>
</body>
</html>`, html)
}
