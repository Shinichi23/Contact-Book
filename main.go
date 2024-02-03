package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Contact struct {
	Name  string
	Email string
	Phone string
	Index int
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/create", createContact)
	http.HandleFunc("/delete", deleteContact)
	http.HandleFunc("/edit", editContact)
	http.HandleFunc("/edit-page", editPage)
	http.HandleFunc("/update", updateContact)
	fmt.Println("Listen & Serve ...")
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Contacts []Contact
	}{
		Contacts: contacts,
	}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}

}

var contacts []Contact

func createContact(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	contact := Contact{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
		Phone: r.FormValue("phone"),
		Index: len(contacts),
	}

	contacts = append(contacts, contact)

	http.Redirect(w, r, "/", http.StatusSeeOther)
	fmt.Println(contacts)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	// Get the index of the contact to delete
	index, err := strconv.Atoi(r.FormValue("index"))
	if err != nil || index < 0 || index >= len(contacts) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}

	// Remove the contact from the slice
	contacts = append(contacts[:index], contacts[index+1:]...)

	// Redirect to show the form again or any other desired page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Get the index of the contact to edit
	index, err := strconv.Atoi(r.FormValue("index"))
	if err != nil || index < 0 || index >= len(contacts) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Redirect to an edit page with the contact details
	http.Redirect(w, r, "/edit-page?index="+strconv.Itoa(index), http.StatusSeeOther)
}

func editPage(w http.ResponseWriter, r *http.Request) {
	// Parse the index parameter from the query string
	indexStr := r.URL.Query().Get("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(contacts) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Retrieve the contact details based on the index
	contact := contacts[index]

	// Render the edit page with the contact details
	tmpl, err := template.ParseFiles("edit.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := struct {
		Contact Contact
	}{
		Contact: contact,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
func updateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Get the index of the contact to update
	index, err := strconv.Atoi(r.FormValue("index"))
	if err != nil || index < 0 || index >= len(contacts) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		log.Println(err)
		return
	}

	// Update the contact details based on the form data
	contacts[index].Name = r.FormValue("name")
	contacts[index].Email = r.FormValue("email")
	contacts[index].Phone = r.FormValue("phone")

	// Redirect to the home page after updating
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
