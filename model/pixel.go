package model

import (
	"encoding/json"
	"github.com/s-ichikawa/piql/scalar"
)

type Pixel struct {
	Id       string          `json:"id"`
	Quantity scalar.Quantity `json:"quantity"`
}

func (p *Pixel) UnmarshalJSON(b []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "quantity":
			p.Quantity = scalar.Quantity(v)
		}
	}
	return nil
}
