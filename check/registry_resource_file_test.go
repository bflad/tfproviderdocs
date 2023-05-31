package check

import (
	"testing"
)

func TestRegistryResourceFileCheck(t *testing.T) {
	testCases := []struct {
		Name            string
		BasePath        string
		Path            string
		ExampleLanguage string
		Options         *RegistryResourceFileOptions
		ExpectError     bool
	}{
		{
			Name:            "valid",
			BasePath:        "testdata/valid-registry-files",
			Path:            "resource.md",
			ExampleLanguage: "terraform",
		},
		{
			Name:            "invalid extension",
			BasePath:        "testdata/invalid-registry-files",
			Path:            "resource_invalid_extension.markdown",
			ExampleLanguage: "terraform",
			ExpectError:     true,
		},
		{
			Name:            "invalid frontmatter",
			BasePath:        "testdata/invalid-registry-files",
			Path:            "resource_invalid_frontmatter.md",
			ExampleLanguage: "terraform",
			ExpectError:     true,
		},
		{
			Name:            "invalid frontmatter with layout",
			BasePath:        "testdata/invalid-registry-files",
			Path:            "resource_with_layout.md",
			ExampleLanguage: "terraform",
			ExpectError:     true,
		},
		{
			Name:            "invalid frontmatter with sidebar_current",
			BasePath:        "testdata/invalid-registry-files",
			Path:            "resource_with_sidebar_current.md",
			ExampleLanguage: "terraform",
			ExpectError:     true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Options == nil {
				testCase.Options = &RegistryResourceFileOptions{}
			}

			if testCase.Options.FileOptions == nil {
				testCase.Options.FileOptions = &FileOptions{
					BasePath: testCase.BasePath,
				}
			}

			got := NewRegistryResourceFileCheck(testCase.Options).Run(testCase.Path, testCase.ExampleLanguage)

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
