# v0.12.0

FEATURES

* check: Support provider-defined function directory structure (#97)

ENHANCEMENTS

* check: Add `-ignore-file-mismatch-functions` and `-ignore-file-missing-functions` options (#97)

# v0.11.1

BUG FIXES

check: Only count HCL files (ignoring CDKTF files) for maximum file limit as the limit is per-language (#78)

# v0.11.0

ENHANCEMENTS

* check: Add `ignore-cdktf-missing-files` flag for providers wishing to iteratively introduce CDKTF documentation

# v0.10.0

NOTES

* all: This Go module and the associated Docker image has been updated to Go 1.19 per the [Go support policy](https://go.dev/doc/devel/release#policy). Any consumers building on earlier Go versions or dependent on earlier Go version functionality may experience errors. (#55)

BREAKING CHANGES

* check: The `ignore-side-navigation-data-sources`, `ignore-side-navigation-resources`, and `require-side-navigation` flags have been removed without replacement. Side navigation functionality has not been necessary for provider documentation since the introduction of the Terraform Registry. (#52)

FEATURES

* all: Released binaries now include `darwin/arm64`, `linux/arm64`, and `windows/arm64` (#54)
* check: Support CDKTF directory structure (#50)

# v0.9.1

BUG FIXES

* check: Accept alternate attribute section byline of `No additional attributes are exported.` with experimental `-enable-contents-check` flag

# v0.9.0

BREAKING CHANGES

* check: Prefer `terraform` code block language over `hcl` in examples with experimental `-enable-contents-check` flag

ENHANCEMENTS

* check: Add `-provider-source` option (support Terraform CLI 0.13 and later `-providers-schema-json` file)

# v0.8.0

ENHANCEMENTS

* check: Add experimental `-enable-contents-check` and `-require-schema-ordering` options

# v0.7.0

ENHANCEMENTS

* check: Match increased registry max number of files limit from 1000 to 2000

BUG FIXES

* check: Allow legacy directory structure without side navigation (use `-require-side-navigation` flag to keep deprecated old behavior)
* check: Return correct valid extensions list with registry directory structure (#38)

# v0.6.0

ENHANCEMENTS

* check: Add `-ignore-file-mismatch-data-sources` option
* check: Add `-ignore-file-mismatch-resources` option
* check: Add `-ignore-file-missing-data-sources` option
* check: Add `-ignore-file-missing-resources` option

# v0.5.3

BUG FIXES

* check: Prevent additional errors when `docs/` contains files outside Terraform Provider documentation

# v0.5.2

BUG FIXES

* check: Prevent `mixed Terraform Provider documentation directory layouts found` error when using `website/docs` and `docs/` contains files outside Terraform Provider documentation

# v0.5.1

Released without changes.

# v0.5.0

ENHANCEMENTS

* check: Verify sidebar navigation for missing links and mismatched link text (if legacy directory structure)

# v0.4.1

BUG FIXES

* check: Only verify valid file extensions at end of path (e.g. support additional periods in guide paths) (#25)

# v0.4.0

ENHANCEMENTS

* check: Accept newline-separated files of allowed subcategories with `-allowed-guide-subcategories-file` and `-allowed-resource-subcategories-file` flags
* check: Improve readability with allowed subcategories values in allowed subcategories frontmatter error

# v0.3.0

ENHANCEMENTS

* check: Verify deprecated `sidebar_current` frontmatter is not present

# v0.2.0

ENHANCEMENTS

* check: Verify number of documentation files for Terraform Registry storage limits
* check: Verify size of documentation files for Terraform Registry storage limits
* check: Verify all known data sources and resources have an associated documentation file (if `-providers-schema-json` is provided)
* check: Verify no extraneous or incorrectly named documentation files exist (if `-providers-schema-json` is provided)

# v0.1.2

BUG FIXES

* Remove extraneous `-''` from version information

# v0.1.1

BUG FIXES

* Fix help formatting of `check` command options

# v0.1.0

FEATURES

* Initial release with `check` command
