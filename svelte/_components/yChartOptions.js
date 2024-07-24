const isoDate = require('./formatter').isoDate;

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
          return isoDate( value );
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
          return isoDate( value );
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

function actionCountsOptions(countPerActionsPerDate) {
  return {
    series: countPerActionsPerDate,
    chart: {
      type: 'heatmap',
      height: 600,
    },
    dataLabels: { enabled: false },
    stroke: {
      curve: 'smooth',
      width: 2
    },
    title: {
      text: 'Count Actions Per Date',
      align: 'left',
      style: { fontSize: '14px' },
    },
    xaxis: {
      type: 'datetime',
      axisBorder: { show: false },
      axisTicks: { show: false },
      labels: {
        formatter: function (value) {
          return isoDate( value );
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
          return isoDate( value );
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
    },
    plotOptions: {
      heatmap: {
        shadeIntensity: 1,
        enableShades: false,
        colorScale: {
          inverse: false,
          min: 0,
          max: 9999999,
          ranges: [
            {
              from: 0,
              to: 10,
              color: '#d1fae5',
            },
            {
              from: 11,
              to: 20,
              color: '#a7f3d0',
            },
            {
              from: 21,
              to: 30,
              color: '#6ee7b7'
            },
            {
              from: 31,
              to: 40,
              color: '#34d399'
            },
            {
              from: 41,
              to: 50,
              color: '#10b981'
            },
            {
              from: 51,
              to: 60,
              color: '#059669'
            },
            {
              from: 61,
              to: 70,
              color: '#047857'
            },
            {
              from: 71,
              to: 80,
              color: '#065f46'
            },
            {
              from: 81,
              to: 90,
              color: '#064e3b'
            },
            {
              from: 91,
              to: 100,
              color: '#022c22'
            },
            {
              from: 101,
              to: 9999999,
              color: '#021716',
              name: 'extreme'
            }
          ]
        }
      }
    }
  };
}

module.exports = {
  httpRequestOptions,
  actionCountsOptions
}