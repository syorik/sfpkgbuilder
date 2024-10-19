package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackage_ToXMLString(t *testing.T) {
	pkg := NewPackage(WithVersion("58.0"))
	pkg.AddType("CustomObject", "Account")

	xmlStr, err := pkg.ToXMLString()

	assert.NoError(t, err)
	assert.Contains(t, xmlStr, `<?xml version="1.0" encoding="UTF-8"?>`)
	assert.Contains(t, xmlStr, `<Package xmlns="http://soap.sforce.com/2006/04/metadata">`)
	assert.Contains(t, xmlStr, "<types>")
	assert.Contains(t, xmlStr, "<members>Account</members>")
	assert.Contains(t, xmlStr, "<name>CustomObject</name>")
	assert.Contains(t, xmlStr, "</types>")
	assert.Contains(t, xmlStr, "<version>58.0</version>")
	assert.Contains(t, xmlStr, "</Package>")

	// Test FromXMLString using the generated XML
	parsedPkg, err := FromXMLString(xmlStr)

	assert.NoError(t, err)
	assert.Equal(t, pkg.Xmlns, parsedPkg.Xmlns)
	assert.Equal(t, pkg.Version, parsedPkg.Version)
	assert.Len(t, parsedPkg.Types, 1)
	assert.Equal(t, pkg.Types[0].Members, parsedPkg.Types[0].Members)
	assert.Equal(t, pkg.Types[0].Name, parsedPkg.Types[0].Name)
}

func TestFromXMLString_Error(t *testing.T) {
	xmlStr := `Invalid XML`

	_, err := FromXMLString(xmlStr)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error unmarshaling XML to Package")
}

func TestPackage_ToXMLString_Error(t *testing.T) {
	pkg := NewPackage(WithVersion("58.0"))
	pkg.AddType("CustomObject", "Invalid\xffMember")

	_, err := pkg.ToXMLString()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid UTF-8 in Type Member")
}

func TestPackage_ToXMLString_NilPackage(t *testing.T) {
	var pkg *Package

	_, err := pkg.ToXMLString()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "package is nil")
}

func TestPackage_AddMember(t *testing.T) {
	pkg := NewPackage()
	pkg.AddMember("CustomObject", "Account")
	pkg.AddMember("CustomObject", "Contact")
	pkg.AddMember("ApexClass", "MyClass")

	assert.Len(t, pkg.Types, 2)
	assert.Equal(t, "CustomObject", pkg.Types[0].Name)
	assert.Equal(t, []string{"Account", "Contact"}, pkg.Types[0].Members)
	assert.Equal(t, "ApexClass", pkg.Types[1].Name)
	assert.Equal(t, []string{"MyClass"}, pkg.Types[1].Members)
}
