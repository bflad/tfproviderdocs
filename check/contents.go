package check

import (
	"fmt"

	"github.com/bflad/tfproviderdocs/check/contents"
)

type ContentsCheck struct {
	Options *ContentsOptions
}

// ContentsOptions represents configuration options for Contents.
type ContentsOptions struct {
	*FileOptions

	Enable                bool
	ProviderName          string
	RequireSchemaOrdering bool
}

func NewContentsCheck(opts *ContentsOptions) *ContentsCheck {
	check := &ContentsCheck{
		Options: opts,
	}

	if check.Options == nil {
		check.Options = &ContentsOptions{}
	}

	if check.Options.FileOptions == nil {
		check.Options.FileOptions = &FileOptions{}
	}

	return check
}

func (check *ContentsCheck) Run(path string) error {
	if !check.Options.Enable {
		return nil
	}

	checkOpts := &contents.CheckOptions{
		ArgumentsSection: &contents.CheckArgumentsSectionOptions{
			RequireSchemaOrdering: check.Options.RequireSchemaOrdering,
		},
		AttributesSection: &contents.CheckAttributesSectionOptions{
			RequireSchemaOrdering: check.Options.RequireSchemaOrdering,
		},
	}

	doc := contents.NewDocument(path, check.Options.ProviderName)

	if err := doc.Parse(); err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}

	if err := doc.Check(checkOpts); err != nil {
		return err
	}

	return nil
}
