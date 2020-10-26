package contents

import (
	"strings"

	"github.com/yuin/goldmark/ast"
)

const (
	walkerSectionUnknown = iota
	walkerSectionTitle
	walkerSectionExample
	walkerSectionArguments
	walkerSectionAttributes
	walkerSectionTimeouts
	walkerSectionImport
)

// Sections represents all expected sections of a resource documentation page
type Sections struct {
	Attributes *AttributesSection
	Arguments  *ArgumentsSection
	Example    *ExampleSection
	Import     *ImportSection
	Timeouts   *TimeoutsSection
	Title      *TitleSection
}

// AttributesSection represents a resource attributes section.
type AttributesSection SchemaAttributeSection

// ArgumentsSection represents a resource arguments section.
type ArgumentsSection SchemaAttributeSection

// ExampleSection represents a resource example code section.
type ExampleSection struct {
	// Children contains further nested sections below this section
	Children []*ExampleSection

	FencedCodeBlocks []*ast.FencedCodeBlock
	Heading          *ast.Heading
	Paragraphs       []*ast.Paragraph
}

// ImportSection represents a resource import section.
type ImportSection struct {
	FencedCodeBlocks []*ast.FencedCodeBlock
	Heading          *ast.Heading
	Paragraphs       []*ast.Paragraph
}

// SchemaAttributeSection represents a schema attribute section
//
// This may represent root or nested lists of arguments or attributes
type SchemaAttributeSection struct {
	// Children contains further nested sections below this section
	Children []*SchemaAttributeSection

	// FencedCodeBlocks contains any found code blocks
	FencedCodeBlocks []*ast.FencedCodeBlock

	// Heading is the root/nested heading for the section
	Heading *ast.Heading

	// Lists is the groupings of per-attribute documentation
	//
	// Some sections may be split these based on Optional versus Required
	Lists []*ast.List

	// SchemaAttributeLists is the groupings of per-attribute documentation
	//
	// Some sections may be split these based on Optional versus Required
	SchemaAttributeLists []*SchemaAttributeList

	// Paragraphs is typically the byline(s) of per-attribute documentation
	//
	// Some sections may be split these based on Optional versus Required
	Paragraphs []*ast.Paragraph
}

// TimeoutsSection represents a resource timeouts section.
type TimeoutsSection struct {
	FencedCodeBlocks []*ast.FencedCodeBlock
	Heading          *ast.Heading
	Lists            []*ast.List
	Paragraphs       []*ast.Paragraph
}

// TitleSection represents the top documentation section
type TitleSection struct {
	FencedCodeBlocks []*ast.FencedCodeBlock
	Heading          *ast.Heading
	Paragraphs       []*ast.Paragraph
}

