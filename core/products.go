package core

import (
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func getProducts(markdownDir string, verbose bool) []string {
	var products []string
	err := filepath.Walk(markdownDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			slog.Error("error accessing path", "path", path, "error", err)
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			product := extractProduct(path)
			if product != "" {
				products = append(products, product)
			}
		}
		return nil
	})
	if err != nil {
		slog.Error("error walking path", "path", markdownDir, "error", err)
		os.Exit(1)
	}

	sort.SliceStable(products, func(i, j int) bool {
		return strings.ToLower(products[i]) < strings.ToLower(products[j])
	})

	if verbose {
		slog.Info("found products", "products", products)
	}

	return products
}

func extractProduct(path string) string {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("error opening file", "path", path, "error", err)
		return ""
	}
	defer file.Close()

	content, err := os.ReadFile(path)
	if err != nil {
		slog.Error("error reading file", "path", path, "error", err)
		return ""
	}

	markdown := goldmark.New(goldmark.WithExtensions(meta.Meta))
	context := parser.NewContext()
	markdown.Parser().Parse(text.NewReader(content), parser.WithContext(context))

	metaData := meta.Get(context)
	if fileType, ok := metaData["filetype"]; ok && fileType == "product" {
		return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	}

	return ""
}
