package check

import (
	"testing"
)

func TestLegacyResourceFileCheck(t *testing.T) {
	testCases := []struct {
		Name        string
		BasePath    string
		Path        string
		Options     *LegacyResourceFileOptions
		ExpectError bool
	}{
		{
			Name:     "valid",
			BasePath: "testdata/valid-legacy-files",
			Path:     "resource.html.markdown",
		},
		{
			Name:        "invalid extension",
			BasePath:    "testdata/invalid-legacy-files",
			Path:        "resource_invalid_extension.txt",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter",
			BasePath:    "testdata/invalid-legacy-files",
			Path:        "resource_invalid_frontmatter.html.markdown",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter without layout",
			BasePath:    "testdata/invalid-legacy-files",
			Path:        "resource_without_layout.html.markdown",
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Options == nil {
				testCase.Options = &LegacyResourceFileOptions{}
			}

			if testCase.Options.FileOptions == nil {
				testCase.Options.FileOptions = &FileOptions{
					BasePath: testCase.BasePath,
				}
			}

			got := NewLegacyResourceFileCheck(testCase.Options).Run(testCase.Path)

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
