package utils

import (
	"net/url"
	"strings"
)

// Form ..
type Form struct {
	url.Values
	Errors errors
}

type errors map[string][]string

// NewForm creates new form
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required ..
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.add(field, "This field cannot be blank")
		}
	}
}

// Valid ..
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (e errors) add(field, message string) {
	e[field] = append(e[field], message)
}
