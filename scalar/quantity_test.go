package scalar_test

import (
	"encoding/json"
	"github.com/s-ichikawa/piql/scalar"
	"testing"
)

type MarshalQuantityTestCase struct {
	test     string
	expected string
}

var marshalQuantityTests = []MarshalQuantityTestCase{
	{"0", "0"},
	{"1", "1"},
	{"1.1", "1.1"},
	{"99999999", "99999999"},
	{"99999999.99999999", "99999999.99999999"},
}

func TestMarshalQuantity(t *testing.T) {
	for _, test := range marshalQuantityTests {
		m := scalar.Quantity(test.test)

		w := TestWriter{}
		m.MarshalGQL(&w)
		if w.V != test.expected {
			t.Errorf(`expexted="%s" actual="%s"`, test.expected, w.V)
		}
	}
}

type UnmarshalQuantityTestCase struct {
	test     json.Number
	expected string
}

var unmarshalQuantityTests = []UnmarshalQuantityTestCase{
	{json.Number("0"), "0"},
	{json.Number("1"), "1"},
	{json.Number("1.1"), "1.1"},
	{json.Number("99999999"), "99999999"},
	{json.Number("99999999.99999999"), "99999999.99999999"},
}

func TestUnmarshalQuantity(t *testing.T) {
	for _, test := range unmarshalQuantityTests {
		q := scalar.Quantity("")
		err := q.UnmarshalGQL(test.test)
		if err != nil {
			t.Error(err)
		}
		if string(q) != test.expected {
			t.Errorf(`expexted="%s" actual="%s"`, test.expected, q)
		}
	}
}
