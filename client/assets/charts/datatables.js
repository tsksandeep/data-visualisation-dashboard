const fetch = require("node-fetch");

var tableDoc = document.createElement('table')
var tableRow = document.createElement('tr')
//var tableHead = document.createElement('th')
var tableData = document.createElement('td')

tableDoc.className = "table table-borderless table-dark custom-table"

function getResponse(url) {
  return fetch(url, {
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    }
  }).then((response) => response.json())
};

function getAllResponse() {
  return Promise.all([getResponse("http://127.0.0.1:3000/data/employee")])
}


// table generator function which generates the table
function generateTable() {
  var table = tableDoc.cloneNode(false)
  var employeeInfo = getAllResponse().then(([data]) => {
    return data;
  })

  for (var i = 0; i < employeeInfo["employee"].length; i++) {
    var tr = tableRow.cloneNode(false)
    for (var key in employeeInfo["employee"][i]) {
      var td = tableData.cloneNode(false)
      td.appendChild(document.createTextNode(employeeInfo["employee"][i][key]))
      tr.appendChild(td)
    }
    table.appendChild(tr)
  }
  return table
}

document.getElementById("dataTable").appendChild(generateTable())