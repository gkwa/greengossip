package core

import _ "embed"

//go:embed instructions.tmpl
var instructionsTmpl string

func getTemplateString() string {
	return instructionsTmpl
}
