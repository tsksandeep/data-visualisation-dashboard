const axios = require('axios');

var tableDoc = document.createElement('table')
var tableRow = document.createElement('tr')
var tableHead = document.createElement('th')
var tableData = document.createElement('td')

tableDoc.className = "table table-borderless table-dark custom-table"

// table generator function which generates the table
function generateTable() {
  axios.get("/data/employee")
    .then(response => {
      var isHeaderAppended = false
      var table = tableDoc.cloneNode(false)

      let employeeInfo = response.data

      for (var i = 0; i < employeeInfo["employee"].length; i++) {
        if (!isHeaderAppended) {
          var tr = tableRow.cloneNode(false)
          for (var key in employeeInfo["employee"][i]) {
            var th = tableHead.cloneNode(false)

            th.appendChild(document.createTextNode(key))
            tr.appendChild(th)
          }
          table.appendChild(tr)
          isHeaderAppended = true
        }

        var tr = tableRow.cloneNode(false)
        for (var key in employeeInfo["employee"][i]) {
          var td = tableData.cloneNode(false)

          td.appendChild(document.createTextNode(employeeInfo["employee"][i][key]))
          tr.appendChild(td)
        }

        table.appendChild(tr)
      }

      document.getElementById("dataTable").appendChild(table)
    })
    .catch(error => {
      console.log(error);
    });
}

generateTable()