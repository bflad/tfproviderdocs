package contents

import (
	"fmt"
	"strings"

	"github.com/bflad/tfproviderdocs/markdown"
)

func (d *Document) checkImportSection() error {
	section := d.Sections.Import

	if section == nil {
		return nil
	}

	heading := section.Heading

	if heading.Level != 2 {
		return fmt.Errorf("import section heading level (%d) should be: 2", heading.Level)
	}

	headingText := string(heading.Text(d.source))
	expectedHeadingText := "Import"

	if headingText != expectedHeadingText {
		return fmt.Errorf("import section heading (%s) should be: %s", headingText, expectedHeadingText)
	}

	for _, fencedCodeBlock := range section.FencedCodeBlocks {
		text := markdown.FencedCodeBlockText(fencedCodeBlock, d.source)

		if !strings.Contains(text, d.ResourceName) {
			return fmt.Errorf("import section code block text should contain resource name: %s", d.ResourceName)
		}
	}

	return nil
}
