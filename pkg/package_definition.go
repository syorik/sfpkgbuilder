package pkg

import (
	"encoding/xml"
	"fmt"
	"strings"
	"unicode/utf8"
)

type Package struct {
	XMLName xml.Name `xml:"Package"`
	Xmlns   string   `xml:"xmlns,attr"`
	Types   []Type   `xml:"types"`
	Version string   `xml:"version"`
}

type Type struct {
	Members []string `xml:"members"`
	Name    string   `xml:"name"`
}

func (p *Package) ToXMLString() (string, error) {
	if p == nil {
		return "", fmt.Errorf("package is nil")
	}
	xmlHeader := `<?xml version="1.0" encoding="UTF-8"?>`

	// Validate UTF-8 encoding for all string fields
	if !utf8.ValidString(p.Xmlns) || !utf8.ValidString(p.Version) {
		return "", fmt.Errorf("invalid UTF-8 in Package fields")
	}
	for _, t := range p.Types {
		if !utf8.ValidString(t.Name) {
			return "", fmt.Errorf("invalid UTF-8 in Type Name")
		}
		for _, m := range t.Members {
			if !utf8.ValidString(m) {
				return "", fmt.Errorf("invalid UTF-8 in Type Member")
			}
		}
	}

	xmlBytes, err := xml.MarshalIndent(p, "", "    ")
	if err != nil {
		return "", fmt.Errorf("error marshaling Package to XML: %w", err)
	}
	return xmlHeader + "\n" + string(xmlBytes), nil
}

func FromXMLString(xmlStr string) (*Package, error) {
	var pkg Package
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))
	err := decoder.Decode(&pkg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling XML to Package: %w", err)
	}
	return &pkg, nil
}
