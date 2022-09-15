package main

import (
	"log"
	"net/http"
	"text/template"
	"web/ascii/asciifunc"
)

// Global varibles for exporting a file

var (
	templateString string = ""
	userString     string = ""
)

// Global varibles for 400-500 Codes
type ViewData struct {
	Number      string
	Description string
}

var notFound ViewData = ViewData{
	Number:      "404",
	Description: "Not found",
}

var badRequest ViewData = ViewData{
	Number:      "400",
	Description: "Bad Request",
}

var serverError ViewData = ViewData{
	Number:      "500",
	Description: "Internal Server Error",
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		t, _ := template.ParseFiles("./ui/html/errors.html")
		t.Execute(w, notFound)
		return
	} else {
		if r.Method != http.MethodGet {
			w.WriteHeader(400)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, badRequest)
			return
		}
		ts, err := template.ParseFiles("./ui/html/home.page.html")
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, serverError)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, serverError)
			return
		}
	}
}

func createAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		inputString := createInputStr(r.FormValue("input"))
		fontTemplate := r.FormValue("font")
		// <<<<for exportin a file
		templateString = r.FormValue("font")
		userString = createInputStr(r.FormValue("input"))
		// >>>>
		err := asciifunc.CreateDict(fontTemplate)
		if !err {
			w.WriteHeader(500)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, serverError)
			return
		}
		outputString := asciifunc.PrintWord(inputString)
		t, _ := template.ParseFiles("./ui/html/home.page.html")

		t.Execute(w, outputString)
	} else {
		w.WriteHeader(400)
		t, _ := template.ParseFiles("./ui/html/errors.html")
		t.Execute(w, badRequest)
		return
	}
}

func exportFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		inputString := userString
		fontTemplate := templateString
		err := asciifunc.CreateDict(fontTemplate)
		if !err {
			w.WriteHeader(500)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, serverError)
			return
		}
		outputString := asciifunc.PrintWord(inputString)
		err = asciifunc.CreateFile(outputString)
		if err {
			w.WriteHeader(500)
			t, _ := template.ParseFiles("./ui/html/errors.html")
			t.Execute(w, serverError)
			return
		}
		filePath := "files/data.txt"
		w.Header().Set("Content-Disposition", "attachment; filename=Data.txt")
		w.Header().Set("Content-Type", "text/plain")
		http.ServeFile(w, r, filePath)
	} else {
		w.WriteHeader(400)
		t, _ := template.ParseFiles("./ui/html/errors.html")
		t.Execute(w, badRequest)
		return
	}
}

func createInputStr(input string) string {
	slice := []rune(input)
	str := ""
	for _, l := range slice {
		if l >= 32 && l <= 126 {
			str = str + string(l)
		}
		if l == 13 || l == 10 {
			str = str + string('\n')
		}
	}
	return str
}
