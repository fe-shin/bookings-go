package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks to see if there's a value for certain fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		if strings.TrimSpace(f.Get(field)) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string) bool {
	fieldValue := f.Get(field)

	if fieldValue == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}

// Valid checks if all of the form fields are valid and without error
func (f *Form) Valid() bool {
	return len((*f).Errors) == 0
}

// MinLength checks to see if a value of a field is longer than a certain length
func (f *Form) MinLength(field string, length int) bool {
	value := f.Get(field)

	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}

	return true
}

// IsEmail checks for valid email address for a form field
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
