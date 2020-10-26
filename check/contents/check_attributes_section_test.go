package contents

import (
	"testing"
)

func TestCheckAttributesSection(t *testing.T) {
	testCases := []struct {
		Name         string
		Path         string
		ProviderName string
		CheckOptions *CheckOptions
		ExpectError  bool
	}{
		{
			Name:         "passing",
			Path:         "testdata/attributes/passing.md",
			ProviderName: "test",
		},
		{
			Name:         "missing byline",
			Path:         "testdata/attributes/missing_byline.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "missing heading",
			Path:         "testdata/attributes/missing_heading.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong byline",
			Path:         "testdata/attributes/wrong_byline.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading level",
			Path:         "testdata/attributes/wrong_heading_level.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading text",
			Path:         "testdata/attributes/wrong_heading_text.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong list order",
			Path:         "testdata/attributes/wrong_list_order.md",
			ProviderName: "test",
		},
		{
			Name:         "wrong list order",
			Path:         "testdata/attributes/wrong_list_order.md",
			ProviderName: "test",
			CheckOptions: &CheckOptions{
				AttributesSection: &CheckAttributesSectionOptions{
					RequireSchemaOrdering: true,
				},
			},
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			doc := NewDocument(testCase.Path, testCase.ProviderName)

			if err := doc.Parse(); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			doc.CheckOptions = testCase.CheckOptions

			got := doc.checkAttributesSection()

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
