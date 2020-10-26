package contents

import (
	"strings"

	"github.com/yuin/goldmark/ast"
)

// SchemaAttributeList represents a schema attribute list
//
// This may represent root or nested lists of arguments or attributes
type SchemaAttributeList struct {
	Items []*SchemaAttributeListItem
}

// SchemaAttributeListItem represents a schema attribute list item
//
// This may represent root or nested lists of arguments or attributes
type SchemaAttributeListItem struct {
	Description string
	ForceNew    bool
	Name        string
	Optional    bool
	Required    bool
	Type        string
}

type SchemaAttributeListItemByName []*SchemaAttributeListItem

func (item SchemaAttributeListItemByName) Len() int           { return len(item) }
func (item SchemaAttributeListItemByName) Swap(i, j int)      { item[i], item[j] = item[j], item[i] }
func (item SchemaAttributeListItemByName) Less(i, j int) bool { return item[i].Name < item[j].Name }

func schemaAttributeListWalker(list *ast.List, source []byte) (*SchemaAttributeList, error) {
	result := &SchemaAttributeList{}

	err := ast.Walk(list, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch node := node.(type) {
		case *ast.ListItem:
			item, err := schemaAttributeListItemWalker(node, source)

			if err != nil {
				return ast.WalkStop, err
			}

			result.Items = append(result.Items, item)

			return ast.WalkContinue, nil
		}

		return ast.WalkContinue, nil
	})

	return result, err
}

func schemaAttributeListItemWalker(listItem *ast.ListItem, source []byte) (*SchemaAttributeListItem, error) {
	result := &SchemaAttributeListItem{}

	// Expected format: `Name` - (Required/Optional[, ForceNew]) Description

	err := ast.Walk(listItem, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch node := node.(type) {
		case *ast.TextBlock:
			text := string(node.Text(source))
			itemParts := strings.SplitN(text, " - ", 2)

			if len(itemParts) != 2 {
				return ast.WalkContinue, nil
			}

			result.Name = itemParts[0]
			fullDescription := itemParts[1]

			if !strings.HasPrefix(fullDescription, "(") {
				result.Description = fullDescription

				return ast.WalkStop, nil
			}

			traitsEndIndex := strings.IndexByte(fullDescription, ')')

			result.Description = fullDescription[traitsEndIndex+1:]

			traits := fullDescription[1:traitsEndIndex]

			for _, trait := range strings.Split(traits, ", ") {
				switch trait {
				case "Boolean", "Number", "String":
					result.Type = trait
				case "Forces new", "Forces new resource":
					result.ForceNew = true
				case "Optional":
					result.Optional = true
				case "Required":
					result.Required = true
				}
			}

			return ast.WalkStop, nil
		}

		return ast.WalkContinue, nil
	})

	return result, err
}
