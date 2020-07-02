package rules

import "encoding/xml"

//Schema , Representacion de un XSD.
type Schema struct {
	XMLFile            string
	XMLName            xml.Name `xml:"schema"`
	Text               string   `xml:",chardata"`
	TargetNamespace    string   `xml:"targetNamespace,attr"`
	ElementFormDefault string   `xml:"elementFormDefault,attr"`
	Xmlns              string   `xml:"xmlns,attr"`
	Tns                string   `xml:"tns,attr"`
	Xsd                string   `xml:"xsd,attr"`
	Com                string   `xml:"com,attr"`
	Import             struct {
		Text           string `xml:",chardata"`
		Namespace      string `xml:"namespace,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"import"`
	Element []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		ComplexType struct {
			Text     string `xml:",chardata"`
			Sequence struct {
				Text    string `xml:",chardata"`
				Element []struct {
					Text      string `xml:",chardata"`
					Name      string `xml:"name,attr"`
					MinOccurs string `xml:"minOccurs,attr"`
					MaxOccurs string `xml:"maxOccurs,attr"`
					Type      string `xml:"type,attr"`
				} `xml:"element"`
			} `xml:"sequence"`
		} `xml:"complexType"`
	} `xml:"element"`
	ComplexType []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Sequence struct {
			Text    string `xml:",chardata"`
			Element []struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"name,attr"`
				Type      string `xml:"type,attr"`
				MinOccurs string `xml:"minOccurs,attr"`
				MaxOccurs string `xml:"maxOccurs,attr"`
			} `xml:"element"`
		} `xml:"sequence"`
		Annotation struct {
			Text          string `xml:",chardata"`
			Documentation string `xml:"documentation"`
		} `xml:"annotation"`
	} `xml:"complexType"`
	SimpleType []struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name,attr"`
		Restriction struct {
			Text      string `xml:",chardata"`
			Base      string `xml:"base,attr"`
			MaxLength struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"maxLength"`
			MinLength struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"minLength"`
			MaxInclusive struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"maxInclusive"`
			FractionDigits struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"fractionDigits"`
			Pattern struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"pattern"`
			TotalDigits struct {
				Text  string `xml:",chardata"`
				Value string `xml:"value,attr"`
			} `xml:"totalDigits"`
		} `xml:"restriction"`
	} `xml:"simpleType"`
}
