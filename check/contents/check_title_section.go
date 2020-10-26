package contents

import (
	"fmt"
	"strings"
)

func (d *Document) checkTitleSection() error {
	section := d.Sections.Title

	if section == nil {
		return fmt.Errorf("missing title section: # Resource: %s", d.ResourceName)
	}

	heading := section.Heading

	if heading.Level != 1 {
		return fmt.Errorf("title section heading level (%d) should be: 1", heading.Level)
	}

	headingText := string(heading.Text(d.source))

	if !strings.HasPrefix(headingText, "Data Source: ") && !strings.HasPrefix(headingText, "Resource: ") {
		return fmt.Errorf("title section heading (%s) should have prefix: \"Data Source: \" or \"Resource: \"", headingText)
	}

	if len(section.FencedCodeBlocks) > 0 {
		return fmt.Errorf("title section code examples should be in Example Usage section")
	}

	return nil
}
