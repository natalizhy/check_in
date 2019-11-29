package main

import (
	"bytes"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
	"github.com/CloudyKit/jet"
	"os"
	"path/filepath"
)

type SignupUser struct {
	Phone    string
	Password string
}

type LoginUser struct {
	Phone    string
	Password string
}

func SignupGet(w http.ResponseWriter, r *http.Request) {
	signupUser := SignupUser{}

	View := jet.NewHTMLSet("./views")
	root, _ := os.Getwd()
	View = jet.NewHTMLSet(filepath.Join(root, "views"))

	tmpl, err := View.GetTemplate("signup.jet.html")

	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	body := &bytes.Buffer{}

	err = tmpl.Execute(body, nil, signupUser)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	w.Write(body.Bytes())
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
}

func LoginGet(w http.ResponseWriter, r *http.Request) {
	loginUser := LoginUser{}

	View := jet.NewHTMLSet("./views")
	root, _ := os.Getwd()
	View = jet.NewHTMLSet(filepath.Join(root, "views"))

	tmpl, err := View.GetTemplate("login.jet.html")

	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	body := &bytes.Buffer{}

	err = tmpl.Execute(body, nil, loginUser)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	w.Write(body.Bytes())
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
}

func main() {

	r := chi.NewRouter()

	r.Get("/signup", SignupGet)
	r.Post("/signup", SignupPost)

	r.Get("/login", LoginGet)
	r.Post("/login", LoginPost)

	log.Fatal(http.ListenAndServe(":8000", r))

}
