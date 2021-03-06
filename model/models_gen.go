// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	fmt "fmt"
	io "io"
	strconv "strconv"

	scalar "github.com/s-ichikawa/piql/scalar"
)

type CallWebhook struct {
	Username   string `json:"username"`
	HashString string `json:"hashString"`
}

type DecrementPixel struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type DeleteGraph struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type DeletePixel struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	Date     string `json:"date"`
}

type DeleteUser struct {
	Username string `json:"username"`
}

type DeleteWebhook struct {
	Username   string `json:"username"`
	HashString string `json:"hashString"`
}

type GetGraph struct {
	Username string     `json:"username"`
	ID       string     `json:"id"`
	Date     *string    `json:"date"`
	Mode     *GraphMode `json:"mode"`
}

type GetGraphs struct {
	Username string `json:"username"`
}

type GetPixel struct {
	Username string `json:"username"`
	ID       string `json:"id"`
	Date     string `json:"date"`
}

type GetWebhooks struct {
	Username string `json:"username"`
}

type IncrementPixel struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type NewGraph struct {
	Username string     `json:"username"`
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Unit     string     `json:"unit"`
	Type     GraphType  `json:"type"`
	Color    GraphColor `json:"color"`
}

type NewPixel struct {
	Username string          `json:"username"`
	ID       string          `json:"id"`
	Date     string          `json:"date"`
	Quantity scalar.Quantity `json:"quantity"`
}

type NewToken struct {
	Username string `json:"username"`
	NewToken string `json:"newToken"`
}

type NewUser struct {
	Token               string `json:"token"`
	Username            string `json:"username"`
	AgreeTermsOfService string `json:"agreeTermsOfService"`
	NotMinor            string `json:"notMinor"`
}

type NewWebhook struct {
	Username string `json:"username"`
	GraphID  string `json:"graphID"`
	Type     string `json:"type"`
}

type NewWebhookResponse struct {
	HashString *string `json:"hashString"`
	Message    string  `json:"message"`
	IsSuccess  bool    `json:"isSuccess"`
}

type UpdateGraph struct {
	Username string     `json:"username"`
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Unit     string     `json:"unit"`
	Color    GraphColor `json:"color"`
}

type GraphColor string

const (
	GraphColorShibafu GraphColor = "shibafu"
	GraphColorMomiji  GraphColor = "momiji"
	GraphColorSora    GraphColor = "sora"
	GraphColorIchou   GraphColor = "ichou"
	GraphColorAjisai  GraphColor = "ajisai"
	GraphColorKuro    GraphColor = "kuro"
)

func (e GraphColor) IsValid() bool {
	switch e {
	case GraphColorShibafu, GraphColorMomiji, GraphColorSora, GraphColorIchou, GraphColorAjisai, GraphColorKuro:
		return true
	}
	return false
}

func (e GraphColor) String() string {
	return string(e)
}

func (e *GraphColor) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GraphColor(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GraphColor", str)
	}
	return nil
}

func (e GraphColor) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GraphMode string

const (
	GraphModeShort GraphMode = "short"
)

func (e GraphMode) IsValid() bool {
	switch e {
	case GraphModeShort:
		return true
	}
	return false
}

func (e GraphMode) String() string {
	return string(e)
}

func (e *GraphMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GraphMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GraphMode", str)
	}
	return nil
}

func (e GraphMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GraphType string

const (
	GraphTypeInt   GraphType = "int"
	GraphTypeFloat GraphType = "float"
)

func (e GraphType) IsValid() bool {
	switch e {
	case GraphTypeInt, GraphTypeFloat:
		return true
	}
	return false
}

func (e GraphType) String() string {
	return string(e)
}

func (e *GraphType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GraphType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GraphType", str)
	}
	return nil
}

func (e GraphType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WebhookType string

const (
	WebhookTypeIncrement WebhookType = "increment"
	WebhookTypeDecrement WebhookType = "decrement"
)

func (e WebhookType) IsValid() bool {
	switch e {
	case WebhookTypeIncrement, WebhookTypeDecrement:
		return true
	}
	return false
}

func (e WebhookType) String() string {
	return string(e)
}

func (e *WebhookType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WebhookType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WebhookType", str)
	}
	return nil
}

func (e WebhookType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
