package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

/*
	表单
*/

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		for k, v := range r.Header {
			fmt.Println(k, " : ", v)
		}

		details := ContactDetails{
			Email:   strings.TrimRight(r.FormValue("email"), " \n"),
			Subject: strings.TrimRight(r.FormValue("subject"), " \n"),
			Message: strings.TrimRight(r.FormValue("message"), " \n"),
		}

		//details
		fmt.Println("email : ", details.Email)
		fmt.Println("subject : ", details.Subject)
		fmt.Println("message : ", details.Message)
		fmt.Printf("detail : %#v\n", details)

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
