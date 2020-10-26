package contents

import (
	"testing"
)

func TestCheckArgumentsSection(t *testing.T) {
	testCases := []struct {
		Name         string
		Path         string
		ProviderName string
		CheckOptions *CheckOptions
		ExpectError  bool
	}{
		{
			Name:         "passing",
			Path:         "testdata/arguments/passing.md",
			ProviderName: "test",
		},
		{
			Name:         "missing heading",
			Path:         "testdata/arguments/missing_heading.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading level",
			Path:         "testdata/arguments/wrong_heading_level.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading text",
			Path:         "testdata/arguments/wrong_heading_text.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong list order",
			Path:         "testdata/arguments/wrong_list_order.md",
			ProviderName: "test",
		},
		{
			Name:         "wrong list order",
			Path:         "testdata/arguments/wrong_list_order.md",
			ProviderName: "test",
			CheckOptions: &CheckOptions{
				ArgumentsSection: &CheckArgumentsSectionOptions{
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

			got := doc.checkArgumentsSection()

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
