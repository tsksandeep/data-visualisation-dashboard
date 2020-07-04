package router

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	"know/handlers/account"
	"know/middleware"
)

// FileSystem is a custom file system handler to handle requests to React routes
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	index := "/index.html"

	f, err := fs.fs.Open(path)
	if os.IsNotExist(err) {
		if f, err = fs.fs.Open(index); err != nil {
			log.Error(err)
			return nil, err
		}
	} else if err != nil {
		log.Error(err)
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if s.IsDir() {
		if _, err = fs.fs.Open(index); err != nil {
			log.Error(err)
			return nil, err
		}
	}

	return f, nil
}

//Router is the wrapper for go chi
type Router struct {
	*chi.Mux
}

//NewRouter creates new router
func NewRouter() *Router {
	return &Router{Mux: chi.NewRouter()}
}

//AddRoutes adds routes to the router
func (router *Router) AddRoutes() {
	accountHandler := account.New()

	router.Group(func(r chi.Router) {
		r.Get("/login", accountHandler.Login)
		r.Get("/register", accountHandler.Register)
		r.Get("/dashboard", accountHandler.Dashboard)
		r.Post("/account/login", accountHandler.PostLogin)
		r.Post("/account/register", accountHandler.PostRegister)
	})

	// set up static file serving
	fs := http.FileServer(FileSystem{fs: http.Dir("./client/")})
	router.With(middleware.UICacheControl).Handle("/*", fs)
}
