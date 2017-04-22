package main

import (
	// "fmt"
	// "io"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHTMLHandler(t *testing.T) {
	// Test a GET on /html
	r := httptest.NewRequest("GET", "/html", nil)

	// w is a response recorder
	w := httptest.NewRecorder()

	// call the endpoint
	htmlHandler(w, r)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	// Test the body
	expectedBody := "<html><body>Hello World!</body></html>"
	actualBody := string(body)

	if actualBody != expectedBody {
		t.Errorf("got %s, want %s", actualBody, expectedBody)
	}

	// Test the status code
	expectedStatusCode := 200
	actualStatusCode := res.StatusCode

	if actualStatusCode != expectedStatusCode {
		t.Errorf("got %d, want %d", actualStatusCode, expectedStatusCode)
	}

	// Test the Content-Type
	expectedContentType := "text/html; charset=utf-8"
	actualContentType := res.Header.Get("Content-Type")

	if actualContentType != expectedContentType {
		t.Errorf("got %v, want %v", actualContentType, expectedContentType)
	}
}

func TestFormattedStringHandler(t *testing.T) {
	// r is a GET request on /format-string
	r := httptest.NewRequest("GET", "/format-string", nil)

	// w is a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	formattedStringHandler(w, r)

	// Get the response
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	// Test the body
	expectedBody := "hello world"
	actualBody := string(body)

	if actualBody != expectedBody {
		t.Errorf("got %s, want %s", actualBody, expectedBody)
	}

	// Test the status code
	expectedStatusCode := 200
	actualStatusCode := res.StatusCode

	if actualStatusCode != expectedStatusCode {
		t.Errorf("got %d, want %d", actualStatusCode, expectedStatusCode)
	}

	// Test the Content-Type
	expectedContentType := "text/plain; charset=utf-8"
	actualContentType := res.Header.Get("Content-Type")

	if actualContentType != expectedContentType {
		t.Errorf("got %v, want %v", actualContentType, expectedContentType)
	}
}

func TestStringHandler(t *testing.T) {
	// r is a GET request on /string endpoint
	r := httptest.NewRequest("GET", "/string", nil)

	// w is a response recorder
	w := httptest.NewRecorder()

	// Call the handler
	stringHandler(w, r)

	// Get the response
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	// Test the body
	expectedBody := "hello world"
	actualBody := string(body)

	if actualBody != expectedBody {
		t.Errorf("got %s, want %s", actualBody, expectedBody)
	}

	// Test the status code
	expectedStatusCode := 200
	actualStatusCode := res.StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("got %d, want %d", actualStatusCode, expectedStatusCode)
	}

	// Test the Content-Type
	expectedContentType := "text/plain; charset=utf-8"
	actualContentType := res.Header.Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("got %v, expected %v", actualContentType, expectedContentType)
	}
}

func TestJSONHandler(t *testing.T) {
	// r is a GET request on /json endpoint
	r := httptest.NewRequest("GET", "/json", nil)
	r.Header.Set("Content-Type", "application/json")

	// w is the response recorder
	w := httptest.NewRecorder()

	// call the handler
	handler := http.HandlerFunc(jsonHandler)
	handler.ServeHTTP(w, r)

	// handle the response
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	// Test the body
	expectedBody := `{"name":"john.doe","email":"john.doe@mail.com"}`
	actualBody := string(body)

	if actualBody != expectedBody {
		t.Errorf("got %s, want %s", actualBody, expectedBody)
	}

	// Test the status code
	expectedStatusCode := 200
	actualStatusCode := res.StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("got %d, want %d", actualStatusCode, expectedStatusCode)
	}

	// Test the Content-Type
	expectedContentType := "application/json; charset=utf-8"
	// Why does this fail to work?

	// If you set w.WriteHeader(http.StatusOK), this will not work somehow
	actualContentType := res.Header.Get("Content-Type")
	// actualContentType := w.HeaderMap["Content-Type"][0]
	if actualContentType != expectedContentType {
		t.Errorf("got %v, expected %v", actualContentType, expectedContentType)
	}

}

func TestFormHandler(t *testing.T) {
	// Construct the form post values
	form := url.Values{}
	form.Add("name", "john.doe")

	// r is a http POST to /form
	r := httptest.NewRequest("POST", "/form", strings.NewReader(form.Encode()))

	// Add a content-type to ensure the request is posted as a form
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// w is the response recorder
	w := httptest.NewRecorder()

	// Call the handler
	formHandler(w, r)

	// parse the response
	res := w.Result()

	// Test the body
	expectedBody := `{"name":"john.doe"}`
	actualBody := w.Body.String() //string(body)

	if actualBody != expectedBody {
		t.Errorf("got %s, want %s", actualBody, expectedBody)
	}

	// Test the status code
	expectedStatusCode := 200
	actualStatusCode := res.StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("got %d, want %d", actualStatusCode, expectedStatusCode)
	}

	// Test the Content-Type
	expectedContentType := "application/json; charset=utf-8"
	// Why does this fail to work?

	// If you set w.WriteHeader(http.StatusOK), this will not work somehow
	actualContentType := res.Header.Get("Content-Type")
	// actualContentType := w.HeaderMap["Content-Type"][0]
	if actualContentType != expectedContentType {
		t.Errorf("got %v, expected %v", actualContentType, expectedContentType)
	}
}

