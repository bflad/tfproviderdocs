package command

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/bflad/tfproviderdocs/check"
	"github.com/mitchellh/cli"
)

type CheckCommandConfig struct {
	AllowedGuideSubcategories    string
	AllowedResourceSubcategories string
	LogLevel                     string
	Path                         string
	RequireGuideSubcategory      bool
	RequireResourceSubcategory   bool
}

// CheckCommand is a Command implementation
type CheckCommand struct {
	Ui cli.Ui
}

func (*CheckCommand) Help() string {
	optsBuffer := bytes.NewBuffer([]byte{})
	opts := tabwriter.NewWriter(optsBuffer, 0, 0, 1, ' ', 0)
	LogLevelFlagHelp(opts)
	fmt.Fprintf(opts, CommandHelpOptionFormat, "-allowed-guide-subcategories", "Comma separated list of allowed guide frontmatter subcategories.")
	fmt.Fprintf(opts, CommandHelpOptionFormat, "-allowed-resource-subcategories", "Comma separated list of allowed data source and resource frontmatter subcategories.")
	fmt.Fprintf(opts, CommandHelpOptionFormat, "-require-guide-subcategory", "Require guide frontmatter subcategory.")
	fmt.Fprintf(opts, CommandHelpOptionFormat, "-require-resource-subcategory", "Require data source and resource frontmatter subcategory.")
	opts.Flush()

	helpText := fmt.Sprintf(`
Usage: tfproviderdocs check [options] [PATH]

  Performs documentation directory and file checks against the given Terraform Provider codebase.

Options:

%s
`, optsBuffer.String())

	return strings.TrimSpace(helpText)
}

func (c *CheckCommand) Name() string { return "check" }

func (c *CheckCommand) Run(args []string) int {
	var config CheckCommandConfig

	flags := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	flags.Usage = func() { c.Ui.Info(c.Help()) }
	LogLevelFlag(flags, &config.LogLevel)
	flags.StringVar(&config.AllowedGuideSubcategories, "allowed-guide-subcategories", "", "")
	flags.StringVar(&config.AllowedResourceSubcategories, "allowed-resource-subcategories", "", "")
	flags.BoolVar(&config.RequireGuideSubcategory, "require-guide-subcategory", false, "")
	flags.BoolVar(&config.RequireResourceSubcategory, "require-resource-subcategory", false, "")

	if err := flags.Parse(args); err != nil {
		flags.Usage()
		return 1
	}

	args = flags.Args()

	if len(args) == 1 {
		config.Path = args[0]
	}

	ConfigureLogging(c.Name(), config.LogLevel)

	directories, err := check.GetDirectories(config.Path)

	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error getting Terraform Provider documentation directories: %s", err))
		return 1
	}

	if len(directories) == 0 {
		if config.Path == "" {
			c.Ui.Error("No Terraform Provider documentation directories found in current path")
		} else {
			c.Ui.Error(fmt.Sprintf("No Terraform Provider documentation directories found in path: %s", config.Path))
		}

		return 1
	}

	var allowedGuideSubcategories, allowedResourceSubcategories []string

	if v := config.AllowedGuideSubcategories; v != "" {
		allowedGuideSubcategories = strings.Split(v, ",")
	}

	if v := config.AllowedResourceSubcategories; v != "" {
		allowedResourceSubcategories = strings.Split(v, ",")
	}

	fileOpts := &check.FileOptions{
		BasePath: config.Path,
	}
	checkOpts := &check.CheckOptions{
		LegacyDataSourceFile: &check.LegacyDataSourceFileOptions{
			FileOptions: fileOpts,
			FrontMatter: &check.FrontMatterOptions{
				AllowedSubcategories: allowedResourceSubcategories,
				RequireSubcategory:   config.RequireResourceSubcategory,
			},
		},
		LegacyGuideFile: &check.LegacyGuideFileOptions{
			FileOptions: fileOpts,
			FrontMatter: &check.FrontMatterOptions{
				AllowedSubcategories: allowedGuideSubcategories,
				RequireSubcategory:   config.RequireGuideSubcategory,
			},
		},
		LegacyIndexFile: &check.LegacyIndexFileOptions{
			FileOptions: fileOpts,
		},
		LegacyResourceFile: &check.LegacyResourceFileOptions{
			FileOptions: fileOpts,
			FrontMatter: &check.FrontMatterOptions{
				AllowedSubcategories: allowedResourceSubcategories,
				RequireSubcategory:   config.RequireResourceSubcategory,
			},
		},
		RegistryDataSourceFile: &check.RegistryDataSourceFileOptions{
			FileOptions: fileOpts,
			FrontMatter: &check.FrontMatterOptions{
				AllowedSubcategories: allowedResourceSubcategories,
				RequireSubcategory:   config.RequireResourceSubcategory,
			},
		},
		RegistryGuideFile: &check.RegistryGuideFileOptions{
			FileOptions: fileOpts,
			FrontMatter: &check.FrontMatterOptions{
				AllowedSubcategories: allowedGuideSubcategories,
				RequireSubcategory:   config.RequireGuideSubcategory,
			},
		},
		RegistryIndexFile: &check.RegistryIndexFileOptions{
			FileOptions: fileOpts,
		},
		RegistryResourceFile: &check.RegistryResourceFileOptions{
			FileOptions: fileOpts,
			FrontMatter: &check.FrontMatterOptions{
				AllowedSubcategories: allowedResourceSubcategories,
				RequireSubcategory:   config.RequireResourceSubcategory,
			},
		},
	}

	if err := check.NewCheck(checkOpts).Run(directories); err != nil {
		c.Ui.Error(fmt.Sprintf("Error checking Terraform Provider documentation: %s", err))
		return 1
	}

	return 0
}

func (c *CheckCommand) Synopsis() string {
	return "Checks Terraform Provider documentation"
}
