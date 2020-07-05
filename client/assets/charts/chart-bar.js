const axios = require('axios');
const chart = require('chart.js');

Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#292b2c';

function generateChart() {
  axios.get("/data/chart/month")
    .then(response => {
      var ctx = document.getElementById("myBarChart");

      let monthData = response.data
      let labels = []
      let data = []

      for (var key in monthData) {
        labels.push(key)
        data.push(monthData[key])
      }

      new chart(ctx, {
        type: 'bar',
        data: {
          labels: labels,
          datasets: [{
            label: "Revenue",
            backgroundColor: "rgba(2,117,216,1)",
            borderColor: "rgba(2,117,216,1)",
            data: data,
          }],
        },
        options: {
          scales: {
            xAxes: [{
              time: {
                unit: 'month'
              },
              gridLines: {
                display: false
              },
              ticks: {
                maxTicksLimit: 6
              }
            }],
            yAxes: [{
              ticks: {
                min: 0,
                max: 100000,
                maxTicksLimit: 5
              },
              gridLines: {
                display: true
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

generateChart()
