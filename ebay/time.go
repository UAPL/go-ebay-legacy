package ebay

import (
	"encoding/xml"
	"time"
)

type Time struct {
	time.Time
}

func (c Time) Parse(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func (c *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}

func (c *Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(c.Time.Format(time.RFC3339Nano), start)
}
