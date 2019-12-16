package check

import (
	"testing"
)

func TestRegistryGuideFileCheck(t *testing.T) {
	testCases := []struct {
		Name        string
		BasePath    string
		Path        string
		Options     *RegistryGuideFileOptions
		ExpectError bool
	}{
		{
			Name:     "valid",
			BasePath: "testdata/valid-registry-files",
			Path:     "guide.md",
		},
		{
			Name:        "invalid extension",
			BasePath:    "testdata/invalid-registry-files",
			Path:        "guide_invalid_extension.markdown",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter",
			BasePath:    "testdata/invalid-registry-files",
			Path:        "guide_invalid_frontmatter.md",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter without layout",
			BasePath:    "testdata/invalid-registry-files",
			Path:        "guide_without_layout.md",
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Options == nil {
				testCase.Options = &RegistryGuideFileOptions{}
			}

			if testCase.Options.FileOptions == nil {
				testCase.Options.FileOptions = &FileOptions{
					BasePath: testCase.BasePath,
				}
			}

			got := NewRegistryGuideFileCheck(testCase.Options).Run(testCase.Path)

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
