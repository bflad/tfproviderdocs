package check

import (
	"testing"
)

func TestGetFileExtension(t *testing.T) {
	testCases := []struct {
		Name   string
		Path   string
		Expect string
	}{
		{
			Name:   "empty path",
			Path:   "",
			Expect: "",
		},
		{
			Name:   "filename with single extension",
			Path:   "file.md",
			Expect: ".md",
		},
		{
			Name:   "filename with multiple extensions",
			Path:   "file.html.markdown",
			Expect: ".html.markdown",
		},
		{
			Name:   "full path with single extensions",
			Path:   "docs/resource/thing.md",
			Expect: ".md",
		},
		{
			Name:   "full path with multiple extensions",
			Path:   "website/docs/r/thing.html.markdown",
			Expect: ".html.markdown",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := GetFileExtension(testCase.Path)
			want := testCase.Expect

			if got != want {
				t.Errorf("expected %s, got %s", want, got)
			}
		})
	}
}

func TestTrimFileExtension(t *testing.T) {
	testCases := []struct {
		Name   string
		Path   string
		Expect string
	}{
		{
			Name:   "empty path",
			Path:   "",
			Expect: "",
		},
		{
			Name:   "filename with single extension",
			Path:   "file.md",
			Expect: "file",
		},
		{
			Name:   "filename with multiple extensions",
			Path:   "file.html.markdown",
			Expect: "file",
		},
		{
			Name:   "full path with single extensions",
			Path:   "docs/resource/thing.md",
			Expect: "thing",
		},
		{
			Name:   "full path with multiple extensions",
			Path:   "website/docs/r/thing.html.markdown",
			Expect: "thing",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := TrimFileExtension(testCase.Path)
			want := testCase.Expect

			if got != want {
				t.Errorf("expected %s, got %s", want, got)
			}
		})
	}
}
