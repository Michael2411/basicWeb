package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Form creates a custom form struct
type Form struct {
	url.Values
	Errors errors
}

// create a new form struct, it will hav
// url.values is built in go
// it return a pointer to a form so that we could use it anywhere
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// check if form is valid
func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}

/*
RequiredFields checks if the specified fields are present and non-empty in the form.

// It takes one or more field names as arguments (variadic) and validates their values.
// Fields that are empty will trigger an error, which can be handled by the caller.
// Parameters:
//
//	fields ...string - A variadic list of field names to validate. which means that the function can accept different number of fields of type string
//
// Example usage:
//
//	form.RequiredFields("firstname", "lastname", "email")
*/
func (f *Form) RequiredFields(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.AddError(field, "This field cannot be empty")
		}
	}
}

// Has , check if form is in POST and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	checkedField := r.Form.Get(field)
	if checkedField == "" {
		f.Errors.AddError(field, "This field cannot be empty")
		return false
	}
	return true
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	checkedField := r.Form.Get(field)
	if len(checkedField) < length {
		f.Errors.AddError(field, fmt.Sprintf("Must be at least %d characters", length))
		return false
	}
	return true
}

// the validation rules made inside a struct which is used in validation like in ValidEmail function
type Input struct {
	Email string `validate:"email"`
}

// This validates that a string value contains a valid email This may not conform to all possibilities of any rfc standard,
// but neither does any email provider accept all possibilities.
func (f *Form) ValidEmail(field string, r *http.Request) bool {
	//intialize the validator
	validate := validator.New(validator.WithRequiredStructEnabled())
	email := r.Form.Get(field)
	input := Input{Email: email}
	err := validate.Struct(input)
	if err != nil {
		f.Errors.AddError(field, "Must be a valid Email")
		return false
	}
	return true

}
