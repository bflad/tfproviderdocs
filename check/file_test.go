package check

import (
	"testing"
)

func TestFullPath(t *testing.T) {
	testCases := []struct {
		Name        string
		FileOptions *FileOptions
		Path        string
		Expect      string
	}{
		{
			Name:        "without base path",
			FileOptions: &FileOptions{},
			Path:        "docs/resource/thing.md",
			Expect:      "docs/resource/thing.md",
		},
		{
			Name: "without base path",
			FileOptions: &FileOptions{
				BasePath: "/full/path/to",
			},
			Path:   "docs/resource/thing.md",
			Expect: "/full/path/to/docs/resource/thing.md",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := testCase.FileOptions.FullPath(testCase.Path)
			want := testCase.Expect

			if got != want {
				t.Errorf("expected %s, got %s", want, got)
			}
		})
	}
}
