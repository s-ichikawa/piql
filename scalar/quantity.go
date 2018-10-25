package scalar

import (
	"encoding/json"
	"fmt"
	"io"
)

type Quantity string

func (q *Quantity) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case json.Number:
		*q = Quantity(v)
	default:
		return fmt.Errorf("quantity must be number, %T given", v)

	}
	return nil
}

func (q Quantity) MarshalGQL(w io.Writer) {
	w.Write([]byte(q))
}
