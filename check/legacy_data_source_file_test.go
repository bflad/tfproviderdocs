package check

import (
	"testing"
)

func TestLegacyDataSourceFileCheck(t *testing.T) {
	testCases := []struct {
		Name        string
		BasePath    string
		Path        string
		Options     *LegacyDataSourceFileOptions
		ExpectError bool
	}{
		{
			Name:     "valid",
			BasePath: "testdata/valid-legacy-files",
			Path:     "data_source.html.markdown",
		},
		{
			Name:        "invalid extension",
			BasePath:    "testdata/invalid-legacy-files",
			Path:        "data_source_invalid_extension.txt",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter",
			BasePath:    "testdata/invalid-legacy-files",
			Path:        "data_source_invalid_frontmatter.html.markdown",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter with sidebar_current",
			BasePath:    "testdata/invalid-legacy-files",
			Path:        "data_source_with_sidebar_current.html.markdown",
			ExpectError: true,
		},
		{
			Name:        "invalid frontmatter without layout",
			BasePath:    "testdata/invalid-legacy-files",
			Path:        "data_source_without_layout.html.markdown",
			ExpectError: true,
		},
		{
			Name:     "warn about frontmatter deprecated layout",
			BasePath: "testdata/valid-legacy-files",
			Path:     "data_source.html.markdown",
			Options: &LegacyDataSourceFileOptions{
				FrontMatter: &FrontMatterOptions{
					WarnDeprecatedFeatures: true,
				},
			},
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Options == nil {
				testCase.Options = &LegacyDataSourceFileOptions{}
			}

			if testCase.Options.FileOptions == nil {
				testCase.Options.FileOptions = &FileOptions{
					BasePath: testCase.BasePath,
				}
			}

			got := NewLegacyDataSourceFileCheck(testCase.Options).Run(testCase.Path)

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
