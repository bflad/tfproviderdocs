package check

import (
	"testing"
)

func TestRegistryFunctionFileCheck(t *testing.T) {
	testCases := []struct {
		Name        string
		BasePath    string
		Path        string
		Options     *RegistryFunctionFileOptions
		ExpectError bool
	}{
		{
			Name:     "valid",
			BasePath: "testdata/valid-registry-files",
			Path:     "function.md",
		},
		{
			Name:        "invalid extension",
			BasePath:    "testdata/invalid-registry-files",
			Path:        "function_invalid_extension.markdown",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter",
			BasePath:    "testdata/invalid-registry-files",
			Path:        "function_invalid_frontmatter.md",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter with layout",
			BasePath:    "testdata/invalid-registry-files",
			Path:        "function_with_layout.md",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter with sidebar_current",
			BasePath:    "testdata/invalid-registry-files",
			Path:        "function_with_sidebar_current.md",
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Options == nil {
				testCase.Options = &RegistryFunctionFileOptions{}
			}

			if testCase.Options.FileOptions == nil {
				testCase.Options.FileOptions = &FileOptions{
					BasePath: testCase.BasePath,
				}
			}

			got := NewRegistryFunctionFileCheck(testCase.Options).Run(testCase.Path)

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
