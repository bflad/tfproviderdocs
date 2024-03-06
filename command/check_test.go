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
		ProviderName    string
		ProviderSource  string
		ProvidersSchema *tfjson.ProviderSchemas
		Expect          []string
	}{
		{
			Name:            "no providers schemas",
			ProviderName:    "test",
			ProvidersSchema: &tfjson.ProviderSchemas{},
			Expect:          nil,
		},
		{
			Name:         "provider name not found",
			ProviderName: "test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": {},
				},
			},
			Expect: nil,
		},
		{
			Name:           "provider source not found",
			ProviderSource: "registry.terraform.io/test/test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"test": {},
				},
			},
			Expect: nil,
		},
		{
			Name:         "provider name found",
			ProviderName: "test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": {},
					"test": {
						DataSourceSchemas: map[string]*tfjson.Schema{
							"test_data_source1": {},
							"test_data_source2": {},
							"test_data_source3": {},
						},
						ResourceSchemas: map[string]*tfjson.Schema{
							"test_resource1": {},
							"test_resource2": {},
							"test_resource3": {},
						},
					},
				},
			},
			Expect: []string{
				"test_data_source1",
				"test_data_source2",
				"test_data_source3",
			},
		},
		{
			Name:           "provider source found",
			ProviderSource: "registry.terraform.io/test/test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"registry.terraform.io/test/incorrect": {},
					"registry.terraform.io/test/test": {
						DataSourceSchemas: map[string]*tfjson.Schema{
							"test_data_source1": {},
							"test_data_source2": {},
							"test_data_source3": {},
						},
						ResourceSchemas: map[string]*tfjson.Schema{
							"test_resource1": {},
							"test_resource2": {},
							"test_resource3": {},
						},
					},
				},
			},
			Expect: []string{
				"test_data_source1",
				"test_data_source2",
				"test_data_source3",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			want := testCase.Expect
			got := providerSchemasDataSources(testCase.ProvidersSchema, testCase.ProviderName, testCase.ProviderSource)

			if !reflect.DeepEqual(want, got) {
				t.Errorf("mismatch:\n\nwant:\n\n%v\n\ngot:\n\n%v\n\n", want, got)
			}
		})
	}
}

func TestProviderSchemasFunctions(t *testing.T) {
	testCases := []struct {
		Name            string
		ProviderName    string
		ProviderSource  string
		ProvidersSchema *tfjson.ProviderSchemas
		Expect          []string
	}{
		{
			Name:            "no providers schemas",
			ProviderName:    "test",
			ProvidersSchema: &tfjson.ProviderSchemas{},
			Expect:          nil,
		},
		{
			Name:         "provider name not found",
			ProviderName: "test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": {},
				},
			},
			Expect: nil,
		},
		{
			Name:           "provider source not found",
			ProviderSource: "registry.terraform.io/test/test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"test": {},
				},
			},
			Expect: nil,
		},
		{
			Name:         "provider name found",
			ProviderName: "test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": {},
					"test": {
						DataSourceSchemas: map[string]*tfjson.Schema{
							"test_data_source1": {},
							"test_data_source2": {},
							"test_data_source3": {},
						},
						ResourceSchemas: map[string]*tfjson.Schema{
							"test_resource1": {},
							"test_resource2": {},
							"test_resource3": {},
						},
						Functions: map[string]*tfjson.FunctionSignature{
							"test_function1": {},
							"test_function2": {},
							"test_function3": {},
						},
					},
				},
			},
			Expect: []string{
				"test_function1",
				"test_function2",
				"test_function3",
			},
		},
		{
			Name:           "provider source found",
			ProviderSource: "registry.terraform.io/test/test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"registry.terraform.io/test/incorrect": {},
					"registry.terraform.io/test/test": {
						DataSourceSchemas: map[string]*tfjson.Schema{
							"test_data_source1": {},
							"test_data_source2": {},
							"test_data_source3": {},
						},
						ResourceSchemas: map[string]*tfjson.Schema{
							"test_resource1": {},
							"test_resource2": {},
							"test_resource3": {},
						},
						Functions: map[string]*tfjson.FunctionSignature{
							"test_function1": {},
							"test_function2": {},
							"test_function3": {},
						},
					},
				},
			},
			Expect: []string{
				"test_function1",
				"test_function2",
				"test_function3",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			want := testCase.Expect
			got := providerSchemasFunctions(testCase.ProvidersSchema, testCase.ProviderName, testCase.ProviderSource)

			if !reflect.DeepEqual(want, got) {
				t.Errorf("mismatch:\n\nwant:\n\n%v\n\ngot:\n\n%v\n\n", want, got)
			}
		})
	}
}

func TestProviderSchemasResources(t *testing.T) {
	testCases := []struct {
		Name            string
		ProviderName    string
		ProviderSource  string
		ProvidersSchema *tfjson.ProviderSchemas
		Expect          []string
	}{
		{
			Name:            "no providers schemas",
			ProviderName:    "test",
			ProvidersSchema: &tfjson.ProviderSchemas{},
			Expect:          nil,
		},
		{
			Name:         "provider name not found",
			ProviderName: "test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": {},
				},
			},
			Expect: nil,
		},
		{
			Name:           "provider source not found",
			ProviderSource: "registry.terraform.io/test/test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"test": {},
				},
			},
			Expect: nil,
		},
		{
			Name:         "provider name found",
			ProviderName: "test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"incorrect": {},
					"test": {
						DataSourceSchemas: map[string]*tfjson.Schema{
							"test_data_source1": {},
							"test_data_source2": {},
							"test_data_source3": {},
						},
						ResourceSchemas: map[string]*tfjson.Schema{
							"test_resource1": {},
							"test_resource2": {},
							"test_resource3": {},
						},
					},
				},
			},
			Expect: []string{
				"test_resource1",
				"test_resource2",
				"test_resource3",
			},
		},
		{
			Name:           "provider source found",
			ProviderSource: "registry.terraform.io/test/test",
			ProvidersSchema: &tfjson.ProviderSchemas{
				Schemas: map[string]*tfjson.ProviderSchema{
					"registry.terraform.io/test/incorrect": {},
					"registry.terraform.io/test/test": {
						DataSourceSchemas: map[string]*tfjson.Schema{
							"test_data_source1": {},
							"test_data_source2": {},
							"test_data_source3": {},
						},
						ResourceSchemas: map[string]*tfjson.Schema{
							"test_resource1": {},
							"test_resource2": {},
							"test_resource3": {},
						},
					},
				},
			},
			Expect: []string{
				"test_resource1",
				"test_resource2",
				"test_resource3",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			want := testCase.Expect
			got := providerSchemasResources(testCase.ProvidersSchema, testCase.ProviderName, testCase.ProviderSource)

			if !reflect.DeepEqual(want, got) {
				t.Errorf("mismatch:\n\nwant:\n\n%v\n\ngot:\n\n%v\n\n", want, got)
			}
		})
	}
}
