package core

import (
	"bytes"
	"log/slog"
	"os"
	"text/template"

	"github.com/atotto/clipboard"
)

func Run(obsidianVaultPath string, verbose bool) {
	products := getProducts(obsidianVaultPath, verbose)

	clipboardContent := getClipboardContent()

	templateContent := fillTemplate(products, clipboardContent)

	writeToClipboard(templateContent)

	slog.Info("Template written to clipboard")
}

func getClipboardContent() string {
	clipboardContent, err := clipboard.ReadAll()
	if err != nil {
		slog.Error("error reading clipboard", "error", err)
		os.Exit(1)
	}
	return clipboardContent
}

func fillTemplate(products []string, clipboardContent string) string {
	tmplStr := getTemplateString()

	tmpl, err := template.New("template").Parse(tmplStr)
	if err != nil {
		slog.Error("error parsing template", "error", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, struct {
		Products  []string
		Clipboard string
	}{
		Products:  products,
		Clipboard: clipboardContent,
	})
	if err != nil {
		slog.Error("error executing template", "error", err)
		os.Exit(1)
	}

	return buf.String()
}

func writeToClipboard(content string) {
	err := clipboard.WriteAll(content)
	if err != nil {
		slog.Error("error writing to clipboard", "error", err)
		os.Exit(1)
	}
}
