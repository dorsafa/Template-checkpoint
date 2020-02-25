package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main(){
	type Email struct {
		Title string
		Subject string
		To string
		Message string
		Greeting string
	}

	emailData := Email{
		Title:    "EMAIL",
		Subject:  "DEMANDE AVIS",
		To:       "dorsaf.ayed@gmail.com",
		Message:  "Hello jksfhggklgjrzkl jzrk mjhgkg",
		Greeting: "sqkjfhjzemljmzlj",
	}

	templ,err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	err1 := templ.Execute(os.Stdout , emailData)
	if err != nil {
		log.Fatal(err1)
	}

}




/*************** ENONCE ***********************

Sometimes we need to create a custom email (Subject,To,Message,greeting),
we can use a golang template to create a generic email.
Create a go struct to prepare data to new template .
Type Email struct { Title String Subject string To string Message String Greeting string }
So prepare a template for this ?
 *********************************************/