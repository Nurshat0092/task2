package main

import (
	"net/url"
	"strings"
)

type form struct {
	url.Values
	errors errors
}

type errors map[string][]string

func newForm(data url.Values) *form {
	return &form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *form) required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.errors.add(field, "This field cannot be blank")
		}
	}
}

func (f *form) valid() bool {
	return len(f.errors) == 0
}

func (e errors) add(field, message string) {
	e[field] = append(e[field], message)
}
