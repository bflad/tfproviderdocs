package check

import (
	"testing"
)

func TestCheck(t *testing.T) {
	testCases := []struct {
		Name        string
		BasePath    string
		Options     *CheckOptions
		ExpectError bool
	}{
		{
			Name:     "valid registry directories",
			BasePath: "testdata/valid-registry-directories",
		},
		{
			Name:     "valid legacy directories",
			BasePath: "testdata/valid-legacy-directories",
		},
		{
			Name:     "valid mixed directories",
			BasePath: "testdata/valid-mixed-directories",
		},
		{
			Name:        "invalid directories",
			BasePath:    "testdata/invalid-directories",
			ExpectError: true,
		},
		{
			Name:        "invalid mixed directories",
			BasePath:    "testdata/invalid-mixed-directories",
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			fileOpts := &FileOptions{
				BasePath: testCase.BasePath,
			}

			if testCase.Options == nil {
				testCase.Options = &CheckOptions{}
			}

			if testCase.Options.LegacyDataSourceFile == nil {
				testCase.Options.LegacyDataSourceFile = &LegacyDataSourceFileOptions{}
			}

			if testCase.Options.LegacyDataSourceFile.FileOptions == nil {
				testCase.Options.LegacyDataSourceFile.FileOptions = fileOpts
			}

			if testCase.Options.LegacyGuideFile == nil {
				testCase.Options.LegacyGuideFile = &LegacyGuideFileOptions{}
			}

			if testCase.Options.LegacyGuideFile.FileOptions == nil {
				testCase.Options.LegacyGuideFile.FileOptions = fileOpts
			}

			if testCase.Options.LegacyIndexFile == nil {
				testCase.Options.LegacyIndexFile = &LegacyIndexFileOptions{}
			}

			if testCase.Options.LegacyIndexFile.FileOptions == nil {
				testCase.Options.LegacyIndexFile.FileOptions = fileOpts
			}

			if testCase.Options.LegacyResourceFile == nil {
				testCase.Options.LegacyResourceFile = &LegacyResourceFileOptions{}
			}

			if testCase.Options.LegacyResourceFile.FileOptions == nil {
				testCase.Options.LegacyResourceFile.FileOptions = fileOpts
			}

			if testCase.Options.RegistryDataSourceFile == nil {
				testCase.Options.RegistryDataSourceFile = &RegistryDataSourceFileOptions{}
			}

			if testCase.Options.RegistryDataSourceFile.FileOptions == nil {
				testCase.Options.RegistryDataSourceFile.FileOptions = fileOpts
			}

			if testCase.Options.RegistryGuideFile == nil {
				testCase.Options.RegistryGuideFile = &RegistryGuideFileOptions{}
			}

			if testCase.Options.RegistryGuideFile.FileOptions == nil {
				testCase.Options.RegistryGuideFile.FileOptions = fileOpts
			}

			if testCase.Options.RegistryIndexFile == nil {
				testCase.Options.RegistryIndexFile = &RegistryIndexFileOptions{}
			}

			if testCase.Options.RegistryIndexFile.FileOptions == nil {
				testCase.Options.RegistryIndexFile.FileOptions = fileOpts
			}

			if testCase.Options.RegistryResourceFile == nil {
				testCase.Options.RegistryResourceFile = &RegistryResourceFileOptions{}
			}

			if testCase.Options.RegistryResourceFile.FileOptions == nil {
				testCase.Options.RegistryResourceFile.FileOptions = fileOpts
			}

			if testCase.Options.SideNavigation == nil {
				testCase.Options.SideNavigation = &SideNavigationOptions{}
			}

			if testCase.Options.SideNavigation.FileOptions == nil {
				testCase.Options.SideNavigation.FileOptions = fileOpts
			}

			directories, err := GetDirectories(testCase.BasePath)

			if err != nil {
				t.Fatalf("error getting directories for path (%s): %s", testCase.BasePath, err)
			}

			got := NewCheck(testCase.Options).Run(directories)

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
