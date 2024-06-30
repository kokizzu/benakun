<script>
  import { onMount } from 'svelte';
  import MainLayout from './_layouts/mainLayout.svelte';
  import { httpRequestOptions, actionCountsOptions } from './_components/yChartOptions';

  let registeredUserTotal     = '#registeredUserTotal';
  let registeredUserToday     = '#registeredUserToday';

  console.log(registeredUserTotal, registeredUserToday);

  let requestsPerDate         = {/* requestsPerDate */};
  let uniqueUserPerDate       = {/* uniqueUserPerDate */};
  let uniqueIpPerDate         = {/* uniqueIpPerDate */};

  let countPerActionsPerDate  = {/* countPerActionsPerDate */};

  let CHDATA_REQUESTS_PER_DATE = [];
  let CHDATA_UNIQUE_USER_PER_DATE = [];
  let CHDATA_UNIQUE_IP_PER_DATE = [];

  /** @typedef {{x: number, y: number}} chartData */

  /** @typedef {Object} heatmap
    * @property {string} name
    * @property {chartData[]} data
    */
  let CHDATA_ACTION_COUNTS = /** @type {heatmap[]} */ ([]);

  let CHART_HTTP_REQUESTS = null;
  let CHART_ACTION_COUNTS = null;

  let chartHttpRequestsElm;
  let chartActionCountsElm;
  let isChartReady = false;

  function loadScript(/** @type {string} */ url, /** @type {Function} */ callback) {
    let script = /** @type {HTMLScriptElement} */ (document.createElement("script"));
    script.type = "text/javascript";
    script.src = url;
    script.onload = () => callback();
    document.getElementsByTagName("head")[0].appendChild(script);
  }

  onMount(() => {
    for (let key in requestsPerDate) {
      CHDATA_REQUESTS_PER_DATE = [...CHDATA_REQUESTS_PER_DATE, {x: (new Date(key).getTime()) / 1000, y: requestsPerDate[key]}];
    }

    for (let key in uniqueUserPerDate) {
      CHDATA_UNIQUE_USER_PER_DATE = [...CHDATA_UNIQUE_USER_PER_DATE, {x: (new Date(key).getTime()) / 1000, y: uniqueUserPerDate[key]}];
    }

    for (let key in uniqueIpPerDate) {
      CHDATA_UNIQUE_IP_PER_DATE = [...CHDATA_UNIQUE_IP_PER_DATE, {x: (new Date(key).getTime()) / 1000, y: uniqueIpPerDate[key]}];
    }

    for (let key in countPerActionsPerDate) {
      let data = /** @type {chartData[]} */ ([]);
      for (let key2 in countPerActionsPerDate[key]) {
        data = [...data, {x: (new Date(key2).getTime()) / 1000, y: countPerActionsPerDate[key][key2]}];
      }
      CHDATA_ACTION_COUNTS = [ ...CHDATA_ACTION_COUNTS, { name: key, data: data } ];
    }

    loadScript('/assets/apexcharts.js', () => {
      isChartReady = true;
      let chdata = httpRequestOptions(
        CHDATA_REQUESTS_PER_DATE,
        CHDATA_UNIQUE_USER_PER_DATE,
        CHDATA_UNIQUE_IP_PER_DATE
      );
      let chdata2 = actionCountsOptions(CHDATA_ACTION_COUNTS);
      // @ts-ignore
			CHART_HTTP_REQUESTS = new ApexCharts( chartHttpRequestsElm, chdata);
      CHART_HTTP_REQUESTS.render();

      // @ts-ignore
      CHART_ACTION_COUNTS = new ApexCharts( chartActionCountsElm, chdata2);
      CHART_ACTION_COUNTS.render();
		});
  })
</script>

<MainLayout>
  <div class="superadmin_dashboard">
    <div class="chart_http_req">
      <div class="chart" bind:this={chartHttpRequestsElm}></div>
    </div>
    <div class="chart_action_counts">
      <div class="chart" bind:this={chartActionCountsElm}></div>
    </div>
  </div>
</MainLayout>

<style>
  .superadmin_dashboard {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  .chart_http_req {
		width: 100%;
		height: 370px;
		border: 1px solid var(--gray-004);
		border-radius: 8px;
		padding: 10px;
    background-color: #FFF;
	}

  .chart_http_req .chart {
		width: 100%;
		height: 350px;
	}

  .chart_action_counts {
    width: 100%;
    height: 630px;
    border: 1px solid var(--gray-004);
    border-radius: 8px;
    padding: 10px;
    background-color: #FFF;
  }

  .chart_action_counts .chart {
    width: 100%;
    height: 610px;
  }
</style>