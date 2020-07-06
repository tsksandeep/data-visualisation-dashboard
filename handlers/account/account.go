package account

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"

	"know/handlers"
	"know/models"
)

var store = sessions.NewCookieStore([]byte("mysession"))

type accountHandler struct {
	stores *models.Stores
}

//New Account handler
func New(stores *models.Stores) handlers.AccountHandler {
	return &accountHandler{
		stores: stores,
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

	accountInfo, err := ah.stores.AccountStore.Get(session.Values["email"].(string))
	if err != nil {
		log.Error("unable to get account")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"firstName": accountInfo.FirstName,
		"lastName":  accountInfo.LastName,
		"email":     accountInfo.Email,
		"password":  accountInfo.Password,
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

	accountInfo, err := ah.stores.AccountStore.Get(session.Values["email"].(string))
	if err != nil {
		log.Error("unable to get account")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"username": accountInfo.FirstName,
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

	accountInfo, err := ah.stores.AccountStore.Get(email)
	if err != nil || accountInfo.Password != password {
		log.Error("unable to get account")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	session, err := store.Get(r, "mysession")
	if err != nil {
		log.Error("unable to get session")
		return
	}

	session.Values["email"] = email
	session.Values["password"] = password

	session.Save(r, w)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	return
}

func (ah *accountHandler) PostRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	accountInfo := models.Account{}

	accountInfo.FirstName = r.Form.Get("firstName")
	accountInfo.LastName = r.Form.Get("lastName")
	accountInfo.Email = r.Form.Get("email")
	accountInfo.Password = r.Form.Get("password")

	err := ah.stores.AccountStore.Save(&accountInfo)
	if err != nil {
		log.Error("unable to save account")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
