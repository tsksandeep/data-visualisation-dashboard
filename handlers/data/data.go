package data

import (
	"encoding/json"
	"net/http"

	"know/handlers"
)

type dataHandler struct{}

//New data handler
func New() handlers.DataHandler {
	return &dataHandler{}
}

func (dh *dataHandler) GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	employeeData := map[string][]map[string]string{}

	employeeData["employee"] = []map[string]string{
		{
			"Name":       "Tiger Nixon",
			"Position":   "System Architect",
			"Office":     "Edinburgh",
			"Age":        "61",
			"Start Date": "2011/04/25",
			"Salary":     "$320,800",
		},
		{
			"Name":       "Cedric Kelly",
			"Position":   "Senior Javascript Developer",
			"Office":     "Edinburgh",
			"Age":        "22",
			"Start Date": "2012/03/29",
			"Salary":     "$433,060",
		},
		{
			"Name":       "Ashton Cox",
			"Position":   "Junior Technical Author",
			"Office":     "San Francisco",
			"Age":        "66",
			"Start Date": "2009/01/12",
			"Salary":     "$86,000",
		},
		{
			"Name":       "Garrett Winters",
			"Position":   "Accountant",
			"Office":     "Tokyo",
			"Age":        "63",
			"Start Date": "2011/07/25",
			"Salary":     "$170,750",
		},
		{
			"Name":       "Tiger Nixon",
			"Position":   "System Architect",
			"Office":     "Edinburgh",
			"Age":        "61",
			"Start Date": "2011/04/25",
			"Salary":     "$320,800",
		},
		{
			"Name":       "Cedric Kelly",
			"Position":   "Senior Javascript Developer",
			"Office":     "Edinburgh",
			"Age":        "22",
			"Start Date": "2012/03/29",
			"Salary":     "$433,060",
		},
		{
			"Name":       "Ashton Cox",
			"Position":   "Junior Technical Author",
			"Office":     "San Francisco",
			"Age":        "66",
			"Start Date": "2009/01/12",
			"Salary":     "$86,000",
		},
		{
			"Name":       "Garrett Winters",
			"Position":   "Accountant",
			"Office":     "Tokyo",
			"Age":        "63",
			"Start Date": "2011/07/25",
			"Salary":     "$170,750",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employeeData)
}

func (dh *dataHandler) GetMonthData(w http.ResponseWriter, r *http.Request) {
	monthData := map[string]int{
		"January":30000,
		"February":20000,
		"March":45000,
		"April":27000,
		"May":15000,
		"June":35000,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monthData)
}

func (dh *dataHandler) GetDayData(w http.ResponseWriter, r *http.Request) {
	dayData := map[string]int{
		"June 1":30000,
		"June 2":20000,
		"June 3":45000,
		"June 4":65000,
		"June 5":15000,
		"June 6":27000,
		"June 7":34000,
		"June 8":45000,
		"June 9":35000,
		"June 10":45000,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dayData)
}

func (dh *dataHandler) GetProfitData(w http.ResponseWriter, r *http.Request) {
	profitData := map[string]int{
		"January":30000,
		"February":20000,
		"March":45000,
		"April":27000,
		"May":15000,
		"June":35000,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profitData)
}