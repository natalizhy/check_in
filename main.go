package main

import (
	"bytes"
	"fmt"
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

var View = jet.NewHTMLSet("./views")
var root, _ = os.Getwd()

func SignupGet(w http.ResponseWriter, r *http.Request) {
	signupUser := SignupUser{}

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
	signupUser := SignupUser{}

	signupUser.Phone = r.FormValue("phone")
	signupUser.Password = r.FormValue("pass")

	fmt.Println(signupUser.Phone, signupUser.Password)

	if signupUser.Phone == "" {
		signupUser.Phone = "Empty phone"
	}else {
		signupUser.Phone = ""
	}
	if signupUser.Password == "" {
		signupUser.Password = "Empty password"
	}else {
		signupUser.Phone = ""
	}

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

func LoginGet(w http.ResponseWriter, r *http.Request){
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

func LoginPost(w http.ResponseWriter, r * http.Request) {
	loginUser := LoginUser{}

	loginUser.Phone = r.FormValue("phone")
	loginUser.Password = r.FormValue("pass")

	fmt.Println(loginUser.Phone, loginUser.Password)

	if loginUser.Phone == "" {
		loginUser.Phone = "Empty phone"
	}else {
		loginUser.Phone = ""
	}
	if loginUser.Password == "" {
		loginUser.Password = "Empty password"
	}else {
		loginUser.Phone = ""
	}

	View = jet.NewHTMLSet(filepath.Join(root, "views"))

	tmpl, err := View.GetTemplate("signup.jet.html")
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

func main() {

	r := chi.NewRouter()

	r.Get("/signup", SignupGet)
	r.Post("/signup", SignupPost)

	r.Get("/login", LoginGet)
	r.Post("/login", LoginPost)

	log.Fatal(http.ListenAndServe(":8000", r))

}
