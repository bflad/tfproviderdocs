package markdown

import (
	"strings"

	"github.com/yuin/goldmark/ast"
)

const (
	FencedCodeBlockLanguageHcl       = "hcl"
	FencedCodeBlockLanguageMissing   = "MISSING"
	FencedCodeBlockLanguageTerraform = "terraform"
)

// FencedCodeBlockLanguage returns the language or "MISSING"
func FencedCodeBlockLanguage(fcb *ast.FencedCodeBlock, source []byte) string {
	if fcb == nil {
		return FencedCodeBlockLanguageMissing
	}

	language := string(fcb.Language(source))

	if language == "" {
		return FencedCodeBlockLanguageMissing
	}

	return language
}

// FencedCodeBlockText returns the text
func FencedCodeBlockText(fcb *ast.FencedCodeBlock, source []byte) string {
	if fcb == nil {
		return ""
	}

	lines := fcb.Lines()
	var builder strings.Builder

	for i := 0; i < lines.Len(); i++ {
		segment := lines.At(i)
		builder.WriteString(string(segment.Value(source)))
	}

	return strings.TrimSpace(builder.String())
}
