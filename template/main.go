package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", indexHandler)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8081", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		log.Println(err)
		return
	}

	//faqs := make([]map[string]interface{}, 0)

	faqs := []map[string]interface{}{

		{"icon": `<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>`,
			"title": "Why Choose GoLan?", "description": "GoLan offers unparalleled performance with built-in concurrency, clean syntax, and a robust standard library that makes development faster and more enjoyable."},

		{"icon": `<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor"> <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" /></svg>`,
			"title": "Learning Curve?", "description": "Designed with developer experience in mind, GoLan has a gentle learning curve. Its intuitive syntax and comprehensive documentation make it accessible for beginners and powerful for experts."},

		{"icon": `<svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor"> <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" /></svg>`,
			"title": " Performance Guarantee?", "description": "GoLan provides near-native performance with its efficient compilation and runtime. Goroutines and channels enable seamless concurrent programming without complex threading models."},
	}

	data := struct {
		Title   string
		TagLine string
		Faqs    []map[string]interface{}
	}{
		Title:   "GoLangBD.COM",
		TagLine: "GoLang: Easy learning curve", //Simple, Fast, Scalable Programming Language
		Faqs:    faqs,
	}

	tmpl.Execute(w, data)
	//fmt.Fprintln(w, ``)
	return
}
