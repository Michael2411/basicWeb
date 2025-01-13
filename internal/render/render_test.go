package render

import (
	models "GO-WEB/internal/Models"
	"net/http"
	"testing"
)

func Test_AddDefaultData(t *testing.T) {

	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flashMessage", "123")
	result := AddDefaultData(&td, r)

	if result.FlashMessage != "123" {
		t.Error("Didn't get any data")
	}
}

func TestRenderTemp(t *testing.T) {
	pathToTemplates = "./../../Templates"
	templateCache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
	app.TemplateCache = templateCache
	var r *http.Request
	var ww myWriter
	r, err = getSession()
	if err != nil {
		t.Error(err)
	}

	err = RenderTemp(&ww, "home.page.tmpl", &models.TemplateData{}, r)
	if err != nil {
		t.Error("Failed to Render Temp")
	}

	err = RenderTemp(&ww, "non-existent.page.tmpl", &models.TemplateData{}, r)
	if err == nil {
		t.Error("rendered tempalte that doesn't exist")
	}
}

func TestNewTemplate(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../Templates"

	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

}

// getSession initializes a new HTTP GET request and attaches a session to its context.
// It simulates a request to a dummy endpoint and loads the session context from the `X-Session` header.
func getSession() (*http.Request, error) {
	// Create a new HTTP GET request targeting the "/dummy" path with no body.
	r, err := http.NewRequest("GET", "/dummy", nil)
	if err != nil {
		// If the request creation fails, return the error to the caller.
		return nil, err
	}

	// Retrieve the current context associated with the HTTP request.
	ctx := r.Context()

	// Load the session context into the current context using the value from the "X-Session" header.
	// The `session.Load` function initializes or retrieves session data based on the header value.
	// Here, the second returned value (likely an error) is ignored.
	ctx, err = session.Load(ctx, r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	// Update the HTTP request to include the modified context.
	r = r.WithContext(ctx)

	// Return the modified request containing the session context.
	return r, nil
}
