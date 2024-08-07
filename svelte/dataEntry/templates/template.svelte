<script>
  /** @typedef {import('../../_components/types/transaction.js').TransactionTemplate} TransactionTemplate */

  import MainLayout from '../../_layouts/mainLayout.svelte';
  import { Icon } from '../../node_modules/svelte-icons-pack/dist';
  import {
    RiSystemAddBoxLine
  } from '../../node_modules/svelte-icons-pack/dist/ri';
  import PopUpDataEntryJournal from '../../_components/PopUpDataEntryJournal.svelte';
  import { DataEntryTransactionEntry } from '../../jsApi.GEN';
    import { notifier } from '../../_components/notifier';

  let transactionTemplate = /** @type TransactionTemplate */ ({/* transactiontemplate */});
  let transactionTplDetails = /** @type any[] */ ([/* transactionTplDetails */]);
  let coas = /** @type any[] */ ([/* coas */]);

  let popUpDataEntryJournal;

  // jika attribut coa adalah childOnly: maka coa pilihan hanya anaknya

  // Nota bagian atas dan bawah adalah business transaction
  // Nota bagian tengah adalah jurnal

  console.log('transactionTemplate', transactionTemplate);
  console.log('transactionTplDetails', transactionTplDetails);

  let isDebit = true;
  let isSales = false;

  let startDate;
  let endDate;

  let coaId = 0;
  let debitIDR = 0;
  let creditIDR = 0;
  let description = '';
  let date = '';
  let heading = 'Add journal for ' + transactionTemplate.name;

  let isSubmitted = false;

  async function Submit() {
    isSubmitted = true;
    const i = {
      coaId: coaId,
      transactionJournals: [
        {
          debitIDR: debitIDR,
          creditIDR: creditIDR,
          descriptions: description,
          date: date,
        }
      ],
      businessTransaction: {
        startDate: startDate,
        endDate: endDate,
      }
    }

    await DataEntryTransactionEntry( // @ts-ignore
      i, /** @returns {Promise<any>} */
      function (o) {
        isSubmitted = false;
        if (o.error) {
          notifier.showError(o.error || 'failed to add journal');
          return
        }

        notifier.showSuccess('journal added !');
        popUpDataEntryJournal.Hide()
      }
    )
  }

  async function SubmitWithSales(payloadsSales) {
    console.log('Submit sales', payloadsSales);
  }
</script>

<PopUpDataEntryJournal
  bind:this={popUpDataEntryJournal}
  bind:isSubmitted
  heading={heading}

  bind:startDate
  bind:endDate

  bind:isDebit
  bind:isSales
  bind:debitIDR
  bind:creditIDR
  bind:description
  bind:date

  OnSubmit={Submit}
  OnSubmitWithSales={SubmitWithSales}
/>

