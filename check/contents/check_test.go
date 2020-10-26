package contents

import (
	"testing"
)

func TestCheck(t *testing.T) {
	testCases := []struct {
		Name         string
		Path         string
		ProviderName string
		ExpectError  bool
	}{
		{
			Name:         "passing",
			Path:         "testdata/full.md",
			ProviderName: "test",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			doc := NewDocument(testCase.Path, testCase.ProviderName)

			if err := doc.Parse(); err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			got := doc.Check(nil)

			if got == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if got != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", got)
			}
		})
	}
}
