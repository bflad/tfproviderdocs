package markdown

import (
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Parse converts a Markdown source into AST and metadata
func Parse(source []byte) (ast.Node, map[string]interface{}) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.New(),
		),
	)

	context := parser.NewContext()
	reader := text.NewReader(source)
	document := markdown.Parser().Parse(reader, parser.WithContext(context))
	metadata := meta.Get(context)

	return document, metadata
}
