package account

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"

	"know/handlers"
)

type accountHandler struct {
	account map[string]map[string]string
}

//New Account handler
func New() handlers.AccountHandler {
	return &accountHandler{
		account: map[string]map[string]string{},
	}
}

func (ah *accountHandler) Login(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./client/login.html")
	if err != nil {
		log.Error(err)
		return
	}
	tmp.Execute(w, nil)
}

func (ah *accountHandler) Register(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./client/register.html")
	if err != nil {
		log.Error(err)
		return
	}
	tmp.Execute(w, nil)
}

func (ah *accountHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./client/dashboard.html")
	if err != nil {
		log.Error(err)
		return
	}
	tmp.Execute(w, nil)
}

func (ah *accountHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./client/index.html")
	if err != nil {
		log.Error(err)
		return
	}
	tmp.Execute(w, nil)
}

func (ah *accountHandler) PostLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Infof("Email : %s Password : %s", email, password)

	if _, ok := ah.account[email]; ok {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}

func (ah *accountHandler) PostRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	firstName := r.Form.Get("firstName")
	lastName := r.Form.Get("lastName")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Infof("FirstName : %s LastName : %s Email : %s Password : %s", firstName, lastName, email, password)

	ah.account[email] = map[string]string{
		"First Name" : firstName,
		"Last Name" : lastName,
		"Password" : password,
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
