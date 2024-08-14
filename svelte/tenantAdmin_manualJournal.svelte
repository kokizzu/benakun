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

  import { Icon } from './node_modules/svelte-icons-pack/dist';
  import { RiSystemAddBoxLine } from './node_modules/svelte-icons-pack/dist/ri';
  import MainLayout from './_layouts/mainLayout.svelte';
  import MasterTable from './_components/MasterTable.svelte';
  import InputCustom from './_components/InputCustom.svelte';
  import { onMount } from 'svelte';
  import { TenantAdminManualJournal } from './jsApi.GEN';
  import { notifier } from './_components/notifier';

  let segments              = /** @type Access */ ({/* segments */});
  let fields                = /** @type Field[] */ ([/* fields */]);
  let pager                 = /** @type PagerOut */ ({/* pager */});
  let user                  = /** @type User */ ({/* user */});
  let transactionJournals   = /** @type any[][] */ ([/* transactionJournals */]);
  let coas                  = /** @type Record<number, string> */ ({/* coas */});
  let transactionTemplates  = /** @type Record<number, string> */ ({/* transactionTemplates */});
  let org                   = /** @type Org */ ({/* org */});

  let isCreatingJournal = false;
  let isSubmitted       = false;

  function dateISOFormat(/** @type number */ dayTo = 0) {
    const dt    = new Date();
    const date  = (dayTo > 0 ? dt.getDate() + dayTo : dt.getDate());
    const month = dt.toLocaleDateString('default', {month: '2-digit'});
    const year  = dt.getFullYear();

    return `${year}-${month}-${date}`;
  }

  let coaId             = 0;
  let transactionTplId  = 0;
  let startDate         = dateISOFormat(0);
  let endDate           = dateISOFormat(4);
  let debitIDR          = 0;
  let creditIDR         = 0;
  let descriptions      = '';
  let date              = dateISOFormat(0);
  let salesCount        = 0;
  let salesPriceIDR     = 0;

  function resetInput() {
    isCreatingJournal = false;
    isSubmitted       = false;
    coaId             = 0;
    transactionTplId  = 0;
    startDate         = dateISOFormat(0);
    endDate           = dateISOFormat(4);
    debitIDR          = 0;
    creditIDR         = 0;
    descriptions      = '';
    date              = dateISOFormat(0);
    salesCount        = 0;
    salesPriceIDR     = 0;
  }

  onMount(() => {
    for (const [k, v] of Object.entries(coas)) {
      coaId = Number(k);
      break;
    }

    for (const [k, v] of Object.entries(transactionTemplates)) {
      transactionTplId = Number(k);
      break;
    }
  });

  async function Submit() {
    isSubmitted = true;
    const i = {
      cmd: 'upsert',
      coaId: coaId,
      transactionTplId: transactionTplId,
      transactionJournals: {
        debitIDR: debitIDR+'',
        creditIDR: creditIDR+'',
        descriptions: descriptions,
        date: date
      },
      businessTransaction: {
        startDate: startDate,
        endDate: endDate
      }
    }
    await TenantAdminManualJournal( // @ts-ignore
      i, /** @returns {Promise<any>} */
      function (/** @type {any} */ o) {
        isSubmitted = false;
        if (o.error) {
          notifier.showError(o.error || 'failed to add journal');
          return
        }

        resetInput();
        notifier.showSuccess('journal added !');
      }
    )
  }
</script>

<MainLayout>
  <div class={!isCreatingJournal ? 'show' : 'hidden'}>
    <MasterTable
      ACCESS={segments}
      REFS={{
        'coaId': coas
      }}
      bind:FIELDS={fields}
      bind:PAGER={pager}
      bind:MASTER_ROWS={transactionJournals}
      
      CAN_EDIT_ROW
      CAN_SEARCH_ROW
      CAN_DELETE_ROW
      CAN_RESTORE_ROW
    >
      {#if user.tenantCode !== ''}
        <button
          class="action_btn"
          title="add journal"
          on:click={() => (isCreatingJournal = true)}
        >
          <Icon
            color="var(--gray-007)"
            size="16"
            src={RiSystemAddBoxLine}
          />
        </button>
      {/if}
    </MasterTable>
  </div>
  <div class={isCreatingJournal ? 'show' : 'hidden'}>
    <header>
      <h1>Entry journal manual</h1>
    </header>
    <div class="data_entry_journal___container">
      <div class="forms_journal">
        <div class="form_date">
          <InputCustom
            type="date"
            className="input_custom"
            id="startDate"
            label="Start Date"
            bind:value={startDate}
          />
          <InputCustom
            id="endDate"
            label="End Date"
            bind:value={endDate}
            type="date"
          />
          <InputCustom
            id="coaId"
            label="CoA (Chart of Account)"
            bind:value={coaId}
            type="select"
            values={coas}
            isObject
          />
          <InputCustom
            id="transactionTplId"
            label="Transaction Template"
            bind:value={transactionTplId}
            type="select"
            values={transactionTemplates}
            isObject
          />
        </div>
        <div class="company_details">
          <div class="company_name">
            <div class="title">
              <span>Company Name</span>
            </div>
            <h5>{org.name || ''}</h5>
          </div>
          <div class="company_contacts">
            <div class="title">
              <span>Company Contacts</span>
            </div>
            <div class="contacts">
              <p>Head Title: {org.headTitle || ''}</p>
              <p>Email: {user.email}</p>
            </div>
          </div>
        </div>
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
            isCreatingJournal = false;
            resetInput();
          }}
        >
          Cancel
        </button>
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

  .hidden {
    display: none;
  }

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

  .data_entry_journal___container .forms_journal {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .data_entry_journal___container .forms_journal .form_date {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr;
    gap: 10px;
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
