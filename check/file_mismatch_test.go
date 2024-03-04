package check

import (
	"testing"
)

func TestFileHasResource(t *testing.T) {
	testCases := []struct {
		Name      string
		File      string
		Resources []string
		Expect    bool
	}{
		{
			Name: "found",
			File: "resource1.md",
			Resources: []string{
				"test_resource1",
				"test_resource2",
			},
			Expect: true,
		},
		{
			Name: "not found",
			File: "resource1.md",
			Resources: []string{
				"test_resource2",
				"test_resource3",
			},
			Expect: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := fileHasResource(testCase.Resources, "test", testCase.File)
			want := testCase.Expect

			if got != want {
				t.Errorf("expected %t, got %t", want, got)
			}
		})
	}
}

func TestFileResourceName(t *testing.T) {
	testCases := []struct {
		Name   string
		File   string
		Expect string
	}{
		{
			Name:   "filename with single extension",
			File:   "file.md",
			Expect: "test_file",
		},
		{
			Name:   "filename with multiple extensions",
			File:   "file.html.markdown",
			Expect: "test_file",
		},
		{
			Name:   "full path with single extensions",
			File:   "docs/resource/thing.md",
			Expect: "test_thing",
		},
		{
			Name:   "full path with multiple extensions",
			File:   "website/docs/r/thing.html.markdown",
			Expect: "test_thing",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := fileResourceName("test", testCase.File)
			want := testCase.Expect

			if got != want {
				t.Errorf("expected %s, got %s", want, got)
			}
		})
	}
}

func TestFileMismatchCheck(t *testing.T) {
	testCases := []struct {
		Name        string
		Files       []string
		Options     *FileMismatchOptions
		ExpectError bool
	}{
		{
			Name: "all found",
			Files: []string{
				"resource1.md",
				"resource2.md",
			},
			Options: &FileMismatchOptions{
				ProviderName: "test",
				ResourceNames: []string{
					"test_resource1",
					"test_resource2",
				},
			},
		},
		{
			Name: "extra file",
			Files: []string{
				"resource1.md",
				"resource2.md",
				"resource3.md",
			},
			Options: &FileMismatchOptions{
				ProviderName: "test",
				ResourceNames: []string{
					"test_resource1",
					"test_resource2",
				},
			},
			ExpectError: true,
		},
		{
			Name: "ignore extra file",
			Files: []string{
				"resource1.md",
				"resource2.md",
				"resource3.md",
			},
			Options: &FileMismatchOptions{
				IgnoreFileMismatch: []string{"test_resource3"},
				ProviderName:       "test",
				ResourceNames: []string{
					"test_resource1",
					"test_resource2",
				},
			},
		},
		{
			Name: "missing file",
			Files: []string{
				"resource1.md",
			},
			Options: &FileMismatchOptions{
				ProviderName: "test",
				ResourceNames: []string{
					"test_resource1",
					"test_resource2",
				},
			},
			ExpectError: true,
		},
		{
			Name: "ignore missing file",
			Files: []string{
				"resource1.md",
			},
			Options: &FileMismatchOptions{
				IgnoreFileMissing: []string{"test_resource2"},
				ProviderName:      "test",
				ResourceNames: []string{
					"test_resource1",
					"test_resource2",
				},
			},
		},
		{
			Name: "no files",
			Options: &FileMismatchOptions{
				ProviderName: "test",
				ResourceNames: []string{
					"test_resource1",
					"test_resource2",
				},
			},
		},
		{
			Name: "no schemas",
			Files: []string{
				"resource1.md",
			},
			Options: &FileMismatchOptions{
				ProviderName: "test",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := NewFileMismatchCheck(testCase.Options).Run(testCase.Files)

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}

func TestResourceHasFile(t *testing.T) {
	testCases := []struct {
		Name         string
		Files        []string
		ResourceName string
		Expect       bool
	}{
		{
			Name: "found",
			Files: []string{
				"resource1.md",
				"resource2.md",
			},
			ResourceName: "test_resource1",
			Expect:       true,
		},
		{
			Name: "not found",
			Files: []string{
				"resource1.md",
				"resource2.md",
			},
			ResourceName: "test_resource3",
			Expect:       false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := resourceHasFile(testCase.Files, "test", testCase.ResourceName)
			want := testCase.Expect

			if got != want {
				t.Errorf("expected %t, got %t", want, got)
			}
		})
	}
}
