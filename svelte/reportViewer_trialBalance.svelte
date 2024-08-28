<script>
  /** @typedef {import('./_components/types/transaction').TransactionJournal} TransactionJournal */

  import MainLayout from './_layouts/mainLayout.svelte';
  import { onMount } from 'svelte';
  import InputBox from './_components/InputBox.svelte';
  import SubmitButton from './_components/SubmitButton.svelte';
  import { ReportViewerTrialBalance } from './jsApi.GEN';
    import { notifier } from './_components/notifier';

  const coaChoices = /** @type Record<string|number, string> */ ({/* coaChoices */});
  let transactionJournals = /** @type TransactionJournal[] */ ([/* transactionJournals */]);

  let DebitIDRTotal = 0;
  let CreditIDRTotal = 0;

  onMount(() => {
    if (transactionJournals && transactionJournals.length > 0) {
      transactionJournals.forEach((trxJournal) => {
        DebitIDRTotal += Number(trxJournal.debitIDR);
        CreditIDRTotal += Number(trxJournal.creditIDR);
      });
    }
  });

  function dateISOFormat() {
    const dt    = new Date();
    const date  = dt.getDate();
    const month = dt.toLocaleDateString('default', {month: '2-digit'});
    const year  = dt.getFullYear();

    return `${year}-${month}-${date}`;
  }

  let dateToFilter = dateISOFormat();
  let isSubmitted = false;

  async function getTrxJournalsByDate() {
    isSubmitted = true;
    const i = {
      transactionJournal: {
        date: dateToFilter
      }
    }

    await ReportViewerTrialBalance( // @ts-ignore
      i, /** @type {import('./jsApi.GEN').ReportViewerTrialBalanceCallback} */
      function(/** @type any */ o) {
        isSubmitted = false;
        if (o.error) {
          console.log(o);
          notifier.showError(o.error);
          return
        }

        console.log('Result', o);

        transactionJournals = o.transactionJournals || [];
        if (transactionJournals.length == 0) {
          notifier.showInfo('No data found at ' + dateToFilter);
        }

        if (transactionJournals && transactionJournals.length > 0) {
          transactionJournals.forEach((trxJournal) => {
            DebitIDRTotal += Number(trxJournal.debitIDR);
            CreditIDRTotal += Number(trxJournal.creditIDR);
          });
        }
      }
    );
  }
</script>

<MainLayout>
  <div>
    <div class="date_filter">
      <InputBox
        id="dateToFilter"
        label="Date"
        bind:value={dateToFilter}
        type="date"
      />
      <SubmitButton
        label="Submit"
        bind:isSubmitted
        on:click={getTrxJournalsByDate}
        isFullWidth={false}
      />
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
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</MainLayout>

<style>
  .date_filter {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: flex-end;
    gap: 12px;
    margin-bottom: 20px;
  }

  .table_root {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 10px;
    padding: 0;
    border: 1px solid var(--gray-003);
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