func TestGetHandler(t *testing.T) {
	// Construct the query
	query := url.Values{}
	query.Set("page", "1")

	// r is a GET request on /get
	r := httptest.NewRequest("GET", "/get", nil)

	// Set the query
	r.URL.RawQuery = query.Encode()
	// 	// Our handler might also expect an API key.
	// 	req.Header.Set("Authorization", "Bearer abc123")

	w := httptest.NewRecorder()

	getHandler(w, r)
	// parse the response
	res := w.Result()

	// Test the body
	expectedBody := `{"page":1}`
	actualBody := w.Body.String() //string(body)

	if actualBody != expectedBody {
		t.Errorf("got %s, want %s", actualBody, expectedBody)
	}

	// Test the status code
	expectedStatusCode := 200
	actualStatusCode := res.StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("got %d, want %d", actualStatusCode, expectedStatusCode)
	}

	// Test the Content-Type
	expectedContentType := "application/json; charset=utf-8"
	// Why does this fail to work?

	// If you set w.WriteHeader(http.StatusOK), this will not work somehow
	actualContentType := res.Header.Get("Content-Type")
	// actualContentType := w.HeaderMap["Content-Type"][0]
	if actualContentType != expectedContentType {
		t.Errorf("got %v, expected %v", actualContentType, expectedContentType)
	}
}

func TestCreateHandler(t *testing.T) {
	// Method1: This is how you test the post json
	// userJson := `{"name": "dennis", "balance": 200}`
	// data := strings.NewReader(userJson) //Convert string to reader
	// r, _ := http.NewRequest("POST", "/login", data)

	// Construct the body json
	payload := []byte(`{"name":"test product","price":11.22}`)

	// r is a http POST to /form
	r := httptest.NewRequest("POST", "/form", bytes.NewBuffer(payload))

	// Add a content-type to ensure the request is posted as a form
	r.Header.Add("Content-Type", "application/json; charset=utf-8")

	// w is the response recorder
	w := httptest.NewRecorder()

	// Call the handler
	postCreateHandler(w, r)

	// parse the response
	res := w.Result()

	// Test the body
	expectedBody := `{"ok":true}`
	actualBody := w.Body.String() //string(body)

	if actualBody != expectedBody {
		t.Errorf("got %s, want %s", actualBody, expectedBody)
	}

	// Test the status code
	expectedStatusCode := http.StatusCreated
	actualStatusCode := res.StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("got %d, want %d", actualStatusCode, expectedStatusCode)
	}
}

func TestRedirectHandler(t *testing.T) {
	// Construct the form post values

	// r is a http GET to /redirect
	r := httptest.NewRequest("GET", "/redirect", nil)

	// w is the response recorder
	w := httptest.NewRecorder()

	// Call the handler
	redirectHandler(w, r)

	// parse the response
	res := w.Result()

	// Test the status code
	expectedStatusCode := 302
	actualStatusCode := res.StatusCode
	if actualStatusCode != expectedStatusCode {
		t.Errorf("got %d, want %d", actualStatusCode, expectedStatusCode)
	}

	// Test the Content-Type
	expectedHeaderLocation := "/profile"
	actualHeaderLocation := res.Header.Get("Location")
	// actualContentType := w.HeaderMap["Content-Type"][0]
	if actualHeaderLocation != expectedHeaderLocation {
		t.Errorf("got %v, expected %v", actualHeaderLocation, expectedHeaderLocation)
	}
}

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/home", handler)
// 	http.ListenAndServe(":8080", nil)
// }

// Testing context
// func TestGetProjectsHandler(t *testing.T) {
//     req, err := http.NewRequest("GET", "/api/users", nil)
//     if err != nil {
//         t.Fatal(err)
//     }

//     rr := httptest.NewRecorder()
//     // e.g. func GetUsersHandler(ctx context.Context, w http.ResponseWriter, r *http.Request)
//     handler := http.HandlerFunc(GetUsersHandler)

//     // Populate the request's context with our test data.
//     ctx := req.Context()
//     ctx = context.WithValue(ctx, "app.auth.token", "abc123")
//     ctx = context.WithValue(ctx, "app.user",
//         &YourUser{ID: "qejqjq", Email: "user@example.com"})

//     // Add our context to the request: note that WithContext returns a copy of
//     // the request, which we must assign.
//     req = req.WithContext(ctx)
//     handler.ServeHTTP(rr, req)

//     // Check the status code is what we expect.
//     if status := rr.Code; status != http.StatusOK {
//         t.Errorf("handler returned wrong status code: got %v want %v",
//             status, http.StatusOK)
//     }
// }
// ctx := context.WithValue(r.Context(), "app.req.id", "12345")

// 		if val, ok := r.Context().Value("app.req.id").(string); !ok {
///	t.Errorf("app.req.id not in request context: got %q", val)
//}
