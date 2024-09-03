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
	return strings.ReplaceAll(`<html lang="en">
<head>
	<style>
		/*
			Josh's Custom CSS Reset
			https://www.joshwcomeau.com/css/custom-css-reset/
			Thanks Josh!
		*/
		*, *::before, *::after {
			box-sizing: border-box;
		}

		* {
			margin: 0;
		}

		body {
			line-height: 1.5;
			-webkit-font-smoothing: antialiased;
		}

		img, picture, video, canvas, svg {
			display: block;
			max-width: 100%;
		}

		input, button, textarea, select {
			font: inherit;
		}

		p, h1, h2, h3, h4, h5, h6 {
			overflow-wrap: break-word;
		}

		/*
			100 Bytes of CSS to look great everywhere, by swyx.
			https://www.swyx.io/css-100-bytes
			Thanks swyx!
		*/
		html {
			max-width: 70ch;
			padding: 3em 1em;
			margin: auto;
			line-height: 1.75;
			font-size: 1.25em;
		}

		/*
			Optional 100 more bytes, by swyx, slightly tweaked.
			https://www.swyx.io/css-100-bytes#optional-100-more-bytes
			Thanks again swyx!
		*/
		h1, h2, h3, h4, h5, h6 {
			margin: 3em 0 1em;
		}

		p,ul,ol {
			margin-bottom: 2em;
			color: #1d1d1d;
			font-family: sans-serif;
		}
	</style>
</head>
<body>
	<article>
		THEARTICLE
	</article>
</body>
</html>`, "THEARTICLE", string(html))
}
