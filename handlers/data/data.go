package data

import (
	"io"
	"strconv"
	"os"
	"encoding/json"
	"net/http"

	"know/handlers"

	"github.com/iancoleman/orderedmap"
	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
)

type dataHandler struct{}

//New data handler
func New() handlers.DataHandler {
	return &dataHandler{}
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

	employeeData.Set("employee", []orderedmap.OrderedMap{
		*employee1,
		*employee2,
		*employee3,
		*employee4,
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
		log.Error(err)
	}
	row := sheet.AddRow()
	row.AddCell().Value = "Done"

	err = file.Save("./excel/excel.xlsx")
	if err != nil {
		log.Error(err)
	}

	Openfile, err := os.Open("./excel/excel.xlsx")
	defer Openfile.Close()
	if err != nil {
		http.Error(w, "File not found.", 404)
		return
	}

	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename=excel.xlsx")
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(w, Openfile) //'Copy' the file to the client
	return
}

func (dh *dataHandler) DownloadYesterday(w http.ResponseWriter, r *http.Request) {}
func (dh *dataHandler) DownloadWeek(w http.ResponseWriter, r *http.Request)      {}
func (dh *dataHandler) DownloadMonth(w http.ResponseWriter, r *http.Request)     {}