func sectionsWalker(document ast.Node, source []byte, resourceName string) (*Sections, error) {
	result := &Sections{}

	var walkerSectionStartingLevel, walkerSection int

	err := ast.Walk(document, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch node := node.(type) {
		case *ast.FencedCodeBlock:
			switch walkerSection {
			case walkerSectionTitle:
				result.Title.FencedCodeBlocks = append(result.Title.FencedCodeBlocks, node)
			case walkerSectionExample:
				result.Example.FencedCodeBlocks = append(result.Example.FencedCodeBlocks, node)
			case walkerSectionArguments:
				result.Arguments.FencedCodeBlocks = append(result.Arguments.FencedCodeBlocks, node)
			case walkerSectionAttributes:
				result.Attributes.FencedCodeBlocks = append(result.Attributes.FencedCodeBlocks, node)
			case walkerSectionTimeouts:
				result.Timeouts.FencedCodeBlocks = append(result.Timeouts.FencedCodeBlocks, node)
			case walkerSectionImport:
				result.Import.FencedCodeBlocks = append(result.Import.FencedCodeBlocks, node)
			}

			return ast.WalkSkipChildren, nil
		case *ast.Heading:
			headingText := string(node.Text(source))
			//fmt.Printf("(walker section level: %d) found heading level %d: %s\n", walkerSectionStartingLevel, node.Level, headingText)

			// Always reset section handling when reaching starting level again
			if node.Level == walkerSectionStartingLevel {
				walkerSection = walkerSectionUnknown
			}

			if result.Title == nil && strings.Contains(headingText, resourceName) {
				result.Title = &TitleSection{
					Heading: node,
				}

				walkerSection = walkerSectionTitle
				walkerSectionStartingLevel = node.Level

				return ast.WalkContinue, nil
			}

			if result.Example == nil && strings.HasPrefix(headingText, "Example") {
				result.Example = &ExampleSection{
					Heading: node,
				}

				walkerSection = walkerSectionExample
				walkerSectionStartingLevel = node.Level

				return ast.WalkContinue, nil
			}

			if result.Arguments == nil && strings.HasPrefix(headingText, "Argument") {
				result.Arguments = &ArgumentsSection{
					Heading: node,
				}

				walkerSection = walkerSectionArguments
				walkerSectionStartingLevel = node.Level

				return ast.WalkContinue, nil
			}

			if result.Attributes == nil && strings.HasPrefix(headingText, "Attribute") {
				result.Attributes = &AttributesSection{
					Heading: node,
				}

				walkerSection = walkerSectionAttributes
				walkerSectionStartingLevel = node.Level

				return ast.WalkContinue, nil
			}

			if result.Timeouts == nil && strings.HasPrefix(headingText, "Timeout") {
				result.Timeouts = &TimeoutsSection{
					Heading: node,
				}

				walkerSection = walkerSectionTimeouts
				walkerSectionStartingLevel = node.Level

				return ast.WalkContinue, nil
			}

			if result.Import == nil && strings.HasPrefix(headingText, "Import") {
				result.Import = &ImportSection{
					Heading: node,
				}

				walkerSection = walkerSectionImport
				walkerSectionStartingLevel = node.Level

				return ast.WalkContinue, nil
			}

			//fmt.Printf("(walker section level: %d) unknown heading level %d: %s\n", walkerSectionStartingLevel, node.Level, headingText)
			walkerSection = walkerSectionUnknown

			return ast.WalkSkipChildren, nil
		case *ast.List:
			switch walkerSection {
			case walkerSectionArguments:
				result.Arguments.Lists = append(result.Arguments.Lists, node)

				schemaAttributeList, err := schemaAttributeListWalker(node, source)

				if err != nil {
					return ast.WalkStop, err
				}

				result.Arguments.SchemaAttributeLists = append(result.Arguments.SchemaAttributeLists, schemaAttributeList)
			case walkerSectionAttributes:
				result.Attributes.Lists = append(result.Attributes.Lists, node)

				schemaAttributeList, err := schemaAttributeListWalker(node, source)

				if err != nil {
					return ast.WalkStop, err
				}

				result.Attributes.SchemaAttributeLists = append(result.Attributes.SchemaAttributeLists, schemaAttributeList)
			case walkerSectionTimeouts:
				result.Timeouts.Lists = append(result.Timeouts.Lists, node)
			}

			return ast.WalkSkipChildren, nil
		case *ast.Paragraph:
			switch walkerSection {
			case walkerSectionTitle:
				result.Title.Paragraphs = append(result.Title.Paragraphs, node)
			case walkerSectionExample:
				result.Example.Paragraphs = append(result.Example.Paragraphs, node)
			case walkerSectionArguments:
				result.Arguments.Paragraphs = append(result.Arguments.Paragraphs, node)
			case walkerSectionAttributes:
				result.Attributes.Paragraphs = append(result.Attributes.Paragraphs, node)
			case walkerSectionTimeouts:
				result.Timeouts.Paragraphs = append(result.Timeouts.Paragraphs, node)
			case walkerSectionImport:
				result.Import.Paragraphs = append(result.Import.Paragraphs, node)
			}

			return ast.WalkSkipChildren, nil
		}

		return ast.WalkContinue, nil
	})

	return result, err
}