<MainLayout>
  <div>
    <h1>Data entry for {transactionTemplate.name}</h1>
  </div>
  <div class="data_entry_template___container">
    <div class="transaction_template_detail">
      <div class="table_container">
        <table>
          <thead>
            <tr>
              <th class="no">#</th>
              <th class="a_row">Action</th>
              <th>Kredit</th>
              <th>Debit</th>
              <th>CoA (Chart of Accounts)</th>
              <th>Auto Sum</th>
              <th>Child Only</th>
              <th>Sales</th>
            </tr>
          </thead>
          <tbody>
            {#if transactionTplDetails && transactionTplDetails.length > 0}
              {#each (transactionTplDetails || []) as trxTplDetail, i (trxTplDetail.id)}
                <tr>
                  <td>{i + 1}</td>
                  <td class="a_row">
                    <div class="actions">
                      <button
                        on:click={() => {
                          isDebit = trxTplDetail.isDebit
                          isSales = (trxTplDetail.attributes || []).includes('sales')
                          coaId = trxTplDetail.coaId
                          if (isSales) {
                            popUpDataEntryJournal.ShowWithSales()
                          } else {
                            popUpDataEntryJournal.Show()
                          }
                        }}
                        class="btn"
                        title="Add journal"
                      >
                        <Icon
                          size="15"
                          color="var(--gray-007)"
                          src={RiSystemAddBoxLine}
                        />
                      </button>
                    </div>
                  </td>
                  <td>{!trxTplDetail.isDebit ? 'Yes' : 'No'}</td>
                  <td>{trxTplDetail.isDebit ? 'Yes' : 'No'}</td>
                  <td>{coas[trxTplDetail.coaId]}</td>
                  <td>{(trxTplDetail.attributes || []).includes('autoSum') ? 'Yes' : 'No'}</td>
                  <td>{(trxTplDetail.attributes || []).includes('childOnly') ? 'Yes' : 'No'}</td>
                  <td>{(trxTplDetail.attributes || []).includes('sales') ? 'Yes' : 'No'}</td>
                </tr>
              {/each}
            {:else}
              <tr>
                <td>0</td>
                <td>no-data</td>
              </tr>
            {/if}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</MainLayout>

<style>
  .data_entry_template___container {
    display: flex;
  }

  .transaction_template_detail {
    display: flex;
    flex-direction: column;
    background-color: #fff;
    border-radius: 10px;
    padding: 0;
    border: 1px solid var(--gray-003);
    overflow: hidden;
    width: 100%;
  }

  .transaction_template_detail .table_container {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .transaction_template_detail .table_container table {
    width: 100%;
    background: #fff;
    box-shadow: none;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0;
    overflow: hidden;
  }

  .transaction_template_detail .table_container table thead {
    box-shadow: none;
    border-bottom: 1px solid var(--gray-003);
  }

  .transaction_template_detail .table_container table thead tr th {
    padding: 12px;
		background-color: var(--gray-001);
		text-transform: capitalize;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-003);
		min-width: fit-content;
		width: auto;
    text-wrap: nowrap;
  }

  .transaction_template_detail .table_container table thead tr th.no {
    width: 30px;
  }

  .transaction_template_detail .table_container table thead tr th.a_row {
    max-width: 100px;
    width: 100px;
  }

  .transaction_template_detail .table_container table thead tr th:last-child {
    border-right: none;
  }

  .transaction_template_detail .table_container table tbody tr td {
    padding: 8px 12px;
  }

	.transaction_template_detail .table_container table tbody tr td {
    padding: 8px 12px;
		border-right: 1px solid var(--gray-004);
		border-bottom: 1px solid var(--gray-004);
  }

	.transaction_template_detail .table_container table tbody tr:last-child td,
	.transaction_template_detail .table_container table tbody tr:last-child th {
		border-bottom: none !important;
	}

  .transaction_template_detail .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .transaction_template_detail .table_container table tbody tr:last-child td,
  .transaction_template_detail .table_container table tbody tr:last-child th {
    border-bottom: none !important;
  }

  .transaction_template_detail .table_container table tbody tr:last-child td:last-child {
    border-right: none !important;
  }

  .transaction_template_detail .table_container table tbody tr td:last-child {
    border-right: none !important;
  }

  .transaction_template_detail .table_container table tbody tr th {
    text-align: center;
    border-right: 1px solid var(--gray-004);
    border-bottom: 1px solid var(--gray-004);
  }

  .transaction_template_detail .table_container table tbody tr td .actions {
    display: flex;
    flex-direction: row;
  }

  .transaction_template_detail .table_container table tbody tr td .actions .btn {
    border: none;
    padding: 6px;
    border-radius: 8px;
    background-color: transparent;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .transaction_template_detail .table_container table tbody tr td .actions .btn:hover {
    background-color: var(--violet-transparent);
  }

  :global(.transaction_template_detail .table_container table tbody tr td .actions .btn:hover svg) {
    fill: var(--violet-005);
  }
</style>