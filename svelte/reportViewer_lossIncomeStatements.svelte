<script>
  /** @typedef {import('./_components/types/transaction').TransactionJournal} TransactionJournal */

  import { onMount } from 'svelte';
  import MainLayout from './_layouts/mainLayout.svelte';

  let transactionJournals = /** @type {TransactionJournal[]} */ ([ /* transactionJournals */ ]);
  const coaChoices = /** @type Record<string|number, string> */ ({/* coaChoices */});

  let formattedTrxJournals = /** @type TransactionJournal[] */ ([]);
  let DebitIDRTotal = 0;
  let CreditIDRTotal = 0;

  onMount(() => {
    if (transactionJournals && transactionJournals.length > 0) {
      let objTrxJournals = /** @type Record<string, TransactionJournal> */ ({})
      transactionJournals.forEach((trxJournal) => {
        DebitIDRTotal += Number(trxJournal.debitIDR);
        CreditIDRTotal += Number(trxJournal.creditIDR);
        if (objTrxJournals[trxJournal.coaId]) {
          objTrxJournals[trxJournal.coaId].debitIDR = Number(trxJournal.debitIDR) + Number(objTrxJournals[trxJournal.coaId].debitIDR);
          objTrxJournals[trxJournal.coaId].creditIDR = Number(trxJournal.creditIDR) + Number(objTrxJournals[trxJournal.coaId].creditIDR);
        } else {
          objTrxJournals[trxJournal.coaId] = trxJournal
        }
      });

      formattedTrxJournals = Object.values(objTrxJournals);
      transactionJournals = []; // prevent memory leak
    }
  })
  
</script>

<MainLayout>
  <div>
    <div class="table_root">
      <div class="table_container">
        <table>
          <thead>
            <tr>
              <th class="no">No</th>
              <th>Name / Nama</th>
              <th>Income / Laba</th>
              <th>Loss / Rugi</th>
            </tr>
          </thead>
          <tbody>
            {#if formattedTrxJournals && formattedTrxJournals.length > 0}
              {#each formattedTrxJournals as trxJournal, idx (trxJournal.id)}
                <tr>
                  <td class="no">{idx + 1}</td>
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
