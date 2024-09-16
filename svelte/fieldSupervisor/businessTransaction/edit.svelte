<script>
  /** @typedef {import('../../_components/types/organization.js').Org} Org */
  /** @typedef {import('../../_components/types/user.js').User} User */
  /** @typedef {import('../../_components/types/transaction.js').TransactionJournal} TransactionJournal */
  /** @typedef {import('../../_components/types/transaction.js').DetailObjectTransaction} DetailObjectTransaction */
  /** @typedef {import('../../_components/types/transaction.js').BusinessTransaction} BusinessTransaction */

  /**
   * @typedef {Object} PayloadTransactionJournals
   * @property {number} id
   * @property {string|number} creditIDR
   * @property {string|number} debitIDR
   * @property {string} descriptions
   * @property {string|Date} date
   * @property {number} coaId
   * @property {number?} salesCount?
   * @property {number?} salesPriceIDR?
   */

  import MainLayout from '../../_layouts/mainLayout.svelte';
  import { FieldSupervisorDashboard } from '../../jsApi.GEN';
  import { notifier } from '../../_components/notifier';
  import InputBox from '../../_components/InputBox.svelte';
  import { onMount } from 'svelte';

  let user = /** @type User */ ({/* user */});
  let org  = /** @type Org */ ({/* org */});
  let transaction = /** @type BusinessTransaction */ ({/* transaction */});
  let transactionJournals = /** @type TransactionJournal[] */ ([/* transactionJournals */]);

  let payloads    = /** @type PayloadTransactionJournals[] */ ([]);
  let isDataReady = false;
  let isSubmitted = false;

  let startDate     = transaction.startDate;
  let endDate       = transaction.endDate;
  
  onMount(() => {
    if (transactionJournals && transactionJournals.length > 0) {
      transactionJournals.forEach((transactionJournal) => {
        let detailObj = /** @type DetailObjectTransaction */ ({
          salesCount: 0,
          salesPriceIDR: 0
        });

        if (transactionJournal.detailObj !== '') {
          detailObj = JSON.parse(transactionJournal.detailObj);
        }
        payloads.push({
          id: transactionJournal.id,
          creditIDR: transactionJournal.creditIDR,
          debitIDR: transactionJournal.debitIDR,
          descriptions: transactionJournal.descriptions,
          date: transactionJournal.date,
          coaId: transactionJournal.coaId,
          salesCount: detailObj.salesCount, //@ts-ignore
          salesPriceIDR: detailObj.salesPriceIDR
        })
      });
    }

    isDataReady = true;
  });

  async function Submit() {
    isSubmitted = true;
    let trxJournalsPayloads = /** @type TransactionJournal[]|any */ ([]);
    for (const payload of payloads) {
      let detailObj = '';
      if (payload.salesCount > 0 || payload.salesPriceIDR > 0) {
        const detail = /** @type DetailObjectTransaction */ ({
          salesCount: payload.salesCount,
          salesPriceIDR: payload.salesPriceIDR+'',
        });
        detailObj = JSON.stringify(detail);
      }

      trxJournalsPayloads.push({
        id: payload.id,
        descriptions: payload.descriptions,
        date: payload.date,
        debit: payload.debitIDR,
        credit: payload.creditIDR,
        detailObj: detailObj
      })
    }

    const i = {
      cmd: 'upsert',
      transactionJournals: trxJournalsPayloads,
      transaction: {
        id: transaction.id,
        startDate: startDate,
        endDate: endDate
      }
    }

    await FieldSupervisorDashboard( // @ts-ignore
      i, /** @returns {Promise<any>} */
      function (/** @type {any} */ o) {
        isSubmitted = false;
        if (o.error) {
          notifier.showError(o.error || 'failed to edit transaction');
          return
        }
        
        notifier.showSuccess('transaction edited !');
      }
    )
  }
</script>

