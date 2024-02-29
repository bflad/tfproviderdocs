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
			Name:     "valid registry directories with cdktf docs",
			BasePath: "testdata/valid-registry-directories-with-cdktf",
		},
		{
			Name:     "valid legacy directories",
			BasePath: "testdata/valid-legacy-directories",
		},
		{
			Name:     "valid legacy directories with cdktf docs",
			BasePath: "testdata/valid-legacy-directories-with-cdktf",
		},
		{
			Name:     "valid mixed directories",
			BasePath: "testdata/valid-mixed-directories",
		},
		{
			Name:        "invalid registry directories",
			BasePath:    "testdata/invalid-registry-directories",
			ExpectError: true,
		},
		{
			Name:        "invalid legacy directories",
			BasePath:    "testdata/invalid-legacy-directories",
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

			if testCase.Options.DataSourceFileMismatch == nil {
				testCase.Options.DataSourceFileMismatch = &FileMismatchOptions{}
			}

			if testCase.Options.DataSourceFileMismatch.FileOptions == nil {
				testCase.Options.DataSourceFileMismatch.FileOptions = fileOpts
			}

			if testCase.Options.DataSourceFileMismatch.ProviderName == "" {
				testCase.Options.DataSourceFileMismatch.ProviderName = "test"
			}

			if testCase.Options.FunctionFileMismatch == nil {
				testCase.Options.FunctionFileMismatch = &FileMismatchOptions{}
			}

			if testCase.Options.FunctionFileMismatch.FileOptions == nil {
				testCase.Options.FunctionFileMismatch.FileOptions = fileOpts
			}

			if testCase.Options.FunctionFileMismatch.ProviderName == "" {
				testCase.Options.FunctionFileMismatch.ProviderName = "test"
			}

			if testCase.Options.LegacyDataSourceFile == nil {
				testCase.Options.LegacyDataSourceFile = &LegacyDataSourceFileOptions{}
			}

			if testCase.Options.LegacyDataSourceFile.FileOptions == nil {
				testCase.Options.LegacyDataSourceFile.FileOptions = fileOpts
			}

			if testCase.Options.LegacyFunctionFile == nil {
				testCase.Options.LegacyFunctionFile = &LegacyFunctionFileOptions{}
			}

			if testCase.Options.LegacyFunctionFile.FileOptions == nil {
				testCase.Options.LegacyFunctionFile.FileOptions = fileOpts
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

			if testCase.Options.ProviderName == "" {
				testCase.Options.ProviderName = "test"
			}

			if testCase.Options.RegistryDataSourceFile == nil {
				testCase.Options.RegistryDataSourceFile = &RegistryDataSourceFileOptions{}
			}

			if testCase.Options.RegistryDataSourceFile.FileOptions == nil {
				testCase.Options.RegistryDataSourceFile.FileOptions = fileOpts
			}

			if testCase.Options.RegistryFunctionFile == nil {
				testCase.Options.RegistryFunctionFile = &RegistryFunctionFileOptions{}
			}

			if testCase.Options.RegistryFunctionFile.FileOptions == nil {
				testCase.Options.RegistryFunctionFile.FileOptions = fileOpts
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

			if testCase.Options.ResourceFileMismatch == nil {
				testCase.Options.ResourceFileMismatch = &FileMismatchOptions{}
			}

			if testCase.Options.ResourceFileMismatch.FileOptions == nil {
				testCase.Options.ResourceFileMismatch.FileOptions = fileOpts
			}

			if testCase.Options.ResourceFileMismatch.ProviderName == "" {
				testCase.Options.ResourceFileMismatch.ProviderName = "test"
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
