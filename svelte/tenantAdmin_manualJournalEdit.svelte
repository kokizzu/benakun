<script>
  /** @typedef {import('./_components/types/master.js').Field} Field */
	/** @typedef {import('./_components/types/access.js').Access} Access */
  /** @typedef {import('./_components/types/master.js').PagerIn} PagerIn */
	/** @typedef {import('./_components/types/master.js').PagerOut} PagerOut */
  /** @typedef {import('./_components/types/user.js').User} User */
  /** @typedef {import('./_components/types/product.js').Product} Product */
  /** @typedef {import('./_components/types/coa.js').CoA} CoA */
  /** @typedef {import('./_components/types/transaction.js').TransactionTemplate} TransactionTemplate */
  /** @typedef {import('./_components/types/transaction.js').TransactionJournal} TransactionJournal */
  /** @typedef {import('./_components/types/transaction.js').DetailObjectTransaction} DetailObjectTransaction */
  /** @typedef {import('./_components/types/transaction.js').TransactionTplDetail} TransactionTplDetail */
  /** @typedef {import('./_components/types/organization.js').Org} Org */

  import MainLayout from './_layouts/mainLayout.svelte';
  import { TenantAdminManualJournal } from './jsApi.GEN';
  import { notifier } from './_components/notifier';

  let segments              = /** @type Access */ ({/* segments */});
  let user                  = /** @type User */ ({/* user */});
  let org                   = /** @type Org */ ({/* org */});
  let coas                  = /** @type Record<number, string> */ ({/* coas */});
  let transactionJournal    = /** @type TransactionJournal */ ({/* transactionJournal */});

  console.log(transactionJournal);

  let isSubmitted = false;

  function dateISOFormat(/** @type string | Date */ dateStr) {
    const dt    = new Date(dateStr);
    const date  = dt.getDate();
    const month = dt.toLocaleDateString('default', {month: '2-digit'});
    const year  = dt.getFullYear();

    return `${year}-${month}-${date}`;
  }

  const detailObj = /** @type DetailObjectTransaction */ (
    JSON.parse(transactionJournal.detailObj || '{"salesCount": 0, "salesPriceIDR": "0"}')
  );

  let debitIDR          = transactionJournal.debitIDR;
  let creditIDR         = transactionJournal.creditIDR;
  let descriptions      = transactionJournal.descriptions;
  let date              = dateISOFormat(transactionJournal.date);
  let salesCount        = detailObj.salesCount;
  let salesPriceIDR     = detailObj.salesPriceIDR;

  async function SubmitEdit() {
    isSubmitted = true;

    if (salesPriceIDR !== 0) {
      detailObj.salesPriceIDR = salesPriceIDR+'';
      detailObj.salesCount = salesCount;      
    }
    const i = {
      cmd: 'upsert',
      transactionJournal: {
        id: transactionJournal.id,
        debitIDR: debitIDR+'',
        creditIDR: creditIDR+'',
        descriptions: descriptions,
        date: date,
        detailObj: salesPriceIDR !== 0 ? JSON.stringify(detailObj) : ''
      },
    }
    await TenantAdminManualJournal( // @ts-ignore
      i, /** @returns {Promise<any>} */
      function (/** @type {any} */ o) {
        isSubmitted = false;
        if (o.error) {
          notifier.showError(o.error || 'failed to edit journal');
          return
        }

        notifier.showSuccess('journal edited !');
      }
    )
  }
</script>

<MainLayout>
  <div>
    <header>
      <h1>Edit journal</h1>
    </header>
    <div class="data_entry_journal___container">
      <div class="details">
        <div class="left">
          <div class="item">
            <div class="title">
              <span>Company Name</span>
            </div>
            <h5>{org.name || ''}</h5>
          </div>
          <div class="item">
            <div class="title">
              <span>Company Contacts</span>
            </div>
            <div class="detail">
              <p>Head Title: {org.headTitle || ''}</p>
              <p>Email: {user.email}</p>
            </div>
          </div>
        </div>
        <div class="right">
          <div class="item">
            <div class="title">
              <span>CoA (Chart of Account)</span>
            </div>
            <div class="detail">
              <p>{coas[transactionJournal.coaId]}</p>
            </div>
          </div>
        </div>
      </div>
      <div class="forms_journal">
        <div class="forms_table">
          <table class="table_transaction_journals">
            <thead>
              <tr>
                <th>Debit (IDR)</th>
                <th>Credit (IDR)</th>
                <th>Description</th>
                <th>Date</th>
                <th>Sales Count</th>
                <th>Sales Price (IDR)</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>
                  <input type="number" min=0 bind:value={debitIDR}/>
                </td>
                <td>
                  <input type="number" min=0 bind:value={creditIDR}/>
                </td>
                <td>
                  <textarea
                    placeholder="Description"
                    bind:value={descriptions}
                    rows="1"
                  />
                </td>
                <td>
                  <input type="date" bind:value={date}/>
                </td>
                <td>
                  <input type="number" min=0 bind:value={salesCount}/>
                </td>
                <td>
                  <input type="number" min=0 bind:value={salesPriceIDR}/>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="actions">
        <button
          class="btn cancel"
          on:click={() => {
            window.location.href = '/tenantAdmin/manualJournal';
          }}
        >
          Cancel
        </button>
        <button
          disabled={isSubmitted}
          class="btn submit"
          on:click={SubmitEdit}
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
  header {
    margin-left: 30px;
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

  .data_entry_journal___container .details {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: start;
    gap: 10px;
  }

  .data_entry_journal___container .details .left,
  .data_entry_journal___container .details .right {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: start;
    gap: 50px;
    padding: 30px 0;
  }

  .data_entry_journal___container .details .item {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .data_entry_journal___container .details .item .title {
    font-size: var(--font-lg);
    border-bottom: 2px solid var(--blue-005);
    padding-bottom: 8px;
  }

  .data_entry_journal___container .details .item h5 {
    font-weight: 600;
    font-size: 16px;
    line-height: 24px;
    color: var(--blue-005);
    margin: 0;
  }

  .data_entry_journal___container .details .item .detail {
    display: flex;
    flex-direction: column;
    gap: 5px;
  }

  .data_entry_journal___container .details .item .detail p {
    margin: 0;
  }

  .data_entry_journal___container .forms_journal {
    display: flex;
    flex-direction: column;
    gap: 10px;
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

  .data_entry_journal___container .actions .btn.cancel {
    background-color: transparent;
  }

  .data_entry_journal___container .actions .btn.cancel:hover {
    background-color: var(--gray-003);
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