<MainLayout>
  <div class="show">
    <header>
      <h1>Edit Business Transaction</h1>
    </header>
    <div class="data_entry_journal___container">
      <div class="forms_journal">
        <div class="form_date">
          <InputBox
            type="date"
            className="input_custom"
            id="startDate"
            label="Start Date"
            bind:value={startDate}
          />
          <InputBox
            id="endDate"
            label="End Date"
            bind:value={endDate}
            type="date"
          />
        </div>
        <div class="company_details">
          <div class="company_name">
            <div class="title">
              <span>Company Name</span>
            </div>
            <h5>{org.name}</h5>
          </div>
          <div class="company_contacts">
            <div class="title">
              <span>Company Contacts</span>
            </div>
            <div class="contacts">
              <p>Head Title: {org.headTitle}</p>
              <p>Email: {user.email}</p>
            </div>
          </div>
        </div>
            <div class="form_item">
              <div class="forms_table">
                <table class="table_transaction_journals">
                  <thead>
                    <tr>
                      <th>Description</th>
                      <th>Sales Count</th>
                      <th>Sales Price (IDR)</th>
                      <th>Date</th>
                      <th>Debit (IDR)</th>
                      <th>Credit (IDR)</th>
                    </tr>
                  </thead>
                  <tbody>
                    {#if isDataReady}
                      {#each (payloads || []) as py, idx}
                        <tr>
                          <td>
                            <textarea
                              placeholder="Description"
                              rows="1"
                              bind:value={py.descriptions}
                            />
                          </td>
                          <td>
                            <input
                              inputmode="numeric" 
                              type="number" min=0
                              bind:value={py.salesCount}
                            />
                          </td>
                          <td>
                            <input
                              inputmode="numeric" 
                              type="number" min=0
                              bind:value={py.salesPriceIDR}
                            />
                          </td>
                          <td>
                            <input type="date" bind:value={py.date}/>
                          </td>
                          <td>
                            <input
                              inputmode="numeric" 
                              type="number" min=0
                              bind:value={py.debitIDR}
                            />
                          </td>
                          <td>
                            <input
                              inputmode="numeric" 
                              type="number" min=0
                              bind:value={py.creditIDR}
                            />
                          </td>
                        </tr>
                      {/each}
                    {/if}
                  </tbody>
                </table>
              </div>
            </div>
      </div>
      <div class="actions">
        <button
          disabled={isSubmitted}
          class="btn submit"
          on:click={Submit}
        >
          {#if !isSubmitted}
            <span>Submit</span>
          {/if}
          {#if isSubmitted}
            <span>Submitting ...</span>
          {/if}
        </button>
      </div>
    </div>
  </div>
</MainLayout>

<style>
  .show {
    display: block;
  }

  header {
    margin-left: 30px;
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

  .data_entry_journal___container {
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 20px;
    border-radius: 8px;
    background-color: #fff;
    border: 1px solid var(--gray-003);
  }

  .data_entry_journal___container .forms_journal {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .data_entry_journal___container .forms_journal .form_date {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 10px;
  }

  .data_entry_journal___container .forms_journal .form_item {
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 15px;
    background-color: var(--gray-001);
    border-radius: 8px;
    border: 1px solid var(--gray-003);
  }

  .data_entry_journal___container .company_details {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: start;
    gap: 100px;
    padding: 30px 0;
  }

  .data_entry_journal___container .company_details .company_name,
  .data_entry_journal___container .company_details .company_contacts {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .data_entry_journal___container .company_details .company_name .title,
  .data_entry_journal___container .company_details .company_contacts .title {
    font-size: var(--font-lg);
    border-bottom: 2px solid var(--blue-005);
    padding-bottom: 8px;
  }

  .data_entry_journal___container .company_details .company_name h5 {
    font-weight: 600;
    font-size: 16px;
    line-height: 24px;
    color: var(--blue-005);
    margin: 0;
  }

  .data_entry_journal___container .company_details .company_contacts .contacts {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .data_entry_journal___container .company_details .company_contacts .contacts p {
    margin: 0;
  }

  .data_entry_journal___container .forms_journal .actions {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    gap: 10px;
  }

  .data_entry_journal___container .forms_journal .actions .btn {
    padding: 8px 13px;
    border-radius: 9999px;
    border: none;
    color: #FFF;
    cursor: pointer;
    font-weight: 600;
    background-color: var(--blue-006);
  }

  .data_entry_journal___container .forms_journal .actions .btn:hover {
    background-color: var(--blue-005);
  }

  .data_entry_journal___container .forms_journal .forms_table {
    overflow-x: auto;
    scrollbar-color: var(--gray-003) transparent;
    scrollbar-width: thin;
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals thead tr th {
    font-weight: normal;
    text-align: left;
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals input,
  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals textarea {
    padding: 10px 12px;
    border-radius: 8px;
    border: 1px solid var(--gray-003);
    width: 100%;
    resize: vertical;
    background-color: #FFF;
  }

  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals input:focus,
  .data_entry_journal___container .forms_journal .forms_table .table_transaction_journals textarea:focus {
    border-color: var(--sky-005);
    outline: 1px solid var(--sky-005);
  }

  .data_entry_journal___container .actions {
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    gap: 10px;
  }

  .data_entry_journal___container .actions .btn {
    padding: 8px 20px;
    border-radius: 9999px;
    border: none;
    cursor: pointer;
    font-weight: 600;
  }

  .data_entry_journal___container .actions .btn.submit {
    color: #FFF;
    background-color: var(--green-006);
  }

  .data_entry_journal___container .actions .btn.submit:hover {
    background-color: var(--green-005);
  }

  .data_entry_journal___container .actions .btn.submit:disabled {
    cursor: not-allowed;
    background-color: var(--gray-003);
    color: var(--gray-007);
  }
</style>