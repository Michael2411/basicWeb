package models

import "GO-WEB/internal/forms"

// Holds Data sent from handlers to Templates
type TemplateData struct {
	StringMap    map[string]string
	IntMap       map[int]int
	FloatMap     map[float32]float32
	Data         map[string]interface{}
	CSRFToken    string
	FlashMessage string
	Warning      string
	Error        string
	Form         *forms.Form
	Start        string
	End          string
}

//holds reservation Data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
