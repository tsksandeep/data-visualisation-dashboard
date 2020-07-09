package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/logdna/logdna-go/logger"

	"know/handlers/account"
	"know/handlers/data"
	"know/middleware"
	"know/models"
)

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
func (router *Router) AddRoutes(stores *models.Stores, log *logger.Logger) {
	accountHandler := account.New(stores, log)
	dataHandler := data.New(log)

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
	}).Use(middleware.UICacheControl)
}
