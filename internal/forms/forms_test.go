package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	formData := url.Values{}
	form := NewForm(formData)
	isValid := form.IsValid()
	if !isValid {
		t.Error("Got invalid form where it should have beeb valid")
	}
}

func TestForm_RequiredFields(t *testing.T) {
	formData := url.Values{}
	form := NewForm(formData)
	form.RequiredFields("a", "b", "c")
	if form.IsValid() {
		t.Error("Form is valid despite required fields are missing")
	}

	formData = url.Values{}
	formData.Add("a", "a")
	formData.Add("b", "b")
	formData.Add("c", "c")
	form = NewForm(formData)
	form.RequiredFields("a", "b", "c")
	if !form.IsValid() {
		t.Error("Form should be valid but got invalid")
	}
}

func TestForm_HasField(t *testing.T) {
	formData := url.Values{}
	formData.Add("b", "12")
	form := NewForm(formData)
	form.Has("a")
	if form.IsValid() {
		t.Error("Form is valid despite required fields are missing")
	}

	formData = url.Values{}
	formData.Add("a", "a")
	form = NewForm(formData)
	form.Has("a")
	if !form.IsValid() {
		t.Error("Form should be valid but got invalid")
	}
}

func TestForm_MinLength(t *testing.T) {
	formData := url.Values{}
	formData.Add("a", "12")
	form := NewForm(formData)
	form.MinLength("a", 3)
	if form.IsValid() {
		t.Error("Form is valid despite length of field exceed required")
	}

	formData = url.Values{}
	formData.Add("a", "1234")
	form = NewForm(formData)
	form.MinLength("a", 3)
	if !form.IsValid() {
		t.Error("Form should be valid but got invalid")
	}
}

func TestForm_ValidEmail(t *testing.T) {
	formData := url.Values{}
	formData.Add("email", "michael@")
	form := NewForm(formData)
	form.ValidEmail("email")
	if form.IsValid() {
		t.Error("Form is valid despite invalid email")
	}

	formData = url.Values{}
	formData.Add("email", "michael@ejada.com")
	form = NewForm(formData)
	form.ValidEmail("email")
	if !form.IsValid() {
		t.Error("Form should be valid but got invalid")
	}
}

func TestError_GetError(t *testing.T) {
	formData := url.Values{}
	formData.Add("email", "michael@")
	form := NewForm(formData)
	form.ValidEmail("email")
	if form.IsValid() {
		t.Error("Form is valid despite invalid email")
	}
	errString := form.Errors.GetError("email")
	if errString == "" {
		t.Error("Expected error but got success")
	}

	formData = url.Values{}
	formData.Add("email", "michael@ejada.com")
	form = NewForm(formData)
	form.ValidEmail("email")
	if !form.IsValid() {
		t.Error("Form should be valid but got invalid")
	}
	errString = form.Errors.GetError("email")
	if errString != "" {
		t.Error("Expected Success but got error")
	}
}
