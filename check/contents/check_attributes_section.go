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
		return fmt.Errorf("missing attribute section: ## Attribute Reference")
	}

	heading := section.Heading

	if heading.Level != 2 {
		return fmt.Errorf("attribute section heading level (%d) should be: 2", heading.Level)
	}

	headingText := string(heading.Text(d.source))
	expectedHeadingTexts := []string{
		"Attribute Reference",
		"Attributes Reference",
	}

	if headingText != expectedHeadingTexts[0] && headingText != expectedHeadingTexts[1] {
		return fmt.Errorf("attribute section heading (%s) should be: %q (preferred) or %q", headingText, expectedHeadingTexts[0], expectedHeadingTexts[1])
	}

	paragraphs := section.Paragraphs
	expectedBylineTexts := []string{
		"This resource exports the following attributes in addition to the arguments above:",
		"In addition to all arguments above, the following attributes are exported:",
		"No additional attributes are exported.",
	}

	switch len(paragraphs) {
	case 0:
		return fmt.Errorf("attribute section byline should be: %q (preferred), %q or %q", expectedBylineTexts[0], expectedBylineTexts[1], expectedBylineTexts[2])
	case 1:
		paragraphText := string(paragraphs[0].Text(d.source))

		if paragraphText != expectedBylineTexts[0] && paragraphText != expectedBylineTexts[1] && paragraphText != expectedBylineTexts[2] {
			return fmt.Errorf("attribute section byline (%s) should be: %q (preferred), %q or %q", paragraphText, expectedBylineTexts[0], expectedBylineTexts[1], expectedBylineTexts[2])
		}
	}

	if checkOpts.RequireSchemaOrdering {
		for _, list := range section.SchemaAttributeLists {
			if !sort.IsSorted(SchemaAttributeListItemByName(list.Items)) {
				return fmt.Errorf("attribute section is not sorted by name")
			}
		}
	}

	return nil
}
