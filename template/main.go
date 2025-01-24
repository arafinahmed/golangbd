package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signupHandler)

	//http.HandleFunc("/signupprocess", formProcess)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8081", nil)

}

func signupHandler(w http.ResponseWriter, r *http.Request) {

	//fmt.Println(r.Method, "****")

	if r.Method == "GET" {

		tmpl, err := template.ParseFiles("pages/signup.html")
		if err != nil {
			log.Println(err)
			return
		}

		tmpl.Execute(w, nil)
		//fmt.Fprintln(w, ``)
		return

	}

	if r.Method == "POST" {

		err := r.ParseMultipartForm(10 << 20) // 10MB Maximum size
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Form fields:", r.MultipartForm.Value)
		fmt.Println("File headers:", r.MultipartForm.File)

		var photoFiles []*multipart.FileHeader

		for fieldName, files := range r.MultipartForm.File {
			if strings.HasPrefix(fieldName, "photos") {
				photoFiles = append(photoFiles, files...)
			}
		}

		if len(photoFiles) == 0 {
			log.Println("No files uploaded")
			http.Error(w, "No files uploaded", http.StatusBadRequest)
			return
		}

		for i, fileHeader := range photoFiles {
			fmt.Printf("Processing file %d: %s\n", i+1, fileHeader.Filename)
			file, err := fileHeader.Open()
			if err != nil {
				log.Printf("Error opening file %s: %v\n", fileHeader.Filename, err)
				continue
			}
			defer file.Close()

			err = fileUpload(file, fileHeader)
			if err != nil {
				log.Printf("Error uploading file %s: %v\n", fileHeader.Filename, err)
				continue
			}
		}

		row := make(map[string]interface{})
		row["error"] = 1
		row["message"] = "ERROR failed"

		js, err := json.Marshal(row)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(js))

		return
		//fmt.Fprintln(w, "OK.....")
		//fmt.Fprintln(w, "Successfully received!")
		//http.Redirect(w, r, "/", http.StatusSeeOther) //303

	}

}

func fileUpload(mpf multipart.File, mfh *multipart.FileHeader) error {

	//file, header, _ := r.FormFile("file")
	defer mpf.Close()
	// create a destination file
	tmpFile := filepath.Join("upload", mfh.Filename)
	dst, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer dst.Close()
	// upload the file to destination path
	_, err = io.Copy(dst, mpf)
	return err
}

func formProcess(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form) // map[string][]string

	// map[ email: [ashik@gmail.com] name:[Ashik] passw:[test321] submitButton:[] username:[ashikn]]

	// for key, val := range r.Form {
	// 	fmt.Println(key, val)
	// }

	name := r.FormValue("name")
	username := r.FormValue("username")

	fmt.Println(name)
	fmt.Println(username)

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
		TagLine: "GoLang: Simple, Fast, Scalable Programming Language", //Simple, Fast, Scalable Programming Language
		Faqs:    faqs,
	}

	tmpl.Execute(w, data)
	//fmt.Fprintln(w, ``)
	return
}
