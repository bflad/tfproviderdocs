package contents

func (d *Document) checkTimeoutsSection() error {
	section := d.Sections.Timeouts

	if section == nil {
		return nil
	}

	return nil
}
