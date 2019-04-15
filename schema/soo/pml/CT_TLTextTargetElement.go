// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"

	"baliance.com/gooxml"
)

type CT_TLTextTargetElement struct {
	// Character Range
	CharRg *CT_IndexRange
	// Paragraph Text Range
	PRg *CT_IndexRange
}

func NewCT_TLTextTargetElement() *CT_TLTextTargetElement {
	ret := &CT_TLTextTargetElement{}
	return ret
}

func (m *CT_TLTextTargetElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CharRg != nil {
		secharRg := xml.StartElement{Name: xml.Name{Local: "p:charRg"}}
		e.EncodeElement(m.CharRg, secharRg)
	}
	if m.PRg != nil {
		sepRg := xml.StartElement{Name: xml.Name{Local: "p:pRg"}}
		e.EncodeElement(m.PRg, sepRg)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLTextTargetElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLTextTargetElement:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "charRg"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "charRg"}:
				m.CharRg = NewCT_IndexRange()
				if err := d.DecodeElement(m.CharRg, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "pRg"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "pRg"}:
				m.PRg = NewCT_IndexRange()
				if err := d.DecodeElement(m.PRg, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TLTextTargetElement %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLTextTargetElement
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLTextTargetElement and its children
func (m *CT_TLTextTargetElement) Validate() error {
	return m.ValidateWithPath("CT_TLTextTargetElement")
}

// ValidateWithPath validates the CT_TLTextTargetElement and its children, prefixing error messages with path
func (m *CT_TLTextTargetElement) ValidateWithPath(path string) error {
	if m.CharRg != nil {
		if err := m.CharRg.ValidateWithPath(path + "/CharRg"); err != nil {
			return err
		}
	}
	if m.PRg != nil {
		if err := m.PRg.ValidateWithPath(path + "/PRg"); err != nil {
			return err
		}
	}
	return nil
}
