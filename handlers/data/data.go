package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"know/handlers"

	"github.com/iancoleman/orderedmap"
	"github.com/logdna/logdna-go/logger"
	"github.com/tealeg/xlsx"
)

type dataHandler struct {
	log *logger.Logger
}

//New data handler
func New(log *logger.Logger) handlers.DataHandler {
	return &dataHandler{
		log: log,
	}
}

func (dh *dataHandler) GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	employeeData := orderedmap.New()

	employee1 := orderedmap.New()
	employee1.Set("Name", "Tiger Nixon")
	employee1.Set("Position", "System Architect")
	employee1.Set("Office", "Edinburgh")
	employee1.Set("Age", "61")
	employee1.Set("Start Date", "2011/04/25")
	employee1.Set("Salary", "$320,800")

	employee2 := orderedmap.New()
	employee2.Set("Name", "Cedric Kelly")
	employee2.Set("Position", "Senior Javascript Developer")
	employee2.Set("Office", "Edinburgh")
	employee2.Set("Age", "22")
	employee2.Set("Start Date", "2012/03/29")
	employee2.Set("Salary", "$433,060")

	employee3 := orderedmap.New()
	employee3.Set("Name", "Ashton Cox")
	employee3.Set("Position", "Junior Technical Author")
	employee3.Set("Office", "San Francisco")
	employee3.Set("Age", "66")
	employee3.Set("Start Date", "2009/01/12")
	employee3.Set("Salary", "$86,000")

	employee4 := orderedmap.New()
	employee4.Set("Name", "Garrett Winters")
	employee4.Set("Position", "Accountant")
	employee4.Set("Office", "Tokyo")
	employee4.Set("Age", "63")
	employee4.Set("Start Date", "2011/07/25")
	employee4.Set("Salary", "$170,750")

	employee5 := orderedmap.New()
	employee5.Set("Name", "Tiger Nixon")
	employee5.Set("Position", "System Architect")
	employee5.Set("Office", "Edinburgh")
	employee5.Set("Age", "61")
	employee5.Set("Start Date", "2011/04/25")
	employee5.Set("Salary", "$320,800")

	employee6 := orderedmap.New()
	employee6.Set("Name", "Cedric Kelly")
	employee6.Set("Position", "Senior Javascript Developer")
	employee6.Set("Office", "Edinburgh")
	employee6.Set("Age", "22")
	employee6.Set("Start Date", "2012/03/29")
	employee6.Set("Salary", "$433,060")

	employee7 := orderedmap.New()
	employee7.Set("Name", "Ashton Cox")
	employee7.Set("Position", "Junior Technical Author")
	employee7.Set("Office", "San Francisco")
	employee7.Set("Age", "66")
	employee7.Set("Start Date", "2009/01/12")
	employee7.Set("Salary", "$86,000")

	employee8 := orderedmap.New()
	employee8.Set("Name", "Garrett Winters")
	employee8.Set("Position", "Accountant")
	employee8.Set("Office", "Tokyo")
	employee8.Set("Age", "63")
	employee8.Set("Start Date", "2011/07/25")
	employee8.Set("Salary", "$170,750")

	employeeData.Set("employee", []orderedmap.OrderedMap{
		*employee1,
		*employee2,
		*employee3,
		*employee4,
		*employee5,
		*employee6,
		*employee7,
		*employee8,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employeeData)
}

func (dh *dataHandler) GetMonthData(w http.ResponseWriter, r *http.Request) {
	monthData := orderedmap.New()

	monthData.Set("January", 30000)
	monthData.Set("February", 20000)
	monthData.Set("March", 45000)
	monthData.Set("April", 27000)
	monthData.Set("May", 15000)
	monthData.Set("June", 35000)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(monthData)
}

func (dh *dataHandler) GetDayData(w http.ResponseWriter, r *http.Request) {
	dayData := orderedmap.New()

	dayData.Set("June 1", 30000)
	dayData.Set("June 2", 20000)
	dayData.Set("June 3", 45000)
	dayData.Set("June 4", 31000)
	dayData.Set("June 5", 15000)
	dayData.Set("June 6", 27000)
	dayData.Set("June 7", 34000)
	dayData.Set("June 8", 45000)
	dayData.Set("June 9", 35000)
	dayData.Set("June 10", 45000)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dayData)
}

func (dh *dataHandler) GetProfitData(w http.ResponseWriter, r *http.Request) {
	profitData := orderedmap.New()

	profitData.Set("January", 30000)
	profitData.Set("February", 20000)
	profitData.Set("March", 45000)
	profitData.Set("April", 27000)
	profitData.Set("May", 15000)
	profitData.Set("June", 35000)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profitData)
}

