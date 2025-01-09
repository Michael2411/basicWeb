package handlers

import (
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
	{"POST Make Reservation", "/makeReservation", "POST", []postData{
		{key: "first_name", value: "Michael"},
		{key: "last_name", value: "Nakhla"},
		{key: "email", value: "minakhl@ejada.com"},
		{key: "phone_number", value: "536764604"},
	}, http.StatusOK},
	{"POST Make Reservation missing validation", "/makeReservation", "POST", []postData{
		{key: "last_name", value: ""},
		{key: "email", value: "minakhl"},
		{key: "phone_number", value: "536764604"},
	}, http.StatusOK},
}

func convertParamsToFormData(params []postData) url.Values {
	values := url.Values{}
	for _, param := range params {
		values.Add(param.key, param.value)
	}
	return values
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
