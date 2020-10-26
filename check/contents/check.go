package contents

type CheckOptions struct {
	ArgumentsSection  *CheckArgumentsSectionOptions
	AttributesSection *CheckAttributesSectionOptions
}

func (d *Document) Check(opts *CheckOptions) error {
	d.CheckOptions = opts

	if err := d.checkTitleSection(); err != nil {
		return err
	}

	if err := d.checkExampleSection(); err != nil {
		return err
	}

	if err := d.checkArgumentsSection(); err != nil {
		return err
	}

	if err := d.checkAttributesSection(); err != nil {
		return err
	}

	if err := d.checkTimeoutsSection(); err != nil {
		return err
	}

	if err := d.checkImportSection(); err != nil {
		return err
	}

	return nil
}
