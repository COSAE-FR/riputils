package helpers

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type TrueIfPresentBool bool

func (b TrueIfPresentBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if b {
		return e.EncodeElement("", start)
	}
	return nil
}

func (b *TrueIfPresentBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	*b = true
	return nil
}

type OnOffBool bool

func (b OnOffBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if b {
		return e.EncodeElement("on", start)
	} else {
		return e.EncodeElement("off", start)
	}
	return nil
}

func (b *OnOffBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	*b = false
	if s == "on" {
		*b = true
	}

	return nil
}

func (b OnOffBool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

type YesNoBool bool

func (b YesNoBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if b {
		return e.EncodeElement("yes", start)
	} else {
		return e.EncodeElement("no", start)
	}
}

func (b *YesNoBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	*b = false
	if s == "yes" {
		*b = true
	}

	return nil
}

func (b YesNoBool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

type CommaSeparatedList []string

func (l *CommaSeparatedList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	if len(s) > 0 {
		*l = strings.Split(s, ",")
	} else {
		*l = []string{}
	}
	return nil
}

func (l CommaSeparatedList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strings.Join(l, ","), start)
}

type SpaceSeparatedList []string

func (l *SpaceSeparatedList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	if len(s) > 0 {
		*l = strings.Split(s, " ")
	} else {
		*l = []string{}
	}
	return nil
}

func (l SpaceSeparatedList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strings.Join(l, " "), start)
}

type SemiColonSeparatedList []string

func (l *SemiColonSeparatedList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	if len(s) > 0 {
		*l = strings.Split(s, ";")
	} else {
		*l = []string{}
	}
	return nil
}

func (l SemiColonSeparatedList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strings.Join(l, ";"), start)
}

type StringPort uint16

func (l *StringPort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	if len(s) > 0 {
		port, err := strconv.ParseUint(s, 10, 16)
		if err != nil {
			return err
		}
		*l = StringPort(port)
	} else {
		*l = 0
	}
	return nil
}

func (l StringPort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(fmt.Sprintf("%d", l), start)
}

type Base64String string

func (l *Base64String) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	if len(s) > 0 {
		data, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			*l = Base64String(s)
		} else {
			*l = Base64String(data)
		}
	} else {
		*l = ""
	}
	return nil
}

func (l Base64String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var toXml string
	if len(l) >= 0 {
		toXml = base64.StdEncoding.EncodeToString([]byte(l))
	}
	return e.EncodeElement(toXml, start)
}