<script>
  /** @typedef {import('./_components/types/transaction').TransactionJournal} TransactionJournal */

  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import { ReportViewerTrialBalance } from './jsApi.GEN';
  import { notifier } from './_components/notifier';
  import { dateISOFormat, loadScript } from './_components/formatter';
    import { trialBalanceOptions } from './_components/yChartOptions';

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
          notifier.showInfo('No data found at ' + startDate);
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
    'Weekly', 'Monthly', 'Quarterly', 'Yearly',
  ]

  let toFilter = filterDatesOptions[0] || 'Weekly';

  $: {
    if (toFilter) {
      switch (toFilter) {
        case 'Weekly':
          startDate = dateISOFormat(-7);
          endDate = dateISOFormat(0);
          break;
        case 'Monthly':
          startDate = dateISOFormat(-30);
          endDate = dateISOFormat(0);
          break;
        case 'Quarterly':
          startDate = dateISOFormat(-90);
          endDate = dateISOFormat(0);
          break;
        case 'Yearly':
          startDate = dateISOFormat(-365);
          endDate = dateISOFormat(0);
          break;
      }
    }
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
  })
</script>

<MainLayout>
  <div>
    <div class="chart_and_filter">
      <div class="chart_container">
        <div class="chart" bind:this={accountActivityElm}></div>
      </div>
      <div class="date_filter">
        <InputBox
          id="dateFilter"
          label="Filter Dates"
          bind:value={toFilter}
          type="combobox"
          values={filterDatesOptions}
          isObject={false}
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
</style>