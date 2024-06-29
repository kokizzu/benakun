<script>
  import { onMount } from 'svelte';
  import MainLayout from './_layouts/mainLayout.svelte';
  import { httpRequestOptions } from './_components/yChartOptions';

  let registeredUserTotal     = '#registeredUserTotal';
  let registeredUserToday     = '#registeredUserToday';

  let requestsPerDate         = {/* requestsPerDate */};
  let uniqueUserPerDate       = {/* uniqueUserPerDate */};
  let uniqueIpPerDate         = {/* uniqueIpPerDate */};

  let countPerActionsPerDate  = {/* countPerActionsPerDate */};

  console.log('registeredUserTotal', registeredUserTotal);
  console.log('registeredUserToday', registeredUserToday);
  console.log('countPerActionsPerDate', countPerActionsPerDate);

  let CHDATA_REQUESTS_PER_DATE = [];
  let CHDATA_UNIQUE_USER_PER_DATE = [];
  let CHDATA_UNIQUE_IP_PER_DATE = [];
  let CHART_HTTP_REQUESTS = null

  let chartHttpRequestsElm;
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

    loadScript('/assets/apexcharts.js', () => {
      isChartReady = true;
      let chdata = httpRequestOptions(
        CHDATA_REQUESTS_PER_DATE,
        CHDATA_UNIQUE_USER_PER_DATE,
        CHDATA_UNIQUE_IP_PER_DATE
      );
      // @ts-ignore
			CHART_HTTP_REQUESTS = new ApexCharts( chartHttpRequestsElm, chdata);
      CHART_HTTP_REQUESTS.render();
		});
  })
</script>

<MainLayout>
  <div>
    <div class="chart_container">
      <div class="chart" bind:this={chartHttpRequestsElm}></div>
    </div>
  </div>
</MainLayout>

<style>
  .chart_container {
		width: 100%;
		height: 370px;
		border: 1px solid var(--gray-004);
		border-radius: 8px;
		padding: 10px;
    background-color: #FFF;
	}

  .chart_container .chart {
		width: 100%;
		height: 350px;
	}
</style>