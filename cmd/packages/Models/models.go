package models

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
}
