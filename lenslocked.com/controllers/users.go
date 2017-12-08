package controllers

import (
	"net/http"
	"gophercise/lenslocked.com/views"
	"fmt"
	"github.com/gorilla/schema"
)

func NewUsers() *Users {
	return &Users {
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}

// This is used to render the form to signup new user accounts
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// This is used to process signup form for creating new user.
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err!= nil {
		panic(err)
	}

	fmt.Fprintln(w, form)
}