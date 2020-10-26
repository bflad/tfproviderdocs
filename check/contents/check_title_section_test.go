package contents

import (
	"testing"
)

func TestCheckTitleSection(t *testing.T) {
	testCases := []struct {
		Name         string
		Path         string
		ProviderName string
		ExpectError  bool
	}{
		{
			Name:         "passing",
			Path:         "testdata/title/passing.md",
			ProviderName: "test",
		},
		{
			Name:         "missing heading",
			Path:         "testdata/title/missing_heading.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "missing heading resource type",
			Path:         "testdata/title/missing_heading_resource_type.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading level",
			Path:         "testdata/title/wrong_heading_level.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong resource in heading",
			Path:         "testdata/title/wrong_resource_in_heading.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong code block section",
			Path:         "testdata/title/wrong_code_block_section.md",
			ProviderName: "test",
			ExpectError:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			doc := NewDocument(testCase.Path, testCase.ProviderName)

			if err := doc.Parse(); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got := doc.checkTitleSection()

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
