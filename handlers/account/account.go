package account

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"

	"know/handlers"
)

var store = sessions.NewCookieStore([]byte("mysession"))

type accountHandler struct {
	account map[string]map[string]string
}

//New Account handler
func New() handlers.AccountHandler {
	return &accountHandler{
		account: map[string]map[string]string{},
	}
}

func (ah *accountHandler) Info(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "mysession")
	if err != nil {
		log.Error("unable to get session")
	}

	if session.Values["email"] == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	firstName := ah.account[session.Values["email"].(string)]["First Name"]
	lastName := ah.account[session.Values["email"].(string)]["Last Name"]
	email := session.Values["email"].(string)
	password := ah.account[session.Values["email"].(string)]["Password"]

	data := map[string]interface{}{
		"firstName": firstName,
		"lastName":  lastName,
		"email":     email,
		"password":  password,
	}

	tmp, err := template.ParseFiles("./client/info.html")
	if err != nil {
		log.Error(err)
		return
	}
	tmp.Execute(w, data)
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
	session, err := store.Get(r, "mysession")
	if err != nil {
		log.Error("unable to get session")
	}

	if session.Values["email"] == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	userName := ah.account[session.Values["email"].(string)]["First Name"]

	data := map[string]interface{}{
		"username": userName,
	}

	tmp, err := template.ParseFiles("./client/dashboard.html")
	if err != nil {
		log.Error(err)
		return
	}
	tmp.Execute(w, data)
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

	if accountInfo, ok := ah.account[email]; ok {
		if accountInfo["Password"] == password {
			session, err := store.Get(r, "mysession")
			if err != nil {
				log.Error("unable to get session")
			}
			session.Values["email"] = email
			session.Values["password"] = password

			session.Save(r, w)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
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
		"First Name": firstName,
		"Last Name":  lastName,
		"Password":   password,
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
