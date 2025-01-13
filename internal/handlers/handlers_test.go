package handlers

import (
	models "GO-WEB/internal/Models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

// Making the struct for test and intilaize with data in a single go
var theTests = []struct {
	testName           string
	testURL            string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"About", "/about", "GET", []postData{}, http.StatusOK},
	{"Kratos Room", "/kratos-room", "GET", []postData{}, http.StatusOK},
	{"Batman Room", "/batman-room", "GET", []postData{}, http.StatusOK},
	{"Get Search Availability Page", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"POST Search Availability Page", "/search-availability", "POST", []postData{
		{key: "start_date", value: "2023-12-12"},
		{key: "end_date", value: "2024-12-12"},
	}, http.StatusOK},
	{"POST Search Availability json", "/search-availability-json", "POST", []postData{}, http.StatusOK},
	{"Contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"GET make reservation", "/makeReservation", "GET", []postData{}, http.StatusOK},
	{"POST Make Reservation missing validation", "/makeReservation", "POST", []postData{
		{key: "last_name", value: ""},
		{key: "email", value: "minakhl"},
		{key: "phone_number", value: "536764604"},
	}, http.StatusOK},
	{"POST Make Reservation", "/makeReservation", "POST", []postData{
		{key: "first_name", value: "Michael"},
		{key: "last_name", value: "Nakhla"},
		{key: "email", value: "minakhl@ejada.com"},
		{key: "phone_number", value: "536764604"},
	}, http.StatusOK},

	{"Summary", "/reservation-summary", "GET", []postData{}, http.StatusOK},
}

func convertParamsToFormData(params []postData) url.Values {
	values := url.Values{}
	for _, param := range params {
		values.Add(param.key, param.value)
	}
	return values
}
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
	ctx, err = app.Session.Load(ctx, r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	// Update the HTTP request to include the modified context.
	r = r.WithContext(ctx)

	// Return the modified request containing the session context.
	return r, nil
}
func TestHandlers(t *testing.T) {

	// routes is basically the mux
	routes := getRoutes()
	//Making a new test server to be able to test routes on it
	testServer := httptest.NewTLSServer(routes)

	// closes after the function finishes
	defer testServer.Close()

	// _ is to ignore the index
	for _, test := range theTests {
		// we have now to type of Test, get and post and each needs different logic
		if test.method == "GET" {
			if test.testName == "Summary" {
				r, _ := getSession()
				reservation := models.Reservation{
					FirstName: "first_name",
					LastName:  "last_name",
					Email:     "email@ejada.com",
					Phone:     "phone_number",
				}
				app.Session.Put(r.Context(), "reservation", reservation)
				app.Session.Put(r.Context(), "startDate", "2024-12-12")
				app.Session.Put(r.Context(), "endDate", "2025-12-12")
			}
			// Client creates a web server and then call GET on it
			resp, err := testServer.Client().Get(testServer.URL + test.testURL) //testServer.URL = base URL of form http://ipaddr:port
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for test %s ,Expected status Code %d but got %d", test.testName, test.expectedStatusCode, resp.StatusCode)
			}
		} else if test.method == "POST" {
			// This function to convert our params to suitable url data
			requestBody := convertParamsToFormData(test.params)
			resp, err := testServer.Client().PostForm(testServer.URL+test.testURL, requestBody) //testServer.URL = base URL of form http://ipaddr:port
			if err != nil {
				t.Error(err)
			}
			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for test %s ,Expected status Code %d but got %d", test.testName, test.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
