package packages

import (
	"encoding/xml"
	"github.com/COSAE-FR/riputils/pfsense/configuration/helpers"
)

type ColumnItem struct {
	Name        string `xml:"fieldname"`
	Description string `xml:"fielddescr"`
}

type SelectOption struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type BaseField struct {
	Name             string                     `xml:"name"`
	Type             string                     `xml:"type"`
	FieldDescription string                     `xml:"fielddescr,cdata"`
	FieldName        string                     `xml:"fieldname"`
	Description      string                     `xml:"description"`
	Size             uint8                      `xml:"size"`
	Required         helpers.TrueIfPresentBool  `xml:"required"`
	Source           string                     `xml:"source,cdata"`
	SourceName       string                     `xml:"source_name"`
	SourceValue      string                     `xml:"source_value"`
	Options          []SelectOption             `xml:"options>option"`
	Multiple         helpers.TrueIfPresentBool  `xml:"multiple"`
	Sethelp          string                     `xml:"sethelp"`
	Enablefields     helpers.CommaSeparatedList `xml:"enablefields"`
	ShowDisableValue string                     `xml:"show_disable_value"`
	Cols             string                     `xml:"cols"`
	Rows             string                     `xml:"rows"`
	Encoding         string                     `xml:"encoding"`
}

type Field struct {
	BaseField
	RowHelper []BaseField `xml:"rowhelper>rowhelperfield"`
}

// Gui is the package GUI definition, located in /usr/local/pkg/*.xml
type Gui struct {
	// Metadata
	XMLName     xml.Name     `xml:"packagegui"`
	Copyright   string       `xml:"copyright,cdata"`
	Name        string       `xml:"name"`
	Title       string       `xml:"title"`
	IncludeFile string       `xml:"include_file"`
	Note string `xml:"note"`
	Tabs        []PackageTab `xml:"tabs>tab"`

	// adddeleteeditpagefields
	ListFields []ColumnItem `xml:"adddeleteeditpagefields>columnitem"`

	// Fields
	Fields []Field `xml:"fields>field"`

	// PHP helpers
	ValidationCommand   string `xml:"custom_php_validation_command"`
	ResyncConfigCommand string `xml:"custom_php_resync_config_command"`
	GlobalFunctions     string `xml:"custom_php_global_functions,cdata"`
	DeinstallCommand    string `xml:"custom_php_deinstall_command"`
	AddCommand string `xml:"custom_php_add_command"`
	DeleteCommand string `xml:"custom_php_delete_command"`
	BeforeFormCommand string `xml:"custom_php_before_form_command"`
	AfterFormCommand string `xml:"custom_php_after_form_command"`
}
