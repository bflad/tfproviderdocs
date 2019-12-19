package command

import (
	"reflect"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
)

func TestAllowedSubcategoriesFile(t *testing.T) {
	testCases := []struct {
		Name        string
		Path        string
		Expect      []string
		ExpectError bool
	}{
		{
			Name: "valid",
			Path: "testdata/allowed-subcategories.txt",
			Expect: []string{
				"Example Subcategory 1",
				"Example Subcategory 2",
				"Example Subcategory 3",
			},
		},
		{
			Name:        "invalid path",
			Path:        "testdata/does-not-exist.txt",
			Expect:      nil,
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got, err := allowedSubcategoriesFile(testCase.Path)
			want := testCase.Expect

			if err == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", err)
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("expected: %v, got: %v", want, got)
			}
		})
	}
}

func TestProviderNameFromPath(t *testing.T) {
	testCases := []struct {
		Name   string
		Path   string
		Expect string
	}{
		{
			Name:   "full path without prefix",
			Path:   "/path/to/test",
			Expect: "",
		},
		{
			Name:   "full path with prefix",
			Path:   "/path/to/terraform-provider-test",
			Expect: "test",
		},
		{
			Name:   "relative path without prefix",
			Path:   "test",
			Expect: "",
		},
		{
			Name:   "relative path with prefix",
			Path:   "terraform-provider-test",
			Expect: "test",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			want := testCase.Expect
			got := providerNameFromPath(testCase.Path)

			if want != got {
				t.Errorf("expected: %s, got: %s", want, got)
			}
		})
	}
}

func TestProviderSchemas(t *testing.T) {
	testCases := []struct {
		Name        string
		Path        string
		ExpectError bool
	}{
		{
			Name: "valid",
			Path: "testdata/valid-providers-schema.json",
		},
		{
			Name:        "invalid path",
			Path:        "testdata/does-not-exist.json",
			ExpectError: true,
		},
		{
			Name:        "invalid json",
			Path:        "testdata/not-providers-schema.json",
			ExpectError: true,
		},
		{
			Name:        "invalid format version",
			Path:        "testdata/invalid-providers-schema-version.json",
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			_, err := providerSchemas(testCase.Path)

			if err == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", err)
			}
		})
	}
}

func TestProviderSchemasDataSources(t *testing.T) {
	testCases := []struct {
		Name            string
		ProvidersSchema *tfjson.ProviderSchemas
		Expect          map[string]*tfjson.Schema
	}{
		{
			Name:            "no providers schemas",
			ProvidersSchema: &tfjson.ProviderSchemas{},
			Expect:          nil,
		},
		{
			Name: "provider not found",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": &tfjson.ProviderSchema{},
				},
			},
			Expect: nil,
		},
		{
			Name: "provider found",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": &tfjson.ProviderSchema{},
					"test": &tfjson.ProviderSchema{
						DataSourceSchemas: map[string]*tfjson.Schema{
							"test_data_source1": &tfjson.Schema{},
							"test_data_source2": &tfjson.Schema{},
							"test_data_source3": &tfjson.Schema{},
						},
						ResourceSchemas: map[string]*tfjson.Schema{
							"test_resource1": &tfjson.Schema{},
							"test_resource2": &tfjson.Schema{},
							"test_resource3": &tfjson.Schema{},
						},
					},
				},
			},
			Expect: map[string]*tfjson.Schema{
				"test_data_source1": &tfjson.Schema{},
				"test_data_source2": &tfjson.Schema{},
				"test_data_source3": &tfjson.Schema{},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			want := testCase.Expect
			got := providerSchemasDataSources(testCase.ProvidersSchema, "test")

			if !reflect.DeepEqual(want, got) {
				t.Errorf("mismatch:\n\nwant:\n\n%v\n\ngot:\n\n%v\n\n", want, got)
			}
		})
	}
}

func TestProviderSchemasResources(t *testing.T) {
	testCases := []struct {
		Name            string
		ProvidersSchema *tfjson.ProviderSchemas
		Expect          map[string]*tfjson.Schema
	}{
		{
			Name:            "no providers schemas",
			ProvidersSchema: &tfjson.ProviderSchemas{},
			Expect:          nil,
		},
		{
			Name: "provider not found",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": &tfjson.ProviderSchema{},
				},
			},
			Expect: nil,
		},
		{
			Name: "provider found",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": &tfjson.ProviderSchema{},
					"test": &tfjson.ProviderSchema{
						DataSourceSchemas: map[string]*tfjson.Schema{
							"test_data_source1": &tfjson.Schema{},
							"test_data_source2": &tfjson.Schema{},
							"test_data_source3": &tfjson.Schema{},
						},
						ResourceSchemas: map[string]*tfjson.Schema{
							"test_resource1": &tfjson.Schema{},
							"test_resource2": &tfjson.Schema{},
							"test_resource3": &tfjson.Schema{},
						},
					},
				},
			},
			Expect: map[string]*tfjson.Schema{
				"test_resource1": &tfjson.Schema{},
				"test_resource2": &tfjson.Schema{},
				"test_resource3": &tfjson.Schema{},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			want := testCase.Expect
			got := providerSchemasResources(testCase.ProvidersSchema, "test")

			if !reflect.DeepEqual(want, got) {
				t.Errorf("mismatch:\n\nwant:\n\n%v\n\ngot:\n\n%v\n\n", want, got)
			}
		})
	}
}
