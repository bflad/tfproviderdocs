package contents

import (
	"fmt"
	"sort"
)

type CheckAttributesSectionOptions struct {
	RequireSchemaOrdering bool
}

func (d *Document) checkAttributesSection() error {
	checkOpts := &CheckAttributesSectionOptions{}

	if d.CheckOptions != nil && d.CheckOptions.AttributesSection != nil {
		checkOpts = d.CheckOptions.AttributesSection
	}

	section := d.Sections.Attributes

	if section == nil {
		return fmt.Errorf("missing attributes section: ## Attributes Reference")
	}

	heading := section.Heading

	if heading.Level != 2 {
		return fmt.Errorf("attributes section heading level (%d) should be: 2", heading.Level)
	}

	headingText := string(heading.Text(d.source))
	expectedHeadingText := "Attributes Reference"

	if headingText != expectedHeadingText {
		return fmt.Errorf("attributes section heading (%s) should be: %s", headingText, expectedHeadingText)
	}

	paragraphs := section.Paragraphs
	expectedBylineText := "In addition to all arguments above, the following attributes are exported:"

	switch len(paragraphs) {
	case 0:
		return fmt.Errorf("attributes section byline should be: %s", expectedBylineText)
	case 1:
		paragraphText := string(paragraphs[0].Text(d.source))

		if paragraphText != expectedBylineText {
			return fmt.Errorf("attributes section byline (%s) should be: %s", paragraphText, expectedBylineText)
		}
	}

	if checkOpts.RequireSchemaOrdering {
		for _, list := range section.SchemaAttributeLists {
			if !sort.IsSorted(SchemaAttributeListItemByName(list.Items)) {
				return fmt.Errorf("attributes section is not sorted by name")
			}
		}
	}

	return nil
}
