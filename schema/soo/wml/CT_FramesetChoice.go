// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"

	"baliance.com/gooxml"
)

type CT_FramesetChoice struct {
	Frameset []*CT_Frameset
	Frame    []*CT_Frame
}

func NewCT_FramesetChoice() *CT_FramesetChoice {
	ret := &CT_FramesetChoice{}
	return ret
}

func (m *CT_FramesetChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Frameset != nil {
		seframeset := xml.StartElement{Name: xml.Name{Local: "w:frameset"}}
		for _, c := range m.Frameset {
			e.EncodeElement(c, seframeset)
		}
	}
	if m.Frame != nil {
		seframe := xml.StartElement{Name: xml.Name{Local: "w:frame"}}
		for _, c := range m.Frame {
			e.EncodeElement(c, seframe)
		}
	}
	return nil
}

func (m *CT_FramesetChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FramesetChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "frameset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "frameset"}:
				tmp := NewCT_Frameset()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Frameset = append(m.Frameset, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "frame"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "frame"}:
				tmp := NewCT_Frame()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Frame = append(m.Frame, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_FramesetChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FramesetChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_FramesetChoice and its children
func (m *CT_FramesetChoice) Validate() error {
	return m.ValidateWithPath("CT_FramesetChoice")
}

// ValidateWithPath validates the CT_FramesetChoice and its children, prefixing error messages with path
func (m *CT_FramesetChoice) ValidateWithPath(path string) error {
	for i, v := range m.Frameset {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Frameset[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Frame {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Frame[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
