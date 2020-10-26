package contents

import (
	"reflect"
	"testing"
)

func TestNewDocument(t *testing.T) {
	testCases := []struct {
		Name           string
		Path           string
		ProviderName   string
		ExpectDocument *Document
	}{
		{
			Name:         "basic",
			Path:         "docs/r/thing.md",
			ProviderName: "test",
			ExpectDocument: &Document{
				ProviderName: "test",
				ResourceName: "test_thing",
				path:         "docs/r/thing.md",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := NewDocument(testCase.Path, testCase.ProviderName)
			want := testCase.ExpectDocument

			if !reflect.DeepEqual(got, want) {
				t.Errorf("expected %#v, got %#v", want, got)
			}
		})
	}
}

func TestDocumentParse(t *testing.T) {
	testCases := []struct {
		Name         string
		Path         string
		ProviderName string
		ExpectError  bool
	}{
		{
			Name:         "empty",
			Path:         "testdata/empty.md",
			ProviderName: "test",
		},
		{
			Name:         "full",
			Path:         "testdata/full.md",
			ProviderName: "test",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			doc := NewDocument(testCase.Path, testCase.ProviderName)

			got := doc.Parse()

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
