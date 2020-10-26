package contents

import (
	"fmt"
	"sort"
)

type CheckArgumentsSectionOptions struct {
	RequireSchemaOrdering bool
}

func (d *Document) checkArgumentsSection() error {
	checkOpts := &CheckArgumentsSectionOptions{}

	if d.CheckOptions != nil && d.CheckOptions.ArgumentsSection != nil {
		checkOpts = d.CheckOptions.ArgumentsSection
	}

	section := d.Sections.Arguments

	if section == nil {
		return fmt.Errorf("missing arguments section: ## Argument Reference")
	}

	heading := section.Heading

	if heading.Level != 2 {
		return fmt.Errorf("arguments section heading level (%d) should be: 2", heading.Level)
	}

	headingText := string(heading.Text(d.source))
	expectedHeadingText := "Argument Reference"

	if headingText != expectedHeadingText {
		return fmt.Errorf("arguments section heading (%s) should be: %s", headingText, expectedHeadingText)
	}

	if checkOpts.RequireSchemaOrdering {
		for _, list := range section.SchemaAttributeLists {
			if !sort.IsSorted(SchemaAttributeListItemByName(list.Items)) {
				return fmt.Errorf("arguments section is not sorted by name")
			}
		}
	}

	return nil
}
