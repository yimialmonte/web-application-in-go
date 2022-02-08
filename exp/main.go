package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name      string
	Dog       string
	Age       int
	Locations []string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name:      "Yelmi Almonte",
		Dog:       "Gau Gau",
		Age:       200,
		Locations: []string{"Santiago", "Puerto Plata", "Samana", "San Pedro"},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	data.Name = "Jhon"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
