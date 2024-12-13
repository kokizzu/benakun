<script>
  /** @typedef {import('./_components/types/transaction').TransactionJournal} TransactionJournal */

  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import { ReportViewerTrialBalance } from './jsApi.GEN';
  import { notifier } from './_components/xNotifier.js';
  import { dateISOFormat, loadScript } from './_components/xformatter.js';
  import { trialBalanceOptions } from './_components/yChartOptions';
  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiArrowsArrowRightSLine, RiArrowsArrowLeftSLine } from './node_modules/svelte-icons-pack/dist/ri';
  import Select from './node_modules/svelte-select';

  const coaChoices = /** @type Record<string|number, string> */ ({/* coaChoices */});
  let transactionJournals = /** @type TransactionJournal[] */ ([/* transactionJournals */]);

  let DebitIDRTotal = 0;
  let CreditIDRTotal = 0;

  let startDate = dateISOFormat(-7);
  let endDate = dateISOFormat(0);

  let CHART_ACCOUNT_ACTIVITY = null;

  let accountActivityElm;
  let isChartReady = false;

  onMount(() => {
    if (transactionJournals && transactionJournals.length > 0) {
      transactionJournals.forEach((trxJournal) => {
        DebitIDRTotal += Number(trxJournal.debitIDR);
        CreditIDRTotal += Number(trxJournal.creditIDR);
      });
    }
  });

  let isSubmitted = false;

  async function getTrxJournalsByDate() {
    isSubmitted = true;
    const i = {startDate, endDate};

    await ReportViewerTrialBalance( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').ReportViewerTrialBalanceCallback} */
      function(/** @type any */ o) {
        isSubmitted = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        transactionJournals = o.transactionJournals || [];
        if (transactionJournals.length == 0) {
          notifier.showInfo('No data found at "' + startDate+'" to "'+endDate+'"');
        }

        DebitIDRTotal = 0;
        CreditIDRTotal = 0;

        if (transactionJournals && transactionJournals.length > 0) {
          transactionJournals.forEach((trxJournal) => {
            DebitIDRTotal += Number(trxJournal.debitIDR);
            CreditIDRTotal += Number(trxJournal.creditIDR);
          });
        }

        const chData = getChartData();
        CHART_ACCOUNT_ACTIVITY.updateSeries([
          { name: 'Transaction Activity', data: chData }
        ])
      }
    );
  }

  const filterDatesOptions = [
    {value: 'weekly', label: 'Weekly'},
    {value: 'monthly', label: 'Monthly'},
    {value: 'quarterly', label: 'Quarterly'},
    {value: 'yearly', label: 'Yearly'},
  ]

  let toFilter = {value: 'weekly', label: 'Weekly'};

  async function HandleSelectFilter(e) {
    switch (e.detail.value) {
      case 'weekly':
        startDate = dateISOFormat(-7);
        endDate = dateISOFormat(0);
        break;
      case 'monthly':
        startDate = dateISOFormat(-30);
        endDate = dateISOFormat(0);
        break;
      case 'quarterly':
        startDate = dateISOFormat(-90);
        endDate = dateISOFormat(0);
        break;
      case 'yearly':
        startDate = dateISOFormat(-365);
        endDate = dateISOFormat(0);
        break;
    }

    await getTrxJournalsByDate();
  }

  /** @returns {Array<Record<string, any>>} */
  function getChartData() {
    let dateMaps = {};
    let chData = [];
    if (transactionJournals && transactionJournals.length > 0) {
      for (const trxJournal of transactionJournals) {
        if (dateMaps[trxJournal.date]) {
          dateMaps[trxJournal.date] += 1;
        } else {
          dateMaps[trxJournal.date] = 1;
        }
      }

      for (let key in dateMaps) {
        chData = [...chData, { x: (new Date(key).getTime()) / 1000, y: dateMaps[key] }];
      }
    }

    return chData;
  }

  onMount(() => {
    loadScript('/assets/apexcharts.js', () => {
      isChartReady = true;
      const chData = getChartData();
      console.log(chData);
      const dataOption = trialBalanceOptions(chData)
      // @ts-ignore
      CHART_ACCOUNT_ACTIVITY = new ApexCharts( accountActivityElm, dataOption);
      CHART_ACCOUNT_ACTIVITY.render();
		});
  });

  function getDateBy(dateISOstr, dayTo) {
    const dt = new Date(dateISOstr);
    dt.setDate(dt.getDate() + dayTo);

    const date = String(dt.getDate()).padStart(2, '0');
    const month = String(dt.getMonth() + 1).padStart(2, '0');
    const year = dt.getFullYear();

    return `${year}-${month}-${date}`;
  }

  async function PrevDate() {
    switch (toFilter.value) {
      case 'weekly':
        startDate = endDate;
        endDate = getDateBy(endDate, -7);
        break;
      case 'monthly':
        startDate = endDate;
        endDate = getDateBy(endDate, -30);
        break;
      case 'quarterly':
        startDate = endDate;
        endDate = getDateBy(endDate, -90);
        break;
      case 'yearly':
        startDate = endDate;
        endDate = getDateBy(endDate, -365);
        break;
    }

    await getTrxJournalsByDate();
  }

  async function NextDate() {
    switch (toFilter.value) {
      case 'weekly':
        startDate = endDate;
        endDate = getDateBy(endDate, 7);
        break;
      case 'monthly':
        startDate = endDate;
        endDate = getDateBy(endDate, 30);
        break;
      case 'quarterly':
        startDate = endDate;
        endDate = getDateBy(endDate, 90);
        break;
      case 'yearly':
        startDate = endDate;
        endDate = getDateBy(endDate, 365);
        break;
    }

    await getTrxJournalsByDate();
  }
