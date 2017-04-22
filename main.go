package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"io/ioutil"
	"net/http/httptest"
	"strconv"
)

// indexHandler returns html string
// Recommended
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html><body>Hello World!</body></html>")
}

func formattedStringHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s", "world")
}

func stringHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	v, _ := json.Marshal(struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{"john.doe", "john.doe@mail.com"})

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(v)
}

// formHandler posts a form with values
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form
	r.ParseForm()

	// You can access the form values from r.Form or r.PostForm
	// r.Form["name"] returns an array of name ["hello"]
	names := r.Form["name"]
	if len(names) == 0 {
		fmt.Println("name is not defined")
		return
	}

	// Marshal an anonymous struct
	v, err := json.Marshal(struct {
		Name string `json:"name"`
	}{names[0]})

	if err != nil {
		// Don't panic here, return error instead
		panic(err)
		return
	}
	// OK
	// w.WriteHeader(http.StatusOK)
	// Is returning JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// If you are just writing string back, use io.WriteString

	w.Write(v)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	q := r.URL.Query()
	// Don't do this unless you want to throw error
	// Somehow in the testing, you will get "" when trying to get the header
	// content-type
	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	page, _ := strconv.Atoi(q.Get("page"))
	if page == 1 {
		io.WriteString(w, `{"page":1}`)
		return
	}
	io.WriteString(w, `{"alive": true}`)
}

type Model struct {
	Name string `json:"name"`
}

// postCreateHandler is the CRUD create handler
func postCreateHandler(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// Posting application is not a form, you need to marshal the json
	// fmt.Println("creating somtething...", r.Form, r.PostForm)
	// var m map[string]interface{}

	// fmt.Println("get body", r.Body)
	//    err = json.NewDecoder(req.Body).Decode(&m)
	var m Model
	m = Model{
		Name: "",
	}

	// body, err := ioutil.ReadAll(r.Body)

	// fmt.Println("this is model", m)
	// err = json.Unmarshal(body, &m)
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		fmt.Println("error decoding", err, m, r.Body)
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"ok":true}`)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/profile", http.StatusFound)
}

func testSomething() {
	r := httptest.NewRequest("GET", "/json", nil)

	// w is the response recorder
	w := httptest.NewRecorder()

	// call the handler
	jsonHandler(w, r)

	// handle the response
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", w)
	fmt.Println(w.HeaderMap["Content-Type"][0], res.Header, body)
}

func main() {

	http.HandleFunc("/html", htmlHandler)
	http.HandleFunc("/format-string", formattedStringHandler)
	http.HandleFunc("/string", stringHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/create", postCreateHandler)
	http.HandleFunc("/redirect", redirectHandler)
	testSomething()

	fmt.Println("listening to port:8080")
	http.ListenAndServe(":8080", nil)
}
