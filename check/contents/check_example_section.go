package contents

import (
	"fmt"
	"strings"

	"github.com/bflad/tfproviderdocs/markdown"
)

type CheckExamplesSectionOptions struct {
	ExpectedCodeBlockLanguage string
}

func (d *Document) checkExampleSection() error {
	checkOpts := &CheckExamplesSectionOptions{
		ExpectedCodeBlockLanguage: markdown.FencedCodeBlockLanguageTerraform,
	}

	if d.CheckOptions != nil && d.CheckOptions.ExamplesSection != nil {
		checkOpts = d.CheckOptions.ExamplesSection
	}

	section := d.Sections.Example

	if section == nil {
		return fmt.Errorf("missing example section: ## Example Usage")
	}

	heading := section.Heading

	if heading.Level != 2 {
		return fmt.Errorf("example section heading level (%d) should be: 2", heading.Level)
	}

	headingText := string(heading.Text(d.source))
	expectedHeadingText := "Example Usage"

	if headingText != expectedHeadingText {
		return fmt.Errorf("example section heading (%s) should be: %s", headingText, expectedHeadingText)
	}

	// CDKTF conversion will leave the original terraform code blocks if unsuccessful
	if checkOpts.ExpectedCodeBlockLanguage != markdown.FencedCodeBlockLanguageTerraform {
		return nil
	}

	for _, fencedCodeBlock := range section.FencedCodeBlocks {
		language := markdown.FencedCodeBlockLanguage(fencedCodeBlock, d.source)

		if language != checkOpts.ExpectedCodeBlockLanguage {
			return fmt.Errorf("example section code block language (%s) should be: ```%s", language, checkOpts.ExpectedCodeBlockLanguage)
		}

		text := markdown.FencedCodeBlockText(fencedCodeBlock, d.source)

		if !strings.Contains(text, d.ResourceName) {
			return fmt.Errorf("example section code block text should contain resource name: %s", d.ResourceName)
		}
	}

	return nil
}
