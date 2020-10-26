package contents

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/bflad/tfproviderdocs/markdown"
	"github.com/yuin/goldmark/ast"
)

type Document struct {
	CheckOptions *CheckOptions
	ProviderName string
	ResourceName string
	Sections     *Sections

	document ast.Node
	metadata map[string]interface{}
	path     string
	source   []byte
}

func NewDocument(path string, providerName string) *Document {
	return &Document{
		ProviderName: providerName,
		ResourceName: resourceName(providerName, filepath.Base(path)),
		path:         path,
	}
}

func (d *Document) Parse() error {
	var err error

	d.source, err = ioutil.ReadFile(d.path)

	if err != nil {
		return fmt.Errorf("error reading file (%s): %w", d.path, err)
	}

	d.document, d.metadata = markdown.Parse(d.source)

	// d.document.Dump(d.source, 1)

	// fmt.Println(d.metadata["page_title"])
	// fmt.Println(d.metadata["description"])

	d.Sections, err = sectionsWalker(d.document, d.source, d.ResourceName)

	if err != nil {
		return fmt.Errorf("error parsing file (%s) sections: %w", d.path, err)
	}

	return nil
}

func resourceName(providerName string, fileName string) string {
	return providerName + "_" + fileName[:strings.IndexByte(fileName, '.')]
}
