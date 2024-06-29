const utcDatetime = require('./formatter').utcDatetime;

function httpRequestOptions(requestsPerDate, uniqueUserPerDate, uniqueIpPerDate) {
  return {
    series: [
      { name: 'Request Per Date', data: requestsPerDate },
      { name: 'Unique User Per Date', data: uniqueUserPerDate },
      { name: 'Unique IP Per Date', data: uniqueIpPerDate },
    ],
    chart: {
      type: 'area',
      height: 340,
    },
    dataLabels: { enabled: false },
    stroke: {
      curve: 'smooth',
      width: 2
    },
    title: {
      text: 'HTTP Requests',
      align: 'left',
      style: { fontSize: '14px' },
    },
    xaxis: {
      type: 'datetime',
      axisBorder: { show: false },
      axisTicks: { show: false },
      labels: {
        formatter: function (value) {
          return utcDatetime( value );
        }
      }
    },
    yaxis: {
      tickAmount: 4,
      floating: false,
      labels: { offsetY: -7, offsetX: 0 },
      axisBorder: { show: false },
      axisTicks: { show: false },
    },
    fill: { opacity: 0.5 },
    tooltip: {
      x: {
        formatter: function (value) {
          return utcDatetime( value );
        },
      },
      fixed: {
        enabled: false,
        position: 'topRight',
      },
    },
    grid: {
      yaxis: {
        lines: {
          offsetX: -30,
        },
      },
      padding: {
        left: 20,
      },
    },
    legend: {
      show: true,
      showForNullSeries: true,
      showForZeroSeries: true,
      position: 'bottom',
      horizontalAlign: 'center', 
      floating: false,
      onItemClick: {
        toggleDataSeries: true
      },
      onItemHover: {
        highlightDataSeries: true
      },
    }
  };
}

module.exports = {
  httpRequestOptions
}