func (dh *dataHandler) DownloadToday(w http.ResponseWriter, r *http.Request) {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Today")
	if err != nil {
		dh.log.Error(fmt.Sprintf("%s : %s", handlers.ReadUserIP(r), err.Error()))
	}

	dh.log.Info(fmt.Sprintf("%s : download today clicked by user", handlers.ReadUserIP(r)))

	dayData := orderedmap.New()

	dayData.Set("Time", "Profit")
	dayData.Set("10 AM", "30000")
	dayData.Set("11 AM", "20000")
	dayData.Set("12 PM", "45000")
	dayData.Set("1 PM", "31000")
	dayData.Set("2 PM", "15000")
	dayData.Set("3 PM", "27000")
	dayData.Set("4 PM", "34000")
	dayData.Set("5 PM", "45000")
	dayData.Set("6 PM", "35000")

	for _, key := range dayData.Keys() {
		row := sheet.AddRow()
		data, ok := dayData.Get(key)
		if !ok {
			dh.log.Error(fmt.Sprintf("%s : error in getting the data", handlers.ReadUserIP(r)))
		}
		row.AddCell().Value = key
		row.AddCell().Value = data.(string)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=today.xlsx")

	err = file.Write(w)
	if err != nil {
		dh.log.Error(fmt.Sprintf("%s : %s", handlers.ReadUserIP(r), err.Error()))
	}
	return
}

func (dh *dataHandler) DownloadYesterday(w http.ResponseWriter, r *http.Request) {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Yesterday")
	if err != nil {
		dh.log.Error(fmt.Sprintf("%s : %s", handlers.ReadUserIP(r), err.Error()))
	}

	dh.log.Info(fmt.Sprintf("%s : download yesterday clicked by user", handlers.ReadUserIP(r)))

	dayData := orderedmap.New()

	dayData.Set("Time", "Profit")
	dayData.Set("10 AM", "23423")
	dayData.Set("11 AM", "44422")
	dayData.Set("12 PM", "43233")
	dayData.Set("1 PM", "31000")
	dayData.Set("2 PM", "15000")
	dayData.Set("3 PM", "24555")
	dayData.Set("4 PM", "53214")
	dayData.Set("5 PM", "45000")
	dayData.Set("6 PM", "35000")

	for _, key := range dayData.Keys() {
		row := sheet.AddRow()
		data, ok := dayData.Get(key)
		if !ok {
			dh.log.Error(fmt.Sprintf("%s : error in getting the data", handlers.ReadUserIP(r)))
		}
		row.AddCell().Value = key
		row.AddCell().Value = data.(string)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=yesterday.xlsx")

	err = file.Write(w)
	if err != nil {
		dh.log.Error(fmt.Sprintf("%s : %s", handlers.ReadUserIP(r), err.Error()))
	}
	return
}

func (dh *dataHandler) DownloadWeek(w http.ResponseWriter, r *http.Request) {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Week")
	if err != nil {
		dh.log.Error(fmt.Sprintf("%s : %s", handlers.ReadUserIP(r), err.Error()))
	}

	dh.log.Info(fmt.Sprintf("%s : download week clicked by user", handlers.ReadUserIP(r)))

	dayData := orderedmap.New()

	dayData.Set("Day", "Profit")
	dayData.Set("June 1", "30000")
	dayData.Set("June 2", "20000")
	dayData.Set("June 3", "45000")
	dayData.Set("June 4", "31000")
	dayData.Set("June 5", "15000")
	dayData.Set("June 6", "27000")
	dayData.Set("June 7", "34000")

	for _, key := range dayData.Keys() {
		row := sheet.AddRow()
		data, ok := dayData.Get(key)
		if !ok {
			dh.log.Error(fmt.Sprintf("%s : error in getting the data", handlers.ReadUserIP(r)))
		}
		row.AddCell().Value = key
		row.AddCell().Value = data.(string)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=week.xlsx")

	err = file.Write(w)
	if err != nil {
		dh.log.Error(fmt.Sprintf("%s : %s", handlers.ReadUserIP(r), err.Error()))
	}
	return
}

func (dh *dataHandler) DownloadMonth(w http.ResponseWriter, r *http.Request) {

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Month")
	if err != nil {
		dh.log.Error(fmt.Sprintf("%s : %s", handlers.ReadUserIP(r), err.Error()))
	}

	dh.log.Info(fmt.Sprintf("%s : download month clicked by user", handlers.ReadUserIP(r)))

	monthData := orderedmap.New()

	monthData.Set("Month", "Profit")
	monthData.Set("January", "30000")
	monthData.Set("February", "20000")
	monthData.Set("March", "45000")
	monthData.Set("April", "27000")
	monthData.Set("May", "15000")
	monthData.Set("June", "35000")

	for _, key := range monthData.Keys() {
		row := sheet.AddRow()
		data, ok := monthData.Get(key)
		if !ok {
			dh.log.Error(fmt.Sprintf("%s : error in getting the data", handlers.ReadUserIP(r)))
		}
		row.AddCell().Value = key
		row.AddCell().Value = data.(string)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=month.xlsx")

	err = file.Write(w)
	if err != nil {
		dh.log.Error(fmt.Sprintf("%s : %s", handlers.ReadUserIP(r), err.Error()))
	}
	return
}
