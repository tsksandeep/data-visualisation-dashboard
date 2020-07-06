package router

import (
	"net/http"
	"os"
	"time"

	log "github.com/ctrlrsf/logdna"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"know/handlers/account"
	"know/handlers/data"
	"know/middleware"
	"know/models"
)

// FileSystem is a custom file system handler to handle requests to React routes
type FileSystem struct {
	fs  http.FileSystem
	log *log.Client
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	index := "/index.html"

	f, err := fs.fs.Open(path)
	if os.IsNotExist(err) {
		if f, err = fs.fs.Open(index); err != nil {
			fs.log.Log(time.Now(), err.Error())
			return nil, err
		}
	} else if err != nil {
		fs.log.Log(time.Now(), err.Error())
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		fs.log.Log(time.Now(), err.Error())
		return nil, err
	}
	if s.IsDir() {
		if _, err = fs.fs.Open(index); err != nil {
			fs.log.Log(time.Now(), err.Error())
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
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(c.Handler)
	return &Router{Mux: r}
}

//AddRoutes adds routes to the router
func (router *Router) AddRoutes(stores *models.Stores, logDNAClient *log.Client) {
	accountHandler := account.New(stores, logDNAClient)
	dataHandler := data.New()

	router.Group(func(r chi.Router) {
		//routes to account handler
		r.Get("/info", accountHandler.Info)
		r.Get("/login", accountHandler.Login)
		r.Get("/register", accountHandler.Register)
		r.Get("/dashboard", accountHandler.Dashboard)
		r.Post("/account/login", accountHandler.PostLogin)
		r.Post("/account/register", accountHandler.PostRegister)

		//routes to data handler
		r.Get("/data/employee", dataHandler.GetEmployeeData)
		r.Get("/data/chart/month", dataHandler.GetMonthData)
		r.Get("/data/chart/day", dataHandler.GetDayData)
		r.Get("/data/chart/profit", dataHandler.GetProfitData)

		//router to data download
		r.Get("/data/download/today", dataHandler.DownloadToday)
		r.Get("/data/download/yesterday", dataHandler.DownloadYesterday)
		r.Get("/data/download/week", dataHandler.DownloadWeek)
		r.Get("/data/download/month", dataHandler.DownloadMonth)

		//default
		r.Get("/*", accountHandler.Welcome)
	})

	// set up static file serving
	fs := http.FileServer(FileSystem{fs: http.Dir("./client/"),
		log: logDNAClient})
	router.With(middleware.UICacheControl).Handle("/*", fs)
}
