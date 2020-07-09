package handlers

import "net/http"

//AccountHandler is for Account Authentication
type AccountHandler interface {
	Info(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Dashboard(w http.ResponseWriter, r *http.Request)
	Welcome(w http.ResponseWriter, r *http.Request)
	PostLogin(w http.ResponseWriter, r *http.Request)
	PostRegister(w http.ResponseWriter, r *http.Request)
}

//DataHandler is for Data serving
type DataHandler interface {
	GetEmployeeData(w http.ResponseWriter, r *http.Request)
	GetMonthData(w http.ResponseWriter, r *http.Request)
	GetDayData(w http.ResponseWriter, r *http.Request)
	GetProfitData(w http.ResponseWriter, r *http.Request)

	DownloadToday(w http.ResponseWriter, r *http.Request)
	DownloadYesterday(w http.ResponseWriter, r *http.Request)
	DownloadWeek(w http.ResponseWriter, r *http.Request)
	DownloadMonth(w http.ResponseWriter, r *http.Request)
}

//ReadUserIP reads IP address from handlers
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}