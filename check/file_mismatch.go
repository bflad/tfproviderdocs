package check

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-multierror"
)

type FileMismatchOptions struct {
	*FileOptions

	IgnoreFileMismatch []string

	IgnoreFileMissing []string

	ProviderName string

	ResourceType string

	ResourceNames []string
}

type FileMismatchCheck struct {
	Options *FileMismatchOptions
}

func NewFileMismatchCheck(opts *FileMismatchOptions) *FileMismatchCheck {
	check := &FileMismatchCheck{
		Options: opts,
	}

	if check.Options == nil {
		check.Options = &FileMismatchOptions{}
	}

	if check.Options.FileOptions == nil {
		check.Options.FileOptions = &FileOptions{}
	}

	return check
}

func (check *FileMismatchCheck) Run(files []string) error {
	if len(files) == 0 {
		log.Printf("[DEBUG] Skipping %s file mismatch checks due to missing file list", check.Options.ResourceType)
		return nil
	}

	if len(check.Options.ResourceNames) == 0 {
		log.Printf("[DEBUG] Skipping %s file mismatch checks, no resources found", check.Options.ResourceType)
		return nil
	}

	var extraFiles []string
	var missingFiles []string

	for _, file := range files {
		if fileHasResource(check.Options.ResourceNames, check.Options.ProviderName, file) {
			continue
		}

		if check.IgnoreFileMismatch(file) {
			continue
		}

		extraFiles = append(extraFiles, file)
	}

	for _, resourceName := range check.Options.ResourceNames {
		if resourceHasFile(files, check.Options.ProviderName, resourceName) {
			continue
		}

		if check.IgnoreFileMissing(resourceName) {
			continue
		}

		missingFiles = append(missingFiles, resourceName)
	}

	var result *multierror.Error

	for _, extraFile := range extraFiles {
		err := fmt.Errorf("matching %s for documentation file (%s) not found, file is extraneous or incorrectly named", check.Options.ResourceType, extraFile)
		result = multierror.Append(result, err)
	}

	for _, missingFile := range missingFiles {
		err := fmt.Errorf("missing documentation file for %s: %s", check.Options.ResourceType, missingFile)
		result = multierror.Append(result, err)
	}

	return result.ErrorOrNil()
}

func (check *FileMismatchCheck) IgnoreFileMismatch(file string) bool {
	for _, ignoreResourceName := range check.Options.IgnoreFileMismatch {
		if ignoreResourceName == fileResourceName(check.Options.ProviderName, file) {
			return true
		}
	}

	return false
}

func (check *FileMismatchCheck) IgnoreFileMissing(resourceName string) bool {
	for _, ignoreResourceName := range check.Options.IgnoreFileMissing {
		if ignoreResourceName == resourceName {
			return true
		}
	}

	return false
}

func fileHasResource(resourceNames []string, providerName, file string) bool {
	for _, name := range resourceNames {
		if name == fileResourceName(providerName, file) {
			return true
		}
	}

	return false
}

func fileResourceName(providerName, fileName string) string {
	resourceSuffix := TrimFileExtension(fileName)

	// providerName is empty for functions
	if providerName == "" {
		return resourceSuffix
	}
	return fmt.Sprintf("%s_%s", providerName, resourceSuffix)
}

func resourceHasFile(files []string, providerName, resourceName string) bool {
	var found bool

	for _, file := range files {
		if fileResourceName(providerName, file) == resourceName {
			found = true
			break
		}
	}

	return found
}
