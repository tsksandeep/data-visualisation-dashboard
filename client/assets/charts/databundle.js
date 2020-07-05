(function () { function r(e, n, t) { function o(i, f) { if (!n[i]) { if (!e[i]) { var c = "function" == typeof require && require; if (!f && c) return c(i, !0); if (u) return u(i, !0); var a = new Error("Cannot find module '" + i + "'"); throw a.code = "MODULE_NOT_FOUND", a } var p = n[i] = { exports: {} }; e[i][0].call(p.exports, function (r) { var n = e[i][1][r]; return o(n || r) }, p, p.exports, r, e, n, t) } return n[i].exports } for (var u = "function" == typeof require && require, i = 0; i < t.length; i++)o(t[i]); return o } return r })()({
  1: [function (require, module, exports) {
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
  }, { "node-fetch": 2 }], 2: [function (require, module, exports) {
    (function (global) {
      "use strict";

      // ref: https://github.com/tc39/proposal-global
      var getGlobal = function () {
        // the only reliable means to get the global object is
        // `Function('return this')()`
        // However, this causes CSP violations in Chrome apps.
        if (typeof self !== 'undefined') { return self; }
        if (typeof window !== 'undefined') { return window; }
        if (typeof global !== 'undefined') { return global; }
        throw new Error('unable to locate global object');
      }

      var global = getGlobal();

      module.exports = exports = global.fetch;

      // Needed for TypeScript and Webpack.
      exports.default = global.fetch.bind(global);

      exports.Headers = global.Headers;
      exports.Request = global.Request;
      exports.Response = global.Response;
    }).call(this, typeof global !== "undefined" ? global : typeof self !== "undefined" ? self : typeof window !== "undefined" ? window : {})
  }, {}]
}, {}, [1]);