</script>

<MainLayout>
  <div>
    <div class="chart_and_filter">
      <div class="chart_container">
        <div class="chart" bind:this={accountActivityElm}></div>
      </div>
      <div class="date_filter">
        <Select
          items={filterDatesOptions}
          bind:value={toFilter}
          on:select={HandleSelectFilter}
        />
        <InputBox
          id="startDate"
          label="Start Date"
          bind:value={startDate}
          type="date"
        />
        <InputBox
          id="endDate"
          label="End Date"
          bind:value={endDate}
          type="date"
        />
        <div class="prev_next">
          <button class="btn" on:click={PrevDate}>
            <Icon
              src={RiArrowsArrowLeftSLine}
              size="20"
            />
            <span>Prev</span>
          </button>
          <button class="btn" on:click={NextDate}>
            <span>Next</span>
            <Icon
              src={RiArrowsArrowRightSLine}
              size="20"
            />
          </button>
        </div>
        <SubmitButton
          label="Submit"
          bind:isSubmitted
          on:click={getTrxJournalsByDate}
          isFullWidth
        />
      </div>
    </div>
    <div class="table_root">
      <div class="table_container">
        <table>
          <thead>
            <tr>
              <th class="no">No. Acc / No. Rek</th>
              <th>Name / Nama</th>
              <th>Debit / Debet</th>
              <th>Credit / Kredit</th>
              <th>Date / Tanggal</th>
            </tr>
          </thead>
          <tbody>
            {#if transactionJournals && transactionJournals.length > 0}
              {#each transactionJournals as trxJournal, idx (trxJournal.id)}
                <tr>
                  <td class="no">{trxJournal.coaId}</td>
                  <td>{coaChoices[trxJournal.coaId]}</td>
                  <td>{trxJournal.debitIDR}</td>
                  <td>{trxJournal.creditIDR}</td>
                  <td>{trxJournal.date}</td>
                </tr>
              {/each}
            {:else}
              <tr>
                <td colspan="4">No data</td>
              </tr>
            {/if}
            <tr>
              <td colspan="2">
                <b>Total</b>
              </td>
              <td>
                <b>{DebitIDRTotal}</b>
              </td>
              <td>
                <b>{CreditIDRTotal}</b>
              </td>
              <td></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</MainLayout>

<style>
  .chart_and_filter {
    display: flex;
    flex-direction: row;
    gap: 20px;
    margin-bottom: 20px;
    align-items: start;
    flex-grow: 1;
    max-width: 100%;
  }
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
  .date_filter {
    width: 30%;
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .date_filter .prev_next {
    display: flex;
    flex-direction: row;
    gap: 10px;
    width: 100%;
  }

  .date_filter .prev_next .btn {
    width: 100%;
    border: none;
    padding: 7px 15px;
    border-radius: 8px;
    background-color: var(--blue-006);
    color: #fff;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    gap: 5px;
    cursor: pointer;
    text-align: center;
    font-weight: 700;
  }

  .date_filter .prev_next .btn:hover {
    background-color: var(--blue-005);
  }

  .table_root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 10px;
    padding: 0;
    border: 1px solid var(--gray-004);
    overflow: hidden;
  }

 .table_root .table_container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .table_root .table_container table {
    width: 100%;
    background: #fff;
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
  }

  .table_root .table_container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .table_root .table_container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .table_root .table_container table thead tr th.no {
    width: 30px;
  }

  .table_root .table_container table thead tr th:last-child {
    border-right: none;
  }

  .table_root .table_container table tbody tr td {
    padding: 8px 12px;
  }

	.table_root .table_container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.table_root .table_container table tbody tr:last-child td,
	.table_root .table_container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .table_root .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table_root .table_container table tbody tr:last-child td,
  .table_root .table_container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .table_root .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .table_root .table_container table tbody tr td:last-child {
    border-right: none !important;
  }

  .table_root .table_container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }

  @media only screen and (max-width : 768px) {
    .chart_and_filter {
      flex-direction: column;
    }

    .date_filter {
      width: 100%;
    }
  }
</style>
