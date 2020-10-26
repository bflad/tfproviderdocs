package contents

import (
	"testing"
)

func TestCheckExampleSection(t *testing.T) {
	testCases := []struct {
		Name         string
		Path         string
		ProviderName string
		ExpectError  bool
	}{
		{
			Name:         "passing",
			Path:         "testdata/example/passing.md",
			ProviderName: "test",
		},
		{
			Name:         "missing code block language",
			Path:         "testdata/example/missing_code_block_language.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "missing heading",
			Path:         "testdata/example/missing_heading.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading level",
			Path:         "testdata/example/wrong_heading_level.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong heading text",
			Path:         "testdata/example/wrong_heading_text.md",
			ProviderName: "test",
			ExpectError:  true,
		},
		{
			Name:         "wrong code block language",
			Path:         "testdata/example/wrong_code_block_language.md",
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

			got := doc.checkExampleSection()

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
