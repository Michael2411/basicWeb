package forms

type errors map[string][]string

func (e errors) AddError(field, message string) {
	// appending the field name and the error message related to it to an entry in the errors object
	e[field] = append(e[field], message)
}

func (e errors) GetError(field string) string {
	errorMessage := e[field]
	if len(errorMessage) == 0 {
		return ""
	} else {
		return errorMessage[0]
	}
}
