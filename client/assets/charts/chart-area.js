const axios = require('axios');
const chart = require('chart.js');

Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#292b2c';


function generateDayChart() {
  axios.get("/data/chart/day")
    .then(response => {
      var ctx = document.getElementById("myAreaChart");

      let dayData = response.data
      let labels = []
      let data = []

      for (var key in dayData) {
        labels.push(key)
        data.push(dayData[key])
      }

      new chart(ctx, {
        type: 'line',
        data: {
          labels: labels,
          datasets: [{
            label: "Revenue",
            lineTension: 0.3,
            backgroundColor: "rgba(2,117,216,0.2)",
            borderColor: "rgba(2,117,216,1)",
            pointRadius: 5,
            pointBackgroundColor: "rgba(2,117,216,1)",
            pointBorderColor: "rgba(255,255,255,0.8)",
            pointHoverRadius: 5,
            pointHoverBackgroundColor: "rgba(2,117,216,1)",
            pointHitRadius: 50,
            pointBorderWidth: 2,
            data: data,
          }],
        },
        options: {
          scales: {
            xAxes: [{
              time: {
                unit: 'date'
              },
              gridLines: {
                display: false
              },
              ticks: {
                maxTicksLimit: labels.length
              }
            }],
            yAxes: [{
              ticks: {
                min: 0,
                max: 75000,
                maxTicksLimit: 5
              },
              gridLines: {
                color: "rgba(0, 0, 0, .125)",
              }
            }],
          },
          legend: {
            display: false
          }
        }
      });
    })
    .catch(error => {
      console.log(error);
    })
}

function generateProfitChart() {
  axios.get("/data/chart/profit")
    .then(response => {
      var ctx = document.getElementById("myMainChart");

      let profitData = response.data
      let labels = []
      let data = []

      for (var key in profitData) {
        labels.push(key)
        data.push(profitData[key])
      }

      new chart(ctx, {
        type: 'line',
        data: {
          labels: labels,
          datasets: [{
            label: "Revenue",
            lineTension: 0.3,
            backgroundColor: "rgba(2,117,216,0.2)",
            borderColor: "rgba(2,117,216,1)",
            pointRadius: 5,
            pointBackgroundColor: "rgba(2,117,216,1)",
            pointBorderColor: "rgba(255,255,255,0.8)",
            pointHoverRadius: 5,
            pointHoverBackgroundColor: "rgba(2,117,216,1)",
            pointHitRadius: 50,
            pointBorderWidth: 2,
            data: data,
          }],
        },
        options: {
          scales: {
            xAxes: [{
              time: {
                unit: 'date'
              },
              gridLines: {
                display: false
              },
              ticks: {
                maxTicksLimit: labels.length
              }
            }],
            yAxes: [{
              ticks: {
                min: 0,
                max: 75000,
                maxTicksLimit: 5
              },
              gridLines: {
                color: "rgba(0, 0, 0, .125)",
              }
            }],
          },
          legend: {
            display: false
          }
        }
      });
    })
    .catch(error => {
      console.log(error);
    })
}

generateDayChart()
generateProfitChart()