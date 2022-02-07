package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name: %s. He is %d", u.Name, u.Age)
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"compute game", "sport", "job", "university"}}
	//fmt.Fprintf(w, "<b>NNNNN bbbbb MMMM<b>")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go contacts")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User = ....
	handleRequest()
}
