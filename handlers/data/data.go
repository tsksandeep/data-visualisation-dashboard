package data

import (
	"encoding/json"
	"net/http"

	"know/handlers"

	log "github.com/sirupsen/logrus"
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

	log.Info(employeeData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employeeData)
}
