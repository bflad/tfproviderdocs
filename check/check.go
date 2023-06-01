package check

import (
	"fmt"
	"sort"

	"github.com/bflad/tfproviderdocs/markdown"
	"github.com/hashicorp/go-multierror"
)

const (
	ResourceTypeDataSource = "data source"
	ResourceTypeResource   = "resource"

	// Terraform Registry Storage Limits
	// https://www.terraform.io/docs/registry/providers/docs.html#storage-limits
	RegistryMaximumNumberOfFiles = 2000
	RegistryMaximumSizeOfFile    = 500000 // 500KB
)

type Check struct {
	Options *CheckOptions
}

type CheckOptions struct {
	DataSourceFileMismatch *FileMismatchOptions

	LegacyDataSourceFile *LegacyDataSourceFileOptions
	LegacyGuideFile      *LegacyGuideFileOptions
	LegacyIndexFile      *LegacyIndexFileOptions
	LegacyResourceFile   *LegacyResourceFileOptions

	ProviderName   string
	ProviderSource string

	RegistryDataSourceFile *RegistryDataSourceFileOptions
	RegistryGuideFile      *RegistryGuideFileOptions
	RegistryIndexFile      *RegistryIndexFileOptions
	RegistryResourceFile   *RegistryResourceFileOptions

	ResourceFileMismatch *FileMismatchOptions

	EnableCdktfCheck bool
}

func NewCheck(opts *CheckOptions) *Check {
	check := &Check{
		Options: opts,
	}

	if check.Options == nil {
		check.Options = &CheckOptions{}
	}

	return check
}

func (check *Check) Run(directories map[string][]string) error {
	if err := InvalidDirectoriesCheck(directories); err != nil {
		return err
	}

	if err := MixedDirectoriesCheck(directories); err != nil {
		return err
	}

	if err := NumberOfFilesCheck(directories); err != nil {
		return err
	}

	var result *multierror.Error

	if files, ok := directories[fmt.Sprintf("%s/%s", RegistryIndexDirectory, RegistryDataSourcesDirectory)]; ok {
		if err := NewFileMismatchCheck(check.Options.DataSourceFileMismatch).Run(files); err != nil {
			result = multierror.Append(result, err)
		}

		if err := NewRegistryDataSourceFileCheck(check.Options.RegistryDataSourceFile).RunAll(files); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if files, ok := directories[fmt.Sprintf("%s/%s", RegistryIndexDirectory, RegistryGuidesDirectory)]; ok {
		if err := NewRegistryGuideFileCheck(check.Options.RegistryGuideFile).RunAll(files); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if files, ok := directories[RegistryIndexDirectory]; ok {
		if err := NewRegistryIndexFileCheck(check.Options.RegistryIndexFile).RunAll(files); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if files, ok := directories[fmt.Sprintf("%s/%s", RegistryIndexDirectory, RegistryResourcesDirectory)]; ok {
		if err := NewFileMismatchCheck(check.Options.ResourceFileMismatch).Run(files); err != nil {
			result = multierror.Append(result, err)
		}

		if err := NewRegistryResourceFileCheck(check.Options.RegistryResourceFile).RunAll(files, markdown.FencedCodeBlockLanguageTerraform); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if check.Options.EnableCdktfCheck {
		for _, cdktfLanguage := range ValidCdktfLanguages {
			if files, ok := directories[fmt.Sprintf("%s/%s/%s/%s", RegistryIndexDirectory, CdktfIndexDirectory, cdktfLanguage, RegistryDataSourcesDirectory)]; ok {
				if err := NewFileMismatchCheck(check.Options.DataSourceFileMismatch).Run(files); err != nil {
					result = multierror.Append(result, err)
				}

				if err := NewRegistryDataSourceFileCheck(check.Options.RegistryDataSourceFile).RunAll(files); err != nil {
					result = multierror.Append(result, err)
				}
			}

			if files, ok := directories[fmt.Sprintf("%s/%s/%s/%s", RegistryIndexDirectory, CdktfIndexDirectory, cdktfLanguage, RegistryResourcesDirectory)]; ok {
				if err := NewFileMismatchCheck(check.Options.ResourceFileMismatch).Run(files); err != nil {
					result = multierror.Append(result, err)
				}

				if err := NewRegistryResourceFileCheck(check.Options.RegistryResourceFile).RunAll(files, cdktfLanguage); err != nil {
					result = multierror.Append(result, err)
				}
			}
		}
	}

	legacyDataSourcesFiles, legacyDataSourcesOk := directories[fmt.Sprintf("%s/%s", LegacyIndexDirectory, LegacyDataSourcesDirectory)]
	legacyResourcesFiles, legacyResourcesOk := directories[fmt.Sprintf("%s/%s", LegacyIndexDirectory, LegacyResourcesDirectory)]

	if legacyDataSourcesOk {
		if err := NewFileMismatchCheck(check.Options.DataSourceFileMismatch).Run(legacyDataSourcesFiles); err != nil {
			result = multierror.Append(result, err)
		}

		if err := NewLegacyDataSourceFileCheck(check.Options.LegacyDataSourceFile).RunAll(legacyDataSourcesFiles); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if files, ok := directories[fmt.Sprintf("%s/%s", LegacyIndexDirectory, LegacyGuidesDirectory)]; ok {
		if err := NewLegacyGuideFileCheck(check.Options.LegacyGuideFile).RunAll(files); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if files, ok := directories[LegacyIndexDirectory]; ok {
		if err := NewLegacyIndexFileCheck(check.Options.LegacyIndexFile).RunAll(files); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if legacyResourcesOk {
		if err := NewFileMismatchCheck(check.Options.ResourceFileMismatch).Run(legacyResourcesFiles); err != nil {
			result = multierror.Append(result, err)
		}

		if err := NewLegacyResourceFileCheck(check.Options.LegacyResourceFile).RunAll(legacyResourcesFiles, markdown.FencedCodeBlockLanguageTerraform); err != nil {
			result = multierror.Append(result, err)
		}
	}

	if check.Options.EnableCdktfCheck {
		for _, cdktfLanguage := range ValidCdktfLanguages {
			if files, ok := directories[fmt.Sprintf("%s/%s/%s/%s", LegacyIndexDirectory, CdktfIndexDirectory, cdktfLanguage, LegacyDataSourcesDirectory)]; ok {
				if err := NewFileMismatchCheck(check.Options.DataSourceFileMismatch).Run(files); err != nil {
					result = multierror.Append(result, err)
				}

				if err := NewLegacyDataSourceFileCheck(check.Options.LegacyDataSourceFile).RunAll(files); err != nil {
					result = multierror.Append(result, err)
				}
			}

			if files, ok := directories[fmt.Sprintf("%s/%s/%s/%s", LegacyIndexDirectory, CdktfIndexDirectory, cdktfLanguage, LegacyResourcesDirectory)]; ok {
				if err := NewFileMismatchCheck(check.Options.ResourceFileMismatch).Run(files); err != nil {
					result = multierror.Append(result, err)
				}

				if err := NewLegacyResourceFileCheck(check.Options.LegacyResourceFile).RunAll(files, cdktfLanguage); err != nil {
					result = multierror.Append(result, err)
				}
			}
		}
	}

	if result != nil {
		sort.Sort(result)
	}

	return result.ErrorOrNil()
}
